<?php

namespace CodeSmells\OopAbusers;

class EmailNotifier
{
    public function sendEmail(string $to, string $subject, string $body): void
    {
        echo "Sending email to $to: $subject\n";
    }
}

class SMSNotifier
{
    public function sendSMS(string $phoneNumber, string $message): void
    {
        echo "Sending SMS to $phoneNumber: $message\n";
    }
}

class NotificationService
{
    public function __construct(
        private readonly EmailNotifier $emailNotifier,
        private readonly SMSNotifier $smsNotifier
    ) {}

    public function notifyByEmail(string $to, string $subject, string $body): void
    {
        $this->emailNotifier->sendEmail($to, $subject, $body);
    }

    public function notifyBySMS(string $phoneNumber, string $message): void
    {
        $this->smsNotifier->sendSMS($phoneNumber, $message);
    }
}
