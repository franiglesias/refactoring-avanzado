# Large class

Clase grande.

## Definición

Una clase contiene muchas propiedades, muchos métodos o muchas líneas de código, acumulando muchas responsabilidades no relacionadas o que pueden responder a necesidades diferentes.

## Ejemplo

```php
<?php

declare(strict_types=1);

namespace RefactoringAvanzado\CodeSmells\Bloaters;

use DateTime;

class UserAccount
{
    private string $name;
    private string $email;
    private string $password;
    private DateTime $lastLogin;
    private int $loginAttempts = 0;
    /** @var array<string> */
    private array $notifications = [];
    private bool $isAdmin;

    public function __construct(string $name, string $email, string $password, bool $isAdmin = false)
    {
        $this->name = $name;
        $this->email = $email;
        $this->password = $password;
        $this->lastLogin = new DateTime();
        $this->isAdmin = $isAdmin;
    }

    // --- Autenticación ---
    public function login(string $password): bool
    {
        if ($this->password === $password) {
            $this->lastLogin = new DateTime();
            $this->loginAttempts = 0;
            echo "Inicio de sesión exitoso\n";
            return true;
        } else {
            $this->loginAttempts++;
            echo "Contraseña incorrecta\n";
            return false;
        }
    }

    public function resetPassword(string $newPassword): void
    {
        $this->password = $newPassword;
        echo "Contraseña actualizada\n";
    }

    // --- Perfil ---
    public function updateEmail(string $newEmail): void
    {
        $this->email = $newEmail;
        echo "Correo actualizado\n";
    }

    public function updateName(string $newName): void
    {
        $this->name = $newName;
        echo "Nombre actualizado\n";
    }

    // --- Notificaciones ---
    public function addNotification(string $message): void
    {
        $this->notifications[] = $message;
    }

    public function getNotifications(): array
    {
        return $this->notifications;
    }

    public function clearNotifications(): void
    {
        $this->notifications = [];
    }

    // --- Administración ---
    public function promoteToAdmin(): void
    {
        $this->isAdmin = true;
    }

    public function revokeAdmin(): void
    {
        $this->isAdmin = false;
    }
}
```

## Ejercicio

Añade soporte para autenticación de dos factores (2FA) y preferencias de notificación.

## Problemas que encontrarás

Tocarás autenticación, estado y notificaciones en una clase inflada, aumentando la probabilidad de romper comportamiento no relacionado.
