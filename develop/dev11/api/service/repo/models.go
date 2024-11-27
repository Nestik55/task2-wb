package repo

import "time"

type Event struct {
	UserID      int       `json:"user_id"`
	Time        time.Time `json:"time"`
	Description string    `json:"description"`
}

type Events []Event
