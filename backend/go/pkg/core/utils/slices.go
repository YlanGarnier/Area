package utils

func Contains[T any](slice []T, f func(T) bool) bool {
	for _, e := range slice {
		if f(e) {
			return true
		}
	}
	return false
}

func Diff[T any](A []T, B []T, diff func(T, T) bool) []T {
	var res []T
	for _, a := range A {
		in := false
		for _, b := range B {
			if !diff(a, b) {
				in = true
				break
			}
		}
		if !in {
			res = append(res, a)
		}
	}
	return res
}
