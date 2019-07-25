package connpass

import (
	"github.com/jarcoal/httpmock"
	"github.com/sue445/condo3/testutil"
	"testing"
)

func Test_fetchSeriesID(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://bpstudy.connpass.com/",
		httpmock.NewStringResponder(200, testutil.ReadTestData("testdata/bpstudy.html")))
	httpmock.RegisterResponder("GET", "https://gocon.connpass.com/",
		httpmock.NewStringResponder(200, testutil.ReadTestData("testdata/gocon.html")))
	httpmock.RegisterResponder("GET", "https://without-series-id.connpass.com/",
		httpmock.NewStringResponder(200, "<html><body>dummy</body></html>"))
	httpmock.RegisterResponder("GET", "https://not-found.connpass.com/",
		httpmock.NewStringResponder(404, ""))

	type args struct {
		groupName string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "bpstudy",
			args: args{
				groupName: "bpstudy",
			},
			want: 1,
		},
		{
			name: "gocon",
			args: args{
				groupName: "gocon",
			},
			want: 312,
		},
		{
			name: "NotFound",
			args: args{
				groupName: "not-found",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "without seriesId",
			args: args{
				groupName: "without-series-id",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fetchSeriesID(tt.args.groupName)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchSeriesID() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("fetchSeriesID() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
