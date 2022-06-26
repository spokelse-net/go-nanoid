package main

import (
	"log"

	"github.com/jaevor/go-nanoid"
)

func main() {
	f, err := nanoid.New(21)

	if err != nil {
		panic(err)
	}

	id1 := f()
	id2 := f()

	log.Printf("ID 1: %s", id1)
	log.Printf("ID 2: %s", id2)
}
