package exchanger

import (
	"github.com/me-io/go-swap/staticMock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYahooApi_Latest(t *testing.T) {
	rate := NewYahooApi()
	rate.Client.Transport = staticMock.NewTestMT()

	rate.Latest(`USD`, `EUR`)
	assert.Equal(t, float64(0.2723), rate.GetValue())
}
