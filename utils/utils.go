package utils

import (
	"fmt"
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
	case define.WMATICAddress:
		return "MATIC"
	case define.WETHAddress:
		return "ETH"
	case define.WBTCAddress:
		return "BTC"
	}
	return address
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
