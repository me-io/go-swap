package redis

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

// todo write more tests
var redisURL = os.Getenv(`REDIS_URL`)

func parse(s string) time.Duration {
	d, _ := time.ParseDuration(s)
	return d
}

func TestStorage_Redis_WrongURL(t *testing.T) {
	storage, err := NewStorage("wrong://wtf")
	if err == nil || storage != nil {
		t.Fail()
	}
}

func TestStorage_Redis_GetEmpty(t *testing.T) {
	storage, _ := NewStorage(redisURL)
	content := storage.Get("MY_KEY")

	assert.EqualValues(t, []byte(""), content)
}

func TestStorage_Redis_GetRateValue(t *testing.T) {
	storage, _ := NewStorage(redisURL)
	storage.Set("MY_KEY", []byte("123456"), parse("5s"))
	content := storage.Get("MY_KEY")

	assert.EqualValues(t, []byte("123456"), content)
}

func TestStorage_Redis_GetExpiredValue(t *testing.T) {
	storage, _ := NewStorage(redisURL)
	storage.Set("MY_KEY", []byte("123456"), parse("1s"))
	time.Sleep(parse("2s"))
	content := storage.Get("MY_KEY")

	assert.EqualValues(t, []byte(""), content)
}
