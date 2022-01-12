package cache

import (
	"go.uber.org/zap"
	"os"
	"sync"
	"time"
)

type Cache struct {
	//translator string
	storage map[string]*cachedItem
	l       sync.Mutex
	Logger  *zap.SugaredLogger
}

// Cacher defines interface for cache implementers.
type Cacher interface {
	Set(word, translation string)
	Get(value string) (translation string)
	Refresher(TTL int, wg *sync.WaitGroup, osShutdown chan os.Signal)
}

func NewCache(size int, logger *zap.SugaredLogger) Cacher {
	return New(size, logger)
}

func New(size int, logger *zap.SugaredLogger) (c *Cache) {
	c = &Cache{
		storage: make(map[string]*cachedItem, size),
		Logger:  logger,
	}
	logger.Info("New cache initialized")

	return c
}

func (c *Cache) Refresher(TTL int, wg *sync.WaitGroup, osShutdown chan os.Signal) {

	go func() {
		<-osShutdown
		for now := range time.Tick(time.Second) {
			c.l.Lock()
			for k, v := range c.storage {
				if now.Unix()-v.lastAccess > int64(TTL) {
					delete(c.storage, k)
				}
			}
			c.l.Unlock()
			wg.Done()
		}
	}()
}

type cachedItem struct {
	word        string
	translation string
	lastAccess  int64
}

func (c *Cache) Get(value string) (translation string) {
	c.l.Lock()
	if item, ok := c.storage[value]; ok {
		translation = item.translation
		item.lastAccess = time.Now().Unix()
		c.Logger.Info("Value got from cache")
	} else {
		c.Logger.Info("Value has not been cached yet")
	}

	c.l.Unlock()
	return
}

func (c *Cache) Set(word, translation string) {
	c.l.Lock()
	var newItem cachedItem
	newItem.word = word
	newItem.translation = translation
	newItem.lastAccess = time.Now().Unix()
	c.storage[word] = &newItem
	c.Logger.Info("New value was added to cache")
	c.l.Unlock()
	return
}

/*
func main() {
	nc := NewCache(10,10)
	nc.Set("mouse", "mysz")
	//nc.Set("mouse", "mysz")
	fmt.Printf("1st %v",nc.Get("mouse"))
	time.Sleep(5 * time.Second)
	fmt.Printf("\n2nd %v",nc.Get("mouse"))
	time.Sleep(10 * time.Second)
	fmt.Printf("\n3d %T",nc.Get("mouse"))
	time.Sleep(11* time.Second)
	fmt.Printf("\nlast %v,  %T",nc.Get("mouse"),nc.Get("mouse"))
	}

*/
