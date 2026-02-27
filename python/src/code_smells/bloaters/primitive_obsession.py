class Order:
    def __init__(
        self,
        customer_name: str,
        customer_email: str,
        address: str,
        total_amount: float,
        currency: str,
    ):
        self.customer_name = customer_name
        self.customer_email = customer_email
        self.address = address
        self.total_amount = total_amount
        self.currency = currency

    def send_invoice(self):
        if '@' not in self.customer_email:
            raise ValueError('Email inválido')
        if not self.address:
            raise ValueError('No se ha indicado dirección')
        if self.total_amount <= 0:
            raise ValueError('El monto debe ser mayor que cero')
        print(f"Factura enviada a {self.customer_name} in {self.address} por {self.total_amount} {self.currency}")
