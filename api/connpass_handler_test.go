package api

import (
	"github.com/gorilla/mux"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/sue445/condo3/model"
	"github.com/sue445/condo3/testutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestConpassHandler(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://gocon.connpass.com/",
		httpmock.NewStringResponder(200, testutil.ReadTestData("../model/connpass/grouppage/testdata/gocon.html")))
	httpmock.RegisterResponder("GET", `=~^https://connpass\.com/api/v1/event/`,
		httpmock.NewStringResponder(200, testutil.ReadTestData("../model/connpass/testdata/gocon.json")))

	// FIXME: race condition error when same responder is called in goroutine
	//httpmock.RegisterResponder("GET", `=~^https://gocon\.connpass\.com/event/`,
	//	httpmock.NewStringResponder(200, testutil.ReadTestData("../model/connpass/eventpage/testdata/gocon_139024.html")))
	httpmock.RegisterResponder("GET", "https://gocon.connpass.com/event/139024/",
		httpmock.NewStringResponder(200, testutil.ReadTestData("../model/connpass/eventpage/testdata/gocon_139024.html")))
	httpmock.RegisterResponder("GET", "https://gocon.connpass.com/event/124530/",
		httpmock.NewStringResponder(200, testutil.ReadTestData("../model/connpass/eventpage/testdata/gocon_124530.html")))

	req, err := http.NewRequest("GET", "/api/connpass/gocon.ics", nil)
	assert.NoError(t, err)

	memcachedConfig := model.MemcachedConfig{Server: os.Getenv("MEMCACHED_SERVER")}
	handler := Handler{MemcachedConfig: &memcachedConfig}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/connpass/{group}.{format}", handler.ConnpassHandler)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "BEGIN:VCALENDAR")
}
