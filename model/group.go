package model

import (
	"fmt"
	"github.com/arran4/golang-ical"
	"github.com/gorilla/feeds"
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
			Description: fmt.Sprintf("開催日時：%s〜%s\n開催場所：%s", e.StartedAt.Format("2006/01/02 15:04"), e.EndedAt.Format("15:04"), e.Address),
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
