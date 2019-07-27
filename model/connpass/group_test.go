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

func TestGetGroupEvents(t *testing.T) {
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetGroupEvents(tt.args.groupName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroupEvents() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), tt.wantEventCount) {
				t.Errorf("GetGroupEvents() = %+v, want %+v", len(got), tt.wantEventCount)
			}
			if !reflect.DeepEqual(got[0], tt.wantEventFirst) {
				t.Errorf("GetGroupEvents() = %+v, want %+v", got[0], tt.wantEventFirst)
			}
		})
	}
}
