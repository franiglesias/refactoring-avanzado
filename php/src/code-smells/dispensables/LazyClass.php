<?php

namespace CodeSmells\Dispensables;

class StringHelper
{
    public function toUpperCase(string $str): string
    {
        return strtoupper($str);
    }
}

class UserProcessor
{
    private StringHelper $helper;

    public function __construct()
    {
        $this->helper = new StringHelper();
    }

    public function processName(string $name): string
    {
        return $this->helper->toUpperCase($name);
    }
}
