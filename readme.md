# Durin client for Go
go-durin provides a simple interface for interacting with the durin database in Go.

## Installation
To install with go get, do:
```shell
go get github.com/chrisvarga/go-durin
```

## Quick start
```go
import (
    "github.com/chrisvarga/go-durin"
    "fmt"
)

func ExampleClient() {
    // Initialize reusable connection.
    ctx, err := Init("localhost:8045")
    if err != nil {
        panic(err)
    }
    defer (*ctx).Close()

    val, err := durin.Set(ctx, "key", "value")
    if err != nil {
        panic(err)
    }
    val, err = durin.Get(ctx, "key")
    if err != nil {
        panic(err)
    }
    fmt.Println("key", val)
    // Output: key value
}
```

