package main

import "testing"

func TestIsPrime(t *testing.T) {
	tests := map[int]bool{-7: false, 0: false, 1: false, 2: true, 3: true, 4: false, 97: true, 99: false}
	for n, want := range tests {
		if got := isPrime(n); got != want {
			t.Errorf("isPrime(%d) = %v; want %v", n, got, want)
		}
	}
}
