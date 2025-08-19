# Functions in Go

- Las funciones pueden recibir múltiples parámetros, entre paréntesis.
- Las funciones pueden tener valores de retorno, entre paréntesis también. Cuando se devuelven esos valores de retorno es posible usar `return` sin nada.
- Los valores de retorno crean esa variable en la función, dependiendo del type del valor, este valor se asignará a cero 0 si es de tipo `integer` o a string vacío "" si es `string`
- Al crear una función, esta puede ser pública o privada. Si el nombre de la función empieza en minúsculas esta será privada. Mientras que si empieza por Mayúsculas será pública.


```go
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
```