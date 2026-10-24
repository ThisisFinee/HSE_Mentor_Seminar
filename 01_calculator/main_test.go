package main

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		name    string
		a, b    float64
		op      string
		want    float64
		wantErr bool
	}{
		{"addition", 2.5, 1.5, "+", 4, false},
		{"subtraction", 2, 5, "-", -3, false},
		{"multiplication", -3, 4, "*", -12, false},
		{"division", 7, 2, "/", 3.5, false},
		{"division by zero", 7, 0, "/", 0, true},
		{"unknown operator", 7, 2, "%", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculate(tt.a, tt.b, tt.op)
			if (err != nil) != tt.wantErr || !tt.wantErr && got != tt.want {
				t.Fatalf("calculate(%v, %v, %q) = %v, %v; want %v, error=%v", tt.a, tt.b, tt.op, got, err, tt.want, tt.wantErr)
			}
		})
	}
}
