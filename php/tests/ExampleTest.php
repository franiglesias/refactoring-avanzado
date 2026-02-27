<?php

namespace Tests;

use PHPUnit\Framework\TestCase;
use Calisthenics\DontUseElse;

class ExampleTest extends TestCase
{
    public function testCalculateDiscountForVip(): void
    {
        $result = DontUseElse::calculateDiscount(100, true);
        $this->assertEquals(80, $result);
    }

    public function testCalculateDiscountForNonVip(): void
    {
        $result = DontUseElse::calculateDiscount(100, false);
        $this->assertEquals(90, $result);
    }

    public function testGetStatusSuccess(): void
    {
        $result = DontUseElse::getStatus(200);
        $this->assertEquals('Success', $result);
    }

    public function testGetStatusNotFound(): void
    {
        $result = DontUseElse::getStatus(404);
        $this->assertEquals('Not Found', $result);
    }

    public function testGetStatusError(): void
    {
        $result = DontUseElse::getStatus(500);
        $this->assertEquals('Error', $result);
    }
}
