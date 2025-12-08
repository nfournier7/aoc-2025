package cond

func If[T any](cond bool, valueTrue, valueFalse T) T {
	if cond {
		return valueTrue
	}
	return valueFalse
}