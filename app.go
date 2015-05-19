package main

import (
	"./handler"
	"./infra/db"
	_ "github.com/zenazn/goji"
	"github.com/zenazn/goji/bind"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"
	"net/http"
	"regexp"
)

var (
	sHandler = &handler.SimpleHandler{}
	eHandler = &handler.ExceptionHandler{}
)

func main() {
	// Database
	db.Connect()

	// Mux
	m := web.New()

	// Routes
	m.Get("/", sHandler.Top)
	m.Get(regexp.MustCompile(`^/(?P<id>\d+)$`), sHandler.Entry)
	m.Get("/entry", http.RedirectHandler("/", 301))
	m.Post("/entry", sHandler.Create)

	// Exception
	m.NotFound(eHandler.NotFound)

	// Serve
	graceful.Serve(bind.Default(), m)
}
