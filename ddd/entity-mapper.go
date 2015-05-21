package ddd

import (
	"time"
)

type EntityMapper struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
}
