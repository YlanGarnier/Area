package utils

func Map[E any, R any](s []E, f func(E) R) []R {
	res := make([]R, len(s))
	for i, v := range s {
		res[i] = f(v)
	}
	return res
}

func MapErr[E any, R any](s []E, f func(E) (R, error)) ([]R, error) {
	res := make([]R, len(s))
	for i, v := range s {
		r, err := f(v)
		if err != nil {
			return nil, err
		}
		res[i] = r
	}
	return res, nil
}
