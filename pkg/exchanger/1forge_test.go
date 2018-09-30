package exchanger

import (
	"github.com/me-io/go-swap/test/staticMock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOneForgeApi_Latest(t *testing.T) {
	rate := NewOneForgeApi(nil)
	assert.Equal(t, rate.name, `1forge`)

	rate.Client.Transport = staticMock.NewTestMT()

	rate.Latest(`EUR`, `EUR`)
	assert.Equal(t, float64(1), rate.GetRateValue())

	rate.Latest(`USD`, `AED`)
	assert.Equal(t, float64(3.675), rate.GetRateValue())
}
