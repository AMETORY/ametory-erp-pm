package objects

import "time"

type ScheduledMessage struct {
	To      string    `json:"to"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}
