package actions

import "time"

type FailedSpamCheckAction struct {
	FailedAt time.Time `json:"failed_at"`
	Reason   string    `json:"reason"`
}
