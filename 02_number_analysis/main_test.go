package main

import "testing"

func TestAnalyzeNumber(t *testing.T) {
	tests := []struct {
		n    int
		want NumberAnalysis
	}{
		{15, NumberAnalysis{"положительное", false, true, true}},
		{-6, NumberAnalysis{"отрицательное", true, true, false}},
		{0, NumberAnalysis{"равно нулю", true, true, true}},
		{7, NumberAnalysis{"положительное", false, false, false}},
	}
	for _, tt := range tests {
		if got := analyzeNumber(tt.n); got != tt.want {
			t.Errorf("analyzeNumber(%d) = %+v; want %+v", tt.n, got, tt.want)
		}
	}
}
