/*
Copyright 2022 jaevor.
License can be found in the LICENSE file.

Original reference: https://github.com/ai/nanoid
*/

package nanoid

import (
	cryptoRand "crypto/rand"
	"errors"
	"math"
	"math/bits"
	mathRand "math/rand"
	"sync"
)

// Default characters (A-Za-z0-9_-).
var defaultCharset = []rune("useandom-26T198340PX75pxJACKVERYMINDBUSHWOLF_GQZbfghjklqvwyzrict")

type generator = func() string

// Creates a new generator for Nano IDs. Recommended length is 21.
// Returns error if length is not between 2 and 255 (inclusive).
// Concurrency safe.
func New(length int) (generator, error) {
	if invalidLength(length) {
		return nil, errInvalidLength
	}

	var mu sync.Mutex

	// Multiplying by 128 and using an offset so
	// that the bytes only have to be refilled
	// every 129th nanoid. This is more efficient.

	b := make([]byte, length*128)
	size := len(b)
	offset := 0
	cryptoRand.Read(b)

	id := make([]rune, length)

	return func() string {
		mu.Lock()
		defer mu.Unlock()

		// If all the bytes in the slice
		// have been used, refill.
		if offset == size {
			cryptoRand.Read(b)
			offset = 0
		}

		for i := 0; i < length; i++ {
			// Index using the offset.
			id[i] = defaultCharset[b[i+offset]&63]
		}

		// Extend the offset.
		offset += length

		return string(id)
	}, nil
}

// TODO: make NewNonSecure faster

// Create a non-secure Nano ID generator.
// Non-secure is faster than secure because it uses pseudorandom numbers.
// Returns error if length is not between 2 and 255 (inclusive).
// Remember to seed using math.Seed().
// Concurrency safe.
func NewNonSecure(length int) (generator, error) {
	if invalidLength(length) {
		return nil, errInvalidLength
	}

	var mu sync.Mutex

	// Reuse.
	id := make([]rune, length)

	return func() string {
		mu.Lock()
		defer mu.Unlock()

		for i := 0; i < length; i++ {
			id[i] = defaultCharset[mathRand.Intn(64)]
		}

		return string(id)
	}, nil
}

// Create a Nano ID generator that uses a custom character set.
// Concurrency safe.
func NewCustom(charset string, length int) (generator, error) {
	if invalidLength(length) {
		return nil, errInvalidLength
	}

	var mu sync.Mutex

	setLen := len(charset)
	runicSet := []rune(charset)

	// Because the custom character-set is not guaranteed to have
	// 64 chars to utilise, we have to calculate a suitable mask.
	// This is 1:1 to the original implementation.
	clz := bits.LeadingZeros32((uint32(setLen) - 1) | 1)
	mask := (2 << (31 - clz)) - 1
	w := (1.6 * float64(mask*length)) / float64(setLen)
	step := int(math.Ceil(w))

	// Will be reusing the same rune and byte slices.
	id := make([]rune, length)
	b := make([]byte, step)

	return func() string {
		mu.Lock()
		defer mu.Unlock()

		for u := 0; ; {
			cryptoRand.Read(b)

			for i := 0; i < step; i++ {
				idx := b[i] & byte(mask)

				if idx < byte(setLen) {
					// id.WriteRune(runicSet[idx])
					id[u] = runicSet[idx]
					u++
					if u == length {
						return string(id)
					}
				}
			}
		}
	}, nil
}

// Create a non-secure Nano ID generator that uses a custom character set.
// Non-secure is faster than secure because it uses pseudorandom numbers.
// Returns error if length is not between 2 and 255 (inclusive).
// Remember to seed using math.Seed().
// Concurrency safe.
func NewCustomNonSecure(charset string, length int) (generator, error) {
	if invalidLength(length) {
		return nil, errInvalidLength
	}

	var mu sync.Mutex

	runicSet := []rune(charset)
	setLen := len(runicSet)

	// var id strings.Builder
	// id.Grow(length)
	id := make([]rune, length)

	return func() string {
		mu.Lock()
		defer mu.Unlock()

		// id.Reset()

		for i := 0; i < length; i++ {
			id[i] = runicSet[mathRand.Intn(setLen)]
		}

		return string(id)
	}, nil
}

// Making life easier.

var errInvalidLength = errors.New("length must be between 2 and 255 (inclusive)")

func invalidLength(length int) bool {
	return length < 2 || length > 255
}
