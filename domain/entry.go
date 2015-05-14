package domain

import (
	"../db"
	"../ddd"
	_ "github.com/k0kubun/pp"
	_ "log"
	"time"
)

type Entry struct {
	ddd.Entity
	Title   string
	Content string
	ThemeId int
}

func FindEntry(id string) Entry {
	var entry Entry
	db.Dbmap.Find(&entry, id).First(&entry)
	return entry
}
