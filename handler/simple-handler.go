package handler

import (
	"encoding/json"
	"github.com/funnythingz/gogogo/domain"
	"github.com/funnythingz/gogogo/domain/entry"
	"github.com/funnythingz/gogogo/helper"
	_ "github.com/k0kubun/pp"
	"github.com/zenazn/goji/web"
	"io"
	"net/http"
	"strconv"
	"unicode/utf8"
)

type SimpleHandler struct{}

// GET
func (h *SimpleHandler) Top(c web.C, w http.ResponseWriter, r *http.Request) {
	//TODO
}

// GET
func (h *SimpleHandler) Entry(c web.C, w http.ResponseWriter, r *http.Request) {
	//TODO
	id, _ := strconv.Atoi(c.URLParams["id"])
	e := entry.Repository.Fetch(id)
	if e.Id == 0 {
		io.WriteString(w, "Not found.")
		return
	}
	response, _ := json.Marshal(e)
	io.WriteString(w, string(response))
}

// POST
func (h *SimpleHandler) Create(c web.C, w http.ResponseWriter, r *http.Request) {
	//TODO

	title := r.FormValue("entry[title]")
	content := r.FormValue("entry[content]")

	// Validation
	errors := []string{}

	if utf8.RuneCountInString(title) <= 0 {
		errors = append(errors, "input Title must be blank.")
	}
	if utf8.RuneCountInString(title) >= 120 {
		errors = append(errors, "input Title max 120")
	}
	if utf8.RuneCountInString(content) <= 0 {
		errors = append(errors, "input Content must be blank.")
	}
	if utf8.RuneCountInString(content) >= 120 {
		errors = append(errors, "input Content max 120")
	}

	if len(errors) > 0 {
		helper.ResultMessageJSON(w, errors)
		return
	}

	resultEntry := entry.Repository.Commit(domain.Entry{
		Title:   title,
		Content: content,
	})

	response, _ := json.Marshal(resultEntry)
	io.WriteString(w, string(response))
}
