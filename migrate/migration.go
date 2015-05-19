package main

import (
	"../infra/db"
	"../infra/table"
	"log"
)

func main() {
	db.Connect()
	reset()
	create()
	migrate()
}

func reset() {
	log.Println(db.Dbmap.DropTableIfExists(&table.Entry{}))
}

func create() {
	log.Println(db.Dbmap.CreateTable(&table.Entry{}))
}

func migrate() {
	log.Println(db.Dbmap.AutoMigrate(&table.Entry{}))
}
