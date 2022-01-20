package define

import (
	"encoding/json"
	"fmt"
	"strings"
)

var C Config

type Config struct {
	BotToken string
}

type Response struct {
	Quote                      string `json:"quote,omitempty"`
	QuoteGasAdjusted           string `json:"quoteGasAdjusted,omitempty"`
	EstimatedGasUsed           string `json:"estimatedGasUsed,omitempty"`
	EstimatedGasUsedQuoteToken string `json:"estimatedGasUsedQuoteToken,omitempty"`
	EstimatedGasUsedUSD        string `json:"estimatedGasUsedUSD,omitempty"`
	GasPriceWei                string `json:"gasPriceWei,omitempty"`
	BlockNumber                string `json:"blockNumber,omitempty"`
	RouteString                string `json:"routeString,omitempty"`
	Message                    string `json:"message,omitempty"`
}

type Alert struct {
	ID      int64
	Name    string
	Address string
	Price   float64
	Notify  string
}

type ScanResponse struct {
	Status  int             `json:"status,string"`
	Message string          `json:"message"`
	Result  json.RawMessage `json:"result"`
}

type Transfer struct {
	BlockNumber  string `json:"blockNumber"`
	TimeStamp    int64  `json:"timeStamp,string"`
	Hash         string `json:"hash"`
	From         string `json:"from"`
	Value        string `json:"value"`
	TokenSymbol  string `json:"tokenSymbol"`
	TokenDecimal int    `json:"tokenDecimal,string"`
}

func (t *Transfer) String() (s string) {
	if n := len(t.Value); n > t.TokenDecimal {
		s = fmt.Sprintf("%s.%s", t.Value[:n-t.TokenDecimal], t.Value[n-t.TokenDecimal:])
	} else {
		s = fmt.Sprintf("0.%s%s", strings.Repeat("0", t.TokenDecimal-n), t.Value)
	}
	return fmt.Sprintf("%s %s", strings.TrimRight(s, "0"), t.TokenSymbol)
}

type PricesResponse struct {
	Message string   `json:"message,omitempty"`
	Prices  []*Price `json:"prices,omitempty"`
}

type Price struct {
	Time  int64   `json:"time"`
	Value float64 `json:"value"`
}

const (
	WBTCAddress = "0x1bfd67037b42cf73acf2047067bd4f2c47d9bfd6"

	WETHAddress = "0x7ceb23fd6bc0add59e62ac25578270cff1b9f619"

	WMATICAddress = "0x0d500b1d8e8ef31e21c99d1db9a6444d3adf1270"

	UNIAddress = "0xb33eaad8d922b1083446dc23f610c2567fb5180f"

	CRVAddress = "0x172370d5cd63279efa6d502dab29171933a610af"

	MANAAddress = "0xa1c57f48f0deb89f569dfbe6e2b7f46d33606fd4"

	LINKAddress = "0x53e0bca35ec356bd5dddfebbd1fc0fd03fabad39"

	MKRAddress = "0x6f7c932e7684666c9fd1d44527765433e01ff61d"

	AVAXAddress = "0x2c89bbc92bd86f8075d1decc58c7f4e0107f286b"

	LUNAAddress = "0x24834bbec7e39ef42f4a75eaf8e5b6486d3f0e57"

	PanShiAddress = "0xa67153e17bb2f4b51b127c3dd3869b7bc3e256c1"

	ZhuGeAddress = "0x1197a5441250d3de76e23a3a1ca1c73949e3d4d8"
)
