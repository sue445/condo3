package grouppage

import (
	"github.com/stretchr/testify/assert"
	"github.com/sue445/condo3/model"
	"os"
	"testing"
)

func TestPageCache_SetAndGet(t *testing.T) {
	page := &Page{
		SeriesID: 312,
		URL:      "https://gocon.connpass.com/",
		Title:    "Go Conference - connpass",
	}

	cache, quit := newPageCache(&model.MemcachedConfig{Server: os.Getenv("MEMCACHED_SERVER")})
	defer quit()

	cache.memcached.Flush(0)

	err := cache.set("gocon", page)
	assert.NoError(t, err)

	actual, err := cache.get("gocon")
	assert.NoError(t, err)
	assert.Equal(t, page, actual)

	actual2, err := cache.get("not-found")
	assert.NoError(t, err)
	assert.Nil(t, actual2)
}
