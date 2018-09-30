package exchanger

import (
	"github.com/me-io/go-swap/test/staticMock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFixerApi_Latest(t *testing.T) {
	rate := NewFixerApi(nil)
	assert.Equal(t, rate.name, `fixer`)

	rate.Client.Transport = staticMock.NewTestMT()

	rate.Latest(`EUR`, `EUR`)
	assert.Equal(t, float64(1), rate.GetRateValue())

	rate.Latest(`EUR`, `USD`)
	assert.Equal(t, float64(3724.305775), rate.GetRateValue())
}
