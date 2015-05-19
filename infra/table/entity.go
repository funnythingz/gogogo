package table

import (
	"time"
)

type Entity struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
}
