package connpass

import (
	"github.com/jarcoal/httpmock"
	"github.com/sue445/condo3/model"
	"github.com/sue445/condo3/testutil"
	"reflect"
	"testing"
	"time"
)

func tp(t time.Time) *time.Time {
	return &t
}

func TestGetGroup(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://gocon.connpass.com/",
		httpmock.NewStringResponder(200, testutil.ReadTestData("testdata/gocon.html")))
	httpmock.RegisterResponder("GET", "http://connpass.com/api/v1/event/?count=100&order=2&series_id=312",
		httpmock.NewStringResponder(200, testutil.ReadTestData("testdata/gocon.json")))

	type args struct {
		groupName string
	}
	tests := []struct {
		name           string
		args           args
		wantEventFirst model.Event
		wantEventCount int
		wantErr        bool
		wantURL        string
		wantTitle      string
	}{
		{
			name: "successful",
			args: args{
				groupName: "gocon",
			},
			wantEventFirst: model.Event{
				Title:     "Go 1.13 Release Party in Tokyo",
				URL:       "https://gocon.connpass.com/event/139024/",
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
			got, err := GetGroup(tt.args.groupName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroup() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if len(got.Events) != tt.wantEventCount {
				t.Errorf("GetGroup().Events count = %+v, want %+v", len(got.Events), tt.wantEventCount)
			}
			if !reflect.DeepEqual(got.Events[0], tt.wantEventFirst) {
				t.Errorf("GetGroup().Events[0] = %+v, want %+v", got.Events[0], tt.wantEventFirst)
			}
			if got.URL != tt.wantURL {
				t.Errorf("GetGroup().URL = %+v, want %+v", got.URL, tt.wantURL)
			}
			if got.Title != tt.wantTitle {
				t.Errorf("GetGroup().Title = %+v, want %+v", got.Title, tt.wantTitle)
			}
		})
	}
}
