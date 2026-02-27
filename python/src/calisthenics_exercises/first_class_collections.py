def calculate_statistics(numbers: list[float]) -> dict:
    """Operaciones sobre una colección sin envolverla en una clase"""
    if not numbers:
        return {'sum': 0, 'average': 0, 'max': 0, 'min': 0}

    total = sum(numbers)
    average = total / len(numbers)
    maximum = max(numbers)
    minimum = min(numbers)

    return {
        'sum': total,
        'average': average,
        'max': maximum,
        'min': minimum,
        'count': len(numbers)
    }
