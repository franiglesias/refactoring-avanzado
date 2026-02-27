<?php

namespace Calisthenics;

class ShoppingCart
{
    /** @var array<string> */
    private array $items = [];

    /**
     * @param array<string> $items
     */
    public function __construct(array $items = [])
    {
        $this->items = $items;
    }

    public function addItem(string $item): void
    {
        $this->items[] = $item;
    }

    /**
     * @return array<string>
     */
    public function getItems(): array
    {
        return $this->items;
    }

    public function getTotal(): int
    {
        return count($this->items);
    }
}
