package main

import "testing"

func TestHelloRefactor(t *testing.T) {
	t.Run("Say custom hello", func(t *testing.T) {
		got := HelloRefactor("Pepe")
		want := "Hello, " + "Pepe"
		assertCorrectString(t, got, want)
	})

	t.Run("Say 'Hello, world' when an empty string is passed to the function", func(t *testing.T) {
		got := HelloRefactor("")
		want := "Hello, " + "world"
		assertCorrectString(t, got, want)
	})
}

func assertCorrectString(t testing.TB, got, want string) {
	// t.Helper() is needed to tell the test suite that this method is a helper
	t.Helper()
	if got != want {
		t.Errorf("\n got %q \n want %q", got, want)
	}
}
