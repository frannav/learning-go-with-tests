package main

import "fmt"

func HelloRefactor(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(HelloRefactor("Paco"))
}
