package main

import (
	"fmt"
	ex "github.com/me-io/go-swap/exchanger"
	"github.com/me-io/go-swap/swap"
)

func main() {
	SwapTest := swap.NewSwap()

	SwapTest.
		AddExchanger(ex.NewGoogleApi(nil)).
		AddExchanger(ex.NewYahooApi(nil)).
		Build()

	euroToUsdRate := SwapTest.Latest("EUR/USD")
	fmt.Println(euroToUsdRate.GetValue())
	fmt.Println(euroToUsdRate.GetDate())
	fmt.Println(euroToUsdRate.GetExchangerName())
}