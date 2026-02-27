<?php

namespace Calisthenics;

class Address
{
    public function __construct(
        public readonly string $street,
        public readonly string $city,
        public readonly string $zip
    ) {}

    public function getCity(): string
    {
        return $this->city;
    }
}

class Customer
{
    public function __construct(
        public readonly string $name,
        public readonly Address $address
    ) {}

    public function getAddress(): Address
    {
        return $this->address;
    }
}

class Order
{
    public function __construct(
        public readonly Customer $customer
    ) {}

    public function getCustomer(): Customer
    {
        return $this->customer;
    }

    public function getCustomerCity(): string
    {
        return $this->customer->getAddress()->getCity();
    }
}
