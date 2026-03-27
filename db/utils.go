package db

import (
	"log"
)

func check_err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
