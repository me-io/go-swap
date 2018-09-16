package swap

import (
	ex "github.com/me-io/go-swap/exchanger"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSwap_New(t *testing.T) {
	SwapTest := NewSwap()
	assert.Equal(t, "*swap.Swap", reflect.TypeOf(SwapTest).String())
}

//func TestSwap_Latest_Error(t *testing.T) {
//	SwapTest := NewSwap()
//	assert.Equal(t, "*swap.Swap", reflect.TypeOf(SwapTest).String())
//	SwapTest.Build()
//	SwapTest.Latest("EUR/USD")
//}

func TestSwap_AddExchanger(t *testing.T) {
	SwapTest := NewSwap()
	SwapTest.
		AddExchanger(ex.NewGoogleApi()).
		AddExchanger(ex.NewYahooApi()).
		Build()
	assert.Equal(t, "*swap.Swap", reflect.TypeOf(SwapTest).String())
}

func TestSwap_Build_Google(t *testing.T) {
	SwapTest := NewSwap()

	SwapTest.
		AddExchanger(ex.NewGoogleApi()).
		Build()

	euroToUsdRate := SwapTest.Latest("EUR/USD")
	assert.Equal(t, float64(1.16), euroToUsdRate.GetValue())
	assert.Equal(t, `GoogleApi`, euroToUsdRate.GetExchangerName())

	// usdToUsdRate := SwapTest.latest("USD/USD")
	// assert.Equal(t, float64(1), usdToUsdRate.GetValue())
}

func TestSwap_Build_Yahoo(t *testing.T) {
	SwapTest := NewSwap()

	SwapTest.
		AddExchanger(ex.NewYahooApi()).
		Build()

	euroToUsdRate := SwapTest.Latest("EUR/USD")
	assert.Equal(t, float64(1.169), euroToUsdRate.GetValue())
	assert.Equal(t, `YahooApi`, euroToUsdRate.GetExchangerName())
}

func TestSwap_Build_Stack_Yahoo_Google(t *testing.T) {
	SwapTest := NewSwap()

	SwapTest.
		AddExchanger(ex.NewYahooApi()).
		AddExchanger(ex.NewGoogleApi()).
		Build()

	euroToUsdRate := SwapTest.Latest("EUR/USD")
	assert.Equal(t, float64(1.169), euroToUsdRate.GetValue())
	assert.Equal(t, `YahooApi`, euroToUsdRate.GetExchangerName())
}
