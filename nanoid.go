/*
Copyright 2022 jaevor.
License can be found in the LICENSE file.

Original reference: https://github.com/ai/nanoid
*/

package nanoid

import (
	crand "crypto/rand"
	"errors"
	"math"
	"math/bits"
	"sync"
)

type generator = func() string

const defaultAlphabetSize = 64

// Default characters (A-Za-z0-9_-).
// Using less memory with [64]byte{...} than []byte(...).
var defaultAlphabet = [defaultAlphabetSize]byte{
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

/*
Creates a new generator for standard Nano IDs.

📝 Recommended (standard) length is 21

⛔ Returns error if length is not, or within 2 and 255.

🧿 Concurrency safe.
*/
func Standard(length int) (generator, error) {
	if invalidLength(length) {
		return nil, errInvalidLength
	}

	// Multiplying to increase the 'buffer' so that .Read()
	// has to be called less, which is more efficient.
	// b holds the random crypto bytes.
	size := length * length * 7
	b := make([]byte, size)
	offset := 0

	crand.Read(b)

	// Since the default alphabet is ASCII, we don't have to use runes.
	// ASCII max is 128, so byte will be perfect.
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
Create a Nano ID generator that uses a custom alphabet.
The alphabet is allowed to contain non-ASCII.

⛔ Returns error if length is not, or within 2 and 255.

🧿 Concurrency safe.
*/
func Custom(alphabet string, length int) (generator, error) {
	if invalidLength(length) {
		return nil, errInvalidLength
	}

	alphabetLen := len(alphabet)
	// Have to use runes because input is
	// not guaranteed to be ASCII.
	runes := []rune(alphabet)

	// Because the custom alphabet is not guaranteed to have
	// 64 chars to utilise, we have to calculate a suitable mask.
	x := uint32(alphabetLen) - 1
	clz := bits.LeadingZeros32(x | 1)
	mask := (2 << (31 - clz)) - 1
	step := int(math.Ceil((1.6 * float64(mask*length)) / float64(alphabetLen)))

	size := step * step * 7
	b := make([]byte, size)

	id := make([]rune, length)

	var idx int
	j := 0

	var mu sync.Mutex

	return func() string {
		mu.Lock()
		defer mu.Unlock()

		for {
			crand.Read(b)

			for i := 0; i < step; i++ {
				idx = int(b[i]) & mask
				if idx < alphabetLen {
					id[j] = runes[idx]
					j++
					if j == length {
						j = 0
						return string(id)
					}
				}
			}
		}
	}, nil
}

var errInvalidLength = errors.New("length for ID is invalid (must be within 2-255)")

func invalidLength(length int) bool {
	return length < 2 || length > 255
}
