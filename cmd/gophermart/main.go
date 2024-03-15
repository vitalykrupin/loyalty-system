package main

import (
	"log"

	"github.com/vitalykrupin/loyalty-system/internal/router"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	router.Route()
	return nil
}
