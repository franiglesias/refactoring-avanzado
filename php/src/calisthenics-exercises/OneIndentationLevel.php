<?php

namespace Calisthenics;

class OneIndentationLevel
{
    /**
     * @param array<array{id: string, items: array<array{price: float}>, customer: array{isVip: bool}}> $orders
     */
    public static function processOrdersWithDiscounts(array $orders): float
    {
        $total = 0;
        foreach ($orders as $order) {
            if (isset($order['items']) && count($order['items']) > 0) {
                foreach ($order['items'] as $item) {
                    if (isset($order['customer']) && $order['customer']['isVip']) {
                        if ($item['price'] > 100) {
                            $total += $item['price'] * 0.8; // gran descuento VIP
                        } else {
                            $total += $item['price'] * 0.9; // pequeño descuento VIP
                        }
                    } else {
                        if ($item['price'] > 100) {
                            $total += $item['price'] * 0.95; // gran descuento regular
                        } else {
                            $total += $item['price']; // sin descuento
                        }
                    }
                }
            }
        }
        return $total;
    }
}
