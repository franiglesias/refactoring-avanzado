def register_user(name: str, email: str, age: int, country: str) -> dict:
    """Función que usa primitivos en lugar de objetos con validación"""
    # Validar email
    if '@' not in email:
        raise ValueError("Email inválido")

    # Validar edad
    if age < 18:
        raise ValueError("Debes ser mayor de edad")

    # Crear usuario
    user = {
        'name': name,
        'email': email,
        'age': age,
        'country': country,
        'status': 'active'
    }

    return user
