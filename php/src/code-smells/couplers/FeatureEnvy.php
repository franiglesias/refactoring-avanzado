<?php

namespace CodeSmells\Couplers;

class Customer
{
    public function __construct(
        public readonly string $name,
        public readonly string $phone,
        public readonly string $email
    ) {}
}

class Invoice
{
    public function __construct(
        private readonly Customer $customer,
        private readonly float $amount
    ) {}

    public function generateInvoiceText(): string
    {
        return "Invoice for: {$this->customer->name}\n" .
               "Phone: {$this->customer->phone}\n" .
               "Email: {$this->customer->email}\n" .
               "Amount: {$this->amount}";
    }
}
