package ejercicios

import (
	"guia4/set"
)

func Interseccion[T comparable](conjuntos ...*set.Set[T]) *set.Set[T] {
	newSet := conjuntos[0]
	for i := 0; i < len(conjuntos); i++ {
		newSet = set.Intersection(conjuntos[i], newSet)
	}
	return newSet
}
