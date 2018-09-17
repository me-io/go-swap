package exchanger

import (
	"github.com/me-io/go-swap/static_mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoogleApi_Latest(t *testing.T) {
	rate := NewGoogleApi()
	rate.Client.Transport = static_mock.NewTestMT()

	rate.Latest(`EUR`, `EUR`)
	assert.Equal(t, float64(1), rate.GetValue())

	rate.Latest(`EUR`, `USD`)
	assert.Equal(t, float64(3.67), rate.GetValue())
}
