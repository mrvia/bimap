# bimap
![testing](https://github.com/mrvia/bimap/actions/workflows/test.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/mrvia/bimap.svg)](https://pkg.go.dev/github.com/mrvia/bimap)

Bi-directional map for Go.

## Installation

Use go get to install the package:

```bash
go get github.com/mrvia/bimap
```

## Usage

```go
import "github.com/mrvia/bimap"

// Create a new bi-directional map
bm := bimap.New(map[string]int{
    "key1": 1,
    "key2": 2,
})

// Add key-value pairs
bm.Put("key3", 3)

// Get value by key
value, ok := bm.GetByKey("key1") // value == 1, ok == true

// Get key by value
key, ok := bm.GetByValue(2) // key == "key2", ok == true
```

## License

This project is licensed under the [MIT License](LICENSE).
