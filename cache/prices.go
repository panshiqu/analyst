package cache

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/panshiqu/analyst/define"
	"github.com/panshiqu/analyst/utils"
)

const MaxCap = 10000

var pricesMutex sync.Mutex
var prices map[string][]*define.Price

func init() {
	prices = make(map[string][]*define.Price)
}

func PrintPrices() {
	pricesMutex.Lock()
	defer pricesMutex.Unlock()

	for k, v := range prices {
		fmt.Println(k, len(v), cap(v))
		for _, vv := range v {
			fmt.Println(vv)
		}
	}
}

func AppendPrice(address string, value float64) {
	pricesMutex.Lock()
	defer pricesMutex.Unlock()

	ps, ok := prices[address]
	if !ok || len(ps) >= MaxCap {
		prices[address] = make([]*define.Price, 0, MaxCap)
		if ok {
			prices[address] = append(prices[address], ps[MaxCap/2:]...)
		}
	}

	prices[address] = append(prices[address], &define.Price{
		Time:  time.Now().Unix(),
		Value: value,
	})
}

func ServePrices(w http.ResponseWriter, r *http.Request) error {
	n := 60
	if s := r.FormValue("n"); s != "" {
		d, err := strconv.Atoi(s)
		if err != nil {
			return utils.Wrap(err)
		}
		n = d
	}

	address, _, err := utils.Symbol2Address(r.FormValue("s"))
	if err != nil {
		return utils.Wrap(err)
	}

	pricesMutex.Lock()
	defer pricesMutex.Unlock()

	ps, ok := prices[address]
	if !ok {
		return utils.Wrap(fmt.Errorf("no %s price", r.FormValue("s")))
	}
	if len(ps) > n {
		ps = ps[len(ps)-n:]
	}

	return utils.WriteJSON(w, define.PricesResponse{Prices: ps})
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err := ServePrices(w, r); err != nil {
		utils.WriteJSON(w, define.PricesResponse{Message: err.Error()})
	}
}

const pricesFile = "/tmp/prices.json"

func SavePrices() error {
	pricesMutex.Lock()
	defer pricesMutex.Unlock()

	data, err := json.Marshal(prices)
	if err != nil {
		return utils.Wrap(err)
	}

	if err := ioutil.WriteFile(pricesFile, data, 0644); err != nil {
		return utils.Wrap(err)
	}

	return nil
}

func LoadPrices() error {
	if _, err := os.Stat(pricesFile); os.IsNotExist(err) {
		return nil
	}

	pricesMutex.Lock()
	defer pricesMutex.Unlock()

	return utils.ReadJSON(pricesFile, &prices)
}
