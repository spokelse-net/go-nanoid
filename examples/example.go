package main

import (
	"log"

	"github.com/jaevor/go-nanoid"
)

func main() {
	canonicID, err := nanoid.Standard(21)
	if err != nil {
		panic(err)
	}

	id1 := canonicID()
	log.Printf("ID 1: %s", id1) // se-jlhSbQbwlviPDFbfGe

	customID, err := nanoid.Custom("0123456789", 12)
	if err != nil {
		panic(err)
	}

	id2 := customID()
	log.Printf("ID 2: %s", id2) // 466568050433
}
