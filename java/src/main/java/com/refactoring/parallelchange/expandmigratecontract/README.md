# Técnica de refactorización: Cambio en Paralelo (Expand–Migrate–Contract)

Este ejercicio te ayuda a practicar la técnica de **Parallel Change** en su forma más pura: expandir la interfaz, migrar los consumidores y contraer la interfaz antigua.

## Escenario

Tenemos una clase `User` con un campo `name` (nombre completo como String único) y **cuatro funciones consumidoras** que dependen de ese campo:

- `formatGreeting(user)` — saludo al usuario
- `formatEmailHeader(user)` — cabecera de email con nombre y dirección
- `formatDisplayName(user)` — nombre para mostrar con ID
- `buildUserSummary(users)` — listado de nombres de usuarios

**Objetivo**: Refactorizar `name` para que sea `firstName` y `lastName`, aplicando cambio en paralelo para no romper nunca los tests.

## Implementación inicial sugerida

```java
public class User {
    private String id;
    private String name;
    private String email;

    public User(String id, String name, String email) {
        this.id = id;
        this.name = name;
        this.email = email;
    }

    public String getId() { return id; }
    public String getName() { return name; }
    public String getEmail() { return email; }
}

public class UserFormatter {
    public String formatGreeting(User user) {
        return "Hello, " + user.getName() + "!";
    }

    public String formatEmailHeader(User user) {
        return user.getName() + " <" + user.getEmail() + ">";
    }

    public String formatDisplayName(User user) {
        return user.getName() + " (#" + user.getId() + ")";
    }

    public String buildUserSummary(List<User> users) {
        return users.stream()
            .map(User::getName)
            .collect(Collectors.joining(", "));
    }
}
```

## Ejercicio: Expand–Migrate–Contract

### Fase 1: EXPAND (Expandir)

Añade los campos `firstName` y `lastName` a la clase `User` **sin eliminar** el campo `name`. Puedes hacer que `name` se calcule a partir de los nuevos campos, o aceptar los tres en el constructor.

```java
public class User {
    private String id;
    private String name;  // Mantener por ahora
    private String firstName;
    private String lastName;
    private String email;

    public User(String id, String name, String email) {
        this.id = id;
        this.name = name;
        // Inicializar firstName y lastName desde name
        String[] parts = name.split(" ", 2);
        this.firstName = parts[0];
        this.lastName = parts.length > 1 ? parts[1] : "";
        this.email = email;
    }

    // O constructor alternativo:
    public User(String id, String firstName, String lastName, String email) {
        this.id = id;
        this.firstName = firstName;
        this.lastName = lastName;
        this.name = firstName + " " + lastName; // Calcular name
        this.email = email;
    }

    public String getFirstName() { return firstName; }
    public String getLastName() { return lastName; }
    // Mantener getName() por ahora
}
```

- Los tests deben seguir pasando sin cambios
- Haz commit

### Fase 2: MIGRATE (Migrar)

Migra las funciones consumidoras **una por una** para que usen `firstName` y/o `lastName` en lugar de `name`. Después de migrar cada función:

1. Actualiza la implementación
2. Actualiza el test correspondiente si es necesario
3. Ejecuta los tests — deben pasar
4. Haz commit

**Orden sugerido:**

1. **`formatGreeting`** — usa solo `firstName`:
```java
public String formatGreeting(User user) {
    return "Hello, " + user.getFirstName() + "!";
}
```

2. **`formatEmailHeader`** — usa `firstName` + `lastName`:
```java
public String formatEmailHeader(User user) {
    return user.getFirstName() + " " + user.getLastName() +
           " <" + user.getEmail() + ">";
}
```

3. **`formatDisplayName`** — usa `firstName` + `lastName`:
```java
public String formatDisplayName(User user) {
    return user.getFirstName() + " " + user.getLastName() +
           " (#" + user.getId() + ")";
}
```

4. **`buildUserSummary`** — usa `lastName`, `firstName`:
```java
public String buildUserSummary(List<User> users) {
    return users.stream()
        .map(u -> u.getLastName() + ", " + u.getFirstName())
        .collect(Collectors.joining("; "));
}
```

### Fase 3: CONTRACT (Contraer)

Una vez que **ningún consumidor** use `name`:

1. Elimina el campo `name` de la clase `User`
2. Elimina el método `getName()`
3. Ajusta el constructor para que solo reciba `firstName` y `lastName`
4. Los tests deben pasar
5. Haz commit

```java
public class User {
    private String id;
    private String firstName;
    private String lastName;
    private String email;

    public User(String id, String firstName, String lastName, String email) {
        this.id = id;
        this.firstName = firstName;
        this.lastName = lastName;
        this.email = email;
    }

    // Solo getters para firstName y lastName
}
```

## Testing

Crea tests que verifiquen cada función antes y después de cada migración:

```java
@Test
public void testFormatGreeting() {
    User user = new User("1", "Ana", "García", "ana@example.com");
    assertEquals("Hello, Ana!", formatter.formatGreeting(user));
}
```

## Criterios de aceptación

- ✅ Cada fase mantiene los tests en verde
- ✅ Cada cambio es pequeño e incremental
- ✅ La migración es segura y reversible en cada paso
- ✅ Al final, el campo `name` ya no existe

## Pasos pequeños y seguros

La clave del ejercicio es hacer cambios **muy pequeños** que siempre mantienen el código funcionando. Nunca deberías tener tests rotos por más de unos segundos.

## Recursos adicionales

- Ver ejercicio equivalente en TypeScript: `typescript/src/refactoring/parallel-change/expand-migrate-contract/`
