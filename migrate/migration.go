package main

import (
	"../db"
	"../mapper"
	"fmt"
	"log"
	"os"
)

func main() {
	db.Connect()

	action := "migrate"
	if len(os.Args) >= 2 {
		action = os.Args[1]
	}

	log.Println(fmt.Sprintf("mode: %s", action))

	switch {
	case action == "reset":
		Reset()
		return
	default:
		Migrate()
		return
	}

}

func Reset() {
	log.Println(db.Dbmap.DropTableIfExists(&mapper.Entry{}))
	Create()
}

func Create() {
	log.Println(db.Dbmap.CreateTable(&mapper.Entry{}))
}

func Migrate() {
	log.Println(db.Dbmap.AutoMigrate(&mapper.Entry{}))
}
