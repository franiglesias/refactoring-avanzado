<?php

namespace CodeSmells\Dispensables;

class Comments
{
    /**
     * Calculate total price with discount
     *
     * @param float $price The original price
     * @param float $discount The discount percentage (0-1)
     * @return float The final price after discount
     */
    public function calculateTotal(float $price, float $discount): float
    {
        // First we multiply the price by the discount
        $discountAmount = $price * $discount;

        // Then we subtract the discount from the original price
        $finalPrice = $price - $discountAmount;

        // Return the final price
        return $finalPrice;
    }
}
