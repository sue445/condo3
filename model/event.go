package model

import (
	"fmt"
	"time"
)

// Event represents event data
type Event struct {
	Title     string
	URL       string
	Address   string
	StartedAt *time.Time
	EndedAt   *time.Time
}

func (e *Event) atomDescription() string {
	return fmt.Sprintf("開催日時：%s〜%s\n開催場所：%s", e.StartedAt.In(JST).Format("2006/01/02 15:04"), e.EndedAt.In(JST).Format("15:04"), e.Address)
}
