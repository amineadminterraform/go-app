package api

import "time"

type SwaggerProject struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	GitPath     string    `json:"git_path"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"` // Use time.Time instead of pgtype.Timestamptz
}
