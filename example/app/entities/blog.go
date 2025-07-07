package entities

import "time"

type Blogs struct {
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	Blog      string    `json:"blog"`
}
