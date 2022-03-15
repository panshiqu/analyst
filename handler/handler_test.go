package handler

import (
	"fmt"
	"testing"

	"github.com/panshiqu/analyst/cache"
)

func TestHandle(t *testing.T) {
	Handle(1, "panshiqu", "Get price btc")
	Handle(1, "panshiqu", "Getprice btc")
	Handle(1, "panshiqu", "Get btc")

	Handle(1, "panshiqu", "Set alert btc")
	Handle(1, "panshiqu", "Setalert btc")
	Handle(1, "panshiqu", "Set btc")

	Handle(1, "panshiqu", "Analyse cost btc")
	Handle(1, "panshiqu", "Analysecost btc")
	Handle(1, "panshiqu", "Analyse btc")
	Handle(1, "panshiqu", "Ana btc")

	if _, err := Handle(1, "panshiqu", "abc btc"); err != nil {
		fmt.Println(err)
	}
}

func TestGetPrice(t *testing.T) {
	fmt.Println(Handle(1, "panshiqu", "get"))

	// fmt.Println(getPrice([]string{"WETH"}))

	// fmt.Println(getPrice([]string{"WMATIC"}))
	// return
	// fmt.Println(getPrice([]string{"WBTC", "0.01"}))
	// fmt.Println(getPrice([]string{"WBTC", "0.1"}))
	// fmt.Println(getPrice([]string{"WBTC", "1"}))

	// fmt.Println(getPrice([]string{"WETH", "0.1"}))
	// fmt.Println(getPrice([]string{"WETH", "1"}))

	// fmt.Println(getPrice([]string{"WMATIC", "1"}))

	// fmt.Println(getPrice([]string{"0xb33eaad8d922b1083446dc23f610c2567fb5180f,18", "1"}))
}

func TestSetAlert(t *testing.T) {
	fmt.Println(Handle(1, "panshiqu", "set btc 428.00 "))
	fmt.Println(Handle(1, "panshiqu", "set eth 300 "))
	// fmt.Println(Handle(1, "panshiqu", "set btc 440 8-13,15-23"))
	cache.PrintAlerts()
	fmt.Println(Handle(1, "panshiqu", "set"))
	fmt.Println(Handle(1, "panshiqu", "set eth"))
	// fmt.Println(Handle(1, "panshiqu", "get btc"))
	// cache.PrintAlerts()
}

func TestAnalyseCost(t *testing.T) {
	fmt.Println(analyseCost("panshiqu", []string{"panshi", "25780001", ""}))
}
