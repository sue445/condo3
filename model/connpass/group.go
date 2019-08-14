package connpass

import (
	"github.com/hkurokawa/go-connpass"
	"github.com/sue445/condo3/model"
	"github.com/sue445/condo3/model/connpass/grouppage"
	"time"
)

// GetGroup returns group detail
func GetGroup(memcachedConfig *model.MemcachedConfig, groupName string, currentTime time.Time) (*model.Group, error) {
	page, err := grouppage.FetchGroupPageWithCache(memcachedConfig, groupName)

	if err != nil {
		return nil, err
	}

	events, err := getEvents(page.SeriesID, currentTime)

	if err != nil {
		return nil, err
	}

	group := model.Group{
		Title:  page.Title,
		URL:    page.URL,
		Events: events,
	}

	updatedAt := group.MaxEventsUpdatedAt()

	if updatedAt != nil {
		group.UpdatedAt = updatedAt
	}

	return &group, nil
}

func getEvents(seriesID int, currentTime time.Time) ([]model.Event, error) {
	query := connpass.Query{
		SeriesId: []int{seriesID},
		Count:    100,
		Order:    connpass.CREATE,
		Time:     getTerms(currentTime, 6, 6),
	}
	result, err := query.Search()

	if err != nil {
		return []model.Event{}, err
	}

	events := []model.Event{}

	for _, resultEvent := range result.Events {
		updatedAt, err := time.ParseInLocation(time.RFC3339, resultEvent.Updated, model.JST)

		if err != nil {
			return []model.Event{}, err
		}

		ev := model.Event{
			Title:     resultEvent.Title,
			URL:       resultEvent.Url,
			Address:   resultEvent.Address,
			UpdatedAt: &updatedAt,
		}

		if resultEvent.Start != "" {
			startedAt, err := time.ParseInLocation(time.RFC3339, resultEvent.Start, model.JST)

			if err != nil {
				return []model.Event{}, err
			}

			ev.StartedAt = &startedAt
		}

		if resultEvent.End != "" {
			endedAt, err := time.ParseInLocation(time.RFC3339, resultEvent.End, model.JST)

			if err != nil {
				return []model.Event{}, err
			}

			ev.EndedAt = &endedAt
		}

		events = append(events, ev)
	}

	return events, nil
}

func getTerms(currentTime time.Time, beforeMonth int, afterMonth int) []connpass.Time {
	currentMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, time.UTC)
	startMonth := currentMonth.AddDate(0, -beforeMonth, 0)

	// NOTE: time1.Before(time2) = time1 < time2
	endMonth := currentMonth.AddDate(0, afterMonth, 1)

	var months []connpass.Time

	for month := startMonth; month.Before(endMonth); month = month.AddDate(0, 1, 0) {
		months = append(months, connpass.Time{Year: month.Year(), Month: int(month.Month())})
	}

	return months
}
