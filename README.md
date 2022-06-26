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
> Fastest, most performant implementation of Nano ID

> No production dependencies

*See [comparison with UUID](https://github.com/ai/nanoid/blob/main/README.md#comparison-with-uuid)*

**[NanoID collison calculator](https://zelark.github.io/nano-id-cc/)**

**Read more [here](https://github.com/ai/nanoid/blob/main/README.md)**


## Example

```go
import (
	"log"
	"math/rand"
	"time"
	"github.com/jaevor/go-nanoid"
)

func main() {
  createNanoid, err := nanoid.New(21)
  if err != nil {
    panic(err)
  }

  id1 := createNanoid()
  id2 := createNanoid()

  log.Printf("ID 1: %s", id1) // p7aXQf7xK3jlfecYGKeRK
  log.Printf("ID 2: %s", id2) // WTGJ1VPh472R_4h--UNL0

  // [!] Remember to seed. 
  rand.Seed(time.Now().Unix())

  createNonSecureNanoid, err := nanoid.NewNonSecure(21)
  if err != nil {
    panic(err)
  }

  id3 := createNonSecureNanoid()
  id4 := createNonSecureNanoid()

  log.Printf("ID 3: %s", id3) // japKZqwnQvllgUQ8lwgkP
  log.Printf("ID 4: %s", id4) // 2jo3VdWZ2LbTB79TKB9je
}

```

---

## Benchmarks
All benchmarks & tests can be found in [nanoid_test.go](./nanoid_test.go).

**14,500,000** Nano IDs in **1300ms** @ `82.5 ns/op`**,** `24 B/op`**,** `1 alloc/op`.
![benchmark](./img/benchmark.png)

---

## Credits & references
- [Original reference](https://github.com/ai/nanoid)
- [Outdated Go implementation (26/06/22)](https://github.com/matoous/go-nanoid)

---

## License
[MIT LICENSE](./LICENSE)
