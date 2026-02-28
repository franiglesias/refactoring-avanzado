# Alternative Classes with Different Interfaces

Clases alternativas con interfaces diferentes.

## Definición

Ocurre cuando dos clases realizan tareas similares o son conceptualmente intercambiables, pero exponen nombres de métodos diferentes. Esto impide el uso de polimorfismo y obliga a los clientes a escribir código condicional para decidir qué método llamar dependiendo de la clase que estén usando.

## Ejemplo

`TextLogger` usa `Log()` mientras que `MessageWriter` usa `Write()`, a pesar de que ambos tienen el mismo propósito.

```go
type TextLogger struct{}

func (t *TextLogger) Log(message string) {
	fmt.Printf("[text] %s\n", message)
}

type MessageWriter struct{}

func (m *MessageWriter) Write(entry string) {
	fmt.Printf("[text] %s\n", entry)
}

func UseAltClasses(choice string, msg string) {
	if choice == "logger" {
		logger := &TextLogger{}
		logger.Log(msg)
	} else {
		writer := &MessageWriter{}
		writer.Write(msg)
	}
}
```

## Ejercicio

Añade logging con marca de tiempo a ambas implementaciones y permite que el cliente pueda intercambiarlas en tiempo de ejecución sin usar condicionales.

## Problemas que encontrarás

Al no compartir una interfaz común, te verás obligado a duplicar lógica en métodos con nombres distintos y a esparcir sentencias `if/else` o `switch` en los clientes, haciendo que cambios simples se vuelvan tediosos y propensos a errores.
