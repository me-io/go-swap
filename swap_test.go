package swap

import (
	ex "github.com/me-io/go-swap/exchanger"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestClientRequestMiddleware(t *testing.T) {
	client := NewSwap()
	assert.Equal(t, "*swap.Swap", reflect.TypeOf(client).String())
}

func TestAll(t *testing.T) {
	SwapTest := NewSwap()

	SwapTest.
		//AddExchanger(ex.NewCurrencyLayerApi(), map[string]string{"access_key": "your-access-key"}).
		AddExchanger(ex.NewGoogleApi(), nil).
		Build()

	rate := SwapTest.latest("EUR/USD")

	println(rate.GetValue())
	println(rate.GetDate())

	//var RateTest = SwapTest.latest("EUR/USD")
	assert.Equal(t, float64(1.16), rate.GetValue())

	// 1.129
	//var value = RateTest.value
	//2016-08-26
	//var date = RateTest.date
}
