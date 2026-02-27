<?php

namespace Calisthenics;

class DontUseElse
{
    public static function calculateDiscount(float $price, bool $isVip): float
    {
        if ($isVip) {
            return $price * 0.8;
        } else {
            return $price * 0.9;
        }
    }

    public static function getStatus(int $code): string
    {
        if ($code === 200) {
            return 'Success';
        } elseif ($code === 404) {
            return 'Not Found';
        } else {
            return 'Error';
        }
    }
}
