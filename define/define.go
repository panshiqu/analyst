package define

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

const (
	WBTCAddress = "0x1bfd67037b42cf73acf2047067bd4f2c47d9bfd6"

	WETHAddress = "0x7ceb23fd6bc0add59e62ac25578270cff1b9f619"

	WMATICAddress = "0x0d500b1d8e8ef31e21c99d1db9a6444d3adf1270"
)
