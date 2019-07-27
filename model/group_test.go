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
				URL: "https://gocon.connpass.com/",
				Events: []Event{
					{
						Title: "Go 1.13 Release Party in Tokyo",
						URL: "https://gocon.connpass.com/event/139024/",
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
