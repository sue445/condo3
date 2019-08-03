package connpass

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

// Page represents connpass group page
type Page struct {
	SeriesID int
	URL      string
	Title    string
}

// FetchGroupPageWithCache returns group page with memcache
func FetchGroupPageWithCache(ctx context.Context, groupName string) (*Page, error) {
	cache := NewPageCache(ctx)

	cached, err := cache.Get(groupName)

	if err != nil {
		log.Printf("[WARN] cache.Get is failed: groupName=%s, err=%+v\n", groupName, err)
	}

	if cached != nil {
		return cached, nil
	}

	page, err := FetchGroupPage(groupName)

	if err != nil {
		return nil, err
	}

	err = cache.Set(groupName, page)

	if err != nil {
		log.Printf("[WARN] cache.Set is failed: groupName=%s, err=%+v\n", groupName, err)
	}

	return page, nil
}

// FetchGroupPage fetch connpass group page
func FetchGroupPage(groupName string) (*Page, error) {
	url := fmt.Sprintf("https://%s.connpass.com/", groupName)

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
