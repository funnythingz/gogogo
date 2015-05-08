package main

import (
	"./db"
	"github.com/shaoshing/train"
	"github.com/zenazn/goji/bind"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"
	"net/http"
	"regexp"
)

func main() {
	db.DbDevelopmentConnect()
	train.ConfigureHttpHandler(nil)

	// Mux
	m := web.New()

	// Routes
	m.Get("/", top)
	m.Get(regexp.MustCompile(`^/(?P<id>\d+)$`), entry)
	m.Get("/new", newEntry)
	m.Get("/entry", http.RedirectHandler("/", 301))
	m.Post("/entry", createEntry)

	// Exception
	m.NotFound(NotFound)

	// Serve
	graceful.Serve(bind.Default(), m)
}
