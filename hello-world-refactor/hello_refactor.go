package main

import "fmt"

const spanishLanguage string = "Spanish"
const frenchLanguage string = "French"
const englishHelloPrefix string = "Hello, "
const spanishHelloPrefix string = "Hola, "
const frenchHelloPrefix string = "Bonjour, "

func HelloRefactor(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanishLanguage:
		prefix = spanishHelloPrefix
	case frenchLanguage:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(HelloRefactor("Paco", ""))
}
