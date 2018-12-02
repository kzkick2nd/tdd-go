package main

import "testing"

func TestMain(t *testing.T) {
	a, b, want := 5, 2, 10
	var five = Doller{a}
	five.times(b)
	if got := five.amount; got != want {
		t.Fatalf("want = %d, got = %d", want, got)
	}
}
