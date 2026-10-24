package main

import "testing"

func TestMultiplicationTable(t *testing.T) {
	table := multiplicationTable(-3)
	if len(table) != 10 {
		t.Fatalf("table length = %d; want 10", len(table))
	}
	if table[0] != "-3 x 1 = -3" || table[9] != "-3 x 10 = -30" {
		t.Errorf("unexpected bounds: first=%q, last=%q", table[0], table[9])
	}
}
