package main

import (
	"log"

	"github.com/jaevor/go-nanoid"
)

func main() {
	canonic, err := nanoid.Standard(21)
	if err != nil {
		panic(err)
	}

	id1 := canonic()
	log.Printf("ID 1: %s", id1) // se-jlhSbQbwlviPDFbfGe

	// Makes sense to use CustomASCII since 0-9 is ascii.
	custom, err := nanoid.CustomASCII("0123456789", 12)
	if err != nil {
		panic(err)
	}

	id2 := custom()
	log.Printf("ID 2: %s", id2) // 466568050433
}
