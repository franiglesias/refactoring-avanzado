# Refused Bequest

Herencia rechazada.

## Definición

Este smell aparece cuando una subclase hereda de una clase base, pero ignora o sobrescribe con excepciones/operaciones vacías gran parte de lo que hereda. Esto sugiere que la relación de herencia no es adecuada ("is-a") o que la jerarquía de clases necesita ser replanteada, posiblemente usando composición o extrayendo una interfaz más pequeña.

## Ejemplo

`ReadOnlyController` hereda la interfaz `Controller` pero no puede (o no debe) implementar comportamientos que modifiquen el estado, dejando métodos vacíos.

```php
<?php

declare(strict_types=1);

namespace CodeSmells\OopAbusers;

interface Resettable
{
    public function reset(): void;
}

interface Controller
{
    public function start(): void;
    public function stop(): void;
}

class BaseController implements Controller, Resettable
{
    public function start(): void
    {
        echo "starting\n";
    }

    public function stop(): void
    {
        echo "stopping\n";
    }

    public function reset(): void
    {
        echo "resetting\n";
    }
}

class ReadOnlyController implements Controller
{
    public function start(): void
    {
        // Implementación vacía - herencia rechazada
    }

    public function stop(): void
    {
        // Implementación vacía - herencia rechazada
    }
}

function demoRefusedBequest(bool $readonly): void
{
    $controller = $readonly ? new ReadOnlyController() : new BaseController();
    $controller->start();
    $controller->stop();
}
```

## Ejercicio

Añade un método de ciclo de vida `pause` a la interfaz `Controller` y haz que `start` y `stop` sean obligatorios con lógica real.

## Problemas que encontrarás

`ReadOnlyController` se verá forzado a implementar métodos que no tienen sentido para su propósito, lo que te obligará a lanzar excepciones o dejar implementaciones vacías que violan el Principio de Sustitución de Liskov.
