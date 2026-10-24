package main

import "testing"

func TestCalculateStatistics(t *testing.T) {
	got, err := calculateStatistics([]int{-3, 0, 4, 5, -2})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := Statistics{Sum: 4, Average: 0.8, Min: -3, Max: 5, Even: 3, Odd: 2}
	if got != want {
		t.Errorf("calculateStatistics() = %+v; want %+v", got, want)
	}
}

func TestCalculateStatisticsRejectsEmptySequence(t *testing.T) {
	if _, err := calculateStatistics(nil); err == nil {
		t.Error("empty sequence must produce an error")
	}
}
