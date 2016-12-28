# goptr

The `goptr` package gives you some basic helpers for working with pointers in
Go.

## Usage

    goptr.Bool(true)        // Returns *bool
    goptr.Byte(byte('a'))   // Returns *byte
    goptr.Float32(123.3)    // Returns *float32
    goptr.Float64(123.3)    // Returns *float64
    goptr.Int(123)          // Returns *int
    goptr.Int8(123)         // Returns *int8
    goptr.Int16(123)        // Returns *int16
    goptr.Int32(123)        // Returns *int32
    goptr.Int64(123)        // Returns *int64
    goptr.Rune(123)         // Returns *rune
    goptr.String("string")  // Returns *string
    goptr.Uint(123)         // Returns *uint
    goptr.Uint8(123)        // Returns *uint8
    goptr.Uint16(123)       // Returns *uint16
    goptr.Uint32(123)       // Returns *uint32
    goptr.Uint64(123)       // Returns *uint64

## Running the tests

    make test

## License

MIT
