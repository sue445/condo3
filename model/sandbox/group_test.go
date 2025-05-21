package sandbox

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/sue445/condo3/model"
	"github.com/sue445/condo3/testutil"
	"testing"
	"time"
)

func tp(t time.Time) *time.Time {
	return &t
}

func TestGetGroup(t *testing.T) {
	httpmock.Activate(t)

	httpmock.RegisterResponder("GET", "https://sue445.github.io/condo3-sandbox/data/test1.json",
		httpmock.NewStringResponder(200, testutil.ReadTestData("testdata/test1.json")))

	type args struct {
		groupName string
	}
	tests := []struct {
		name string
		args args
		want *model.Group
	}{
		{
			name: "Successful",
			args: args{
				groupName: "test1",
			},
			want: &model.Group{
				Title:     "test1",
				URL:       "https://sue445.github.io/condo3-sandbox/data/test1.json",
				UpdatedAt: tp(time.Date(2019, 8, 14, 16, 0, 0, 0, model.JST)),
				Events: []model.Event{
					{
						Title:       "test1 event1",
						URL:         "https://sue445.github.io/condo3-sandbox/data/test1.json",
						Address:     "東京都",
						UpdatedAt:   tp(time.Date(2019, 8, 14, 16, 0, 0, 0, model.JST)),
						PublishedAt: tp(time.Date(2019, 8, 14, 16, 0, 0, 0, model.JST)),
						StartedAt:   nil,
						EndedAt:     nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetGroup(tt.args.groupName)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
