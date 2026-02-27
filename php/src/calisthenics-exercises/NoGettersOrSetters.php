<?php

namespace Calisthenics;

class BankAccount
{
    private float $balance;

    public function __construct(float $initialBalance)
    {
        $this->balance = $initialBalance;
    }

    public function getBalance(): float
    {
        return $this->balance;
    }

    public function setBalance(float $balance): void
    {
        $this->balance = $balance;
    }
}
