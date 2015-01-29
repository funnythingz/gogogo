package main

import (
	"./db"
	"github.com/zenazn/goji"
)

func main() {
	dbmap.DbConnect()

	goji.Get("/", top)
	goji.Get("/:year/:month/:basename", article)

	goji.Serve()
}
