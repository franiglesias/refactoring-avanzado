# Técnica de refactorización: Wrap

La técnica **Wrap** consiste en envolver una dependencia problemática manteniendo su interfaz existente, pero mejorando su implementación interna (añadiendo validación, logging, retry, caching, etc.) sin romper el código que la usa.

**Clave:** La interfaz pública NO cambia. Los clientes siguen llamando igual, pero ganan funcionalidad.

## Escenario

Tenemos `LegacyEmailService` (una dependencia externa que NO podemos modificar) que se usa directamente en todo el código. El servicio es limitado: no valida emails, no tiene retry, no tiene logging, no maneja errores bien.

Queremos mejorar la funcionalidad SIN cambiar todas las llamadas existentes.

## Código inicial sugerido

```java
public class LegacyEmailService {
    public String sendEmail(String to, String subject, String body) {
        // Servicio legacy que no podemos modificar
        System.out.println("Sending email to " + to);
        return "Message ID: " + System.currentTimeMillis();
    }
}

public class NotificationService {
    private LegacyEmailService emailService = new LegacyEmailService();

    public void notifyWelcome(String email) {
        emailService.sendEmail(email, "Welcome", "Welcome to our service!");
    }

    public void notifyPasswordReset(String email) {
        emailService.sendEmail(email, "Password Reset", "Reset your password");
    }
}
```

## Ejercicio: Aplicar técnica WRAP

**Objetivo:** Crear un wrapper que mantiene la misma interfaz (`sendEmail(to, subject, body)`) pero añade:
- Validación de emails
- Logging de operaciones
- Sanitización del contenido
- Manejo de errores mejorado

### Pasos sugeridos

1. **Crear el wrapper** con la misma interfaz que `LegacyEmailService`:

```java
public class EmailServiceWrapper {
    private LegacyEmailService legacyService;

    public EmailServiceWrapper(LegacyEmailService legacyService) {
        this.legacyService = legacyService;
    }

    public String sendEmail(String to, String subject, String body) {
        return legacyService.sendEmail(to, subject, body);
    }
}
```

2. **Instanciar el wrapper** y reemplazarlo en `NotificationService`. Simula gestión de dependencias (puedes tener una instancia global de `EmailServiceWrapper`).

3. **Añadir validación** dentro del wrapper (sin cambiar la interfaz). Los errores de validación se lanzan como excepciones.

```java
private void validateEmail(String email) {
    if (email == null || !email.contains("@")) {
        throw new IllegalArgumentException("Invalid email: " + email);
    }
}
```

4. **Añadir logging**. Para el ejercicio basta con usar `System.out.println`, pero puedes introducir un Logger propio.

```java
public String sendEmail(String to, String subject, String body) {
    validateEmail(to);
    System.out.println("Validating email: " + to);
    System.out.println("Sending email to: " + to);

    String result = legacyService.sendEmail(to, subject, body);

    System.out.println("Email sent successfully: " + result);
    return result;
}
```

5. **Migrar los puntos de uso** al wrapper uno por uno.

6. **Añadir más funcionalidad** según necesites (retry, sanitización, plantillas, etc.)

### Criterios de aceptación

- ✅ La interfaz pública (`sendEmail(to, subject, body)`) NO cambia
- ✅ Los clientes no necesitan modificarse (solo cambiar la instancia usada)
- ✅ El wrapper añade funcionalidad (validación, logging, etc.)
- ✅ El servicio legacy sigue siendo usado internamente
- ✅ Puedes migrar punto por punto sin romper nada

## Estructura final sugerida

```java
public class EmailServiceWrapper {
    private LegacyEmailService legacyService;

    public EmailServiceWrapper(LegacyEmailService legacyService) {
        this.legacyService = legacyService;
    }

    public String sendEmail(String to, String subject, String body) {
        validateEmail(to);
        log("Validating email: " + to);

        String sanitizedBody = sanitize(body);

        log("Sending email to: " + to);
        String result = legacyService.sendEmail(to, subject, sanitizedBody);

        log("Email sent: " + result);
        return result;
    }

    private void validateEmail(String email) {
        if (email == null || !email.contains("@")) {
            throw new IllegalArgumentException("Invalid email: " + email);
        }
    }

    private String sanitize(String text) {
        return text.replaceAll("<script>", "");
    }

    private void log(String message) {
        System.out.println("[EmailService] " + message);
    }
}
```

## Testing

Puedes crear tests para verificar:
- Validación de emails inválidos
- Logging de operaciones
- Que el servicio legacy es llamado correctamente
- Manejo de errores

## Recursos adicionales

- Ver ejercicio equivalente en TypeScript: `typescript/src/refactoring/parallel-change/wrap-change/`
