package cache

import (
	"go.uber.org/zap"
	"sync"
	"time"
)

type cachedItem struct {
	word        string
	translation string
	lastAccess  time.Time
}

type Cache struct {
	//translator string
	storage  map[string]*cachedItem
	l        sync.Mutex
	logger   *zap.SugaredLogger
	cacheTTL time.Duration

	chanStop chan struct{}
	chanDone chan struct{}
}

// CacheInterface defines interface for cache implementers.
type CacheInterface interface {
	Set(word, translation string)
	Get(value string) (translation string)
}

// Creates new cache of size size and TTL of cacheTTL, which starts goroutine to auto delete unused values
func NewCache(size int, cacheTTL time.Duration, logger *zap.SugaredLogger) *Cache {
	c := Cache{
		storage:  make(map[string]*cachedItem, size),
		logger:   logger,
		cacheTTL: cacheTTL,
		chanStop: make(chan struct{}),
		chanDone: make(chan struct{}),
	}
	// routine for removing unused values
	go c.Refresher()
	logger.Info("New cache initialized")
	return &c
}

// closes cache
func (c *Cache) Close() {
	close(c.chanStop)
	<-c.chanDone
	c.logger.Info("Cache closed")
}

func (c *Cache) Refresher() {
	defer close(c.chanDone)
	chanTick := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-chanTick.C:
			c.Cleanup()
		case <-c.chanStop:
			chanTick.Stop()
			c.logger.Info("leaving refresher")
			return
		}
	}
}

func (c *Cache) Cleanup() {
	c.l.Lock()
	for k, v := range c.storage {
		if time.Since(v.lastAccess) > c.cacheTTL {
			delete(c.storage, k)
		}
	}
	c.l.Unlock()
}

func (c *Cache) Get(value string) (translation string) {
	c.l.Lock()
	if item, ok := c.storage[value]; ok {
		translation = item.translation
		item.lastAccess = time.Now()
		c.logger.Info("Value got from cache")
	} else {
		c.logger.Info("Value has not been cached yet")
	}

	c.l.Unlock()
	return
}

func (c *Cache) Set(word, translation string) {
	c.l.Lock()
	var newItem cachedItem
	newItem.word = word
	newItem.translation = translation
	newItem.lastAccess = time.Now()
	c.storage[word] = &newItem
	c.logger.Info("New value was added to cache")
	c.l.Unlock()
	return
}
