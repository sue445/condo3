package api

import (
	"github.com/gorilla/mux"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/sue445/condo3/testutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestDoorkeeperHandler(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.doorkeeper.jp/groups/trbmeetup/events?since=2019-05-03&sort=published_at&until=2019-10-30",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, testutil.ReadTestData(filepath.Join("..", "model", "doorkeeper", "testdata", "events.json")))
			resp.Header.Set("X-Ratelimit", `{"name":"authenticated API","period":300,"limit":300,"remaining":299,"until":"2019-07-29T15:15:00Z"}`)
			return resp, nil
		},
	)
	httpmock.RegisterResponder("GET", "https://api.doorkeeper.jp/groups/trbmeetup",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, testutil.ReadTestData(filepath.Join("..", "model", "doorkeeper", "testdata", "group-ja.json")))
			resp.Header.Set("X-Ratelimit", `{"name":"authenticated API","period":300,"limit":300,"remaining":299,"until":"2019-07-29T15:15:00Z"}`)
			return resp, nil
		},
	)

	req, err := http.NewRequest("GET", "/api/doorkeeper/trbmeetup.ics", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/doorkeeper/{group}.{format}", DoorkeeperHandler)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "BEGIN:VCALENDAR")
}
