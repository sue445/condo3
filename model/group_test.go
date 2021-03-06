package model

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func tp(t time.Time) *time.Time {
	return &t
}

func TestGroup_ToIcal(t *testing.T) {
	goconIcal := `BEGIN:VCALENDAR
VERSION:2.0
CALSCALE:GREGORIAN
METHOD:PUBLISH
PRODID:-//sue445//condo3.appspot.com//JA
X-WR-CALDESC:Go Conference - connpass\nhttps://gocon.connpass.com/
X-WR-CALNAME:Go Conference - connpass
X-WR-TIMEZONE:UTC
BEGIN:VEVENT
DESCRIPTION:https://gocon.connpass.com/event/139024/
DTEND:20190823T130000Z
DTSTART:20190823T103000Z
LOCATION:東京都港区六本木6-10-1 (六本木ヒルズ森タワー18F)
SUMMARY:Go 1.13 Release Party in Tokyo
UID:https://gocon.connpass.com/event/139024/
URL:https://gocon.connpass.com/event/139024/
END:VEVENT
END:VCALENDAR
`

	tokyurubykaigiIcal := `BEGIN:VCALENDAR
VERSION:2.0
CALSCALE:GREGORIAN
METHOD:PUBLISH
PRODID:-//sue445//condo3.appspot.com//JA
X-WR-CALDESC:TokyuRubyKaigi | Doorkeeper\nhttps://tokyu-rubykaigi.doorkeepe
 r.jp/
X-WR-CALNAME:TokyuRubyKaigi | Doorkeeper
X-WR-TIMEZONE:UTC
BEGIN:VEVENT
DESCRIPTION:https://tokyu-rubykaigi.doorkeeper.jp/events/91543
DTEND:20190629T103000Z
DTSTART:20190629T050000Z
LOCATION:〒106-0032 東京都港区六本木4-1-4 黒崎ビル4階
SUMMARY:TokyuRubyKaigi13 一般参加者募集(LT発表者は登録不要です)
UID:https://tokyu-rubykaigi.doorkeeper.jp/events/91543
URL:https://tokyu-rubykaigi.doorkeeper.jp/events/91543
END:VEVENT
END:VCALENDAR
`

	type fields struct {
		Title  string
		URL    string
		Events []Event
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "generate gocon ical",
			fields: fields{
				Title: "Go Conference - connpass",
				URL:   "https://gocon.connpass.com/",
				Events: []Event{
					{
						Title:     "Go 1.13 Release Party in Tokyo",
						URL:       "https://gocon.connpass.com/event/139024/",
						Address:   "東京都港区六本木6-10-1 (六本木ヒルズ森タワー18F)",
						StartedAt: tp(time.Date(2019, 8, 23, 19, 30, 0, 0, JST)),
						EndedAt:   tp(time.Date(2019, 8, 23, 22, 0, 0, 0, JST)),
					},
				},
			},
			want: strings.ReplaceAll(goconIcal, "\n", "\r\n"),
		},
		{
			name: "mojibake ics",
			fields: fields{
				Title: "TokyuRubyKaigi | Doorkeeper",
				URL:   "https://tokyu-rubykaigi.doorkeeper.jp/",
				Events: []Event{
					{
						Title:     "TokyuRubyKaigi13 一般参加者募集(LT発表者は登録不要です)",
						URL:       "https://tokyu-rubykaigi.doorkeeper.jp/events/91543",
						Address:   "〒106-0032 東京都港区六本木4-1-4 黒崎ビル4階",
						StartedAt: tp(time.Date(2019, 6, 29, 5, 00, 0, 0, time.UTC)),
						EndedAt:   tp(time.Date(2019, 6, 29, 10, 30, 0, 0, time.UTC)),
					},
				},
			},
			want: strings.ReplaceAll(tokyurubykaigiIcal, "\n", "\r\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Group{
				Title:  tt.fields.Title,
				URL:    tt.fields.URL,
				Events: tt.fields.Events,
			}

			got := g.ToIcal()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGroup_ToAtom(t *testing.T) {
	goconAtom := `<?xml version="1.0" encoding="UTF-8"?>
<feed xmlns="http://www.w3.org/2005/Atom">
  <title>Go Conference - connpass</title>
  <id>https://gocon.connpass.com/</id>
  <link href="https://gocon.connpass.com/"></link>
  <updated>2019-07-25T22:24:00+09:00</updated>
  <entry>
    <title>Go 1.13 Release Party in Tokyo</title>
    <id>https://gocon.connpass.com/event/139024/</id>
    <link rel="alternate" href="https://gocon.connpass.com/event/139024/"></link>
    <published>2019-07-10T12:01:10+09:00</published>
    <updated>2019-07-25T22:24:00+09:00</updated>
    <summary type="html">開催日時：2019/08/23 19:30〜22:00&#xA;開催場所：東京都港区六本木6-10-1 (六本木ヒルズ森タワー18F)</summary>
  </entry>
</feed>`

	tokyurubykaigiAtom := `<?xml version="1.0" encoding="UTF-8"?>
<feed xmlns="http://www.w3.org/2005/Atom">
  <title>Tokyo Rubyist Meetup | Doorkeeper</title>
  <id>https://trbmeetup.doorkeeper.jp/</id>
  <link href="https://trbmeetup.doorkeeper.jp/"></link>
  <updated>2018-05-11T09:07:44+09:00</updated>
  <entry>
    <title>900K records per second with Ruby, Java, and JRuby</title>
    <id>https://trbmeetup.doorkeeper.jp/events/28319</id>
    <link rel="alternate" href="https://trbmeetup.doorkeeper.jp/events/28319"></link>
    <published>2015-07-14T08:48:29+09:00</published>
    <updated>2018-05-11T09:07:44+09:00</updated>
    <summary type="html">開催日時：2015/08/13 19:00〜22:00&#xA;開催場所：東京都渋谷区神泉町8-16 渋谷ファーストプレイス8F</summary>
  </entry>
</feed>`

	type fields struct {
		Title     string
		URL       string
		UpdatedAt *time.Time
		Events    []Event
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "generate gocon atom",
			fields: fields{
				Title:     "Go Conference - connpass",
				URL:       "https://gocon.connpass.com/",
				UpdatedAt: tp(time.Date(2019, 7, 25, 22, 24, 0, 0, JST)),
				Events: []Event{
					{
						Title:       "Go 1.13 Release Party in Tokyo",
						URL:         "https://gocon.connpass.com/event/139024/",
						Address:     "東京都港区六本木6-10-1 (六本木ヒルズ森タワー18F)",
						UpdatedAt:   tp(time.Date(2019, 7, 25, 22, 24, 0, 0, JST)),
						PublishedAt: tp(time.Date(2019, 7, 10, 12, 01, 10, 0, JST)),
						StartedAt:   tp(time.Date(2019, 8, 23, 19, 30, 0, 0, JST)),
						EndedAt:     tp(time.Date(2019, 8, 23, 22, 0, 0, 0, JST)),
					},
				},
			},
			want: goconAtom,
		},
		{
			name: "doorkeeper atom",
			fields: fields{
				Title:     "Tokyo Rubyist Meetup | Doorkeeper",
				URL:       "https://trbmeetup.doorkeeper.jp/",
				UpdatedAt: tp(time.Date(2018, 5, 11, 0, 7, 44, 270000000, time.UTC)),
				Events: []Event{
					{
						Title:       "900K records per second with Ruby, Java, and JRuby",
						URL:         "https://trbmeetup.doorkeeper.jp/events/28319",
						Address:     "東京都渋谷区神泉町8-16 渋谷ファーストプレイス8F",
						PublishedAt: tp(time.Date(2015, 7, 13, 23, 48, 29, 463000000, time.UTC)),
						UpdatedAt:   tp(time.Date(2018, 5, 11, 0, 7, 44, 270000000, time.UTC)),
						StartedAt:   tp(time.Date(2015, 8, 13, 10, 0, 0, 0, time.UTC)),
						EndedAt:     tp(time.Date(2015, 8, 13, 13, 0, 0, 0, time.UTC)),
					},
				},
			},
			want: tokyurubykaigiAtom,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Group{
				Title:     tt.fields.Title,
				URL:       tt.fields.URL,
				Events:    tt.fields.Events,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			got, err := g.ToAtom()

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGroup_ToJSON(t *testing.T) {
	goconJSON := `{"title":"Go Conference - connpass","url":"https://gocon.connpass.com/","updated_at":"2019-07-25T22:24:00+09:00","events":[{"title":"Go 1.13 Release Party in Tokyo","url":"https://gocon.connpass.com/event/139024/","address":"東京都港区六本木6-10-1 (六本木ヒルズ森タワー18F)","updated_at":"2019-07-25T22:24:00+09:00","published_at":"2019-07-10T12:01:10+09:00","started_at":"2019-08-23T19:30:00+09:00","ended_at":"2019-08-23T22:00:00+09:00"}]}`

	tokyorubyistJSON := `{"title":"Tokyo Rubyist Meetup | Doorkeeper","url":"https://trbmeetup.doorkeeper.jp/","updated_at":"2018-05-11T00:07:44.27Z","events":[{"title":"900K records per second with Ruby, Java, and JRuby","url":"https://trbmeetup.doorkeeper.jp/events/28319","address":"東京都渋谷区神泉町8-16 渋谷ファーストプレイス8F","updated_at":"2018-05-11T00:07:44.27Z","published_at":"2015-07-13T23:48:29.463Z","started_at":"2015-08-13T10:00:00Z","ended_at":"2015-08-13T13:00:00Z"}]}`

	type fields struct {
		Title     string
		URL       string
		UpdatedAt *time.Time
		Events    []Event
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "generate gocon json",
			fields: fields{
				Title:     "Go Conference - connpass",
				URL:       "https://gocon.connpass.com/",
				UpdatedAt: tp(time.Date(2019, 7, 25, 22, 24, 0, 0, JST)),
				Events: []Event{
					{
						Title:       "Go 1.13 Release Party in Tokyo",
						URL:         "https://gocon.connpass.com/event/139024/",
						Address:     "東京都港区六本木6-10-1 (六本木ヒルズ森タワー18F)",
						UpdatedAt:   tp(time.Date(2019, 7, 25, 22, 24, 0, 0, JST)),
						PublishedAt: tp(time.Date(2019, 7, 10, 12, 01, 10, 0, JST)),
						StartedAt:   tp(time.Date(2019, 8, 23, 19, 30, 0, 0, JST)),
						EndedAt:     tp(time.Date(2019, 8, 23, 22, 0, 0, 0, JST)),
					},
				},
			},
			want: goconJSON,
		},
		{
			name: "doorkeeper json",
			fields: fields{
				Title:     "Tokyo Rubyist Meetup | Doorkeeper",
				URL:       "https://trbmeetup.doorkeeper.jp/",
				UpdatedAt: tp(time.Date(2018, 5, 11, 0, 7, 44, 270000000, time.UTC)),
				Events: []Event{
					{
						Title:       "900K records per second with Ruby, Java, and JRuby",
						URL:         "https://trbmeetup.doorkeeper.jp/events/28319",
						Address:     "東京都渋谷区神泉町8-16 渋谷ファーストプレイス8F",
						PublishedAt: tp(time.Date(2015, 7, 13, 23, 48, 29, 463000000, time.UTC)),
						UpdatedAt:   tp(time.Date(2018, 5, 11, 0, 7, 44, 270000000, time.UTC)),
						StartedAt:   tp(time.Date(2015, 8, 13, 10, 0, 0, 0, time.UTC)),
						EndedAt:     tp(time.Date(2015, 8, 13, 13, 0, 0, 0, time.UTC)),
					},
				},
			},
			want: tokyorubyistJSON,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Group{
				Title:     tt.fields.Title,
				URL:       tt.fields.URL,
				Events:    tt.fields.Events,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			got, err := g.ToJSON()

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_maxTime(t *testing.T) {
	type args struct {
		times []time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "returns first",
			args: args{
				times: []time.Time{
					time.Date(2019, 1, 3, 12, 0, 0, 0, time.UTC),
					time.Date(2019, 1, 2, 12, 0, 0, 0, time.UTC),
					time.Date(2019, 1, 1, 12, 0, 0, 0, time.UTC),
				},
			},
			want: time.Date(2019, 1, 3, 12, 0, 0, 0, time.UTC),
		},
		{
			name: "returns last",
			args: args{
				times: []time.Time{
					time.Date(2019, 1, 1, 12, 0, 0, 0, time.UTC),
					time.Date(2019, 1, 2, 12, 0, 0, 0, time.UTC),
					time.Date(2019, 1, 3, 12, 0, 0, 0, time.UTC),
				},
			},
			want: time.Date(2019, 1, 3, 12, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxTime(tt.args.times)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGroup_MaxEventsUpdatedAt(t *testing.T) {
	type fields struct {
		Title     string
		URL       string
		UpdatedAt *time.Time
		Events    []Event
	}
	tests := []struct {
		name   string
		fields fields
		want   *time.Time
	}{
		{
			name: "has events",
			fields: fields{
				Events: []Event{
					{UpdatedAt: tp(time.Date(2019, 1, 1, 12, 0, 0, 0, time.UTC))},
					{UpdatedAt: tp(time.Date(2019, 1, 2, 12, 0, 0, 0, time.UTC))},
					{UpdatedAt: tp(time.Date(2019, 1, 3, 12, 0, 0, 0, time.UTC))},
				},
			},
			want: tp(time.Date(2019, 1, 3, 12, 0, 0, 0, time.UTC)),
		},
		{
			name: "has events",
			fields: fields{
				Events: []Event{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			group := &Group{
				Title:     tt.fields.Title,
				URL:       tt.fields.URL,
				UpdatedAt: tt.fields.UpdatedAt,
				Events:    tt.fields.Events,
			}

			got := group.MaxEventsUpdatedAt()

			assert.Equal(t, tt.want, got)
		})
	}
}
