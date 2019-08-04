package api

import (
	"github.com/gorilla/mux"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/sue445/condo3/testutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConpassHandler(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://gocon.connpass.com/",
		httpmock.NewStringResponder(200, testutil.ReadTestData("../model/connpass/testdata/gocon.html")))
	httpmock.RegisterResponder("GET", `=~^https://connpass\.com/api/v1/event/`,
		httpmock.NewStringResponder(200, testutil.ReadTestData("../model/connpass/testdata/gocon.json")))

	req, err := http.NewRequest("GET", "/api/connpass/gocon.ics", nil)
	assert.NoError(t, err)

	a := Handler{}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/connpass/{group}.{format}", a.ConnpassHandler)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "BEGIN:VCALENDAR")
}
