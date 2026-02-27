from dataclasses import dataclass


@dataclass
class Order:
    items: list[dict]
    customer_type: str


def process_orders(orders: list[Order]) -> float:
    total_revenue = 0.0
    for order in orders:
        for item in order.items:
            price = item['price']
            quantity = item['quantity']
            subtotal = price * quantity
            if order.customer_type == 'VIP':
                discount = subtotal * 0.1
                subtotal = subtotal - discount
            total_revenue += subtotal
    return total_revenue
