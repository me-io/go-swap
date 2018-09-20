package exchanger

import (
	"github.com/me-io/go-swap/pkg/staticMock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYahooApi_Latest(t *testing.T) {
	rate := NewYahooApi(nil)
	rate.Client.Transport = staticMock.NewTestMT()

	rate.Latest(`USD`, `EUR`)
	assert.Equal(t, float64(0.2723), rate.GetValue())
}
