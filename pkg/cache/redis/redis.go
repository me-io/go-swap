package redis

import (
	r "github.com/go-redis/redis"
	"time"
)

var prefix = "_SWAP_CACHE_"

//Storage mechanism for caching strings in memory
type Storage struct {
	client *r.Client
}

//NewStorage creates a new redis storage
func NewStorage(url string) (*Storage, error) {
	var (
		opts *r.Options
		err  error
	)

	if opts, err = r.ParseURL(url); err != nil {
		return nil, err
	}

	return &Storage{
		client: r.NewClient(opts),
	}, nil
}

//Get a cached content by key
func (s Storage) Get(key string) []byte {
	val, _ := s.client.Get(prefix + key).Bytes()
	return val
}

//Set a cached content by key
func (s Storage) Set(key string, content []byte, duration time.Duration) {
	s.client.Set(prefix+key, content, duration)
}
