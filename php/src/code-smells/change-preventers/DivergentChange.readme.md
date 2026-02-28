# Divergent change

## Definición

Una clase tiene múltiples razones para cambiar, lo que normalmente indica que se ocupa de muchas responsabilidades que se deberían separar en clases especialistas más pequeñas.

## Ejemplo

ProfileManager maneja validación, persistencia, exportación y envío de emails—múltiples razones para cambiar concentradas en una sola clase.

```php
<?php

declare(strict_types=1);

namespace CodeSmells\ChangePreventers;

class User
{
    public function __construct(
        public string $id,
        public string $name,
        public string $email
    ) {}
}

class ProfileManager
{
    /** @var array<string, User> */
    private array $store = [];

    public function register(User $user): void
    {
        if (empty(trim($user->name))) {
            throw new \InvalidArgumentException('invalid name');
        }
        if (!str_contains($user->email, '@')) {
            throw new \InvalidArgumentException('invalid email');
        }
        $this->store[$user->id] = $user;
    }

    public function updateEmail(string $id, string $newEmail): void
    {
        if (!str_contains($newEmail, '@')) {
            throw new \InvalidArgumentException('invalid email');
        }
        $u = $this->store[$id] ?? null;
        if (!$u) {
            throw new \InvalidArgumentException('not found');
        }
        $this->store[$id] = new User($u->id, $u->name, $newEmail);
    }

    public function exportAsJson(): string
    {
        return json_encode(array_values($this->store));
    }

    public function exportAsCsv(): string
    {
        $rows = ['id,name,email'];
        foreach ($this->store as $u) {
            $rows[] = "{$u->id},{$u->name},{$u->email}";
        }
        return implode("\n", $rows);
    }

    public function sendWelcomeEmail(User $user): string
    {
        return "Welcome {$user->name}! Sent to {$user->email}";
    }
}
```

## Ejercicio

Añade un número de teléfono con validación, inclúyelo en las exportaciones y envía un SMS.

## Problemas que encontrarás

Tocarás validación, almacenamiento, exportAsJson/Csv y mensajería en un solo lugar, demostrando cómo un cambio fuerza ediciones en responsabilidades no relacionadas.
