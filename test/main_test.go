package main

import (
	"testing"
)

func TestSuma(t *testing.T) {
	tables := []struct {
		a int
		b int
		r int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{3, 2, 5},
		{25, 26, 51}, //Resultados
	}

	for _, table := range tables {
		total := Suma(table.a, table.b)
		if total != table.r {
			t.Errorf("[!] Suma(%d, %d) fue incorrecta, got: %d, want: %d.", table.a, table.b, total, table.r)
		}
	}
}

func TestMax(t *testing.T) {
	tablas := []struct {
		a int
		b int
		n int
	}{
		{4, 2, 4},
		{3, 1, 3},
	}

	for _, v := range tablas {
		max := GetMax(v.a, v.b)
		if max != v.n {
			t.Errorf("[!] GetMax fue incorrector, tenia %d, se esperaba %d", max, v.n)
		}
	}

}
