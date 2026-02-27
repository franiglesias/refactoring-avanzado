class Address:
    def __init__(self, street: str, city: str):
        self.street = street
        self.city = city


class Customer:
    def __init__(self, name: str, address: Address):
        self.name = name
        self.address = address


class Order:
    def __init__(self, customer: Customer, total: float):
        self.customer = customer
        self.total = total


def format_order_summary(order: Order) -> str:
    """Violación de la ley de Demeter: múltiples puntos por línea"""
    return f"Pedido de {order.customer.name} en {order.customer.address.city}"
