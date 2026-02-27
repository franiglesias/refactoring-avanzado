class Invoice:
    def __init__(
        self,
        customer_name: str,
        customer_street: str,
        customer_city: str,
        customer_zip: str,
    ):
        self.customer_zip = customer_zip
        self.customer_street = customer_street
        self.customer_city = customer_city
        self.customer_name = customer_name

    def print(self) -> str:
        return (
            f"Factura para: {self.customer_name}\n"
            f"Dirección: {self.customer_street}, {self.customer_city}, {self.customer_zip}"
        )


class ShippingLabel:
    def __init__(
        self,
        customer_name: str,
        customer_street: str,
        customer_city: str,
        customer_zip: str,
    ):
        self.customer_zip = customer_zip
        self.customer_city = customer_city
        self.customer_street = customer_street
        self.customer_name = customer_name

    def print(self) -> str:
        return (
            f"Enviar a: {self.customer_name}\n"
            f"{self.customer_street}, {self.customer_city}, {self.customer_zip}"
        )
