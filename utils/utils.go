package utils

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/panshiqu/analyst/define"
)

func Symbol2Address(symbol string) (string, int64, error) {
	if strings.HasPrefix(symbol, "0x") {
		ss := strings.Split(symbol, ",")
		if len(ss) == 1 {
			return ss[0], 18, nil
		}
		decimals, err := strconv.ParseInt(ss[1], 10, 64)
		if err != nil {
			return "", 18, Wrap(err)
		}
		return ss[0], decimals, nil
	}

	switch strings.ToUpper(symbol) {
	case "LUNA":
		return define.LUNAAddress, 18, nil
	case "AVAX":
		return define.AVAXAddress, 18, nil
	case "MKR":
		return define.MKRAddress, 18, nil
	case "LINK":
		return define.LINKAddress, 18, nil
	case "MANA":
		return define.MANAAddress, 18, nil
	case "CRV":
		return define.CRVAddress, 18, nil
	case "UNI":
		return define.UNIAddress, 18, nil
	case "MATIC", "WMATIC":
		return define.WMATICAddress, 18, nil
	case "ETH", "WETH":
		return define.WETHAddress, 18, nil
	case "BTC", "WBTC":
		return define.WBTCAddress, 8, nil
	default:
		return "", 18, Wrap(fmt.Errorf("unsupported symbol: %s", symbol))
	}
}

func Address2Symbol(address string) string {
	switch address {
	case define.LUNAAddress:
		return "LUNA"
	case define.AVAXAddress:
		return "AVAX"
	case define.MKRAddress:
		return "MKR"
	case define.LINKAddress:
		return "LINK"
	case define.MANAAddress:
		return "MANA"
	case define.CRVAddress:
		return "CRV"
	case define.UNIAddress:
		return "UNI"
	case define.WMATICAddress:
		return "MATIC"
	case define.WETHAddress:
		return "ETH"
	case define.WBTCAddress:
		return "BTC"
	}
	return address
}

func Name2Address(name string) (string, error) {
	if len(name) == 42 && strings.HasPrefix(name, "0x") {
		return strings.ToLower(name), nil
	}

	switch strings.ToUpper(name) {
	case "PANSHI", "PANSHIQU":
		return define.PanShiAddress, nil
	case "ZHUGE", "ZHUGEFEIA":
		return define.ZhuGeAddress, nil
	default:
		return "", Wrap(fmt.Errorf("unsupported name: %s", name))
	}
}

func IsDisableNotification(notify string) bool {
	if notify == "" {
		return true
	}
	if notify == "-" {
		return false
	}
	h := time.Now().Hour()
	ss := strings.Split(notify, ",")
	for _, v := range ss {
		var a, b int
		if n, err := fmt.Sscanf(v, "%d-%d", &a, &b); err != nil || n != 2 {
			return true
		}
		if h >= a && h < b {
			return false
		}
	}
	return true
}

func MapAdd(m map[string]*big.Int, k string, v *big.Int) {
	if _, ok := m[k]; !ok {
		m[k] = v
	} else {
		m[k] = new(big.Int).Add(m[k], v)
	}
}

func Neg(v string) string {
	return fmt.Sprintf("-%s", v)
}

func Avg(a *big.Int, b *big.Int, d int) float64 {
	return float64(new(big.Int).Div(new(big.Int).Mul(a, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(d)), nil)), b).Int64()) / 1000000
}
