package eventpage

import (
	"github.com/stretchr/testify/assert"
	"github.com/sue445/condo3/model"
	"os"
	"testing"
)

func TestPageCache_SetAndGet(t *testing.T) {
	page := &Page{
		PublishDatetime: "2019-07-10T12:01:10",
	}

	cache, quit := newPageCache(&model.MemcachedConfig{Server: os.Getenv("MEMCACHED_SERVER")})
	defer quit()

	cache.memcached.Flush(0)

	err := cache.set("https://gocon.connpass.com/event/139024/", page)
	assert.NoError(t, err)

	actual, err := cache.get("https://gocon.connpass.com/event/139024/")
	assert.NoError(t, err)
	assert.Equal(t, page, actual)

	actual2, err := cache.get("https://not-found.connpass.com/000000/")
	assert.NoError(t, err)
	assert.Nil(t, actual2)
}
