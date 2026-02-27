class BankAccount:
    """Clase con getters y setters expuestos"""

    def __init__(self, account_number: str, balance: float):
        self._account_number = account_number
        self._balance = balance

    def get_balance(self) -> float:
        return self._balance

    def set_balance(self, balance: float):
        self._balance = balance

    def get_account_number(self) -> str:
        return self._account_number


def transfer_money(from_account: BankAccount, to_account: BankAccount, amount: float):
    """Operación que manipula el estado interno de los objetos desde afuera"""
    if from_account.get_balance() >= amount:
        from_account.set_balance(from_account.get_balance() - amount)
        to_account.set_balance(to_account.get_balance() + amount)
        return True
    return False
