package model

import (
	"fmt"
	"strings"
	"time"
)

// Event represents event data
type Event struct {
	Title       string
	URL         string
	Address     string
	UpdatedAt   *time.Time
	PublishedAt *time.Time
	StartedAt   *time.Time
	EndedAt     *time.Time
}

func (e *Event) atomDescription() string {
	var lines []string

	if e.StartedAt != nil || e.EndedAt != nil {
		term := "開催日時："

		if e.StartedAt != nil && e.EndedAt != nil {
			term += fmt.Sprintf("%s〜%s", e.StartedAt.In(JST).Format("2006/01/02 15:04"), e.EndedAt.In(JST).Format("15:04"))
		} else if e.StartedAt != nil {
			term += fmt.Sprintf("%s〜", e.StartedAt.In(JST).Format("2006/01/02 15:04"))
		} else {
			term += fmt.Sprintf("〜%s", e.EndedAt.In(JST).Format("2006/01/02 15:04"))
		}

		lines = append(lines, term)
	}

	lines = append(lines, fmt.Sprintf("開催場所：%s", e.Address))

	return strings.Join(lines, "\n")
}
