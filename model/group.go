package model

import (
	"github.com/gorilla/feeds"
	"github.com/lestrrat-go/ical"
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
	UpdatedAt time.Time
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
	feed := &feeds.Feed{
		Title: g.Title,
		Link:  &feeds.Link{Href: g.URL},
		Items: []*feeds.Item{},
	}

	for _, e := range g.Events {
		item := feeds.Item{
			Title:       e.Title,
			Link:        &feeds.Link{Href: e.URL},
			Description: e.atomDescription(),
			Id:          e.URL,
		}
		feed.Items = append(feed.Items, &item)
	}

	atom, err := feed.ToAtom()

	if err != nil {
		return "", err
	}

	return atom, nil
}

// ApplyUpdatedAt apply UpdatedAt from Events
func (g *Group) ApplyUpdatedAt() {
	var times []time.Time

	for _, event := range g.Events {
		times = append(times, event.UpdatedAt)
	}

	g.UpdatedAt = maxTime(times)
}

func maxTime(times []time.Time) time.Time {
	sort.Slice(times, func(i, j int) bool {
		return times[i].After(times[j])
	})

	return times[0]
}
