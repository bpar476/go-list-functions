package listfns

// Map takes a slice and uses the mapper function to transform it
// into a new slice of the same length with each value being
// the corresponding value in the source slice, transformed by
// the mapper function
func Map[T any, U any](source []T, mapper func(T) U) []U {
	out := make([]U, len(source))

	for i, val := range source {
		out[i] = mapper(val)
	}

	return out
}

// Fold takes a slice and iterates over every element to fold
// the slice into a different value. The folder function is passed
// an accumulator and the value of the current iteration. The
// accumulator starts with the given initial value.
func Fold[T any, U any](source []T, initial U, folder func(acc U, current T) U) U {
	result := initial

	for _, val := range source {
		result = folder(result, val)
	}

	return result
}

// Reduce is a special kind of fold that folds the given slice
// into a single value of the same type
func Reduce[T any](source []T, reducer func(acc, current T) T) T {
	var initial T

	return Fold(source, initial, reducer)
}

// Contains checks if a given target element exists in a given slice
func Contains[T comparable](source []T, target T) bool {
	for _, v := range source {
		if v == target {
			return true
		}
	}

	return false
}
