package main

import "testing"

func TestSoma(t *testing.T) {
	resultado := Soma(3, 2)
	esperado := 5
	if resultado != esperado {
		t.Errorf("soma(3,2) = %d; Esperado %d", resultado, esperado)
	}

}
