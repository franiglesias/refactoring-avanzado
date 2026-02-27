<?php

namespace Calisthenics;

class DontUseAbbreviations
{
    public static function calcTtl(int $qty, float $prc, float $dsc): float
    {
        $subtl = $qty * $prc;
        $dscAmt = $subtl * $dsc;
        return $subtl - $dscAmt;
    }

    /**
     * @param array<array{nm: string, qty: int, prc: float}> $itms
     */
    public static function prcssOrdr(array $itms): float
    {
        $ttl = 0;
        foreach ($itms as $itm) {
            $ttl += $itm['qty'] * $itm['prc'];
        }
        return $ttl;
    }
}
