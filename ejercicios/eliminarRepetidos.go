package ejercicios

import (
	"guia4/set"
)

/*
Escribir una función que reciba un arreglo de elementos
comparables y elimine los repetidos.
*/
func EliminarRepetidos[T comparable](arreglo []T) []T {

	conjunto := set.NewSet(arreglo...)

	return conjunto.Values()

}
