package entry

import (
	"../../db"
	"../../domain"
	"../../mapper"
	_ "github.com/k0kubun/pp"
)

type Repository struct{}

func (r *Repository) Commit(entry domain.Entry) {
	em := mapper.EntryMapper{}
	em.New(entry)
	em.Commit()
	r.Fetch(em.Id)
}

func (r *Repository) Fetch(id int) domain.Entry {
	em := mapper.EntryMapper{}
	em.Fetch(id)
	return domain.Entry{
		ddd.Entity{
			Id: em.Id,
		},
		Title:   em.Title,
		Content: em.Content,
		domain.Theme{
			ddd.Entity{
				Id: em.ThemeId,
			},
		},
	}
}

var (
	repo = &Repository{}
)

func Commit(entry domain.Entry) {
	repo.Commit(entry)
}

func Fetch(id int) domain.Entry {
	return repo.Fetch(id)
}
