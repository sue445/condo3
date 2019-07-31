package connpass

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/sue445/condo3/model"
	"github.com/sue445/condo3/testutil"
	"google.golang.org/appengine/aetest"
	"testing"
	"time"
)

func tp(t time.Time) *time.Time {
	return &t
}

func TestGetGroup(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	assert.NoError(t, err)
	defer done()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://gocon.connpass.com/",
		httpmock.NewStringResponder(200, testutil.ReadTestData("testdata/gocon.html")))
	httpmock.RegisterResponder("GET", "http://connpass.com/api/v1/event/?count=100&order=3&series_id=312",
		httpmock.NewStringResponder(200, testutil.ReadTestData("testdata/gocon.json")))

	type args struct {
		groupName string
	}
	tests := []struct {
		name           string
		args           args
		wantEventFirst model.Event
		wantEventCount int
		wantURL        string
		wantTitle      string
		wantAddress    string
	}{
		{
			name: "successful",
			args: args{
				groupName: "gocon",
			},
			wantEventFirst: model.Event{
				Title:     "Go 1.13 Release Party in Tokyo",
				URL:       "https://gocon.connpass.com/event/139024/",
				Address:   "東京都港区六本木6-10-1 (六本木ヒルズ森タワー18F)",
				StartedAt: tp(time.Date(2019, 8, 23, 19, 30, 0, 0, time.Local)),
				EndedAt:   tp(time.Date(2019, 8, 23, 22, 0, 0, 0, time.Local)),
			},
			wantEventCount: 27,
			wantURL:        "https://gocon.connpass.com/",
			wantTitle:      "Go Conference - connpass",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetGroup(ctx, tt.args.groupName)

			assert.NoError(t, err)
			assert.Equal(t, tt.wantEventCount, len(got.Events))
			assert.Equal(t, tt.wantEventFirst, got.Events[0])
			assert.Equal(t, tt.wantURL, got.URL)
			assert.Equal(t, tt.wantTitle, got.Title)
		})
	}
}
