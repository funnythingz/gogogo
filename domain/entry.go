package domain

import (
	"../ddd"
)

type Entry struct {
	ddd.Entity
	Title   string
	Content string
}
