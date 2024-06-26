package funcutils

import (
	"slices"
)

func MapFunc[S ~[]E, E any, F any](s S, f func(E) F) []F {
	r := make([]F, len(s), len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

func IndexFunc[S ~[]E, E any](s S, f func(E) bool) (int, bool) {
	r := slices.IndexFunc(s, f)
	if r == -1 {
		return 0, false // that's the go idiom, right? embrace the zero value, if ..; ok { etc.
	}
	return r, true
}

func GroupByFunc[S ~[]E, E any, F comparable](s S, f func(E) F) map[F][]E {
	r := make(map[F][]E)
	for _, e := range s {
		fe := f(e)
		r[fe] = append(r[fe], e)
	}
	return r
}

func FindFunc[S ~[]E, E any](s S, f func(E) bool) (E, bool) {
	i := slices.IndexFunc(s, f)
	var zero E
	if i == -1 {
		return zero, false
	}
	return s[i], true
}

func RemoveIf[S ~[]E, E any](s S, f func(E) bool) []E {
	var r []E
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}

func RemoveIfNot[S ~[]E, E any](s S, f func(E) bool) []E {
	fn := func(e E) bool {
		return f(e)
	}
	return RemoveIf(s, fn)
}

func Identity[E any](e E) E {
	return e
}
