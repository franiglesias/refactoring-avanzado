package com.refactoring.codesmells.couplers;

/**
 * Code Smell: Inappropriate Intimacy [Intimidad inapropiada]
 * Manager y Team están demasiado acoplados - Manager conoce demasiado sobre
 * la estructura interna de Team y Budget.
 *
 * Manager manipula directamente el presupuesto y el nombre del equipo,
 * violando el principio de encapsulación.
 */
public class InappropriateIntimacy {

    public static void main(String[] args) {
        Team team = demoInappropriateIntimacy();
        System.out.println("Demo completado con equipo: " + team);
    }

    public static class Team {
        private String name;
        private Budget budget;
        private Manager manager;

        public Team(String name, double budget) {
            this.budget = new Budget(budget);
            this.name = name;
        }

        public void assignManager(Manager m) {
            this.manager = m;
            m.assignTeam(this);
        }

        public void raiseBudget(double amount) {
            this.budget = new Budget(this.budget.getAmount() + amount);
        }

        public void rename(String newName) {
            this.name = newName;
        }

        public Budget getBudget() {
            return budget;
        }

        public String getName() {
            return name;
        }
    }

    public static class Manager {
        private String name;
        private Team team;

        public Manager(String name) {
            this.name = name;
        }

        public void assignTeam(Team t) {
            if (this.team != null) {
                throw new IllegalStateException("Team already assigned");
            }
            this.team = t;
        }

        public void raiseTeamBudget(double amount) {
            if (this.team != null) {
                this.team.raiseBudget(amount);
            }
        }

        public void renameTeam(String newName) {
            if (this.team != null) {
                this.team.rename(newName);
            }
        }

        public String getName() {
            return name;
        }
    }

    public static class Budget {
        private double amount;

        public Budget(double amount) {
            if (amount < 900) {
                throw new IllegalArgumentException("Budget too low");
            }
            if (amount > 3000) {
                throw new IllegalArgumentException("Budget too high");
            }
            this.amount = amount;
        }

        public void raise(double amount) {
            if (this.amount + amount > 3000) {
                throw new IllegalArgumentException("Budget exceeded");
            }
            this.amount += amount;
        }

        public double getAmount() {
            return amount;
        }
    }

    public static Team demoInappropriateIntimacy() {
        Team t = new Team("Core", 1000);
        Manager m = new Manager("Alice");
        m.assignTeam(new Team("Frontend", 2000));
        t.assignManager(m);
        m.raiseTeamBudget(200);
        m.renameTeam("Platform");
        return t;
    }
}
