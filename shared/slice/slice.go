package slice

import (
	"fmt"
	"strings"
)

func DeleteAtIndex[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func PrintArray(slice []string) {
	var builder strings.Builder

	builder.WriteString("[")

	builder.WriteString(strings.Join(slice, ", "))

	builder.WriteString("]")

	fmt.Println(builder.String())
}

func FindIndex[T any](slice []T, pred func(e T) bool) int {
	for i := range slice {
		value := slice[i]
		if pred(value) {
			return i
		}
	}
	return int(-1)
}