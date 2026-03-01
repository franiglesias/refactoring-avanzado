# Data Class

Clase de datos.

## Definición

Una clase que solo contiene datos (getters y setters) sin comportamiento propio. Toda la lógica relacionada con esos datos está en otras clases, violando el principio de que los datos y el comportamiento deben estar juntos.

## Ejemplo

```java
public static class UserRecord {
    public String id;
    public String name;
    public String email;
    public Date createdAt;

    public UserRecord(String id, String name, String email, Date createdAt) {
        this.id = id;
        this.name = name;
        this.email = email;
        this.createdAt = createdAt;
    }
}

public static class UserService {
    public UserRecord createUser(String name, String email) {
        if (!email.contains("@")) {
            throw new IllegalArgumentException("Invalid email");
        }

        return new UserRecord(UUID.randomUUID().toString(), name, email, new Date());
    }

    public void updateUserEmail(UserRecord user, String newEmail) {
        if (!newEmail.contains("@")) {
            throw new IllegalArgumentException("Invalid email");
        }
        user.email = newEmail;
    }
}

public static class UserReportGenerator {
    public String generateUserSummary(UserRecord user) {
        SimpleDateFormat sdf = new SimpleDateFormat("dd/MM/yyyy");
        return String.format("User %s (%s) created on %s",
            user.name, user.email, sdf.format(user.createdAt));
    }
}
```

## Ejercicio

Añade validación de formato de nombre y generación de nombre de usuario único.

## Problemas que encontrarás

La lógica relacionada con usuarios seguirá esparcida en múltiples clases, haciendo difícil encontrar y modificar comportamientos relacionados.
