package connpass

import (
	"github.com/hkurokawa/go-connpass"
	"github.com/sue445/condo3/model"
	"time"
)

// GetGroupEvents returns group events
func GetGroupEvents(groupName string) ([]model.Event, error) {
	page, err := FetchGroupPage(groupName)

	if err != nil {
		return []model.Event{}, err
	}

	query := connpass.Query{SeriesId: []int{page.SeriesID}, Count: 100, Order: connpass.START}
	result, err := query.Search()

	if err != nil {
		return []model.Event{}, err
	}

	events := []model.Event{}

	for _, resultEvent := range result.Events {
		ev := model.Event{
			Title: resultEvent.Title,
			URL:   resultEvent.Url,
		}

		if resultEvent.Start != "" {
			startedAt, err := time.Parse(time.RFC3339, resultEvent.Start)

			if err != nil {
				return []model.Event{}, err
			}

			ev.StartedAt = &startedAt
		}

		if resultEvent.End != "" {
			endedAt, err := time.Parse(time.RFC3339, resultEvent.End)

			if err != nil {
				return []model.Event{}, err
			}

			ev.EndedAt = &endedAt
		}

		events = append(events, ev)
	}

	return events, nil
}
