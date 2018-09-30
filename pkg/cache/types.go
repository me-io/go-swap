package cache

import "time"

// Storage ... Cache Storage Interface
type Storage interface {
	Get(key string) []byte
	Set(key string, content []byte, duration time.Duration)
}
