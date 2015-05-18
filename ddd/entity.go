package ddd

import (
	"time"
)

type Entity struct {
	Identity
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
