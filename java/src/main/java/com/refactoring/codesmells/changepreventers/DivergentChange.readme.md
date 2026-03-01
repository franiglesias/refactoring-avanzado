# Divergent Change

Cambio divergente.

## Definición

Una clase se modifica por diferentes razones no relacionadas. Cada vez que hay un cambio de diferente naturaleza, hay que tocar la misma clase, lo que viola el principio de responsabilidad única.

## Ejemplo

```java
public static class ProfileManager {
    private Map<String, User> store = new HashMap<>();

    // Gestión de usuarios
    public void register(User user) {
        if (user.name == null || user.name.trim().isEmpty()) {
            throw new IllegalArgumentException("invalid name");
        }
        if (!user.email.contains("@")) {
            throw new IllegalArgumentException("invalid email");
        }
        store.put(user.id, user);
    }

    public void updateEmail(String id, String newEmail) {
        if (!newEmail.contains("@")) {
            throw new IllegalArgumentException("invalid email");
        }
        User u = store.get(id);
        if (u == null) {
            throw new IllegalArgumentException("not found");
        }
        u.email = newEmail;
        store.put(id, u);
    }

    // Exportación de datos
    public String exportAsJson() {
        List<String> users = store.values().stream()
            .map(u -> String.format("{\"id\":\"%s\",\"name\":\"%s\",\"email\":\"%s\"}",
                u.id, u.name, u.email))
            .collect(Collectors.toList());
        return "[" + String.join(",", users) + "]";
    }

    public String exportAsCsv() {
        List<String> rows = new ArrayList<>();
        rows.add("id,name,email");
        for (User u : store.values()) {
            rows.add(String.format("%s,%s,%s", u.id, u.name, u.email));
        }
        return String.join("\n", rows);
    }

    // Notificaciones
    public String sendWelcomeEmail(User user) {
        return String.format("Welcome %s! Sent to %s", user.name, user.email);
    }
}
```

## Ejercicio

Añade soporte para exportar a XML y enviar notificaciones SMS.

## Problemas que encontrarás

Cada nuevo formato de exportación o canal de comunicación requiere modificar esta clase, acumulando responsabilidades no relacionadas.
