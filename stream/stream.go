package stream

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

func All[T any](ts []T, f func(T) bool) bool {
	for i := range ts {
		if !f(ts[i]) {
			return false
		}
	}
	return true
}

func ForEach[T any](ts []T, f func(T)) {
	for i := range ts {
		f(ts[i])
	}
}

func Count[T any](ts []T, f func(T) bool) int32 {
	var count int32 = 0
	for i := range ts {
		if f(ts[i]) {
			count++
		}
	}
	return count
}

func SumFn[T any](ts []T, f func(T) uint64) uint64 {
	sum := uint64(0)
	for i := range ts {
		sum += f(ts[i])
	}
	return sum
}

func Filter[T any](ts []T, f func(T) bool) []T {
	fts := []T{}
	for i := range ts {
		if f(ts[i]) {
			fts = append(fts, ts[i])
		}
	}
	return fts
}