package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/jaevor/go-nanoid"
)

func main() {
	f, err := nanoid.Standard(21)

	if err != nil {
		panic(err)
	}

	id1 := f()
	log.Printf("ID 1: %s", id1)

	rand.Seed(time.Now().Unix())

	f2, err := nanoid.StandardNonSecure(21)

	if err != nil {
		panic(err)
	}

	id2 := f2()
	log.Printf("ID 2: %s", id2)

	f3, err := nanoid.Custom("0123456789", 12)
	if err != nil {
		panic(err)
	}

	id3 := f3()
	log.Printf("ID 3: %s", id3)
}
