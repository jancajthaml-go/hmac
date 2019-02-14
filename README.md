## HMAC - SHA 256

[![Go Report Card](https://goreportcard.com/badge/jancajthaml-go/hmac256)](https://goreportcard.com/report/jancajthaml-go/hmac256)

### Usage ###

```
import "github.com/jancajthaml-go/hmac256"

hmac256.Digest([]byte("XXXX"), []byte("xxxx"))
```

### Performance ###

```
BenchmarkSmall-4          1000000  1348 ns/op  480 B/op  6 allocs/op
BenchmarkLarge-4          1000000  1586 ns/op  480 B/op  6 allocs/op
BenchmarkSmallParallel-4  2000000   784 ns/op  480 B/op  6 allocs/op
BenchmarkLargeParallel-4  2000000   805 ns/op  480 B/op  6 allocs/op
```

verify your performance by running `make benchmark`

## Resources

* [Wikipedia - HMAC algorithm](https://en.wikipedia.org/wiki/HMAC)
