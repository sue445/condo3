package doorkeeper

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/sue445/condo3/model"
	"github.com/sue445/condo3/testutil"
	"net/http"
	"path/filepath"
	"testing"
	"time"
)

func tp(t time.Time) *time.Time {
	return &t
}

func TestGetGroup(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.doorkeeper.jp/groups/trbmeetup/events?since=2014-06-01&sort=published_at&until=2016-06-01",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, testutil.ReadTestData(filepath.Join("testdata", "events.json")))
			resp.Header.Set("X-Ratelimit", `{"name":"authenticated API","period":300,"limit":300,"remaining":299,"until":"2019-07-29T15:15:00Z"}`)
			return resp, nil
		},
	)
	httpmock.RegisterResponder("GET", "https://api.doorkeeper.jp/groups/trbmeetup",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, testutil.ReadTestData(filepath.Join("testdata", "group-ja.json")))
			resp.Header.Set("X-Ratelimit", `{"name":"authenticated API","period":300,"limit":300,"remaining":299,"until":"2019-07-29T15:15:00Z"}`)
			return resp, nil
		},
	)

	type args struct {
		accessToken string
		groupName   string
		currentTime time.Time
	}
	tests := []struct {
		name           string
		args           args
		wantEventFirst model.Event
		wantEventCount int
		wantURL        string
		wantTitle      string
	}{
		{
			name: "successful",
			args: args{
				accessToken: "xxxxxxxxx",
				groupName:   "trbmeetup",
				currentTime: time.Date(2015, 6, 1, 0, 0, 0, 0, time.UTC),
			},
			wantEventFirst: model.Event{
				Title:     "900K records per second with Ruby, Java, and JRuby",
				URL:       "https://trbmeetup.doorkeeper.jp/events/28319",
				Address:   "東京都渋谷区神泉町8-16 渋谷ファーストプレイス8F",
				StartedAt: tp(time.Date(2015, 8, 13, 10, 0, 0, 0, time.UTC)),
				EndedAt:   tp(time.Date(2015, 8, 13, 13, 0, 0, 0, time.UTC)),
			},
			wantEventCount: 1,
			wantURL:        "https://trbmeetup.doorkeeper.jp/",
			wantTitle:      "Tokyo Rubyist Meetup | Doorkeeper",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetGroup(tt.args.accessToken, tt.args.groupName, tt.args.currentTime)

			assert.NoError(t, err)
			assert.NotNil(t, got)

			if got != nil {
				assert.Equal(t, tt.wantEventCount, len(got.Events))
				assert.Equal(t, tt.wantEventFirst, got.Events[0])
				assert.Equal(t, tt.wantURL, got.URL)
				assert.Equal(t, tt.wantTitle, got.Title)
			}
		})
	}
}
