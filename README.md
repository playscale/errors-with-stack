# Errors With Stack for Go

This is an extremely simple, lightweight, drop-in replacement that
shadows the go `errors` package. If this package is in scope when
`errors.New()` is called, it will return a standard Go Error type
with the error message payload including a stack trace.

Printing stack traces can be toggled at runtime by invoking
`errors.EnableStackTraces()` and `errors.DisableStackTraces()`
respectively. It is safe to call this from separate threads.

## Usage

```go
package main

import(
    "fmt"
    "github.com/Justin-Randall/errors-with-stack"
)

func exampleErrorsWithStack() {
    errorWithStack := errors.New("This error should contain a stack Trace")
    fmt.Println(errorWithStack.Error())

    errors.DisableStackTraces()

    plainOldGoError := errors.New("This is just a plain old golang error")
    fmt.Println(plainOldGoError.Error())
}

func main() {
    exampleErrorsWithStack()
}
```

When run, you should see:

```shell
$ go run main.go

This error should contain a stack Trace
Stack Trace:
goroutine 1 [running]:
runtime/debug.Stack(0x0, 0x0, 0x44)
        /go/src/runtime/debug/stack.go:24 +0xa4
github.com/Justin-Randall/errors-with-stack.New(0x4d4927, 0x27, 0xc0000c7f48, 0x4410de)
        /User/github.com/Justin-Randall/errors-with-stack/errorswithstack.go:22 +0x7b
main.exampleErrorsWithStack()
        /User/github.com/Justin-Randall/errors-with-stack/example/main.go:10 +0x41
main.main()
        /User/github.com/Justin-Randall/errors-with-stack/example/main.go:20 +0x27

This is just a plain old golang error
```

## License

This package is released under an MIT license. See the LICENSE file
included in this repo.
