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

func TestNew(t *testing.T) {
	t.Run("general", func(t *testing.T) {
		f, err := nanoid.New(21)
		assert.NoError(t, err, "should be no error")
		id := f()
		assert.Len(t, id, 21, "should return the same length as the ID specified length")
		t.Log(id)
	})

	t.Run("negative ID length", func(t *testing.T) {
		_, err := nanoid.New(-1)
		assert.Error(t, err, "should error if passed ID length is negative")
	})

}

func TestNewNonSecure(t *testing.T) {
	t.Run("general", func(t *testing.T) {
		f, err := nanoid.NewNonSecure(21)
		assert.NoError(t, err, "should be no error")
		id := f()
		assert.Len(t, id, 21, "should return the same length as the ID specified length")
		t.Log(id)
	})
}

func TestNewCustom(t *testing.T) {
	t.Run("general", func(t *testing.T) {
		f, err := nanoid.NewCustom("abcdef", 21)
		id := f()
		assert.NoError(t, err, "should be no error")
		assert.Len(t, id, 21, "should return the same length as the ID specified length")
		t.Log(id)
	})
}

// Not always perfect?*
// Usually it is within 0.01, like 0.0157... or like 0.0112, the lowest I've gotten.
func TestFlatDistribution(t *testing.T) {
	tries := 10_000

	set := "0123456789" // 10.
	length := len(set)
	hits := make(map[rune]int)

	f, err := nanoid.NewCustom(set, length)
	if err != nil {
		panic(err)
	}

	for i := 0; i < tries; i++ {
		id := f()
		if err != nil {
			panic(err)
		}
		for _, r := range id {
			hits[r]++
		}
	}

	for _, count := range hits {
		require.InEpsilon(t, length*tries/len(set), count, 0.03 /* 0.01 */, "should have flat-distribution")
	}
}

func TestCollisions(t *testing.T) {
	tries := 50_000

	used := make(map[string]bool)
	f, err := nanoid.New(21)
	if err != nil {
		panic(err)
	}

	for i := 0; i < tries; i++ {
		id := f()
		require.False(t, used[id], "shouldn't be any colliding IDs")
		used[id] = true
	}
}

func BenchmarkNanoID(b *testing.B) {
	f, err := nanoid.New(21)
	if err != nil {
		panic(err)
	}

	for n := 0; n < b.N; n++ {
		f()
	}
}

func BenchmarkNonSecureNanoID(b *testing.B) {
	f, err := nanoid.NewNonSecure(21)
	if err != nil {
		panic(err)
	}

	for n := 0; n < b.N; n++ {
		f()
	}
}
