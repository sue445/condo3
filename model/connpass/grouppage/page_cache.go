package grouppage

import (
	"encoding/json"
	"github.com/memcachier/mc"
	"github.com/sue445/condo3/model"
	"time"
)

const (
	keyPrefix = "connpass.grouppage.pageCache-v1-"
)

// pageCache represents page cache
type pageCache struct {
	memcached *mc.Client
}

// Quit must be called when finalize
type Quit func()

// newPageCache returns new pageCache instance
func newPageCache(memcachedConfig *model.MemcachedConfig) (*pageCache, Quit) {
	memcached := mc.NewMC(memcachedConfig.Server, memcachedConfig.Username, memcachedConfig.Password)
	return &pageCache{memcached: memcached}, memcached.Quit
}

// get returns value from memcache
func (c *pageCache) get(key string) (*Page, error) {
	value, _, _, err := c.memcached.Get(keyPrefix + key)

	if err != nil {
		if err == mc.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}

	var page Page
	err = json.Unmarshal([]byte(value), &page)

	if err != nil {
		return nil, err
	}

	return &page, nil
}

// set sets value to memcache
func (c *pageCache) set(key string, page *Page) error {
	bytes, err := json.Marshal(page)

	if err != nil {
		return err
	}

	expiration := time.Hour * 24 // 1 day
	_, err = c.memcached.Set(keyPrefix+key, string(bytes), 0, uint32(expiration.Seconds()), 0)

	if err != nil {
		return err
	}

	return nil
}
