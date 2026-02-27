<?php

namespace CodeSmells\OopAbusers;

class Calculator
{
    private ?float $result = null;

    public function add(float $a, float $b): void
    {
        $this->result = $a + $b;
    }

    public function multiply(float $a, float $b): void
    {
        $this->result = $a * $b;
    }

    public function getResult(): ?float
    {
        return $this->result;
    }
}
