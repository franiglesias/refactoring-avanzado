from typing import Literal


PaymentType = Literal['CREDIT_CARD', 'DEBIT_CARD', 'PAYPAL', 'CRYPTO']


def process_payment(amount: float, payment_type: PaymentType) -> str:
    """Uso de switch/case (if-elif) en lugar de polimorfismo"""
    if payment_type == 'CREDIT_CARD':
        fee = amount * 0.03
        total = amount + fee
        return f"Procesando tarjeta de crédito por ${total:.2f} (incluye comisión de ${fee:.2f})"
    elif payment_type == 'DEBIT_CARD':
        fee = amount * 0.01
        total = amount + fee
        return f"Procesando tarjeta de débito por ${total:.2f} (incluye comisión de ${fee:.2f})"
    elif payment_type == 'PAYPAL':
        fee = amount * 0.04
        total = amount + fee
        return f"Procesando PayPal por ${total:.2f} (incluye comisión de ${fee:.2f})"
    elif payment_type == 'CRYPTO':
        fee = amount * 0.02
        total = amount + fee
        return f"Procesando criptomoneda por ${total:.2f} (incluye comisión de ${fee:.2f})"
    else:
        raise ValueError(f"Tipo de pago no soportado: {payment_type}")
