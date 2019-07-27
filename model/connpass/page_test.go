package connpass

import (
	"github.com/jarcoal/httpmock"
	"github.com/sue445/condo3/testutil"
	"reflect"
	"testing"
)

func Test_FetchGroupPage(t *testing.T) {
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
		want    *Page
		wantErr bool
	}{
		{
			name: "bpstudy",
			args: args{
				groupName: "bpstudy",
			},
			want: &Page{
				SeriesID: 1,
				URL:      "https://bpstudy.connpass.com/",
				Title:    "BPStudy - connpass",
			},
		},
		{
			name: "gocon",
			args: args{
				groupName: "gocon",
			},
			want: &Page{
				SeriesID: 312,
				URL:      "https://gocon.connpass.com/",
				Title:    "Go Conference - connpass",
			},
		},
		{
			name: "NotFound",
			args: args{
				groupName: "not-found",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "without seriesId",
			args: args{
				groupName: "without-series-id",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FetchGroupPage(tt.args.groupName)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchGroupPage() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchGroupPage() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
