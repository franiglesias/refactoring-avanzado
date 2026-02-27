def validate_and_format(input_str: str) -> str:
    """Esta función hace demasiadas cosas y debería dividirse"""
    # Validar
    if not input_str:
        raise ValueError("La entrada no puede estar vacía")
    if len(input_str) < 3:
        raise ValueError("La entrada es demasiado corta")
    if len(input_str) > 100:
        raise ValueError("La entrada es demasiado larga")

    # Limpiar
    cleaned = input_str.strip()
    cleaned = cleaned.lower()
    cleaned = ''.join(c for c in cleaned if c.isalnum() or c.isspace())

    # Formatear
    words = cleaned.split()
    formatted = ' '.join(words)
    formatted = formatted.capitalize()

    return formatted
