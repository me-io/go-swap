package swap

import (
	p "github.com/me-io/go-swap/provider"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestClientRequestMiddleware(t *testing.T) {
	client := NewSwap()
	assert.Equal(t, "*swap.Swap", reflect.TypeOf(client).String())
	//client.Add();
}

func TestAll(t *testing.T) {
	SwapTest := NewSwap()

	SwapTest.
		//Add(p.NewCurrencyLayerApi(), map[string]string{"access_key": "your-access-key"}).
		Add(p.NewGoogleApi(), nil).
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
