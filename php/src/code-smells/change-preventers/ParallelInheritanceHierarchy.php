<?php

namespace CodeSmells\ChangePreventers;

// Vehicle hierarchy
abstract class Vehicle
{
    abstract public function getType(): string;
}

class Car extends Vehicle
{
    public function getType(): string
    {
        return "Car";
    }
}

class Truck extends Vehicle
{
    public function getType(): string
    {
        return "Truck";
    }
}

// Parallel Driver hierarchy
abstract class Driver
{
    abstract public function drive(): string;
}

class CarDriver extends Driver
{
    public function drive(): string
    {
        return "Driving a car";
    }
}

class TruckDriver extends Driver
{
    public function drive(): string
    {
        return "Driving a truck";
    }
}
