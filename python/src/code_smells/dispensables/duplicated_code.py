def calculate_employee_salary(base_salary: float, years_of_service: int) -> float:
    """Calcular salario de empleado con bonificación por antigüedad"""
    if years_of_service < 1:
        bonus = 0
    elif years_of_service < 3:
        bonus = base_salary * 0.05
    elif years_of_service < 5:
        bonus = base_salary * 0.10
    else:
        bonus = base_salary * 0.15

    total = base_salary + bonus
    return round(total, 2)


def calculate_contractor_payment(hourly_rate: float, hours_worked: float, years_of_service: int) -> float:
    """Calcular pago de contratista con bonificación por antigüedad"""
    base_payment = hourly_rate * hours_worked

    # Código duplicado de bonificación
    if years_of_service < 1:
        bonus = 0
    elif years_of_service < 3:
        bonus = base_payment * 0.05
    elif years_of_service < 5:
        bonus = base_payment * 0.10
    else:
        bonus = base_payment * 0.15

    total = base_payment + bonus
    return round(total, 2)
