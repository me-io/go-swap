package swap

import (
	p "github.com/me-io/go-swap/provider"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestClientRequestMiddleware(t *testing.T) {
	client := NewBuilder()
	assert.Equal(t, "*swap.Builder", reflect.TypeOf(client).String())
	//client.Add();
}

func TestAll(t *testing.T) {
	BuilderTest := NewBuilder()

	BuilderTest.
		Add(p.NewGoogleApi(), map[string]string{"access_key": "your-access-key"}).
		Add(p.NewCurrencyLayerApi(), map[string]string{"access_key": "your-access-key"}).
		Add(p.NewGoogleApi(), map[string]string{"access_key": "your-access-key"}).
		Add(p.NewCurrencyLayerApi(), map[string]string{"access_key": "your-access-key"}).
		Build()

	//BuilderTest = NewBuilder()
	BuilderTest.
		Add(p.NewGoogleApi(), map[string]string{"access_key": "your-access-key"}).
		Add(p.NewCurrencyLayerApi(), map[string]string{"access_key": "your-access-key"}).
		Build()

	rate := BuilderTest.latest("EUR/USD")

	println(rate.GetValue())

	//var RateTest = SwapTest.latest("EUR/USD")
	//assert.Equal(t, "", reflect.TypeOf(RateTest).Name())

	// 1.129
	//var value = RateTest.value
	//2016-08-26
	//var date = RateTest.date
}
