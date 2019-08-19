package sandbox

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sue445/condo3/model"
	"io/ioutil"
	"net/http"
	"time"
)

type group struct {
	Title     string  `json:"title"`
	URL       string  `json:"url"`
	UpdatedAt string  `json:"updated_at"`
	Events    []event `json:"events"`
}

type event struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Address     string `json:"address"`
	UpdatedAt   string `json:"updated_at"`
	PublishedAt string `json:"published_at"`
	StartedAt   string `json:"started_at"`
	EndedAt     string `json:"ended_at"`
}

// GetGroup returns group detail
func GetGroup(groupName string) (*model.Group, error) {
	url := fmt.Sprintf("https://sue445.github.io/condo3-sandbox/data/%s.json", groupName)

	resp, err := http.Get(url)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return nil, errors.New(resp.Status)
	}

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res group
	err = json.Unmarshal(byteArray, &res)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	group := model.Group{
		Title: res.Title,
		URL:   res.URL,
	}

	if res.UpdatedAt != "" {
		updatedAt, err := time.ParseInLocation(time.RFC3339, res.UpdatedAt, model.JST)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		group.UpdatedAt = &updatedAt
	}

	for _, e := range res.Events {
		event, err := e.toEvent()

		if err != nil {
			return nil, errors.WithStack(err)
		}

		group.Events = append(group.Events, *event)
	}

	return &group, nil
}

func (e *event) toEvent() (*model.Event, error) {
	event := model.Event{
		Title:   e.Title,
		URL:     e.URL,
		Address: e.Address,
	}

	if e.UpdatedAt != "" {
		updatedAt, err := time.ParseInLocation(time.RFC3339, e.UpdatedAt, model.JST)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		event.UpdatedAt = &updatedAt
	}

	if e.PublishedAt != "" {
		publishedAt, err := time.ParseInLocation(time.RFC3339, e.PublishedAt, model.JST)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		event.PublishedAt = &publishedAt
	}

	if e.StartedAt != "" {
		startedAt, err := time.ParseInLocation(time.RFC3339, e.StartedAt, model.JST)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		event.StartedAt = &startedAt
	}

	if e.EndedAt != "" {
		endedAt, err := time.ParseInLocation(time.RFC3339, e.EndedAt, model.JST)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		event.EndedAt = &endedAt
	}

	return &event, nil
}
