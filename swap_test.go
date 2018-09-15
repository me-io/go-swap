package swap

import (
	ex "github.com/meabed/go-swap/exchanger"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSwap_New(t *testing.T) {
	SwapTest := NewSwap()
	assert.Equal(t, "*swap.Swap", reflect.TypeOf(SwapTest).String())
}
func TestSwap_AddExchanger(t *testing.T) {
	SwapTest := NewSwap()
	SwapTest.
		AddExchanger(ex.NewGoogleApi(), nil).
		AddExchanger(ex.NewYahooApi(), nil).
		Build()
	assert.Equal(t, "*swap.Swap", reflect.TypeOf(SwapTest).String())
}

func TestSwap_BuildGoogle(t *testing.T) {
	SwapTest := NewSwap()

	SwapTest.
		AddExchanger(ex.NewGoogleApi(), nil).
		Build()

	euroToUsdRate := SwapTest.latest("EUR/USD")
	assert.Equal(t, float64(1.16), euroToUsdRate.GetValue())

	// usdToUsdRate := SwapTest.latest("USD/USD")
	// assert.Equal(t, float64(1), usdToUsdRate.GetValue())
}

func TestSwap_BuildYahoo(t *testing.T) {
	SwapTest := NewSwap()

	SwapTest.
		AddExchanger(ex.NewYahooApi(), nil).
		Build()

	euroToUsdRate := SwapTest.latest("EUR/USD")
	assert.Equal(t, float64(1.169), euroToUsdRate.GetValue())
}
