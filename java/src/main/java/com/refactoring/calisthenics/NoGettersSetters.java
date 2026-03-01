package com.refactoring.calisthenics;

/**
 * Object Calisthenics: No Getters and Setters
 * Exponer getters y setters viola la encapsulación y permite que
 * la lógica de negocio se disperse fuera de la clase.
 *
 * El método pay() manipula el balance externamente usando getter/setter,
 * cuando debería ser responsabilidad de BankAccount.
 */
public class NoGettersSetters {

    public static void main(String[] args) {
        BankAccount account = new BankAccount(1000.0);
        System.out.println("Balance inicial: " + account.getBalance());

        pay(account, 100.0);
        System.out.println("Balance después de pago: " + account.getBalance());

        BankAccount target = new BankAccount(500.0);
        account.transfer(200.0, target);
        System.out.println("Balance después de transferencia: " + account.getBalance());
        System.out.println("Balance cuenta destino: " + target.getBalance());
    }

    public static class BankAccount {
        private double balance;

        public BankAccount(double initialBalance) {
            this.balance = initialBalance;
        }

        public double getBalance() {
            return balance;
        }

        public void setBalance(double value) {
            if (value < 0) {
                throw new IllegalArgumentException("Negative balance");
            }
            this.balance = value;
        }

        public void transfer(double amount, BankAccount to) {
            this.setBalance(this.getBalance() - amount);
            to.setBalance(to.getBalance() + amount);
        }
    }

    public static void pay(BankAccount account, double amount) {
        account.setBalance(account.getBalance() - amount);
    }
}
