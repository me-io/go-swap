package exchanger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYahooApi_Latest(t *testing.T) {
	rateApi := NewYahooApi(nil)
	rateApi.Latest(`EUR`, `EUR`, nil)
	assert.Equal(t, float64(1), rateApi.GetValue())
}
