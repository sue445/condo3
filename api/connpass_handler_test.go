package api

import (
	"github.com/gorilla/mux"
	"github.com/jarcoal/httpmock"
	"github.com/sue445/condo3/testutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestConpassHandler(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://gocon.connpass.com/",
		httpmock.NewStringResponder(200, testutil.ReadTestData("../model/connpass/testdata/gocon.html")))
	httpmock.RegisterResponder("GET", "http://connpass.com/api/v1/event/?count=100&order=2&series_id=312",
		httpmock.NewStringResponder(200, testutil.ReadTestData("../model/connpass/testdata/gocon.json")))

	req, err := http.NewRequest("GET", "/api/connpass/gocon.ics", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/connpass/{group}.{format}", ConnpassHandler)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}

	expected := "BEGIN:VCALENDAR"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf(
			"unexpected body: got (%v) contains (%v)",
			rr.Body.String(),
			expected,
		)
	}
}
