package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/panshiqu/analyst/cache"
	"github.com/panshiqu/analyst/define"
	"github.com/panshiqu/analyst/utils"
)

var atom int32

func PerMinute() {
	defer atomic.AddInt32(&atom, -1)

	if atomic.AddInt32(&atom, 1) != 1 {
		return
	}

	log.Println(Handle(0, "", "get btc"))
	log.Println(Handle(0, "", "get eth"))
	log.Println(Handle(0, "", "get matic"))
	log.Println(Handle(0, "", "get link"))
	log.Println(Handle(0, "", "get uni"))
}

func Handle(id int64, name, text string) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover", err)
		}
	}()

	params := strings.Split(text, " ")
	t := strings.ToLower(text)

	switch {
	case strings.HasPrefix(t, "get price"):
		return getPrice(params[2:])
	case strings.HasPrefix(t, "getprice"):
		return getPrice(params[1:])
	case strings.HasPrefix(t, "get"):
		return getPrice(params[1:])

	case strings.HasPrefix(t, "set alert"):
		return setAlert(id, name, params[2:])
	case strings.HasPrefix(t, "setalert"):
		return setAlert(id, name, params[1:])
	case strings.HasPrefix(t, "set"):
		return setAlert(id, name, params[1:])

	case strings.HasPrefix(t, "analyse cost"):
		return analyseCost(name, params[2:])
	case strings.HasPrefix(t, "analysecost"):
		return analyseCost(name, params[1:])
	case strings.HasPrefix(t, "analyse"):
		return analyseCost(name, params[1:])
	case strings.HasPrefix(t, "ana"):
		return analyseCost(name, params[1:])
	}

	return "", utils.Wrap(fmt.Errorf("unsupported command: %s", text))
}

func getPrice(params []string) (string, error) {
	symbol := "BTC"
	if len(params) > 0 {
		symbol = params[0]
	}

	address, decimals, err := utils.Symbol2Address(symbol)
	if err != nil {
		return "", utils.Wrap(err)
	}

	var amount string
	switch address {
	case define.WBTCAddress:
		amount = "0.01"
	case define.WETHAddress:
		amount = "0.1"
	default:
		amount = "1"
	}
	if len(params) > 1 {
		amount = params[1]
	}

	iAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return "", utils.Wrap(err)
	}

	resp, err := http.Get(fmt.Sprintf("http://localhost:8000?address=%s&decimals=%d&amount=%s", address, decimals,
		new(big.Int).Mul(big.NewInt(int64(iAmount*10000)), new(big.Int).Exp(big.NewInt(10), big.NewInt(decimals-4), nil)).String()))
	if err != nil {
		return "", utils.Wrap(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", utils.Wrap(err)
	}

	r := &define.Response{}
	if err := json.Unmarshal(body, r); err != nil {
		return "", utils.Wrap(err)
	}

	if r.Message != "" {
		return r.Message, nil
	}

	price, err := strconv.ParseFloat(r.Quote, 64)
	if err != nil {
		return "", utils.Wrap(err)
	}

	if err := cache.NotifyAlerts(address, price/iAmount); err != nil {
		return "", utils.Wrap(err)
	}

	return fmt.Sprintf("%s\n\nPrice: %s\nGas: %s", r.RouteString, r.Quote, r.EstimatedGasUsedUSD), nil
}

func setAlert(id int64, name string, params []string) (string, error) {
	if len(params) == 0 {
		return "success", cache.NotifyCurrentAlerts(name, "")
	}

	address, _, err := utils.Symbol2Address(params[0])
	if err != nil {
		return "", utils.Wrap(err)
	}

	if len(params) == 1 {
		return "success", cache.NotifyCurrentAlerts(name, address)
	}

	price, err := strconv.ParseFloat(params[1], 64)
	if err != nil {
		return "", utils.Wrap(err)
	}

	var notify string
	if len(params) > 2 {
		notify = params[2]
	}

	return "success", cache.HandleAlert(id, name, address, price, notify)
}

func analyseCost(name string, params []string) (string, error) {
	if len(params) > 0 && params[0] != "" {
		name = params[0]
	}

	address, err := utils.Name2Address(name)
	if err != nil {
		return "", utils.Wrap(err)
	}

	start := 0
	if len(params) > 1 && params[1] != "" {
		n, err := strconv.Atoi(params[1])
		if err != nil {
			return "", utils.Wrap(err)
		}
		start = n
	}

	end := 99999999
	if len(params) > 2 && params[2] != "" {
		n, err := strconv.Atoi(params[2])
		if err != nil {
			return "", utils.Wrap(err)
		}
		end = n
	}

	resp, err := http.Get(fmt.Sprintf("https://api.polygonscan.com/api?module=account&action=tokentx&address=%s&startblock=%d&endblock=%d&page=1&offset=10000&sort=asc", address, start, end))
	if err != nil {
		return "", utils.Wrap(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", utils.Wrap(err)
	}

	r := &define.ScanResponse{}
	if err := json.Unmarshal(body, r); err != nil {
		return "", utils.Wrap(err)
	}

	if r.Status != 1 || !strings.HasPrefix(r.Message, "OK") {
		return fmt.Sprintf("Status: %d\nMessage: %s", r.Status, r.Message), nil
	}

	var ts []*define.Transfer
	if err := json.Unmarshal(r.Result, &ts); err != nil {
		return "", utils.Wrap(err)
	}

	if len(ts) < 2 {
		return "nothing", nil
	}

	var buf bytes.Buffer
	m := make(map[string]*big.Int)
	for i := 1; i < len(ts); i++ {
		a, b := ts[i-1], ts[i]
		if a.Hash != b.Hash {
			continue
		}

		if a.From != address {
			a, b = b, a
		}

		symbol := b.TokenSymbol
		decimal := b.TokenDecimal
		p, q := a.Value, b.Value
		if a.TokenSymbol != "USDC" {
			symbol = a.TokenSymbol
			decimal = a.TokenDecimal
			p, q = utils.Neg(q), utils.Neg(p)
		}

		us := fmt.Sprintf("usd%s", symbol)
		bs := fmt.Sprintf("btc%s", symbol)

		usd, ok := new(big.Int).SetString(p, 10)
		if !ok {
			return "", utils.Wrap(fmt.Errorf("big.Int.SetString %s", p))
		}

		btc, ok := new(big.Int).SetString(q, 10)
		if !ok {
			return "", utils.Wrap(fmt.Errorf("big.Int.SetString %s", q))
		}

		// polygon 0x92cae7576fe8c3165c7c113b8328bbc23d641ec9ddc28dc5472bd187cf4cb0fb
		if a.TokenSymbol == b.TokenSymbol {
			ts[i].Value = new(big.Int).Add(usd, btc).String()
			ts[i-1].Value = "0"
			continue
		}

		utils.MapAdd(m, us, usd)
		utils.MapAdd(m, bs, btc)

		fmt.Fprintf(&buf, "%s %s\n%s > %s\n%.6f / %.6f\n\n", time.Unix(a.TimeStamp, 0).Format("01-02 15:04"),
			a.BlockNumber, a, b, utils.Avg(usd, btc, decimal), utils.Avg(m[us], m[bs], decimal))

		i++
	}

	return buf.String(), nil
}
