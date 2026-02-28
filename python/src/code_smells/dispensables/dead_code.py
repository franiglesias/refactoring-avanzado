"""
Dead Code Code Smell Example

Code that is never executed or used. This includes:
- Unreachable code after return statements
- Unused variables
- Unused functions
- Unused constants

Refactoring suggestion: Delete the dead code
"""

# Unused constant
THE_ANSWER_TO_EVERYTHING = 42


def format_currency(amount: float) -> str:
    """Unused function - never called"""
    return f"${amount:.2f}"


def active_function(value: int) -> int:
    if value < 0:
        return 0
        never_runs = value * -1  # Unreachable code after return
        print('This will never be printed', never_runs)

    temp = value * 2  # Unused variable

    return value + 1


def demo_dead_code() -> str:
    result = active_function(5)
    # format_currency is never called - dead code
    return format_currency(result)
