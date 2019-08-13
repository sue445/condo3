package model

import (
	"encoding/xml"
	"github.com/lestrrat-go/ical"
	"golang.org/x/tools/blog/atom"
	"sort"
	"time"
)

const (
	icalTimeFormat = "20060102T150405Z"
)

// Group represents group info
type Group struct {
	Title     string
	URL       string
	UpdatedAt *time.Time
	Events    []Event
}

// ToIcal return ical formatted group
func (g *Group) ToIcal() string {
	cal := ical.New()
	cal.AddProperty("PRODID", "-//sue445//condo3.appspot.com//JA")
	cal.AddProperty("METHOD", "PUBLISH")
	cal.AddProperty("CALSCALE", "GREGORIAN")
	cal.AddProperty("X-WR-CALDESC", g.Title+"\n"+g.URL)
	cal.AddProperty("X-WR-CALNAME", g.Title)
	cal.AddProperty("X-WR-TIMEZONE", "UTC")

	for _, e := range g.Events {
		event := ical.NewEvent()
		event.AddProperty("UID", e.URL)
		event.AddProperty("SUMMARY", e.Title)
		event.AddProperty("DESCRIPTION", e.URL)
		event.AddProperty("URL", e.URL)
		event.AddProperty("LOCATION", e.Address)

		if e.StartedAt != nil {
			event.AddProperty("DTSTART", e.StartedAt.UTC().Format(icalTimeFormat))
		}
		if e.EndedAt != nil {
			event.AddProperty("DTEND", e.EndedAt.UTC().Format(icalTimeFormat))
		}

		cal.AddEntry(event)
	}

	return cal.String()
}

// ToAtom return atom formatted group
func (g *Group) ToAtom() (string, error) {
	feed := atom.Feed{
		Title: g.Title,
		ID:    g.URL,
		Link: []atom.Link{
			{Href: g.URL},
		},
	}

	if g.UpdatedAt != nil {
		feed.Updated = atom.Time(g.UpdatedAt.In(JST))
	}

	for _, event := range g.Events {
		entry := atom.Entry{
			Title: event.Title,
			Link: []atom.Link{
				{Href: event.URL, Rel: "alternate"},
			},
			ID: event.URL,
			Summary: &atom.Text{
				Type: "html",
				Body: event.atomDescription(),
			},
		}

		if event.UpdatedAt != nil {
			entry.Updated = atom.Time(event.UpdatedAt.In(JST))
		}

		if event.PublishedAt != nil {
			entry.Published = atom.Time(event.PublishedAt.In(JST))
		}

		feed.Entry = append(feed.Entry, &entry)
	}

	data, err := xml.MarshalIndent(&feed, "", "  ")
	if err != nil {
		return "", err
	}

	return xml.Header + string(data), nil
}

// MaxEventsUpdatedAt returns max UpdatedAt in Events
func (g *Group) MaxEventsUpdatedAt() *time.Time {
	if len(g.Events) == 0 {
		return nil
	}

	var times []time.Time

	for _, event := range g.Events {
		times = append(times, *event.UpdatedAt)
	}

	t := maxTime(times)
	return &t
}

func maxTime(times []time.Time) time.Time {
	sort.Slice(times, func(i, j int) bool {
		return times[i].After(times[j])
	})

	return times[0]
}
