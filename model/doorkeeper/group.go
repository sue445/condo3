package doorkeeper

import (
	"fmt"
	"github.com/sue445/condo3/model"
	"github.com/sue445/go-doorkeeper"
	"golang.org/x/sync/errgroup"
	"time"
)

// GetGroup returns group detail
func GetGroup(accessToken string, groupName string, currentTime time.Time) (*model.Group, error) {
	client := doorkeeper.NewClient(accessToken)

	var eg errgroup.Group

	group := model.Group{
		Events: []model.Event{},
	}

	eg.Go(func() error {
		g, _, err := client.GetGroup(groupName)

		if err != nil {
			return err
		}

		group.Title = fmt.Sprintf("%s | Doorkeeper", g.Name)
		group.URL = g.PublicURL

		return nil
	})

	eg.Go(func() error {
		since := currentTime.AddDate(-1, 0, 0)
		until := currentTime.AddDate(1, 0, 0)
		params := doorkeeper.GetEventsParams{
			Sort:  doorkeeper.SortByPublishedAt(),
			Since: &since,
			Until: &until,
		}
		events, _, err := client.GetGroupEvents(groupName, &params)

		if err != nil {
			return err
		}

		for _, ev := range events {
			event := model.Event{
				Title:     ev.Title,
				URL:       ev.PublicURL,
				Address:   ev.Address,
				UpdatedAt: ev.UpdatedAt,
				StartedAt: &ev.StartsAt,
				EndedAt:   &ev.EndsAt,
			}
			group.Events = append(group.Events, event)
		}

		return nil
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	updatedAt := group.MaxEventsUpdatedAt()

	if updatedAt != nil {
		group.UpdatedAt = updatedAt
	}

	return &group, nil
}
