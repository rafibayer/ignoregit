package util

func Dedup[T comparable](s []T) []T {
	m := map[T]struct{}{}
	for _, e := range s {
		m[e] = struct{}{}
	}

	r := make([]T, 0, len(m))
	for e := range m {
		r = append(r, e)
	}

	return r
}
