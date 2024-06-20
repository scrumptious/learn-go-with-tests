package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello", func(t *testing.T) {
		got := Hello(Greeting{name: "Lud"})
		want := "Hello, Lud"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello(Greeting{})
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello(Greeting{
			name:     "Carlos",
			language: "Spanish",
		})
		want := "Hola, Carlos"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	//helper to make sure we're reported the right line when test fails
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
