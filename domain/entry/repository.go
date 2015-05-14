package entry

import (
	"../../db"
	"../../domain"
	_ "github.com/k0kubun/pp"
)

type Repository struct{}

func (r *Repository) Commit(entry domain.Entry) domain.Entry {
	db.Dbmap.NewRecord(entry)
	db.Dbmap.Create(&entry)
	return entry
}

func (r *Repository) Fetch(id int) domain.Entry {
	entry := domain.Entry{}
	db.Dbmap.Find(&entry, id).First(&entry)
	return entry
}

var (
	repo = &Repository{}
)

func Commit(iro domain.Iro) domain.Iro {
	return repo.Commit(iro)
}

func Fetch(id int) domain.Iro {
	return repo.Fetch(id)
}
