package exchanger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYahooApi_Latest(t *testing.T) {
	rate := NewYahooApi()
	rate.Latest(`EUR`, `EUR`)
	assert.Equal(t, float64(1), rate.GetValue())

	rate.Latest(`EUR`, `USD`)
	assert.Equal(t, float64(1.1631), rate.GetValue())

	rate.Latest(`USD`, `EUR`)
	assert.Equal(t, float64(0.8598), rate.GetValue())
}
