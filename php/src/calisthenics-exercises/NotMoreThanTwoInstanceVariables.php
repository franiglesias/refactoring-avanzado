<?php

namespace Calisthenics;

class NotMoreThanTwoInstanceVariables
{
    public function __construct(
        public readonly string $name,
        public readonly string $email,
        public readonly string $phone,
        public readonly string $address
    ) {}
}
