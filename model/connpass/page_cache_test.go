package connpass

import (
	"github.com/stretchr/testify/assert"
	"github.com/sue445/condo3/model"
	"testing"
)

func TestPageCache_SetAndGet(t *testing.T) {
	page := &Page{
		SeriesID: 312,
		URL:      "https://gocon.connpass.com/",
		Title:    "Go Conference - connpass",
	}

	cache, quit := NewPageCache(&model.MemcachedConfig{Server: "127.0.0.1:11211"})
	defer quit()

	cache.memcached.Flush(0)

	err := cache.Set("gocon", page)
	assert.NoError(t, err)

	actual, err := cache.Get("gocon")
	assert.NoError(t, err)
	assert.Equal(t, page, actual)

	actual2, err := cache.Get("not-found")
	assert.NoError(t, err)
	assert.Nil(t, actual2)
}
