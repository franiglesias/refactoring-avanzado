<?php

namespace CodeSmells\Couplers;

class Person
{
    public string $name;
    public int $age;
    private float $salary;

    public function __construct(string $name, int $age, float $salary)
    {
        $this->name = $name;
        $this->age = $age;
        $this->salary = $salary;
    }

    public function getSalary(): float
    {
        return $this->salary;
    }
}

class TaxCalculator
{
    public function calculateTax(Person $person): float
    {
        // Accessing too much of Person's internal data
        $baseTax = $person->getSalary() * 0.2;

        if ($person->age > 65) {
            $baseTax *= 0.8;
        }

        return $baseTax;
    }
}
