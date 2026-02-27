<?php

namespace CodeSmells\Couplers;

class AddressData
{
    public function __construct(
        public readonly string $street,
        public readonly string $city
    ) {}

    public function getCity(): string
    {
        return $this->city;
    }
}

class CustomerData
{
    public function __construct(
        public readonly string $name,
        public readonly AddressData $address
    ) {}

    public function getAddress(): AddressData
    {
        return $this->address;
    }
}

class OrderData
{
    public function __construct(
        public readonly CustomerData $customer
    ) {}

    public function getCustomer(): CustomerData
    {
        return $this->customer;
    }

    public function getCustomerCity(): string
    {
        return $this->customer->getAddress()->getCity();
    }
}
