package connpass

import (
	"github.com/hkurokawa/go-connpass"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/sue445/condo3/model"
	"github.com/sue445/condo3/testutil"
	"os"
	"testing"
	"time"
)

func tp(t time.Time) *time.Time {
	return &t
}

func TestGetGroup(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	currentTime := time.Date(2019, 8, 2, 0, 0, 0, 0, time.UTC)

	httpmock.RegisterResponder("GET", "https://gocon.connpass.com/",
		httpmock.NewStringResponder(200, testutil.ReadTestData("grouppage/testdata/gocon.html")))
	httpmock.RegisterResponder("GET", "https://connpass.com/api/v1/event/?count=100&order=3&series_id=312&ym=201902%2C201903%2C201904%2C201905%2C201906%2C201907%2C201908%2C201909%2C201910%2C201911%2C201912%2C202001%2C202002",
		httpmock.NewStringResponder(200, testutil.ReadTestData("testdata/gocon.json")))

	// FIXME: race condition error when same responder is called in goroutine
	//httpmock.RegisterResponder("GET", `=~^https://gocon\.connpass\.com/event/`,
	//	httpmock.NewStringResponder(200, testutil.ReadTestData("eventpage/testdata/gocon_139024.html")))
	httpmock.RegisterResponder("GET", "https://gocon.connpass.com/event/139024/",
		httpmock.NewStringResponder(200, testutil.ReadTestData("eventpage/testdata/gocon_139024.html")))
	httpmock.RegisterResponder("GET", "https://gocon.connpass.com/event/124530/",
		httpmock.NewStringResponder(200, testutil.ReadTestData("eventpage/testdata/gocon_124530.html")))

	memcachedConfig := model.MemcachedConfig{Server: os.Getenv("MEMCACHED_SERVER")}

	type args struct {
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
		wantUpdatedAt  *time.Time
	}{
		{
			name: "successful",
			args: args{
				groupName:   "gocon",
				currentTime: currentTime,
			},
			wantEventFirst: model.Event{
				Title:       "Go 1.13 Release Party in Tokyo",
				URL:         "https://gocon.connpass.com/event/139024/",
				Address:     "東京都港区六本木6-10-1 (六本木ヒルズ森タワー18F)",
				UpdatedAt:   tp(time.Date(2019, 7, 25, 22, 24, 0, 0, model.JST)),
				PublishedAt: tp(time.Date(2019, 7, 10, 12, 01, 10, 0, model.JST)),
				StartedAt:   tp(time.Date(2019, 8, 23, 19, 30, 0, 0, model.JST)),
				EndedAt:     tp(time.Date(2019, 8, 23, 22, 0, 0, 0, model.JST)),
			},
			wantEventCount: 2,
			wantURL:        "https://gocon.connpass.com/",
			wantTitle:      "Go Conference - connpass",
			wantUpdatedAt:  tp(time.Date(2019, 7, 25, 22, 24, 0, 0, model.JST)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetGroup(&memcachedConfig, tt.args.groupName, tt.args.currentTime)

			assert.NoError(t, err)
			assert.NotNil(t, got)

			if got != nil {
				assert.Equal(t, tt.wantEventCount, len(got.Events))
				assert.Equal(t, tt.wantEventFirst, got.Events[0])
				assert.Equal(t, tt.wantURL, got.URL)
				assert.Equal(t, tt.wantTitle, got.Title)
				assert.Equal(t, tt.wantUpdatedAt, got.UpdatedAt)
			}
		})
	}
}

func Test_getTerms(t *testing.T) {
	type args struct {
		currentTime time.Time
		beforeMonth int
		afterMonth  int
	}
	tests := []struct {
		name string
		args args
		want []connpass.Time
	}{
		{
			name: "successful",
			args: args{
				currentTime: time.Date(2019, 2, 2, 0, 0, 0, 0, time.UTC),
				beforeMonth: 2,
				afterMonth:  3,
			},
			want: []connpass.Time{
				{Year: 2018, Month: 12, Date: 0},
				{Year: 2019, Month: 1, Date: 0},
				{Year: 2019, Month: 2, Date: 0},
				{Year: 2019, Month: 3, Date: 0},
				{Year: 2019, Month: 4, Date: 0},
				{Year: 2019, Month: 5, Date: 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getTerms(tt.args.currentTime, tt.args.beforeMonth, tt.args.afterMonth)

			assert.Equal(t, tt.want, got)
		})
	}
}
