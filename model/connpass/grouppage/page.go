package grouppage

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sue445/condo3/logger"
	"github.com/sue445/condo3/model"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var (
	log = logger.NewLogger()
)

// Page represents connpass group page
type Page struct {
	SeriesID int    `json:"series_id"`
	URL      string `json:"url"`
	Title    string `json:"title"`
}

// FetchGroupPageWithCache returns group page with memcache
func FetchGroupPageWithCache(memcachedConfig *model.MemcachedConfig, groupName string) (*Page, error) {
	cache, quit := newPageCache(memcachedConfig)
	defer quit()

	cached, err := cache.get(groupName)

	if err != nil {
		log.Warnf("cache.get is failed: groupName=%s, err=%+v", groupName, err)
	}

	if cached != nil {
		return cached, nil
	}

	page, err := fetchGroupPage(groupName)

	if err != nil {
		return nil, err
	}

	err = cache.set(groupName, page)

	if err != nil {
		log.Warnf("cache.set is failed: groupName=%s, err=%+v", groupName, err)
	}

	return page, nil
}

// fetchGroupPage fetch connpass group page
func fetchGroupPage(groupName string) (*Page, error) {
	url := fmt.Sprintf("https://%s.connpass.com/", groupName)

	resp, err := http.Get(url)

	if err != nil {
		return nil, errors.WithStack(err)
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

	seriesID, err := getSeriesID(body)

	if err != nil {
		return nil, err
	}

	title, err := getTitle(body)

	if err != nil {
		return nil, err
	}

	page := &Page{
		SeriesID: seriesID,
		URL:      url,
		Title:    title,
	}

	return page, nil
}

func getSeriesID(body string) (int, error) {
	re := regexp.MustCompile("//connpass.com/series/(\\d+)/")
	matched := re.FindStringSubmatch(body)

	if matched == nil {
		return 0, errors.New("NotFound SeriesID")
	}

	seriesID, _ := strconv.Atoi(matched[1])
	return seriesID, nil
}

func getTitle(body string) (string, error) {
	re := regexp.MustCompile("(?i)<title>(.+)</title>")
	matched := re.FindStringSubmatch(body)

	if matched == nil {
		return "", errors.New("NotFound Title")
	}

	return strings.TrimSpace(matched[1]), nil
}
