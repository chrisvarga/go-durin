# Durin client for Go
go-durin provides a simple interface for interacting with the durin database in Go.

## Installation
To install with go get, do:
```shell
go get github.com/chrisvarga/go-durin
```

## Quickstart
```go
import (
    "github.com/chrisvarga/go-durin"
    "fmt"
)

func ExampleClient() {
    val, err := durin.Durin("set key value")
    if err != nil {
        panic(err)
    }
    val, err = durin.Durin("get key")
    if err != nil {
        panic(err)
    }
    fmt.Println("key", val)
    // Output: key value
}
```

