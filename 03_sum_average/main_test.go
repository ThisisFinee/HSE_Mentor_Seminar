package main

import "testing"

func TestSumAndAverage(t *testing.T) {
	tests := []struct {
		n           int
		wantSum     int
		wantAverage float64
		wantErr     bool
	}{
		{1, 1, 1, false},
		{4, 10, 2.5, false},
		{0, 0, 0, true},
		{-3, 0, 0, true},
	}
	for _, tt := range tests {
		sum, average, err := sumAndAverage(tt.n)
		if (err != nil) != tt.wantErr || !tt.wantErr && (sum != tt.wantSum || average != tt.wantAverage) {
			t.Errorf("sumAndAverage(%d) = (%d, %v, %v); want (%d, %v, error=%v)", tt.n, sum, average, err, tt.wantSum, tt.wantAverage, tt.wantErr)
		}
	}
}
