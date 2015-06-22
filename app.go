package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/funnythingz/gogogo/db"
	"github.com/funnythingz/gogogo/handler"
	"github.com/goji/glogrus"
	_ "github.com/zenazn/goji"
	"github.com/zenazn/goji/bind"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	"net/http"
	"regexp"
)

var (
	sHandler = handler.SimpleHandler{}
	eHandler = handler.ExceptionHandler{}
)

func main() {
	// Database
	db.Connect()

	// Mux
	m := web.New()

	// Logger
	m.Abandon(middleware.Logger)
	logr := logrus.New()
	logr.Formatter = new(logrus.JSONFormatter)
	m.Use(glogrus.NewGlogrus(logr, "gogogo"))

	// Routes
	m.Get("/", sHandler.Top)
	m.Get(regexp.MustCompile(`^/entries/(?P<id>\d+)$`), sHandler.Entry)
	m.Get("/entries", http.RedirectHandler("/", 404))
	m.Post("/entries", sHandler.Create)

	// Exception
	m.NotFound(eHandler.NotFound)

	// Serve
	graceful.Serve(bind.Default(), m)
}
