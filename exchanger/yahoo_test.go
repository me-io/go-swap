package exchanger

import (
	"github.com/me-io/go-swap/static_mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYahooApi_Latest(t *testing.T) {
	rate := NewYahooApi()
	rate.Client.Transport = static_mock.NewTestMT()

	rate.Latest(`USD`, `EUR`)
	assert.Equal(t, float64(0.2723), rate.GetValue())
}
