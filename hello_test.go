package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello", func(t *testing.T) {
		got := Hello("Lud")
		want := "Hello, Lud"

		if got != want {
			t.Errorf("Wanted %q, got %q", want, got)
		}
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
