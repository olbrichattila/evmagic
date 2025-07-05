package actions

import "time"

type FailedProfanityCheckAction struct {
	FailedAt time.Time `json:"failed_at"`
	Reason   string    `json:"reason"`
}
