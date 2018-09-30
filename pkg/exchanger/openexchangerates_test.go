package exchanger

import (
	"github.com/me-io/go-swap/test/staticMock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenExchangeRatesApi_Latest(t *testing.T) {
	rate := NewOpenExchangeRatesApi(nil)
	assert.Equal(t, rate.name, `openexchangerates`)

	rate.Client.Transport = staticMock.NewTestMT()

	rate.Latest(`EUR`, `EUR`)
	assert.Equal(t, float64(1), rate.GetRateValue())

	rate.Latest(`USD`, `AED`)
	assert.Equal(t, float64(3.6571), rate.GetRateValue())
}
