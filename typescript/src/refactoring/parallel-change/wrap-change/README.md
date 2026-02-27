# Técnica de refactorización: Wrap

La técnica **Wrap** consiste en envolver una dependencia problemática manteniendo su interfaz existente, pero mejorando su implementación interna (añadiendo validación, logging, retry, caching, etc.) sin romper el código que la usa.

**Clave:** La interfaz pública NO cambia. Los clientes siguen llamando igual, pero ganan funcionalidad.

## Escenario

Tenemos `LegacyEmailService` (una dependencia externa que NO podemos modificar) que se usa directamente en todo el código. El servicio es limitado: no valida emails, no tiene retry, no tiene logging, no maneja errores bien.

Queremos mejorar la funcionalidad SIN cambiar todas las llamadas existentes.

## Ejecutar tests

```shell
npm run test -- src/refactoring/parallel-change/wrap-change/wrap-change.test.ts
```

## Código actual

En `wrap-change.ts` existe:

- `LegacyEmailService` - servicio externo rígido
- `notifyWelcome`, `notifyPasswordReset`, `notifyOrderConfirmation` - usan el servicio directamente

## Ejercicio: Aplicar técnica WRAP

**Objetivo:** Crear un wrapper que mantiene la misma interfaz (`sendEmail(to, subject, body)`) pero añade:

- Validación de emails
- Logging de operaciones
- Sanitización del contenido
- Plantillas
- Manejo de errores mejorado

### Pasos sugeridos

1. **Crear el wrapper** con la misma interfaz que `LegacyEmailService`:

```typescript
class EmailServiceWrapper {
  private legacyService: LegacyEmailService;

  constructor(legacyService: LegacyEmailService) {
    this.legacyService = legacyService;
  }

  sendEmail(to: string, subject: string, body: string): string {
    return this.legacyService.sendEmail(to, subject, body)
  }
}
```

2. **Instanciar el wrapper y reemplazarlo** en `wrap-change.ts`. En este caso simula algo parecido a gestión de dependencias (tienes una instancia global de `EmailServiceWrapper`), pero podrías reemplazarlo en cada punto de uso.

2. **Añadir validación** dentro del wrapper (sin cambiar la interfaz). Para el ejercicio nos basta una validación simple, pero debe estar cubierta por nuevos tests (puedes usar TDD). Los errores de validación se lanzan.

3. **Añadir logging**. Para el ejercicio nos basta con usar `console.log`, pero puedes introducir un Logger propio si quieres. En vitest puedes usar algo como esto para montar tu test:

```typescript
it('should log actions', () => {
  const consoleSpy = vi.spyOn(console, 'log').mockImplementation(() => {
  })

  notifyWelcome('test@example.com')

  expect(consoleSpy).toHaveBeenCalledWith('Email test@example.com is valid')
  expect(consoleSpy).toHaveBeenCalledWith('Sending email to test@example.com')
  consoleSpy.mockRestore()
}); 
```

4. **Migrar los puntos de uso** al wrapper uno por uno.

5. **Añadir más funcionalidad** según necesites (retry, sanitización, plantillas, etc)

### Criterios de aceptación

- ✅ La interfaz pública (`sendEmail(to, subject, body)`) NO cambia
- ✅ Los clientes no necesitan modificarse (solo cambiar la instancia usada)
- ✅ El wrapper añade funcionalidad (validación, logging, etc.)
- ✅ El servicio legacy sigue siendo usado internamente
- ✅ Puedes migrar punto por punto sin romper nada
