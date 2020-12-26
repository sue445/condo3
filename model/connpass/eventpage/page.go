package eventpage

import (
	"errors"
	"github.com/getsentry/sentry-go"
	"github.com/sue445/condo3/model"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// Page represents connpass group page
type Page struct {
	PublishDatetime string `json:"publish_datetime"`
}

// FetchEventPageWithCache returns event page with memcache
func FetchEventPageWithCache(memcachedConfig *model.MemcachedConfig, url string) (*Page, error) {
	cache, quit := newPageCache(memcachedConfig)
	defer quit()

	cached, err := cache.get(url)

	if err != nil {
		sentry.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetLevel(sentry.LevelWarning)
			scope.SetExtras(map[string]interface{}{
				"url": url,
			})
			sentry.CaptureException(err)
		})
	}

	if cached != nil {
		return cached, nil
	}

	page, err := fetchEventPage(url)

	if err != nil {
		return nil, err
	}

	err = cache.set(url, page)

	if err != nil {
		sentry.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetLevel(sentry.LevelWarning)
			scope.SetExtras(map[string]interface{}{
				"url": url,
			})
			sentry.CaptureException(err)
		})
	}

	return page, nil
}

// fetchEventPage fetch connpass event page
func fetchEventPage(url string) (*Page, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return nil, errors.New(resp.Status)
	}

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	body := string(byteArray)

	publishDatetime, err := getPublishDate(body)
	if err != nil {
		return nil, err
	}

	page := &Page{
		PublishDatetime: publishDatetime,
	}

	return page, nil
}

func getPublishDate(body string) (string, error) {
	re := regexp.MustCompile("\"publish_datetime\": \"(\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2})\"")
	matched := re.FindStringSubmatch(body)

	if matched == nil {
		return "", errors.New("NotFound publish_datetime")
	}

	return strings.TrimSpace(matched[1]), nil
}
