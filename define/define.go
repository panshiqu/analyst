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

const (
	WBTCAddress = "0x1BFD67037B42Cf73acF2047067bd4F2C47D9BfD6"

	WETHAddress = "0x7ceB23fD6bC0adD59E62ac25578270cFf1b9f619"

	WMATICAddress = "0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270"
)
