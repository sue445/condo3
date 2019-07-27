package controller

import (
	"net/http"
	"regexp"
	"strconv"
)

const (
	contentTypeAtom = "application/atom+xml; charset=utf-8"
	contentTypeIcs  = "text/calendar; charset=utf-8"
)

func errorStatusCode(err error) int {
	re := regexp.MustCompile("^(\\d{3})")
	matched := re.FindStringSubmatch(err.Error())

	if matched == nil {
		return 500
	}

	statusCode, _ := strconv.Atoi(matched[1])
	return statusCode
}

func setContentType(w http.ResponseWriter, contentType string) {
	w.Header().Set("Content-Type", contentType)
}
