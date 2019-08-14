package eventpage

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/sue445/condo3/testutil"
	"testing"
)

func Test_fetchEventPage(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://gocon.connpass.com/event/139024/",
		httpmock.NewStringResponder(200, testutil.ReadTestData("testdata/gocon_139024.html")))
	httpmock.RegisterResponder("GET", "https://without-publish-date.connpass.com/000000/",
		httpmock.NewStringResponder(200, "<html><body>dummy</body></html>"))
	httpmock.RegisterResponder("GET", "https://not-found.connpass.com/000000/",
		httpmock.NewStringResponder(404, ""))

	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    *Page
		wantErr bool
	}{
		{
			name: "https://gocon.connpass.com/event/139024/",
			args: args{
				url: "https://gocon.connpass.com/event/139024/",
			},
			want: &Page{
				PublishDatetime: "2019-07-10T12:01:10",
			},
		},
		{
			name: "NotFound",
			args: args{
				url: "https://not-found.connpass.com/000000/",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "without publish_date",
			args: args{
				url: "https://without-publish-date.connpass.com/000000/",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fetchEventPage(tt.args.url)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
