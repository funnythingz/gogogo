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

	switch {
	case env == "production":
		DbConnect("production")
		return
	case env == "development":
		DbConnect("development")
		Dbmap.LogMode(true)
		return
	}

	log.Println(fmt.Sprintf("mode: %s", env))
}
