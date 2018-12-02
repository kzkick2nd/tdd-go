package main

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{0, 1, 1},
		{-1, -1, -2},
		{-3, 2, -1},
	}
	for _, tt := range tests {
		if got := Sum(tt.a, tt.b); got != tt.want {
			t.Fatalf("want = %d, got = %d", tt.want, got)
		}
	}
}
