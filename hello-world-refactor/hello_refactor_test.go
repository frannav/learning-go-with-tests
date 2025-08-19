package main

import "testing"

func TestHelloRefactor(t *testing.T) {
	got := HelloRefactor("Pepe")
	want := "Hello, Pepe"

	if got != want {
		t.Errorf("\n got %q \n want %q", got, want)
	}
}
