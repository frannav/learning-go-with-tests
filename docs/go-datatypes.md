# DataTypes in Go

```go
// --- Tipos básicos predefinidos ---

// Booleano
bool // true, false

// Numéricos
uint8, uint16, uint32, uint64, uint, uintptr   // enteros sin signo
int8, int16, int32, int64, int                 // enteros con signo
float32, float64                               // punto flotante IEEE 754
complex64, complex128                          // números complejos: float32+float32i, float64+float64i
byte = uint8  // alias para uint8
rune = int32  // alias para int32 (punto de código Unicode)

// Cadenas
string // secuencia inmutable de bytes

// --- Tipos compuestos ---
// Array
[n]T       // array de longitud n y elementos de tipo T

// Slice
[]T        // descriptor dinámico sobre un array subyacente de elementos T

// Struct
struct { ... } // colección de campos nombrados

// Puntero
*T         // apunta a un valor de tipo T

// Función
func(params) (results) // firma de función

// Interface
interface { ... }      // conjunto de métodos o restricciones de tipo

// Mapa
map[K]V    // colección desordenada de pares clave-valor

// Canal
chan T     // bidireccional
<-chan T   // solo recepción
chan<- T   // solo envío
```