package api

import (
	"time"
)

type SwaggerProject struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	GitPath     string    `json:"git_path"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"` // Use time.Time instead of pgtype.Timestamptz
}
type SwaggerArgoWorkflow struct {
	ID          int64              `json:"id"`
	Name        string             `json:"name"`
	Path        string             `json:"path"`
	Description string        `json:"description"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CreatedAt   time.Time  `json:"created_at"`
}

type SwaggerEnvLayer struct {
	ID               int64              `json:"id"`
	EnvironmentID    int64              `json:"environment_id"`
	S3Path           string             `json:"s3_path"`
	UpdatedAt        time.Time `json:"updated_at"`
	CreatedAt        time.Time `json:"created_at"`
	ProcessID        int64              `json:"process_id"`
	CurrentRequestID  int64        `json:"current_request_id"`
}

type SwaggerProcess struct {
	ID         int64              `json:"id"`
	ArgoID     int64              `json:"argo_id"`
	Name       string             `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	TemplateID int64              `json:"template_id"`
}


type SwaggerProjectEnvironment struct {
	ID          int64              `json:"id"`
	GitBranch   string             `json:"git_branch"`
	ProjectID   int64              `json:"project_id"`
	Description string        `json:"description"`
	UpdatedAt   time.Time 	`json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type SwaggerRequest struct {
	ID        int64              `json:"id"`
	LayerID   int64              `json:"layer_id"`
	Source    string             `json:"source"`
	Payload   string             `json:"payload"`
	Status    string             `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type SwaggerTemplate struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"`
}