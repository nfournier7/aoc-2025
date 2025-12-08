package number

import "strconv"

type Uint interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64
}

func ParseUint[T Uint](str string) T {
	var zero T
	var bitSize int
	switch any(zero).(type) {
	case uint8:
		bitSize = 8
	case uint16:
		bitSize = 16
	case uint32:
		bitSize = 32
	case uint64:
		bitSize = 64
	}

	v, err := strconv.ParseUint(str, 10, bitSize)
	if err != nil {
		panic(err)
	}

	return T(v)
}

type SInt interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
}

func ParseInt[T SInt](str string) T {
	var zero T
	var bitSize int
	switch any(zero).(type) {
	case int8:
		bitSize = 8
	case int16:
		bitSize = 16
	case int32:
		bitSize = 32
	case int64:
		bitSize = 64
	}

	v, err := strconv.ParseInt(str, 10, bitSize)
	if err != nil {
		panic(err)
	}

	return T(v)
}