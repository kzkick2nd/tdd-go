package main

import "testing"

func TestMain(t *testing.T) {
	a, b, want := 5, 2, 10
	var d Doller
	d.amount = a
	d.times(b)
	if got := d.amount; got != want {
		t.Fatalf("want = %d, got = %d", want, got)
	}
}
