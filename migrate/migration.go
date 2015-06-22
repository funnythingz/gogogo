package main

import (
	"fmt"
	"github.com/funnythingz/gogogo/db"
	"github.com/funnythingz/gogogo/mapper"
	"log"
	"os"
)

func main() {
	db.Connect()

	action := "migrate"
	if len(os.Args) >= 2 {
		action = os.Args[2]
	}

	log.Println(fmt.Sprintf("mode: %s", action))

	switch {
	case action == "migrate":
		Migrate()
		return
	case action == "reset":
		Reset()
		return
	}
}

func Reset() {
	Drop()
	Create()
}

func Create() {
	db.Dbmap.CreateTable(&mapper.Entry{})
}

func Migrate() {
	db.Dbmap.AutoMigrate(&mapper.Entry{})
}

func Drop() {
	db.Dbmap.DropTableIfExists(&mapper.Entry{})
}
