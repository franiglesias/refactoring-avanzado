export class BankAccount {
  private _balance: number

  constructor(initialBalance: number = 0) {
    this._balance = initialBalance
  }

  get balance(): number {
    return this._balance
  }

  set balance(value: number) {
    if (value < 0) throw new Error('Negative')
    this._balance = value
  }

  transfer(amount: number, to: BankAccount) {
    this.balance -= amount
    to.balance += amount
  }
}

// Ejemplo de uso:

export function pay(account: BankAccount, amount: number) {
  account.balance = account.balance - amount // lógica externa usando getter/setter
}
