package entities

type Blogs struct {
	Id int64 `json:"id"`
	// Check why entity does not work with time.Time
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
	Blog      string `json:"blog"`
	Banned    bool   `json:"banned"`
}
