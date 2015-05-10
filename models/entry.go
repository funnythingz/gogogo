package model

import (
	"../db"
	_ "github.com/k0kubun/pp"
	_ "log"
	"time"
)

type Entry struct {
	Id        int
	Title     string
	Content   string
	ThemeId   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FindEntry(id string) Entry {
	var entry Entry
	db.Dbmap.Find(&entry, id).First(&entry)
	return entry
}
