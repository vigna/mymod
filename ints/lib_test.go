package ints

import "testing"

func TestAdd(t *testing.T) {
	if Add(5, 3) != 8 {
		t.Fatal("La somma non funziona")
	}
}

func TestAddNegativi(t *testing.T) {
	if Add(-5, -3) != -8 {
		t.Fatal("La somma non funziona coi negativi")
	}
}

func TestProd(t *testing.T) {
	if Prod(5, 3) != 15 {
		t.Fatal("Il prodotto non funziona")
	}
}

