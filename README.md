# **go-nanoid**

[![Build Status](https://github.com/jaevor/go-nanoid/workflows/test/badge.svg)](https://github.com/jaevor/go-nanoid/actions)
[![Build Status](https://github.com/jaevor/go-nanoid/workflows/lint/badge.svg)](https://github.com/jaevor/go-nanoid/actions)
[![GitHub Issues](https://img.shields.io/github/issues/jaevor/go-nanoid.svg)](https://github.com/jaevor/go-nanoid/issues)
[![Go Ref](https://pkg.go.dev/badge/github.com/jaevor/go-nanoid)](https://pkg.go.dev/github.com/jaevor/go-nanoid)

This module is a Go implementation of [nanoid](https://github.com/ai/nanoid).

Features of the nanoid spec are:
> Use of hardware random generator

> Has larger alphabet than UUID (`A-Za-z0-9_-`), so ID size is reduced from 36 to 21 characters

> Faster than UUID


*See [comparison with UUID](https://github.com/ai/nanoid/blob/main/README.md#comparison-with-uuid)*


**Read more [here](https://github.com/ai/nanoid/blob/main/README.md)**


## Example

```go
import (
	"log"
	"github.com/jaevor/go-nanoid"
)

func main() {
  createNanoid, err := nanoid.New(21)
  if err != nil {
    panic(err)
  }

  id1 := createNanoid()
  id2 := createNanoid()

  log.Printf("ID 1: %s", id1)
  // ID 1: p7aXQf7xK3jlfecYGKeRK
  log.Printf("ID 2: %s", id2)
  // ID 2: WTGJ1VPh472R_4h--UNL0
}

```

---

## Benchmarks
All benchmarks & tests can be found in [nanoid_test.go](./nanoid_test.go).

6,000,000 Nano IDs in 1500ms.
![benchmark](./img/benchmark.png)

---

## Credits & references
- [Original reference](https://github.com/ai/nanoid)
- [Outdated Go implementation (26/06/22)](https://github.com/matoous/go-nanoid)

---

## License
[MIT LICENSE](./LICENSE)
