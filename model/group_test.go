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

	jst, _ := time.LoadLocation("Asia/Tokyo")

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
						StartedAt: tp(time.Date(2019, 8, 23, 19, 30, 0, 0, jst)),
						EndedAt:   tp(time.Date(2019, 8, 23, 22, 0, 0, 0, jst)),
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
	goconAtom := `<?xml version="1.0" encoding="UTF-8"?><feed xmlns="http://www.w3.org/2005/Atom">
  <title>Go Conference - connpass</title>
  <id>https://gocon.connpass.com/</id>
  <updated></updated>
  <link href="https://gocon.connpass.com/"></link>
  <entry>
    <title>Go 1.13 Release Party in Tokyo</title>
    <updated></updated>
    <id>https://gocon.connpass.com/event/139024/</id>
    <link href="https://gocon.connpass.com/event/139024/" rel="alternate"></link>
    <summary type="html">開催日時：2019/08/23 19:30〜22:00&#xA;開催場所：東京都港区六本木6-10-1 (六本木ヒルズ森タワー18F)</summary>
  </entry>
</feed>`

	tokyurubykaigiAtom := `<?xml version="1.0" encoding="UTF-8"?><feed xmlns="http://www.w3.org/2005/Atom">
  <title>TokyuRubyKaigi | Doorkeeper</title>
  <id>https://tokyu-rubykaigi.doorkeeper.jp/</id>
  <updated></updated>
  <link href="https://tokyu-rubykaigi.doorkeeper.jp/"></link>
  <entry>
    <title>TokyuRubyKaigi13 一般参加者募集(LT発表者は登録不要です)</title>
    <updated></updated>
    <id>https://tokyu-rubykaigi.doorkeeper.jp/events/91543</id>
    <link href="https://tokyu-rubykaigi.doorkeeper.jp/events/91543" rel="alternate"></link>
    <summary type="html">開催日時：2019/06/29 14:00〜19:30&#xA;開催場所：〒106-0032 東京都港区六本木4-1-4 黒崎ビル4階</summary>
  </entry>
</feed>`

	jst, _ := time.LoadLocation("Asia/Tokyo")

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
			name: "generate gocon atom",
			fields: fields{
				Title: "Go Conference - connpass",
				URL:   "https://gocon.connpass.com/",
				Events: []Event{
					{
						Title:     "Go 1.13 Release Party in Tokyo",
						URL:       "https://gocon.connpass.com/event/139024/",
						Address:   "東京都港区六本木6-10-1 (六本木ヒルズ森タワー18F)",
						StartedAt: tp(time.Date(2019, 8, 23, 19, 30, 0, 0, jst)),
						EndedAt:   tp(time.Date(2019, 8, 23, 22, 0, 0, 0, jst)),
					},
				},
			},
			want: goconAtom,
		},
		{
			name: "doorkeeper atom",
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
			want: tokyurubykaigiAtom,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Group{
				Title:  tt.fields.Title,
				URL:    tt.fields.URL,
				Events: tt.fields.Events,
			}
			got, err := g.ToAtom()

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
