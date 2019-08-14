package eventpage

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// Page represents connpass group page
type Page struct {
	PublishDatetime string `json:"publish_datetime"`
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
