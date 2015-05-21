package mapper

import (
	"../ddd"
)

type EntryMapper struct {
	ddd.EntityMapper
	Title   string
	Content string
	ThemeId int
}

func (m *Entry) New(entry domain.Entry) {
	m.Title = entry.Title
	m.Content = entry.Content
	m.ThemeId = entry.Theme.Id
}

func (m *Entry) Commit() {
	db.Dbmap.NewRecord(m)
	db.Dbmap.Create(&m)
}

func (m *Entry) Fetch(id int) {
	db.Dbmap.Find(&m, id).First(&m)
}
