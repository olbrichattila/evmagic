package actions

import "time"

type FailedCheckAction struct {
	BlogID    int64     `json:"blog_id"`
	CheckType string    `json:"check_type"`
	FailedAt  time.Time `json:"failed_at"`
	Reason    string    `json:"reason"`
}
