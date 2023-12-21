# jplog

my `slog` handler

### Usage

```go
log := jplog.New()
// which is shorthand for:
log := slog.New(jplog.Handler())
```

### Output

```go
package main

import "github.com/jpillora/jplog"

func main() {
	log := jplog.New().WithGroup("myapp")
	log.Info("a msg", "hello", "world", "foo", "bar")
	log.Warn("a msg", "hello", "world")
	log.Error("a msg", "hello", 42)
}
```

```
$ go run main.go
7:38:33PM 21/12/2023  INFO myapp a msg hello=world foo=bar
7:38:33PM 21/12/2023  WARN myapp a msg hello=world
7:38:33PM 21/12/2023 ERROR myapp a msg hello=42
# and it has colors
```