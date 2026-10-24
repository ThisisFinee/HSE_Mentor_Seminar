package main

import "testing"

func TestCountSigns(t *testing.T) {
	got := countSigns([]int{-5, 0, 4, 1, -2, 0})
	want := SignCounts{Positive: 2, Negative: 2, Zero: 2}
	if got != want {
		t.Errorf("countSigns() = %+v; want %+v", got, want)
	}
}
