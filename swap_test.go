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
func TestSwap_AddExchanger(t *testing.T) {
	SwapTest := NewSwap()
	SwapTest.
		AddExchanger(ex.NewGoogleApi(), nil).
		AddExchanger(ex.NewGoogleApi(), nil).
		AddExchanger(ex.NewGoogleApi(), nil).
		Build()
	assert.Equal(t, "*swap.Swap", reflect.TypeOf(SwapTest).String())
}

func TestSwap_Build(t *testing.T) {
	SwapTest := NewSwap()

	SwapTest.
		//AddExchanger(ex.NewCurrencyLayerApi(), map[string]string{"access_key": "your-access-key"}).
		AddExchanger(ex.NewGoogleApi(), nil).
		Build()

	rate := SwapTest.latest("EUR/USD")

	//println(rate.GetValue())
	//println(rate.GetDate())

	//var RateTest = SwapTest.latest("EUR/USD")
	assert.Equal(t, float64(1.16), rate.GetValue())
}
