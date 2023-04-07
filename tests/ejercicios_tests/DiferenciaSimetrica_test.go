package tests

import (
	"guia4/ejercicios"
	"guia4/set"
	"testing"
)

func TestDiferenciaSimetrica(t *testing.T) {
	s1 := set.NewSet(1, 2, 3, 4)
	s2 := set.NewSet(3, 4, 5, 6)
	s3 := ejercicios.DiferenciaSimetrica(s1, s2)
	if s3.Size() != 4 {
		t.Errorf("Tamaño %d, debería ser: %d", s3.Size(), 4)
	}
}

func TestNilpotnecia(t *testing.T) {
	s1 := set.NewSet(1, 2, 3, 4)
	s3 := ejercicios.DiferenciaSimetrica(s1, s1)
	if s3.Size() != 0 {
		t.Errorf("Tamaño %d, debería ser: %d", s3.Size(), 0)
	}
}

func TestAsociativa(t *testing.T) {
	sA := set.NewSet(1, 2, 3, 4)
	sB := set.NewSet(3, 4, 5, 6)
	sC := set.NewSet(5, 6, 7, 8)
	sAB := ejercicios.DiferenciaSimetrica(sA, sB)
	sAB_C := ejercicios.DiferenciaSimetrica(sAB, sC)
	sBC := ejercicios.DiferenciaSimetrica(sB, sC)
	sA_BC := ejercicios.DiferenciaSimetrica(sA, sBC)
	if sAB_C.Size() != sA_BC.Size() {
		t.Errorf("LOS SET NO SON IGUALES %v, %v", sAB_C, sA_BC)
	}
}

func TestConmutativa(t *testing.T) {
	sA := set.NewSet(1, 2, 3, 4)
	sB := set.NewSet(3, 4, 5, 6)
	sAB := ejercicios.DiferenciaSimetrica(sA, sB)
	sBA := ejercicios.DiferenciaSimetrica(sB, sA)
	if sAB.Size() != sBA.Size() {
		t.Errorf("LOS SET NO SON IGUALES %v, %v", sAB, sBA)
	}
}

func TestNulo(t *testing.T) {
	sA := set.NewSet(1, 2, 3, 4)
	sA_nill := ejercicios.DiferenciaSimetrica(sA, nil)
	if sA_nill.Size() != sA.Size() {
		t.Errorf("LOS SET NO SON IGUALES %v, %v", sA, sA_nill)
	}
}
