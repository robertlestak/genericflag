# genericflag

The stdlib [flag](https://pkg.go.dev/flag) package is great, but only if you know what flags you need to parse ahead of time. If you want to parse flags that you don't know about until runtime, you're out of luck.

`genericflag` provides a way to method to parse arbitrary flags not defined at compile time.

This is useful if writing wrapper utilities around other programs that take flags, and you need to consume / modify a subset of defined flags, and pass the rest through to the underlying program.

## Usage

The `genericflag` package exposes a `FlagSet` struct. This can be initialized with `NewFlagSet`, which takes a name for the flag set. The `FlagSet` struct has a `Parse` method, which takes a slice of strings to parse as flags.

Once parsed, the `FlagSet` struct has a `Flags` field, which is a map of strings to slices of strings. The keys of the map are the flag names, and the values are the values of the flags. If a flag is passed without a value, the value will be an empty slice.

The `FlagSet` also has an optional `Expected` field, which is a slice of strings. If this field is set, then the `Parse` method will only parse the flags with the names defined in the `Expected` field, and all remaining flags will be available through the `Args` function.

## Example

### Basic Parsing of All Flags

```go
package main

import (
	"os"

	"github.com/robertlestak/genericflag"
)

func main() {
	fs := genericflag.NewFlagSet("example")
	if err := fs.Parse(os.Args[1:]); err != nil {
		panic(err)
	}
	for k, v := range fs.Flags {
		for _, val := range v {
			println(k, val)
		}
	}
}
```

```bash
$ go run main.go -foo bar -baz qux
foo bar
baz qux
```

### Parsing Only Expected Flags

```go
package main

import (
	"os"

	"github.com/robertlestak/genericflag"
)

func main() {
	fs := genericflag.NewFlagSet("example")
    fs.Expected = []string{"foo", "bar"}
	if err := fs.Parse(os.Args[1:]); err != nil {
		panic(err)
	}
    println("parsed flags:")
	for k, v := range fs.Flags {
		for _, val := range v {
			println(k, val)
		}
	}
    println("unparsed args:")
    // now print out the remaining args
    for _, arg := range fs.Args() {
        println(arg)
    }
}
```

```bash
$ go run main.go -foo bar -bar foo -baz qux
parsed flags:
foo bar
bar foo
unparsed args:
-baz
qux
```
