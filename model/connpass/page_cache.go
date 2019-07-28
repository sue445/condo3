package connpass

import (
	"context"
	"encoding/json"
	"google.golang.org/appengine/memcache"
	"time"
)

const (
	keyPrefix  = "PageCache-v1-"
	expiration = 24 * time.Hour // 1 day
)

// PageCache represents page cache
type PageCache struct {
	ctx context.Context
}

// NewPageCache returns new PageCache instance
func NewPageCache(ctx context.Context) *PageCache {
	return &PageCache{ctx: ctx}
}

// Get returns value from memcache
func (p *PageCache) Get(key string) (*Page, error) {
	item, err := memcache.Get(p.ctx, keyPrefix+key)

	if err != nil {
		if err == memcache.ErrCacheMiss {
			return nil, nil
		}
		return nil, err
	}

	data := &Page{}
	err = json.Unmarshal(item.Value, data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// Set sets value to memcache
func (p *PageCache) Set(key string, data *Page) error {
	byte, err := json.Marshal(data)

	if err != nil {
		return err
	}

	item := &memcache.Item{
		Key:        keyPrefix + key,
		Value:      byte,
		Expiration: expiration,
	}

	err = memcache.Set(p.ctx, item)

	if err != nil {
		return err
	}

	return nil
}
