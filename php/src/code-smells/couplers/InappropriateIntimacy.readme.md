# Inappropriate Intimacy

Intimidad inapropiada.

## Definición

Dos clases tienen intimidad inapropiada cuando pueden acceder y manipular el estado interno de la otra.

## Ejemplo

Team y Manager exponen y modifican el estado interno del otro, creando un acoplamiento fuerte y diseños frágiles.

```php
<?php

declare(strict_types=1);

namespace CodeSmells\Couplers;

class Team
{
    private string $name;
    private Budget $budget;
    private ?Manager $manager = null;

    public function __construct(string $name, int $budget)
    {
        $this->budget = new Budget($budget);
        $this->name = $name;
    }

    public function assignManager(Manager $m): void
    {
        $this->manager = $m;
        $m->assignTeam($this);
    }

    public function raiseBudget(int $amount): void
    {
        $this->budget = new Budget($this->budget->amount + $amount);
    }

    public function rename(string $newName): void
    {
        $this->name = $newName;
    }
}

class Manager
{
    private string $name;
    private ?Team $team = null;

    public function __construct(string $name)
    {
        $this->name = $name;
    }

    public function assignTeam(Team $t): void
    {
        if ($this->team) {
            throw new \RuntimeException('Team already assigned');
        }
        $this->team = $t;
    }

    public function raiseTeamBudget(int $amount): void
    {
        if ($this->team) {
            $this->team->raiseBudget($amount);
        }
    }

    public function renameTeam(string $newName): void
    {
        if ($this->team) {
            $this->team->rename($newName);
        }
    }
}

class Budget
{
    public int $amount;

    public function __construct(int $amount)
    {
        if ($amount < 900) {
            throw new \InvalidArgumentException('Budget too low');
        }
        if ($amount > 3000) {
            throw new \InvalidArgumentException('Budget too high');
        }
        $this->amount = $amount;
    }

    public function raise(int $amount): void
    {
        if ($this->amount + $amount > 3000) {
            throw new \RuntimeException('Budget exceeded');
        }
        $this->amount += $amount;
    }
}

function demoInappropriateIntimacy(): Team
{
    $t = new Team('Core', 1000);
    $m = new Manager('Alice');
    $m->assignTeam(new Team('Frontend', 2000));
    $t->assignManager($m);
    $m->raiseTeamBudget(200);
    $m->renameTeam('Platform');
    return $t;
}
```

## Ejercicio

Añade una traza de auditoría cuando cambien los presupuestos y aplica reglas de presupuesto mínimo.

## Problemas que encontrarás

Como Team y Manager tocan libremente los campos del otro, tendrás que esparcir comprobaciones y registros en muchos lugares, aumentando el acoplamiento y las regresiones.
