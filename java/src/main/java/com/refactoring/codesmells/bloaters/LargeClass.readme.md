# Large Class

Clase grande.

## Definición

Una clase contiene muchas propiedades, muchos métodos o muchas líneas de código, acumulando muchas responsabilidades no relacionadas o que pueden responder a necesidades diferentes.

## Ejemplo

```java
public static class UserAccount {
    private String name;
    private String email;
    private String password;
    private Date lastLogin;
    private int loginAttempts = 0;
    private List<String> notifications = new ArrayList<>();
    private boolean isAdmin;

    // --- Autenticación ---
    public boolean login(String password) {
        if (this.password.equals(password)) {
            this.lastLogin = new Date();
            this.loginAttempts = 0;
            System.out.println("Inicio de sesión exitoso");
            return true;
        } else {
            this.loginAttempts++;
            System.out.println("Contraseña incorrecta");
            return false;
        }
    }

    public void resetPassword(String newPassword) {
        this.password = newPassword;
        System.out.println("Contraseña actualizada");
    }

    // --- Perfil ---
    public void updateEmail(String newEmail) {
        this.email = newEmail;
        System.out.println("Correo actualizado");
    }

    // --- Notificaciones ---
    public void addNotification(String message) {
        this.notifications.add(message);
    }

    // --- Administración ---
    public void promoteToAdmin() {
        this.isAdmin = true;
    }
}
```

## Ejercicio

Añade soporte para autenticación de dos factores (2FA) y preferencias de notificación.

## Problemas que encontrarás

Tocarás autenticación, estado y notificaciones en una clase inflada, aumentando la probabilidad de romper comportamiento no relacionado.
