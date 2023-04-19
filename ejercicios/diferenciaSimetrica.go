package ejercicios

import (
	"guia4/set"
)

func DiferenciaSimetrica[T comparable](s1, s2 *set.Set[T]) *set.Set[T] {
	if s1 == nil && s2 != nil {
		return s2
	} else if s1 != nil && s2 == nil {
		return s1
	}
	return set.Union(set.Difference(s2, s1), set.Difference(s1, s2))
}
