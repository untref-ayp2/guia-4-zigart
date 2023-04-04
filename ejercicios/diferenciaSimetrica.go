package ejercicios

import (
	"guia4/set"
)

func DiferenciaSimetrica[T comparable](s1, s2 *set.Set[T]) *set.Set[T] {
	result := set.NewSet[T]()
	firstArray := s1.Values()
	for _, value := range firstArray {
		if s1.Contains(value) {
			result.Add(value)
		}
	}

	return result
}
