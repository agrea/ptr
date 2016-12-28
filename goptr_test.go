package goptr

import "testing"

func TestBool(t *testing.T) {
	var b bool = true

	if *Bool(b) != b {
		t.Error("did not get a pointer back")
	}
}

func TestByte(t *testing.T) {
	b := []byte{'a'}

	if *Byte(b[0]) != b[0] {
		t.Error("did not get a pointer back")
	}
}

func TestFloat32(t *testing.T) {
	var f float32 = 123.12

	if *Float32(f) != f {
		t.Error("did not get a pointer back")
	}
}

func TestFloat64(t *testing.T) {
	var f float64 = 123.12

	if *Float64(f) != f {
		t.Error("did not get a pointer back")
	}
}

func TestInt(t *testing.T) {
	var i int = 123

	if *Int(i) != i {
		t.Error("did not get a pointer back")
	}
}

func TestInt8(t *testing.T) {
	var i int8 = 123

	if *Int8(i) != i {
		t.Error("did not get a pointer back")
	}
}

func TestInt16(t *testing.T) {
	var i int16 = 123

	if *Int16(i) != i {
		t.Error("did not get a pointer back")
	}
}

func TestInt32(t *testing.T) {
	var i int32 = 123

	if *Int32(i) != i {
		t.Error("did not get a pointer back")
	}
}

func TestInt64(t *testing.T) {
	var i int64 = 123

	if *Int64(i) != i {
		t.Error("did not get a pointer back")
	}
}

func TestUint(t *testing.T) {
	var i uint = 123

	if *Uint(i) != i {
		t.Error("did not get a pointer back")
	}
}

func TestUint8(t *testing.T) {
	var i uint8 = 123

	if *Uint8(i) != i {
		t.Error("did not get a pointer back")
	}
}

func TestUint16(t *testing.T) {
	var i uint16 = 123

	if *Uint16(i) != i {
		t.Error("did not get a pointer back")
	}
}

func TestUint32(t *testing.T) {
	var i uint32 = 123

	if *Uint32(i) != i {
		t.Error("did not get a pointer back")
	}
}

func TestUint64(t *testing.T) {
	var i uint64 = 123

	if *Uint64(i) != i {
		t.Error("did not get a pointer back")
	}
}

func TestRune(t *testing.T) {
	var r rune = 1

	if *Rune(r) != r {
		t.Error("did not get a pointer back")
	}
}

func TestString(t *testing.T) {
	var i string = "What was the question to life, universe and everything?"

	if *String(i) != i {
		t.Error("did not get a pointer back")
	}
}
