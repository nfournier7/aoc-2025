package number

import "strconv"

func ParseUint32(strValue string, base int, bitSize int) uint32 {
	value, err := strconv.ParseUint(strValue, base, bitSize)
	if err != nil {
		panic(err)
	}
	return uint32(value)
}

func ParseUint64(strValue string, base int, bitSize int) uint64 {
	value, err := strconv.ParseUint(strValue, base, bitSize)
	if err != nil {
		panic(err)
	}
	return value
}

func ParseInt32(valueStr string) int32 {
	value, err := strconv.ParseInt(valueStr, 10, 0)
	if err != nil {
		panic(err)
	}
	return int32(value)
}