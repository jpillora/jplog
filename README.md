# jplog

my `slog` handler

:warning: this API may change

### Usage

```go
log := jplog.New(os.Stdout)
// which is shorthand for:
log := slog.New(jplog.Handler(os.Stdout))
```

### Output

```go
package main

import "github.com/jpillora/jplog"

func main() {
	log := jplog.New(os.Stdout).WithGroup("myapp")
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