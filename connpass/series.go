package connpass

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

// fetchSeriesId returns seriesID from groupName
func fetchSeriesID(groupName string) (int, error) {
	url := fmt.Sprintf("https://%s.connpass.com/", groupName)

	resp, err := http.Get(url)

	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return 0, errors.New(resp.Status)
	}

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	body := string(byteArray)

	re := regexp.MustCompile("//connpass.com/series/(\\d+)/")
	matched := re.FindStringSubmatch(body)

	if matched == nil {
		return 0, errors.New("NotFound seriesID")
	}

	seriesID, _ := strconv.Atoi(matched[1])

	return seriesID, nil
}
