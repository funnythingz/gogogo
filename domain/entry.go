package domain

import (
	"github.com/funnythingz/gogogo/ddd"
)

type Entry struct {
	ddd.Entity
	Title   string
	Content string
}
