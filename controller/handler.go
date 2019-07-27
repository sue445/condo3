package controller

import (
	"regexp"
	"strconv"
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
