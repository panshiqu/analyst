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

const (
	WBTCAddress = "0x1bfd67037b42cf73acf2047067bd4f2c47d9bfd6"

	WETHAddress = "0x7ceb23fd6bc0add59e62ac25578270cff1b9f619"

	WMATICAddress = "0x0d500b1d8e8ef31e21c99d1db9a6444d3adf1270"

	PanShiAddress = "0xa67153e17bb2f4b51b127c3dd3869b7bc3e256c1"

	ZhuGeAddress = "0x1197a5441250d3de76e23a3a1ca1c73949e3d4d8"
)
