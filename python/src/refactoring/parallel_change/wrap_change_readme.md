# Técnica de refactorización: Wrap

La técnica **Wrap** consiste en envolver una dependencia problemática manteniendo su interfaz existente, pero mejorando su implementación interna (añadiendo validación, logging, retry, caching, etc.) sin romper el código que la usa.

**Clave:** La interfaz pública NO cambia. Los clientes siguen llamando igual, pero ganan funcionalidad.

## Escenario

Tenemos `LegacyEmailService` (una dependencia externa que NO podemos modificar) que se usa directamente en todo el código. El servicio es limitado: no valida emails, no tiene retry, no tiene logging, no maneja errores bien.

Queremos mejorar la funcionalidad SIN cambiar todas las llamadas existentes.

## Ejecutar tests

```shell
pytest src/refactoring/parallel_change/ -k wrap -v
```

## Código actual

En `wrap_change.py` existe:

- `LegacyEmailService` - servicio externo rígido
- `notify_welcome`, `notify_password_reset`, `notify_order_confirmation` - usan el servicio directamente

## Ejercicio: Aplicar técnica WRAP

**Objetivo:** Crear un wrapper que mantiene la misma interfaz (`send_email(to, subject, body)`) pero añade:

- Validación de emails
- Logging de operaciones
- Sanitización del contenido
- Plantillas
- Manejo de errores mejorado

### Pasos sugeridos

1. **Crear el wrapper** con la misma interfaz que `LegacyEmailService`:

```python
class EmailServiceWrapper:
    def __init__(self, legacy_service: LegacyEmailService) -> None:
        self._legacy_service = legacy_service

    def send_email(self, to: str, subject: str, body: str) -> str:
        return self._legacy_service.send_email(to, subject, body)
```

2. **Instanciar el wrapper y reemplazarlo** en `wrap_change.py`. En este caso simula algo parecido a gestión de dependencias (tienes una instancia global de `EmailServiceWrapper`), pero podrías reemplazarlo en cada punto de uso.

2. **Añadir validación** dentro del wrapper (sin cambiar la interfaz). Para el ejercicio nos basta una validación simple, pero debe estar cubierta por nuevos tests (puedes usar TDD). Los errores de validación se lanzan.

3. **Añadir logging**. Para el ejercicio nos basta con usar `print()` o logging del módulo estándar. En pytest puedes capturar logs así:

```python
def test_should_log_actions(caplog):
    import logging
    caplog.set_level(logging.INFO)

    notify_welcome('test@example.com')

    assert 'Email test@example.com is valid' in caplog.text
    assert 'Sending email to test@example.com' in caplog.text
```

4. **Migrar los puntos de uso** al wrapper uno por uno.

5. **Añadir más funcionalidad** según necesites (retry, sanitización, plantillas, etc)

### Criterios de aceptación

- ✅ La interfaz pública (`send_email(to, subject, body)`) NO cambia
- ✅ Los clientes no necesitan modificarse (solo cambiar la instancia usada)
- ✅ El wrapper añade funcionalidad (validación, logging, etc.)
- ✅ El servicio legacy sigue siendo usado internamente
- ✅ Puedes migrar punto por punto sin romper nada
