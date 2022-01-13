package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/panshiqu/analyst/define"
	"github.com/panshiqu/analyst/utils"
)

func Handle(name, text string) (string, error) {
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
		return setAlert(name, params[2:])
	case strings.HasPrefix(t, "setalert"):
		return setAlert(name, params[1:])
	case strings.HasPrefix(t, "set"):
		return setAlert(name, params[1:])

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

	return fmt.Sprintf("%s\n\nPrice: %s\nGas: %s", r.RouteString, r.Quote, r.EstimatedGasUsedUSD), nil
}

func setAlert(name string, params []string) (string, error) {
	return "coming soon", nil
}

func analyseCost(name string, params []string) (string, error) {
	return "coming soon...", nil
}
