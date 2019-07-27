package model

import "time"

// Event represents event data
type Event struct {
	Title     string
	URL       string
	StartedAt *time.Time
	EndedAt   *time.Time
}
