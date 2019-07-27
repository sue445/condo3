package model

import (
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
PRODID:-//Arran Ubels//Golang ICS library
METHOD:PUBLISH
BEGIN:VEVENT
UID:https://gocon.connpass.com/event/139024/
DESCRIPTION:Go 1.13 Release Party in Tokyo
URL:https://gocon.connpass.com/event/139024/
LOCATION:東京都港区六本木6-10-1 (六本木ヒルズ森タワー18F
 )
DTSTART:20190823T103000Z
DTEND:20190823T130000Z
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
						StartedAt: tp(time.Date(2019, 8, 23, 19, 30, 0, 0, time.Local)),
						EndedAt:   tp(time.Date(2019, 8, 23, 22, 0, 0, 0, time.Local)),
					},
				},
			},
			want: strings.ReplaceAll(goconIcal, "\n", "\r\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Group{
				Title:  tt.fields.Title,
				URL:    tt.fields.URL,
				Events: tt.fields.Events,
			}
			if got := g.ToIcal(); got != tt.want {
				t.Errorf("Group.ToIcal() = %v, want %v", got, tt.want)
			}
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

	type fields struct {
		Title  string
		URL    string
		Events []Event
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
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
						StartedAt: tp(time.Date(2019, 8, 23, 19, 30, 0, 0, time.Local)),
						EndedAt:   tp(time.Date(2019, 8, 23, 22, 0, 0, 0, time.Local)),
					},
				},
			},
			want: goconAtom,
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
			if (err != nil) != tt.wantErr {
				t.Errorf("Group.ToAtom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Group.ToAtom() = %v, want %v", got, tt.want)
			}
		})
	}
}
