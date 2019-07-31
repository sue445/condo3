package doorkeeper

import (
	"fmt"
	"github.com/sue445/condo3/model"
	"github.com/sue445/go-doorkeeper"
	"os"
	"time"
)

// GetGroup returns group detail
func GetGroup(groupName string, currentTime time.Time) (*model.Group, error) {
	client := doorkeeper.NewClient(os.Getenv("DOORKEEPER_ACCESS_TOKEN"))

	doorkeeperGroup, _, err := client.GetGroup(groupName)

	if err != nil {
		return nil, err
	}

	since := currentTime.AddDate(-1, 0, 0)
	until := currentTime.AddDate(1, 0, 0)
	params := doorkeeper.GetEventsParams{
		Sort:  doorkeeper.SortByPublishedAt,
		Since: &since,
		Until: &until,
	}
	doorkeeperEvents, _, err := client.GetGroupEvents(groupName, &params)

	if err != nil {
		return nil, err
	}

	group := model.Group{
		Title:  fmt.Sprintf("%s | Doorkeeper", doorkeeperGroup.Name),
		URL:    doorkeeperGroup.PublicURL,
		Events: []model.Event{},
	}

	for _, ev := range doorkeeperEvents {
		event := model.Event{
			Title:     ev.Title,
			URL:       ev.PublicURL,
			Address:   ev.Address,
			StartedAt: &ev.StartsAt,
			EndedAt:   &ev.EndsAt,
		}
		group.Events = append(group.Events, event)
	}

	return &group, nil
}
