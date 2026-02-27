<?php

namespace CodeSmells\ChangePreventers;

class User
{
    public function __construct(
        private string $name,
        private string $email,
        private string $password
    ) {}

    // Database operations
    public function save(): void
    {
        echo "Saving user to database...\n";
    }

    public function delete(): void
    {
        echo "Deleting user from database...\n";
    }

    // Validation
    public function validate(): bool
    {
        return !empty($this->email) && !empty($this->password);
    }

    // Email operations
    public function sendWelcomeEmail(): void
    {
        echo "Sending welcome email to {$this->email}...\n";
    }

    public function sendPasswordResetEmail(): void
    {
        echo "Sending password reset email to {$this->email}...\n";
    }

    // Business logic
    public function hashPassword(): void
    {
        $this->password = password_hash($this->password, PASSWORD_DEFAULT);
    }
}
