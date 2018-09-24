package memory

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// todo write more tests
func parse(s string) time.Duration {
	d, _ := time.ParseDuration(s)
	return d
}

func TestStorage_Memory_GetEmpty(t *testing.T) {
	storage := NewStorage()
	content := storage.Get("MY_KEY")

	assert.EqualValues(t, []byte(""), content)
}

func TestStorage_Memory_GetValue(t *testing.T) {
	storage := NewStorage()
	storage.Set("MY_KEY", []byte("123456"), parse("5s"))
	content := storage.Get("MY_KEY")

	assert.EqualValues(t, []byte("123456"), content)
}

func TestStorage_Memory_GetExpiredValue(t *testing.T) {
	storage := NewStorage()
	storage.Set("MY_KEY", []byte("123456"), parse("1s"))
	time.Sleep(parse("1s200ms"))
	content := storage.Get("MY_KEY")

	assert.EqualValues(t, []byte(""), content)
}
