// Tests & benchmarks
package nanoid_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/jaevor/go-nanoid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	// This is so that math/rand can work correctly.
	// Only needed for non-secure. Users will have to do this themselves.
	rand.Seed(time.Now().Unix())
}

func TestStandard(t *testing.T) {
	t.Run("general", func(t *testing.T) {
		f, err := nanoid.Standard(21)
		assert.NoError(t, err, "should be no error")
		id := f()
		assert.Len(t, id, 21, "should return the same length as the ID specified length")
		t.Log(id)
	})

	t.Run("negative ID length", func(t *testing.T) {
		_, err := nanoid.Standard(-1)
		assert.Error(t, err, "should error if passed ID length is negative")
	})

	t.Run("invalid length (256)", func(t *testing.T) {
		_, err := nanoid.Standard(256)
		assert.Error(t, err, "should error if length > 255")
	})

	t.Run("invalid length (1)", func(t *testing.T) {
		_, err := nanoid.Standard(1)
		assert.Error(t, err, "should error if length < 2")
	})
}

func TestNonSecure(t *testing.T) {
	t.Run("general", func(t *testing.T) {
		f, err := nanoid.StandardNonSecure(21)
		assert.NoError(t, err, "should be no error")
		id := f()
		assert.Len(t, id, 21, "should return the same length as the ID specified length")
		t.Log(id)
	})
}

func TestCustom(t *testing.T) {
	t.Run("general", func(t *testing.T) {
		f, err := nanoid.Custom("abcdef", 21)
		id := f()
		assert.NoError(t, err, "should be no error")
		assert.Len(t, id, 21, "should return the same length as the ID specified length")
		t.Log(id)
	})
}

func TestFlatDistribution(t *testing.T) {
	tries := 500_000

	set := "0123456789" // 10.
	length := len(set)
	hits := make(map[rune]int)

	f, err := nanoid.Custom(set, length)
	if err != nil {
		panic(err)
	}

	for i := 0; i < tries; i++ {
		id := f()
		for _, r := range id {
			hits[r]++
		}
	}

	for _, count := range hits {
		require.InEpsilon(t, length*tries/len(set), count, 0.01, "should have flat-distribution")
	}
}

func TestCollisions(t *testing.T) {
	tries := 500_000

	used := make(map[string]bool)
	f, err := nanoid.Standard(8)
	if err != nil {
		panic(err)
	}

	for i := 0; i < tries; i++ {
		id := f()
		require.False(t, used[id], "shouldn't be any colliding IDs")
		used[id] = true
	}
}

func Benchmark8NanoID(b *testing.B) {
	f, err := nanoid.Standard(8)
	if err != nil {
		panic(err)
	}

	for n := 0; n < b.N; n++ {
		f()
	}
}

func Benchmark21NanoID(b *testing.B) {
	f, err := nanoid.Standard(21)
	if err != nil {
		panic(err)
	}

	for n := 0; n < b.N; n++ {
		f()
	}
}

func Benchmark36NanoID(b *testing.B) {
	f, err := nanoid.Standard(36)
	if err != nil {
		panic(err)
	}

	for n := 0; n < b.N; n++ {
		f()
	}
}

func Benchmark255NanoID(b *testing.B) {
	f, err := nanoid.Standard(255)
	if err != nil {
		panic(err)
	}

	for n := 0; n < b.N; n++ {
		f()
	}
}

func BenchmarkNonSecureNanoID(b *testing.B) {
	f, err := nanoid.StandardNonSecure(21)
	if err != nil {
		panic(err)
	}

	for n := 0; n < b.N; n++ {
		f()
	}
}
