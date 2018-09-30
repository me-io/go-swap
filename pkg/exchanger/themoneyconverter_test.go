package exchanger

import (
	"github.com/me-io/go-swap/test/staticMock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTheMoneyConverterApi_Latest(t *testing.T) {
	rate := NewTheMoneyConverterApi(nil)
	assert.Equal(t, rate.name, `themoneyconverter`)

	rate.Client.Transport = staticMock.NewTestMT()

	rate.Latest(`EUR`, `EUR`)
	assert.Equal(t, float64(1), rate.GetRateValue())

	rate.Latest(`USD`, `AED`)
	assert.Equal(t, float64(3.6725), rate.GetRateValue())
}
