package main

import "testing"

func TestSumEvenDivisibleByThree(t *testing.T) {
	tests := []struct {
		n       int
		want    int
		wantErr bool
	}{
		{5, 0, false},
		{6, 6, false},
		{20, 36, false},
		{0, 0, true},
	}
	for _, tt := range tests {
		got, err := sumEvenDivisibleByThree(tt.n)
		if (err != nil) != tt.wantErr || !tt.wantErr && got != tt.want {
			t.Errorf("sumEvenDivisibleByThree(%d) = %d, %v; want %d, error=%v", tt.n, got, err, tt.want, tt.wantErr)
		}
	}
}
