<?php

namespace Calisthenics;

class UserManager
{
    /**
     * @param array{name: string, email: string, password: string} $userData
     */
    public function createUser(array $userData): void
    {
        // Validar datos
        if (empty($userData['name'])) {
            throw new \InvalidArgumentException('Name is required');
        }
        if (empty($userData['email'])) {
            throw new \InvalidArgumentException('Email is required');
        }
        if (!filter_var($userData['email'], FILTER_VALIDATE_EMAIL)) {
            throw new \InvalidArgumentException('Invalid email');
        }
        if (empty($userData['password'])) {
            throw new \InvalidArgumentException('Password is required');
        }
        if (strlen($userData['password']) < 8) {
            throw new \InvalidArgumentException('Password too short');
        }

        // Hash password
        $hashedPassword = password_hash($userData['password'], PASSWORD_DEFAULT);

        // Guardar usuario (simulated)
        echo "User created: {$userData['name']}\n";
    }

    public function updateUser(string $userId, array $userData): void
    {
        // Similar validation and update logic
        echo "User updated: $userId\n";
    }

    public function deleteUser(string $userId): void
    {
        // Delete logic
        echo "User deleted: $userId\n";
    }

    public function sendWelcomeEmail(string $email): void
    {
        // Email sending logic
        echo "Welcome email sent to: $email\n";
    }

    public function sendPasswordResetEmail(string $email): void
    {
        // Password reset email logic
        echo "Password reset email sent to: $email\n";
    }
}
