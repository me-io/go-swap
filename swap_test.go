package swap

import (
	ex "github.com/me-io/go-swap/exchanger"
	"github.com/me-io/go-swap/staticMock"
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

	g := ex.NewGoogleApi()
	g.Client.Transport = staticMock.NewTestMT()
	y := ex.NewGoogleApi()
	y.Client.Transport = staticMock.NewTestMT()

	SwapTest.
		AddExchanger(g).
		AddExchanger(y).
		Build()
	assert.Equal(t, "*swap.Swap", reflect.TypeOf(SwapTest).String())
}

func TestSwap_Build_Google(t *testing.T) {
	SwapTest := NewSwap()

	g := ex.NewGoogleApi()
	g.Client.Transport = staticMock.NewTestMT()

	SwapTest.
		AddExchanger(g).
		Build()

	euroToUsdRate := SwapTest.Latest("EUR/USD")
	assert.Equal(t, float64(3.67), euroToUsdRate.GetValue())
	assert.Equal(t, `googleApi`, euroToUsdRate.GetExchangerName())

	usdToUsdRate := SwapTest.Latest("USD/USD")
	assert.Equal(t, float64(1), usdToUsdRate.GetValue())
	assert.Equal(t, `googleApi`, euroToUsdRate.GetExchangerName())
}

func TestSwap_Build_Yahoo(t *testing.T) {
	SwapTest := NewSwap()

	y := ex.NewYahooApi()
	y.Client.Transport = staticMock.NewTestMT()

	SwapTest.
		AddExchanger(y).
		Build()

	euroToUsdRate := SwapTest.Latest("EUR/USD")
	assert.Equal(t, float64(0.2723), euroToUsdRate.GetValue())
	assert.Equal(t, `yahooApi`, euroToUsdRate.GetExchangerName())
}

func TestSwap_Build_Stack_Google_Yahoo(t *testing.T) {
	SwapTest := NewSwap()
	g := ex.NewGoogleApi()
	g.Client.Transport = staticMock.NewTestMT()
	y := ex.NewGoogleApi()
	y.Client.Transport = staticMock.NewTestMT()

	SwapTest.
		AddExchanger(g).
		AddExchanger(y).
		Build()

	euroToUsdRate := SwapTest.Latest("EUR/USD")
	assert.Equal(t, float64(3.67), euroToUsdRate.GetValue())
	assert.Equal(t, `googleApi`, euroToUsdRate.GetExchangerName())
}
