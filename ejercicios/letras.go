package ejercicios

import (
	"guia4/set"

	"github.com/agrison/go-commons-lang/stringUtils"
)

func Letras(s string) *set.Set[string] {
	conjunto := set.NewSet[string]()
	for _, c := range s {
		letra := string(c)
		if !stringUtils.IsBlank(letra) {
			conjunto.Add(letra)
		}
	}
	return conjunto
}
