package api

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sue445/condo3/logger"
	"github.com/sue445/condo3/model"
	"net/http"
	"regexp"
	"strconv"
)

var (
	log      = logger.NewLogger()
	errorLog = logger.NewErrorLogger()
)

// Handler manages API handler
type Handler struct {
	DoorkeeperAccessToken string
	MemcachedConfig       *model.MemcachedConfig
}

func errorStatusCode(err error) int {
	re := regexp.MustCompile("^(\\d{3})")
	matched := re.FindStringSubmatch(err.Error())

	if matched == nil {
		return 500
	}

	statusCode, _ := strconv.Atoi(matched[1])
	return statusCode
}

func renderError(w http.ResponseWriter, err error) {
	errorLog.Errorf("%+v", errors.WithStack(err))
	w.WriteHeader(errorStatusCode(err))
	fmt.Fprint(w, err)
}

func renderGroup(w http.ResponseWriter, group *model.Group, format string) {
	switch format {
	case "ics":
		setContentType(w, "text/calendar; charset=utf-8")
		writeAPIResponse(w, group.ToIcal())
	case "atom":
		atom, err := group.ToAtom()

		if err != nil {
			renderError(w, err)
			return
		}

		setContentType(w, "application/atom+xml; charset=utf-8")
		writeAPIResponse(w, atom)
	default:
		message := fmt.Sprintf("Unknown format: %s", format)
		log.Warn(message)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, message)
	}
}

func setContentType(w http.ResponseWriter, contentType string) {
	w.Header().Set("Content-Type", contentType)
}

func writeAPIResponse(w http.ResponseWriter, body string) {
	enableFrontendCache(w, len(body))
	fmt.Fprint(w, body)
}

func enableFrontendCache(w http.ResponseWriter, contentLength int) {
	expirationSeconds := 60 * 10 // 10 min

	// c.f. https://cloud.google.com/cdn/docs/caching?hl=ja
	w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", expirationSeconds))
	w.Header().Set("Content-Length", strconv.Itoa(contentLength))
}
