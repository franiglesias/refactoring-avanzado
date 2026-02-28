# Técnica de refactorización: Wrap

La técnica **Wrap** consiste en envolver una dependencia problemática manteniendo su interfaz existente, pero mejorando su implementación interna (añadiendo validación, logging, retry, caching, etc.) sin romper el código que la usa.

**Clave:** La interfaz pública NO cambia. Los clientes siguen llamando igual, pero ganan funcionalidad.

## Escenario

Tenemos `LegacyEmailService` (una dependencia externa que NO podemos modificar) que se usa directamente en todo el código. El servicio es limitado: no valida emails, no tiene retry, no tiene logging, no maneja errores bien.

Queremos mejorar la funcionalidad SIN cambiar todas las llamadas existentes.

## Ejecutar tests

```shell
go test ./refactoring/parallel_change/wrap_change
```

## Código actual

En `wrap_change.go` existe:

- `LegacyEmailService` - servicio externo rígido
- `NotifyWelcome`, `NotifyPasswordReset`, `NotifyOrderConfirmation` - usan el servicio directamente

## Ejercicio: Aplicar técnica WRAP

**Objetivo:** Crear un wrapper que mantiene la misma interfaz (`SendEmail(to, subject, body)`) pero añade:

- Validación de emails
- Logging de operaciones
- Sanitización del contenido
- Plantillas
- Manejo de errores mejorado

### Pasos sugeridos

1. **Crear el wrapper** con la misma interfaz que `LegacyEmailService`:

```go
type EmailServiceWrapper struct {
	legacyService *LegacyEmailService
}

func NewEmailServiceWrapper(legacyService *LegacyEmailService) *EmailServiceWrapper {
	return &EmailServiceWrapper{legacyService: legacyService}
}

func (w *EmailServiceWrapper) SendEmail(to, subject, body string) (string, error) {
	return w.legacyService.SendEmail(to, subject, body)
}
```

2. **Instanciar el wrapper y reemplazarlo** en `wrap_change.go`. En este caso simula algo parecido a gestión de dependencias (tienes una instancia global de `EmailServiceWrapper`), pero podrías reemplazarlo en cada punto de uso.

3. **Añadir validación** dentro del wrapper (sin cambiar la interfaz). Para el ejercicio nos basta una validación simple, pero debe estar cubierta por nuevos tests (puedes usar TDD). Los errores de validación se retornan.

4. **Añadir logging**. Para el ejercicio nos basta con usar `log.Println`, pero puedes introducir un Logger propio si quieres.

5. **Migrar los puntos de uso** al wrapper uno por uno.

6. **Añadir más funcionalidad** según necesites (retry, sanitización, plantillas, etc)

### Criterios de aceptación

- La interfaz pública (`SendEmail(to, subject, body)`) NO cambia
- Los clientes no necesitan modificarse (solo cambiar la instancia usada)
- El wrapper añade funcionalidad (validación, logging, etc.)
- El servicio legacy sigue siendo usado internamente
- Puedes migrar punto por punto sin romper nada
