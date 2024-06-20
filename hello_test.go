package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Lud")
	want := "Hello Lud"

	if got != want {
		t.Errorf("Wanted %q, got %q", want, got)
	}
}
