package main

import "fmt"

const englishHelloPrefix string = "Hello, "

func HelloRefactor(name string) string {
	if name == "" {
		name = "world"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(HelloRefactor("Paco"))
}
