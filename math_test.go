package main

import "testing"

func TestSome(t *testing.T) {
	total := Soma(15, 13)

	if total != 30 {
		t.Errorf("Resultado da soma e invalido: Resultado %d. Resultado esperado: %d", total, 30)
	}
}
