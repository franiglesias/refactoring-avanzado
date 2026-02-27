<?php

namespace CodeSmells\Dispensables;

class DeadCode
{
    public function activeMethod(): string
    {
        return "This method is used";
    }

    public function unusedMethod(): string
    {
        return "This method is never called";
    }

    public function calculate(int $value): int
    {
        $unusedVariable = 42;
        return $value * 2;
    }
}
