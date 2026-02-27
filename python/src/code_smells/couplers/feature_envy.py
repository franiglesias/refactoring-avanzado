class Product:
    def __init__(self, name: str, base_price: float, discount_rate: float):
        self.name = name
        self.base_price = base_price
        self.discount_rate = discount_rate


class Order:
    def calculate_product_total(self, product: Product, quantity: int) -> float:
        """Este método está más interesado en Product que en Order (Feature Envy)"""
        discount = product.base_price * product.discount_rate
        final_price = product.base_price - discount
        return final_price * quantity
