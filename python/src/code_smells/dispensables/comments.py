"""
Comments Code Smell Example

Excessive comments often indicate code that is hard to understand.
Comments should explain WHY, not WHAT. The code should be self-explanatory.

Refactoring suggestion: Extract Method, Rename Method, Introduce Assertion
"""


# Esta función suma dos números y devuelve el resultado.
# Toma el parámetro a que es un número y el parámetro b que también es un número.
# Luego usa el operador más para calcular la suma de a y b.
# Finalmente, devuelve esa suma al invocador de esta función.
def add(a: int, b: int) -> int:
    # Declara una variable llamada result que contendrá la suma de a y b
    result = a + b  # calcula la suma agregando a y b
    # Devuelve el resultado a quien haya llamado a esta función
    return result  # fin de la función


# Ejemplo de uso de este código con mal olor: llamar a una función trivial que no debería necesitar comentarios
def demo_comments_smell() -> int:
    return add(2, 3)
