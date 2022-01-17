package cache_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"translateapp/internal/cache"
	"translateapp/internal/logging"
)

func TestCache_Get(t *testing.T) {
	size := 1
	TTL := 1 * time.Second
	testCache := cache.NewCache(size, TTL, logging.DefaultLogger())
	testCache.Set("mouse", "mysz")
	expectedOutput := "mysz"
	assert.Equal(t, expectedOutput, testCache.Get("mouse"))
	testCache.Close()
}

func TestCache_Cleanup(t *testing.T) {
	size := 1
	TTL := 2 * time.Second
	testCache := cache.NewCache(size, TTL, logging.DefaultLogger())
	testCache.Set("mouse", "mysz")
	time.Sleep(5 * time.Second)
	assert.Equal(t, "", testCache.Get("mouse"))
	testCache.Close()
}
