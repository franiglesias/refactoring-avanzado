<?php

namespace CodeSmells\Bloaters;

class Invoice
{
    private readonly string $customerName;
    private readonly string $customerCity;
    private readonly string $customerStreet;
    private readonly string $customerZip;

    public function __construct(
        string $customerName,
        string $customerStreet,
        string $customerCity,
        string $customerZip
    ) {
        $this->customerName = $customerName;
        $this->customerStreet = $customerStreet;
        $this->customerCity = $customerCity;
        $this->customerZip = $customerZip;
    }

    public function print(): string
    {
        return "Factura para: {$this->customerName}\n" .
               "Dirección: {$this->customerStreet}, {$this->customerCity}, {$this->customerZip}";
    }
}

class ShippingLabel
{
    private readonly string $customerName;
    private readonly string $customerStreet;
    private readonly string $customerCity;
    private readonly string $customerZip;

    public function __construct(
        string $customerName,
        string $customerStreet,
        string $customerCity,
        string $customerZip
    ) {
        $this->customerName = $customerName;
        $this->customerStreet = $customerStreet;
        $this->customerCity = $customerCity;
        $this->customerZip = $customerZip;
    }

    public function print(): string
    {
        return "Enviar a: {$this->customerName}\n" .
               "{$this->customerStreet}, {$this->customerCity}, {$this->customerZip}";
    }
}
