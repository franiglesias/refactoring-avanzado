<?php

namespace CodeSmells\OopAbusers;

class SwitchStatements
{
    public function calculateShipping(string $shippingType, float $weight): float
    {
        switch ($shippingType) {
            case 'standard':
                return $weight * 1.5;
            case 'express':
                return $weight * 3.0;
            case 'overnight':
                return $weight * 5.0;
            default:
                throw new \InvalidArgumentException("Unknown shipping type: $shippingType");
        }
    }

    public function getShippingTime(string $shippingType): int
    {
        switch ($shippingType) {
            case 'standard':
                return 5;
            case 'express':
                return 2;
            case 'overnight':
                return 1;
            default:
                throw new \InvalidArgumentException("Unknown shipping type: $shippingType");
        }
    }
}
