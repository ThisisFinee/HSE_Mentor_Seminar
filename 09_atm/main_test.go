package main

import (
	"math"
	"testing"
)

func TestProcessTransaction(t *testing.T) {
	tests := []struct {
		name    string
		balance float64
		action  int
		amount  float64
		want    float64
		wantErr bool
	}{
		{"show balance", 100, showBalance, 0, 100, false},
		{"deposit and round", 10.1, deposit, 0.2, 10.3, false},
		{"withdraw", 100, withdraw, 25.5, 74.5, false},
		{"withdraw exact balance", 10, withdraw, 10, 0, false},
		{"insufficient funds", 10, withdraw, 10.01, 0, true},
		{"zero deposit", 10, deposit, 0, 0, true},
		{"negative amount", 10, withdraw, -1, 0, true},
		{"invalid action", 10, 99, 0, 0, true},
		{"negative balance", -1, showBalance, 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := processTransaction(tt.balance, tt.action, tt.amount)
			if (err != nil) != tt.wantErr || !tt.wantErr && math.Abs(got-tt.want) > 1e-9 {
				t.Fatalf("processTransaction(%v, %d, %v) = %v, %v; want %v, error=%v", tt.balance, tt.action, tt.amount, got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestProcessTransactionRejectsNonFiniteAmount(t *testing.T) {
	if _, err := processTransaction(10, deposit, math.Inf(1)); err == nil {
		t.Error("infinite amount must be rejected")
	}
}
