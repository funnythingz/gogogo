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

	if env == "production" {
		DbConnect("production")
	} else {
		DbConnect("development")
		Dbmap.LogMode(true)
	}

	log.Println(fmt.Sprintf("mode: %s", env))
}
