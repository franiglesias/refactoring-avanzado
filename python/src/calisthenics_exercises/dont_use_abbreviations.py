def calc_ttl_prc(itms: list[dict]) -> float:
    """Función con abreviaciones en nombres"""
    ttl = 0.0
    for itm in itms:
        prc = itm['prc']
        qty = itm['qty']
        ttl += prc * qty
    return ttl


class Usr:
    """Clase con abreviaciones"""

    def __init__(self, usrnm: str, eml: str, ph: str):
        self.usrnm = usrnm
        self.eml = eml
        self.ph = ph

    def get_full_inf(self) -> str:
        return f"{self.usrnm} ({self.eml}, {self.ph})"
