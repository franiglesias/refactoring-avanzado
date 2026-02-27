<?php

namespace CodeSmells\OopAbusers;

class Bird
{
    public function fly(): string
    {
        return "Flying...";
    }

    public function eat(): string
    {
        return "Eating...";
    }

    public function layEgg(): string
    {
        return "Laying egg...";
    }
}

class Penguin extends Bird
{
    public function fly(): string
    {
        throw new \Exception("Penguins can't fly!");
    }

    public function swim(): string
    {
        return "Swimming...";
    }
}
