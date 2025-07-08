package entities

import "time"

type BlogCheck struct {
	TableName any       `tableName:"blog_checks"`
	BlogId    int64     `json:"blog_id"`
	CheckType string    `json:"check_type"`
	Reason    string    `json:"reason"`
	CreatedAt time.Time `json:"created_at"`
}
