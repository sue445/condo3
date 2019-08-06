package model

import "time"

// Event represents event data
type Event struct {
	Title     string
	URL       string
	Address   string
	StartedAt *time.Time
	EndedAt   *time.Time
}
