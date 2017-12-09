# ptr

The `ptr` package gives you some basic helpers for working with pointers in Go.

## Usage

    ptr.Bool(true)        // Returns *bool
    ptr.Byte(byte('a'))   // Returns *byte
    ptr.Float32(123.3)    // Returns *float32
    ptr.Float64(123.3)    // Returns *float64
    ptr.Int(123)          // Returns *int
    ptr.Int8(123)         // Returns *int8
    ptr.Int16(123)        // Returns *int16
    ptr.Int32(123)        // Returns *int32
    ptr.Int64(123)        // Returns *int64
    ptr.Rune(123)         // Returns *rune
    ptr.String("string")  // Returns *string
    ptr.Uint(123)         // Returns *uint
    ptr.Uint8(123)        // Returns *uint8
    ptr.Uint16(123)       // Returns *uint16
    ptr.Uint32(123)       // Returns *uint32
    ptr.Uint64(123)       // Returns *uint64

## Running the tests

    go test -v ./...

## License

MIT
