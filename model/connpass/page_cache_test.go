package connpass

import (
	"github.com/stretchr/testify/assert"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/memcache"
	"testing"
)

func TestPageCache_SetAndGet(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	assert.NoError(t, err)
	defer done()

	memcache.Flush(ctx)

	page := &Page{
		SeriesID: 312,
		URL:      "https://gocon.connpass.com/",
		Title:    "Go Conference - connpass",
	}

	cache := NewPageCache(ctx)

	err = cache.Set("gocon", page)
	assert.NoError(t, err)

	actual, err := cache.Get("gocon")
	assert.NoError(t, err)
	assert.Equal(t, page, actual)

	actual2, err := cache.Get("not-found")
	assert.NoError(t, err)
	assert.Nil(t, actual2)
}
