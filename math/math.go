package amath

func Abs(value int32) int32 {
	if value < 0 {
		value *= -1
	}
	return value
}

func IsNegative(value int32) bool {
	return value < 0
}

func IsPositive(value int32) bool {
	return value > 0
}

func Max(a int32, b int32) int32 {
	if b > a {
		return b
	}
	return a
}

func Min(a int32, b int32) int32 {
	if b < a {
		return b
	}
	return a
}
