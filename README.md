# Learning Go with Tests 🏗️

## Index

TODO

## Introducción

- Lenguaje creado en Google para resolver problemas de escalabilidad.
- Go es un lenguaje compilado y de tipado estático.
- Muy eficiente en términos de memoria y uso de la CPU.
- Incorpora la concurrencia en el lenguaje mediante `goroutines` y `channels`

### Instalación

- Mac users:
```bash
brew install go
```

- Linting:
    - Go ya incorpora un lint `go fmt`
    - Existen lints con ciertas mejoras respecto default. Como `golangci-lint`, para instalarlo:
        - `brew install golangci-lint`


### Módulos Go

Módulos introducidos a partir de la versión 1.11. Introducidos para resolver diversos problemas como la gestión de dependencias, builds, etc.

- Crear nuevo modulo:
```bash
go mod init module-name
```
Esto creará un archivo `module-name.mod` que contiene la version de Go así como las dependencias necesarias.

- Para más información sobre módulos:
```bash
go help mod init
go help mododules
go help mod
```

### Testing

Para escribir un test crearemos una función pero con algunas reglas:
- El nombre del fichero tiene que ser `example_test.go`
- El nombre de la función tiene que empezar por `Test`
- La función de test solo aceptará un argumento `t *testing.T` que se importará el paquete de `testing`

### Docs

Go tiene una herramienta incorporada que te permite examinar cualquier paquete instalado en tu sistema o módulo en el que estás trabajando. 

```bash
go doc fmt
```

### Apuntes

Iré añadiendo diferentes apuntes que voy haciendo.

- [DataTypes](docs/go-datatypes.md)




