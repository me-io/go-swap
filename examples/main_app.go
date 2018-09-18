package main

import (
	"fmt"
	"github.com/me-io/go-swap"
	ex "github.com/me-io/go-swap/exchanger"
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
