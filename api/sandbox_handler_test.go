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

func TestSandboxHandler(t *testing.T) {
	httpmock.Activate(t)

	httpmock.RegisterResponder("GET", "https://sue445.github.io/condo3-sandbox/data/test1.json",
		httpmock.NewStringResponder(200, testutil.ReadTestData("../model/sandbox/testdata/test1.json")))

	req, err := http.NewRequest("GET", "/api/sandbox/test1.ics", nil)
	assert.NoError(t, err)

	handler := Handler{}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/sandbox/{group}.{format}", handler.SandboxHandler)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "BEGIN:VCALENDAR")
}
