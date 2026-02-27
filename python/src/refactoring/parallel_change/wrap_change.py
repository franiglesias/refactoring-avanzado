# Dependencia externa que NO podemos modificar (simulada)
# En el mundo real, esto podría ser una librería de terceros o un servicio legacy
class LegacyEmailService:
    def send_email(self, to: str, subject: str, body: str) -> str:
        """Implementación rígida y limitada
        Problema: no valida emails, no soporta retry, no tiene logging, etc.
        """
        return f"Email sent to {to}, subject: {subject}, body: {body}"


# Código de nuestra aplicación que usa directamente el servicio legacy
legacy_service = LegacyEmailService()


def notify_welcome(user_email: str) -> str:
    return legacy_service.send_email(user_email, 'Welcome!', 'Thanks for joining our app.')


def notify_password_reset(user_email: str) -> str:
    return legacy_service.send_email(user_email, 'Reset your password', 'Click the link to reset...')


def notify_order_confirmation(user_email: str, order_id: str) -> str:
    return legacy_service.send_email(
        user_email,
        'Order Confirmation',
        f'Your order {order_id} has been confirmed.'
    )
