package grouppage

import (
	"encoding/json"
	"github.com/memcachier/mc"
	"github.com/sue445/condo3/model"
	"time"
)

const (
	keyPrefix = "connpass.PageCache-v1-"
)

// PageCache represents page cache
type PageCache struct {
	memcached *mc.Client
}

// Quit must be called when finalize
type Quit func()

// NewPageCache returns new PageCache instance
func NewPageCache(memcachedConfig *model.MemcachedConfig) (*PageCache, Quit) {
	memcached := mc.NewMC(memcachedConfig.Server, memcachedConfig.Username, memcachedConfig.Password)
	return &PageCache{memcached: memcached}, memcached.Quit
}

// Get returns value from memcache
func (p *PageCache) Get(key string) (*Page, error) {
	value, _, _, err := p.memcached.Get(keyPrefix + key)

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

// Set sets value to memcache
func (p *PageCache) Set(key string, page *Page) error {
	bytes, err := json.Marshal(page)

	if err != nil {
		return err
	}

	expiration := time.Hour * 24 // 1 day
	_, err = p.memcached.Set(keyPrefix+key, string(bytes), 0, uint32(expiration.Seconds()), 0)

	if err != nil {
		return err
	}

	return nil
}
