package airhousePublicApi

import "time"

type WebHookJSONDataResponse struct {
	ID         string      `json:"id"`
	ObjectType string      `json:"object_type"`
	Event      string      `json:"event"`
	Property   Property    `json:"property"`
	Body       interface{} `json:"body"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}
