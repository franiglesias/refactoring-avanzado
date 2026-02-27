<?php

namespace CodeSmells\Bloaters;

class LongParameterList
{
    public function createUser(
        string $firstName,
        string $lastName,
        string $email,
        string $phone,
        string $address,
        string $city,
        string $zipCode,
        string $country
    ): void {
        echo "User created: $firstName $lastName\n";
        echo "Email: $email\n";
        echo "Phone: $phone\n";
        echo "Address: $address, $city, $zipCode, $country\n";
    }
}
