class Product:
    """Clase con más de 2 variables de instancia"""

    def __init__(
        self,
        product_id: str,
        name: str,
        price: float,
        category: str,
        stock: int,
        supplier: str,
        description: str
    ):
        self.product_id = product_id
        self.name = name
        self.price = price
        self.category = category
        self.stock = stock
        self.supplier = supplier
        self.description = description

    def display(self) -> str:
        return (
            f"{self.name} ({self.product_id})\n"
            f"Precio: ${self.price}\n"
            f"Categoría: {self.category}\n"
            f"Stock: {self.stock}\n"
            f"Proveedor: {self.supplier}\n"
            f"Descripción: {self.description}"
        )
