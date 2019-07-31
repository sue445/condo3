package connpass

import (
	"context"
	"github.com/hkurokawa/go-connpass"
	"github.com/sue445/condo3/model"
	"time"
)

// GetGroup returns group detail
func GetGroup(ctx context.Context, groupName string) (*model.Group, error) {
	page, err := FetchGroupPageWithCache(ctx, groupName)

	if err != nil {
		return nil, err
	}

	events, err := getEvents(page.SeriesID)

	if err != nil {
		return nil, err
	}

	return &model.Group{
		Title:  page.Title,
		URL:    page.URL,
		Events: events,
	}, nil
}

func getEvents(seriesID int) ([]model.Event, error) {
	query := connpass.Query{SeriesId: []int{seriesID}, Count: 100, Order: connpass.CREATE}
	result, err := query.Search()

	if err != nil {
		return []model.Event{}, err
	}

	events := []model.Event{}

	for _, resultEvent := range result.Events {
		ev := model.Event{
			Title:   resultEvent.Title,
			URL:     resultEvent.Url,
			Address: resultEvent.Address,
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
