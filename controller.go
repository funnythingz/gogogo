package main

import (
	"./db"
	"./helper"
	"./models"
	"fmt"
	_ "github.com/goji/param"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"github.com/yosssi/ace"
	"github.com/zenazn/goji/web"
	_ "log"
	"net/http"
	"strconv"
)

func top(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl, _ := ace.Load("views/layouts/layout", "views/top", nil)
	err := tpl.Execute(w, map[string]string{"Title": "Welcome"})

	helper.InternalServerErrorCheck(err, w)
}

func entry(c web.C, w http.ResponseWriter, r *http.Request) {

	var entry model.Entry

	if db.Dbmap.Find(&entry, c.URLParams["id"]).RecordNotFound() {
		NotFound(w, r)
		return
	}

	p := bluemonday.UGCPolicy()
	htmlContent := p.Sanitize(string(blackfriday.MarkdownCommon([]byte(entry.Content))))

	tpl, _ := ace.Load("views/layouts/layout", "views/view", nil)
	err := tpl.Execute(w, map[string]string{"Title": entry.Title, "HtmlContent": htmlContent})

	helper.InternalServerErrorCheck(err, w)
}

func newEntry(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl, _ := ace.Load("views/layouts/layout", "views/new", nil)
	err := tpl.Execute(w, nil)

	helper.InternalServerErrorCheck(err, w)
}

func postEntry(c web.C, w http.ResponseWriter, r *http.Request) {

	title := r.FormValue("entry[title]")
	content := r.FormValue("entry[content]")
	themeId, _ := strconv.Atoi(r.FormValue("entry[theme_id]"))

	entry := model.Entry{
		Title:   title,
		Content: content,
		ThemeId: themeId,
	}

	db.Dbmap.NewRecord(entry)
	db.Dbmap.Create(&entry)

	// TODO: validation
	//if err != nil || len(greet.Message) > 140 {
	//    http.Error(w, err.Error(), http.StatusBadRequest)
	//    return
	//}

	url := fmt.Sprintf("/%d", entry.Id)
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tpl, _ := ace.Load("views/layouts/layout", "views/404", nil)
	tpl.Execute(w, nil)
}
