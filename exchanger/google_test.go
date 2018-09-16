package exchanger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoogleApi_Latest(t *testing.T) {
	rate := NewGoogleApi()
	rate.Latest(`EUR`, `EUR`)
	assert.Equal(t, float64(1), rate.GetValue())

	rate.Latest(`EUR`, `USD`)
	assert.Equal(t, float64(1.16), rate.GetValue())

	rate.Latest(`USD`, `EUR`)
	assert.Equal(t, float64(0.86), rate.GetValue())
}
