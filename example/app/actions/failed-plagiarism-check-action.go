package actions

import "time"

type FailedPlagiarismCheckAction struct {
	FailedAt time.Time `json:"failed_at"`
	Reason   string    `json:"reason"`
}
