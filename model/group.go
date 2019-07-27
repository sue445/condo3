package model

import (
	"github.com/arran4/golang-ical"
)

// Group represents group info
type Group struct {
	Title  string
	URL    string
	Events []Event
}

// ToIcal return ical formatted group
func (g *Group) ToIcal() string {
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodPublish)

	for _, e := range g.Events {
		event := cal.AddEvent(e.URL)

		event.SetDescription(e.Title)
		event.SetURL(e.URL)
		event.SetLocation(e.Address)

		if e.StartedAt != nil {
			event.SetStartAt(*e.StartedAt)
		}
		if e.EndedAt != nil {
			event.SetEndAt(*e.EndedAt)
		}
	}

	return cal.Serialize()
}
