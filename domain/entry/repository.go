package entry

import (
	"../../ddd"
	"../../domain"
	"../../mapper"
	_ "github.com/k0kubun/pp"
)

type EntryRepository struct{}

func (r *EntryRepository) Commit(entry domain.Entry) domain.Entry {
	em := mapper.Entry{}
	em.Map(entry)
	em.Commit()
	return r.Fetch(em.Id)
}

func (r *EntryRepository) Fetch(id int) domain.Entry {
	em := mapper.Entry{}
	em.Fetch(id)
	return domain.Entry{
		Entity: ddd.Entity{
			Id: em.Id,
		},
		Title:   em.Title,
		Content: em.Content,
	}
}

var (
	Repository = &EntryRepository{}
)
