package calisthenics_exercises

import "fmt"

// Regla 6: No usar getters y setters (ni propiedades públicas)
// Exponer la estructura interna de los objetos genera acoplamiento y dificulta
// la evolución del código.

// Ejercicio: Elimina los getters y setters, moviendo el comportamiento al objeto.

type BankAccount struct {
	balance float64
}

func NewBankAccount(initialBalance float64) *BankAccount {
	return &BankAccount{balance: initialBalance}
}

func (b *BankAccount) GetBalance() float64 {
	return b.balance
}

func (b *BankAccount) SetBalance(value float64) error {
	if value < 0 {
		return fmt.Errorf("Negative")
	}
	b.balance = value
	return nil
}

func (b *BankAccount) Transfer(amount float64, to *BankAccount) error {
	if err := b.SetBalance(b.GetBalance() - amount); err != nil {
		return err
	}
	if err := to.SetBalance(to.GetBalance() + amount); err != nil {
		return err
	}
	return nil
}

// Ejemplo de uso con mal olor:

func Pay(account *BankAccount, amount float64) error {
	return account.SetBalance(account.GetBalance() - amount) // lógica externa usando getter/setter
}
