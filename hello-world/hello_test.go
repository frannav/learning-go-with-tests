package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	/* := Es una declaración "corta"
	e infiere el tipo de dato
	además de declarar y asignar */

	/* La otra forma de declarar y asignar variables
	es var my-variable string = "my variable" */

	got := Hello()
	// var got string = Hello()

	want := "Hello, world"
	// var want string = "Hello, world"

	if got != want {
		// t.Errorf("got %q want %q", got, want)
		t.Errorf("\n got %q \n want %q", got, want)
	}
}
