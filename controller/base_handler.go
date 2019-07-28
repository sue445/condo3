package controller

import (
	"fmt"
	"github.com/sue445/condo3/model"
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

func renderGroup(w http.ResponseWriter, group *model.Group, format string) {
	switch format {
	case "ics":
		w.WriteHeader(http.StatusOK)
		setContentType(w, contentTypeIcs)
		fmt.Fprint(w, group.ToIcal())
	case "atom":
		atom, err := group.ToAtom()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err)
			return
		}

		setContentType(w, contentTypeAtom)
		fmt.Fprint(w, atom)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func setContentType(w http.ResponseWriter, contentType string) {
	w.Header().Set("Content-Type", contentType)
}
