package goptr

// Bool returns a pointer to a boolean.
func Bool(b bool) *bool {
	return &b
}

// Float32 returns a pointer to a float.
func Float32(f float32) *float32 {
	return &f
}

// Float64 returns a pointer to a float.
func Float64(f float64) *float64 {
	return &f
}

// Int returns a pointer to an int.
func Int(i int) *int {
	return &i
}

// Int32 returns a pointer to an int.
func Int32(i int32) *int32 {
	return &i
}

// Int64 returns a pointer to an int.
func Int64(i int64) *int64 {
	return &i
}

// String returns a pointer to a string.
func String(s string) *string {
	return &s
}
