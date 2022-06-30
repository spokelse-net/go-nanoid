/*
Copyright 2022 jaevor.
License can be found in the LICENSE file.

Original reference: https://github.com/ai/nanoid
*/

package nanoid

import (
	crand "crypto/rand"
	"encoding/binary"
	"errors"
	"math"
	"math/bits"
	mrand "math/rand"
	"sync"
)

const alphabetSize = 64

// Default characters (A-Za-z0-9_-).
// Using less memory with [64]byte{...} than []byte(...).
var defaultAlphabet = [alphabetSize]byte{
	'a', 'b', 'c', 'd',
	'e', 'f', 'g', 'h',
	'i', 'j', 'k', 'l',
	'm', 'n', 'o', 'p',
	'q', 'r', 's', 't',
	'u', 'v', 'w', 'x',
	'y', 'z', 'A', 'B',
	'C', 'D', 'E', 'F',
	'G', 'H', 'I', 'J',
	'K', 'L', 'M', 'N',
	'O', 'P', 'Q', 'R',
	'S', 'T', 'U', 'V',
	'W', 'X', 'Y', 'Z',
	'0', '1', '2', '3',
	'4', '5', '6', '7',
	'8', '9', '-', '_',
}

type generator = func() string

/*
Creates a new generator for canonical Nano IDs.

📝 Recommended (standard) length is 21

Returns error if length is not between 2 and 255 (inclusive).

Concurrency safe.
*/
func Standard(length int) (generator, error) {
	if invalidLength(length) {
		return nil, errInvalidLength
	}

	// Multiplying to increase the 'buffer' so that .Read()
	// has to be called less, which is more efficient.
	// b holds the random crypto bytes.
	b := make([]byte, length*length*6)
	size := len(b)
	offset := 0

	crand.Read(b)

	// Since the default alphabet is ASCII,
	// we don't have to use runes here. ASCII max is
	// 128, so byte will be perfect.
	// id := make([]rune, length)
	id := make([]byte, length)

	var mu sync.Mutex

	return func() string {
		mu.Lock()
		defer mu.Unlock()

		// If all the bytes in the slice
		// have been used, refill.
		if offset == size {
			crand.Read(b)
			offset = 0
		}

		for i := 0; i < length; i++ {
			/*
				"It is incorrect to use bytes exceeding the alphabet size.
				The following mask reduces the random byte in the 0-255 value
				range to the 0-63 value range. Therefore, adding hacks such
				as empty string fallback or magic numbers is unneccessary because
				the bitmask trims bytes down to the alphabet size (64).""
			*/
			// Index using the offset.
			id[i] = defaultAlphabet[b[i+offset]&63]
		}

		// Extend the offset.
		offset += length

		return string(id)
	}, nil
}

/*
Create a non-secure Nano ID generator.
Non-secure is faster than secure because it uses pseudorandom numbers.

Returns error if length is not between 2 and 255 (inclusive).

⚠ Remember to seed using rand.Seed().

Concurrency safe.
*/
func StandardNonSecure(length int) (generator, error) {
	if invalidLength(length) {
		return nil, errInvalidLength
	}

	// b holds pseudorandom bytes.
	b := make([]byte, length*length*6)
	size := len(b)
	offset := 0

	/*
		"Read generates len(p) random bytes from the default Source and
		writes them into p... It always returns len(p) and a **nil error**."
	*/
	mrand.Read(b)

	// Reuse.
	id := make([]byte, length)

	var mu sync.Mutex

	return func() string {
		mu.Lock()
		defer mu.Unlock()

		if offset == size {
			// Refill b.
			mrand.Read(b)
			offset = 0
		}

		for i := 0; i < length; i++ {
			id[i] = defaultAlphabet[prng(alphabetSize)]
		}

		// offset += length

		return string(id)
	}, nil
}

/*
Create a Nano ID generator that uses a custom alphabet.

Concurrency safe.
*/
func Custom(alphabet string, length int) (generator, error) {
	if invalidLength(length) {
		return nil, errInvalidLength
	}

	setLen := len(alphabet)
	// Have to use runes because input is
	// not guaranteed to be ASCII.
	runicSet := []rune(alphabet)

	// Because the custom alphabet is not guaranteed to have
	// 64 chars to utilise, we have to calculate a suitable mask.
	// This code below is 1:1 to the og impl.
	clz := bits.LeadingZeros32((uint32(setLen) - 1) | 1)
	mask := (2 << (31 - clz)) - 1
	w := (1.6 * float64(mask*length)) / float64(setLen)
	step := int(math.Ceil(w))

	// Will be reusing the same rune and byte slices.
	id := make([]rune, length)
	b := make([]byte, step*step*6)

	var mu sync.Mutex

	return func() string {
		mu.Lock()
		defer mu.Unlock()

		for u := 0; ; {
			crand.Read(b)

			for i := 0; i < step; i++ {
				idx := b[i] & byte(mask)

				if idx < byte(setLen) {
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

/*
Create a non-secure Nano ID generator that uses a custom alphabet.
Non-secure is faster than secure because it uses pseudorandom numbers.

Returns error if length is not between 2 and 255 (inclusive).

⚠️ Remember to seed using rand.Seed().

Concurrency safe.
*/
func CustomNonSecure(alphabet string, length int) (generator, error) {
	if invalidLength(length) {
		return nil, errInvalidLength
	}

	runicSet := []rune(alphabet)
	setLen := len(runicSet)

	id := make([]rune, length)

	return func() string {
		for i := 0; i < length; i++ {
			id[i] = runicSet[prng(setLen)]
		}

		return string(id)
	}, nil
}

var prngmu sync.Mutex
var tainer uint64

// This is slow but still better than math.Intn()
func prng(n int) int {
	prngmu.Lock()

	// xorshift.
	tainer ^= tainer << 13
	tainer ^= tainer >> 7
	tainer ^= tainer << 17

	prngmu.Unlock()

	k := int(tainer)

	if k < 0 {
		k = -k
	}

	return k % n
}

var errInvalidLength = errors.New("length for ID is too large or small")

func invalidLength(length int) bool {
	return length < 2 || length > 255
}

func init() {
	b := make([]byte, 8)
	crand.Read(b)
	tainer = binary.BigEndian.Uint64(b)
}
