package actions

import "time"

type BlogCheckAction struct {
	CreatedAt time.Time `json:"created_at"`
	BlogID    int64     `json:"blog"`
}
