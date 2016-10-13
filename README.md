# goptr

The `goptr` package gives you some basic helpers for working with pointers in
Go.

## Usage

    goptr.Bool(true)        // Returns *bool
    goptr.Float32(123.3)    // Returns *float32
    goptr.Float64(123.3)    // Returns *float64
    goptr.Int(123)          // Returns *int
    goptr.Int32(123)        // Returns *int32
    goptr.Int64(123)        // Returns *int64
    goptr.String("string")  // Returns *string

## Running the tests

    make test

## License

MIT
