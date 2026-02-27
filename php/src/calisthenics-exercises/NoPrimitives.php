<?php

namespace Calisthenics;

class NoPrimitives
{
    public function __construct(
        public readonly string $name,
        public readonly string $email,
        public readonly int $age
    ) {}

    public function getName(): string
    {
        return $this->name;
    }

    public function getEmail(): string
    {
        return $this->email;
    }

    public function getAge(): int
    {
        return $this->age;
    }
}
