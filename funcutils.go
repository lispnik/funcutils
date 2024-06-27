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
		// that's the go idiom, right? embrace the zero value, if ..; ok { etc.
		return 0, false
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

func GroupBy[S ~[]E, E comparable](s S) map[E][]E {
	return GroupByFunc(s, Identity[E])
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

// not in Lisp, fuck it

// DifferenceFunc f(s1) - f(s2)
func DifferenceFunc[S ~[]E, E any, F comparable](s1, s2 S, f func(E) F) []E {
	m := make(map[F]bool)
	for _, e := range s2 {
		m[f(e)] = true
	}
	var r []E
	for _, e := range s1 {
		if !m[f(e)] {
			r = append(r, e)
		}
	}
	return r
}

// SymmetricDifferenceFunc f(s1) - f(s2) | f(s2) - f(s1)
func SymmetricDifferenceFunc[S ~[]E, E any, F comparable](s1, s2 S, f func(E) F) []E {
	return append(DifferenceFunc(s1, s2, f), DifferenceFunc(s2, s1, f)...)
}

// for the love of all creatures great and small

func Difference[S ~[]E, E comparable](s1, s2 S) []E {
	return DifferenceFunc(s1, s2, Identity[E])
}

func SymmetricDifference[S ~[]E, E comparable](s1, s2 S) []E {
	return SymmetricDifferenceFunc(s1, s2, Identity[E])
}

func AdjoinFunc[S ~[]E, E any](s S, e E, f func(E) bool) []E {
	if ok := slices.ContainsFunc(s, f); !ok {
		return append(s, e)
	}
	return s
}

func MemberFunc[S ~[]E, E comparable](s S, e E) bool {
	return slices.Contains(s, e)
}

// TODO MemberIfFunc, MemberIfNotFunc

// TODO do big/small slice test for great efficiency:

func UnionFunc[S ~[]E, E any, F comparable](s1, s2 S, f func(E) F) []E {
	m := make(map[F]bool)
	var r []E
	for _, e := range s1 {
		fe := f(e)
		if !m[fe] {
			r = append(r, e)
			m[fe] = true
		}
	}
	for _, e := range s2 {
		if !m[f(e)] {
			r = append(r, e)
		}
	}
	return r
}

func Union[S ~[]E, E comparable](s1, s2 S) []E {
	return UnionFunc(s1, s2, Identity[E])
}

func IntersectionFunc[S ~[]E, E any, F comparable](s1, s2 S, f func(E) F) []E {
	m := make(map[F]bool)
	var r []E
	for _, e := range s1 {
		m[f(e)] = true
	}
	for _, e := range s2 {
		if m[f(e)] {
			r = append(r, e)
		}
	}
	return r
}

func Intersection[S ~[]E, E comparable](s1, s2 S) []E {
	return IntersectionFunc(s1, s2, Identity[E])
}
