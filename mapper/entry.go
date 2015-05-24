package mapper

import (
	"../db"
	"../ddd"
	"../domain"
)

type Entry struct {
	ddd.EntityMapper
	Title   string
	Content string
}

func (m *Entry) Map(entry domain.Entry) {
	m.Title = entry.Title
	m.Content = entry.Content
}

func (m *Entry) Commit() {
	db.Dbmap.NewRecord(m)
	db.Dbmap.Create(&m)
}

func (m *Entry) Fetch(id int) {
	db.Dbmap.Find(&m, id).First(&m)
}
