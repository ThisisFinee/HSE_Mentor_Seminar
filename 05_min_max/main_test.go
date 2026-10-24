package main

import "testing"

func TestMinMax(t *testing.T) {
	tests := []struct {
		values           []int
		wantMin, wantMax int
		wantErr          bool
	}{
		{[]int{4, -2, 9, 0, 9}, -2, 9, false},
		{[]int{5}, 5, 5, false},
		{nil, 0, 0, true},
	}
	for _, tt := range tests {
		min, max, err := minMax(tt.values)
		if (err != nil) != tt.wantErr || !tt.wantErr && (min != tt.wantMin || max != tt.wantMax) {
			t.Errorf("minMax(%v) = (%d, %d, %v)", tt.values, min, max, err)
		}
	}
}
