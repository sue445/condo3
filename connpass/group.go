package connpass

import (
	"github.com/hkurokawa/go-connpass"
	"github.com/sue445/condo3/event"
	"time"
)

// GetGroupEvents returns group events
func GetGroupEvents(groupName string) ([]event.Event, error) {
	seriesID, err := fetchSeriesID(groupName)

	if err != nil {
		return []event.Event{}, err
	}

	query := connpass.Query{SeriesId: []int{seriesID}, Count: 100, Order: connpass.START}
	result, err := query.Search()

	if err != nil {
		return []event.Event{}, err
	}

	events := []event.Event{}

	for _, resultEvent := range result.Events {
		ev := event.Event{
			Title: resultEvent.Title,
			URL:   resultEvent.Url,
		}

		if resultEvent.Start != "" {
			startedAt, err := time.Parse(time.RFC3339, resultEvent.Start)

			if err != nil {
				return []event.Event{}, err
			}

			ev.StartedAt = &startedAt
		}

		if resultEvent.End != "" {
			endedAt, err := time.Parse(time.RFC3339, resultEvent.End)

			if err != nil {
				return []event.Event{}, err
			}

			ev.EndedAt = &endedAt
		}

		events = append(events, ev)
	}

	return events, nil
}
