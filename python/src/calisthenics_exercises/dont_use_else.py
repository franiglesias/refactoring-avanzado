from typing import Literal


def shipping_cost(weight_kg: float, destination: Literal['DOMESTIC', 'INTERNATIONAL']) -> float:
    if destination == 'DOMESTIC':
        if weight_kg <= 1:
            return 5
        elif weight_kg <= 5:
            return 10
        else:
            return 20
    else:
        if weight_kg <= 1:
            return 15
        elif weight_kg <= 5:
            return 25
        else:
            return 40
