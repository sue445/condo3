package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEvent_atomDescription(t *testing.T) {
	type fields struct {
		Title     string
		URL       string
		Address   string
		StartedAt *time.Time
		EndedAt   *time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "with StartedAt and EndedAt",
			fields: fields{
				Address:   "〒106-0032 東京都港区六本木4-1-4 黒崎ビル4階",
				StartedAt: tp(time.Date(2019, 6, 29, 5, 00, 0, 0, time.UTC)),
				EndedAt:   tp(time.Date(2019, 6, 29, 10, 30, 0, 0, time.UTC)),
			},
			want: "開催日時：2019/06/29 14:00〜19:30\n開催場所：〒106-0032 東京都港区六本木4-1-4 黒崎ビル4階",
		},
		{
			name: "only StartedAt",
			fields: fields{
				Address:   "〒106-0032 東京都港区六本木4-1-4 黒崎ビル4階",
				StartedAt: tp(time.Date(2019, 6, 29, 5, 00, 0, 0, time.UTC)),
				EndedAt:   nil,
			},
			want: "開催日時：2019/06/29 14:00〜\n開催場所：〒106-0032 東京都港区六本木4-1-4 黒崎ビル4階",
		},
		{
			name: "only EndedAt",
			fields: fields{
				Address:   "〒106-0032 東京都港区六本木4-1-4 黒崎ビル4階",
				StartedAt: nil,
				EndedAt:   tp(time.Date(2019, 6, 29, 10, 30, 0, 0, time.UTC)),
			},
			want: "開催日時：〜2019/06/29 19:30\n開催場所：〒106-0032 東京都港区六本木4-1-4 黒崎ビル4階",
		},
		{
			name: "without StartedAt and EndedAt",
			fields: fields{
				Address:   "〒106-0032 東京都港区六本木4-1-4 黒崎ビル4階",
				StartedAt: nil,
				EndedAt:   nil,
			},
			want: "開催場所：〒106-0032 東京都港区六本木4-1-4 黒崎ビル4階",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Event{
				Title:     tt.fields.Title,
				URL:       tt.fields.URL,
				Address:   tt.fields.Address,
				StartedAt: tt.fields.StartedAt,
				EndedAt:   tt.fields.EndedAt,
			}

			got := e.atomDescription()
			assert.Equal(t, tt.want, got)
		})
	}
}
