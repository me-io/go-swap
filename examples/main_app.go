// +build ignore

package main

import (
	"fmt"
	ex "github.com/me-io/go-swap/pkg/exchanger"
	"github.com/me-io/go-swap/pkg/swap"
)

func main() {
	SwapTest := swap.NewSwap()

	SwapTest.
		AddExchanger(ex.NewGoogleApi(nil)).
		AddExchanger(ex.NewYahooApi(nil)).
		Build()

	euroToUsdRate := SwapTest.Latest("EUR/USD")
	fmt.Println(euroToUsdRate.GetRateValue())
	fmt.Println(euroToUsdRate.GetRateDateTime())
	fmt.Println(euroToUsdRate.GetExchangerName())
}
