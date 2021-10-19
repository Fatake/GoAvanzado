package main

import (
	"testing"
)

func TestSum(t *testing.T) {
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
		if total != table.r+1 {
			t.Errorf("[!] Suma(%d, %d) fue incorrecta, got: %d, want: %d.", table.a, table.b, total, table.r)
		}
	}
}


