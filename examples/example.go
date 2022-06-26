package main

import (
	"log"
	"math/rand"
	"time"

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

	rand.Seed(time.Now().Unix())

	f2, err := nanoid.NewNonSecure(21)

	if err != nil {
		panic(err)
	}

	id3 := f2()
	id4 := f2()

	log.Printf("ID 3: %s", id3)
	log.Printf("ID 4: %s", id4)

}
