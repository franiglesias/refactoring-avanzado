<?php

namespace CodeSmells\Bloaters;

class PrimitiveObsession
{
    /**
     * @param array{street: string, city: string, zipCode: string, country: string} $address
     */
    public static function formatAddress(array $address): string
    {
        return "{$address['street']}, {$address['city']}, {$address['zipCode']}, {$address['country']}";
    }

    /**
     * @param array{street: string, city: string, zipCode: string, country: string} $address
     */
    public static function validateAddress(array $address): bool
    {
        return !empty($address['street']) &&
               !empty($address['city']) &&
               !empty($address['zipCode']) &&
               !empty($address['country']);
    }
}
