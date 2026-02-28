# Data Class

Clase de datos.

## Definición

Una clase de datos es aquella que solo contiene campos y métodos para acceder a ellos (getters/setters), sin poseer lógica de negocio propia. Esto suele derivar en modelos de dominio anémicos, donde el comportamiento está disperso en otros servicios o clases que manipulan estos datos.

## Ejemplo

`UserRecord` es una clase que solo almacena datos, mientras que la validación y la lógica de creación residen en `UserService`.

```php
<?php

declare(strict_types=1);

namespace CodeSmells\Dispensables;

use DateTime;

class UserRecord
{
    public function __construct(
        public string $id,
        public string $name,
        public string $email,
        public DateTime $createdAt
    ) {
    }
}

class UserService
{
    public function createUser(string $name, string $email): UserRecord
    {
        if (!str_contains($email, '@')) {
            throw new \InvalidArgumentException('Invalid email');
        }

        return new UserRecord(uniqid(), $name, $email, new DateTime());
    }

    public function updateUserEmail(UserRecord $user, string $newEmail): void
    {
        if (!str_contains($newEmail, '@')) {
            throw new \InvalidArgumentException('Invalid email');
        }
        $user->email = $newEmail;
    }
}

class UserReportGenerator
{
    public function generateUserSummary(UserRecord $user): string
    {
        return "User {$user->name} ({$user->email}) created on {$user->createdAt->format('Y-m-d')}";
    }
}

function demoDataClass(): string
{
    $service = new UserService();
    $report = new UserReportGenerator();
    $user = $service->createUser('Lina', 'lina@example.com');
    $service->updateUserEmail($user, 'lina+news@example.com');
    return $report->generateUserSummary($user);
}
```

## Ejercicio

Implementa reglas de dominio adicionales, como requerir verificación de email o restringir el registro a ciertos dominios (ej. `company.com`).

## Problemas que encontrarás

Tendrás que modificar múltiples servicios y lugares que manipulan `UserRecord`. Esto demuestra cómo separar el comportamiento de los datos provoca que cambios simples se dispersen ampliamente por el código (Shotgun Surgery).
