package main

import (
	"../db"
	"../domain"
	"log"
)

func main() {
	db.Connect()
	reset()
	create()
	migrate()
}

func reset() {
	log.Println(db.Dbmap.DropTableIfExists(&domain.Entry{}))
}

func create() {
	log.Println(db.Dbmap.CreateTable(&domain.Entry{}))
}

func migrate() {
	log.Println(db.Dbmap.AutoMigrate(&domain.Entry{}))
}
