# **go-nanoid**

[![Build Status](https://github.com/jaevor/go-nanoid/workflows/tests/badge.svg)](https://github.com/jaevor/go-nanoid/actions)
[![Build Status](https://github.com/jaevor/go-nanoid/workflows/lint/badge.svg)](https://github.com/jaevor/go-nanoid/actions)
[![GitHub Issues](https://img.shields.io/github/issues/jaevor/go-nanoid.svg)](https://github.com/jaevor/go-nanoid/issues)
[![Go Ref](https://pkg.go.dev/badge/github.com/jaevor/go-nanoid)](https://pkg.go.dev/github.com/jaevor/go-nanoid)

This module is a Go implementation of [nanoid](https://github.com/ai/nanoid).

Features of the nanoid spec are:
> Use of hardware random generator

> Uses a larger charset (`A-Za-z0-9_-`) than UUID

> Faster than UUID

> URL friendly

Features of this specific implementation are:
> Fastest and most performant implementation of Nano ID ([benchmarks](#benchmarks))

> Prefetches random bytes in advance

> Uses optimal memory

> No production dependencies

*See [comparison of Nano ID and UUID (V4)](https://github.com/ai/nanoid/blob/main/README.md#comparison-with-uuid)*

**[NanoID collison calculator](https://zelark.github.io/nano-id-cc/)**

**Read more [here](https://github.com/ai/nanoid/blob/main/README.md)**

---

## Example

```go
import (
	"log"
	"math/rand"
	"time"
	"github.com/jaevor/go-nanoid"
)

func main() {
  createNanoid, err := nanoid.Standard(21)
  if err != nil {
    panic(err)
  }

  id1 := createNanoid()
  log.Printf("ID 1: %s", id1) // p7aXQf7xK3jlfecYGKeRK

  // [!] Remember to seed. 
  rand.Seed(time.Now().Unix())

  createNonSecureNanoid, err := nanoid.StandardNonSecure(21)
  if err != nil {
    panic(err)
  }

  id2 := createNonSecureNanoid()
  log.Printf("ID 2: %s", id2) // japKZqwnQvllgUQ8lwgkP

  createCustomNanoid, err := nanoid.Custom("0123456789", 12)
  if err != nil {
    panic(err)
  }

  id3 := f3()
  log.Printf("ID 3: %s", id3) // 462855288020
}

```
---

## Notes
Developed in Go 1.18.3

Remember to `rand.Seed(...)` before using the non-secure generators.

**The generation of non-secure Nano IDs are not as fast as they could be yet**

---

## Benchmarks
All benchmarks & tests can be found in [nanoid_test.go](./nanoid_test.go).

These are all benchmarks of the `Standard` Nano ID generator

| # of characters & # of IDs | benchmark screenshot |
| -------------------------- | ---------- |
| 8, ~19,120,000             | <img src="img/benchmark-8.png">   |
| 21, ~14,500,000            | <img src="img/benchmark-21.png">  |
| 36, ~11,000,000            | <img src="img/benchmark-36.png">  |
| 255, ~2,300,000            | <img src="img/benchmark-255.png"> |

---

## Credits & references
- [Original reference](https://github.com/ai/nanoid)
- [Outdated (by 2+ years) Go implementation](https://github.com/matoous/go-nanoid)

---

## License
[MIT License](./LICENSE)
