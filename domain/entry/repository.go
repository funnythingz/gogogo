package entry

import (
	"github.com/funnythingz/gogogo/ddd"
	"github.com/funnythingz/gogogo/domain"
	"github.com/funnythingz/gogogo/mapper"
	_ "github.com/k0kubun/pp"
)

type EntryRepository struct{}

func (r *EntryRepository) Commit(entry domain.Entry) domain.Entry {
	me := mapper.Entry{}
	me.Map(entry)
	me.Commit()
	return r.Fetch(me.Id)
}

func (r *EntryRepository) Fetch(id int) domain.Entry {
	me := mapper.Entry{}
	me.Fetch(id)
	return domain.Entry{
		Entity: ddd.Entity{
			Id: me.Id,
		},
		Title:   me.Title,
		Content: me.Content,
	}
}

var (
	Repository = &EntryRepository{}
)
