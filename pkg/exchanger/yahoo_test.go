package exchanger

import (
	"github.com/me-io/go-swap/test/staticMock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYahooApi_Latest(t *testing.T) {
	rate := NewYahooApi(nil)
	assert.Equal(t, rate.name, `yahoo`)

	rate.Client.Transport = staticMock.NewTestMT()

	rate.Latest(`USD`, `EUR`)
	assert.Equal(t, float64(0.272272), rate.GetValue())
}
