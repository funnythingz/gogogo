package db

import (
	"fmt"
	"log"
	"os"
)

func Connect() {
	env := "development"
	if len(os.Args) >= 2 {
		env = os.Args[1]
	}

	log.Println(fmt.Sprintf("mode: %s", env))

	switch {
	case env == "production":
		DbConnect("production")
		return
	case env == "development":
		DbConnect("development")
		Dbmap.LogMode(true)
		return
	default:
		DbConnect("development")
		Dbmap.LogMode(true)
		return
	}
}
