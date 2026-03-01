# Inappropriate Intimacy

Intimidad inapropiada.

## Definición

Dos clases están demasiado acopladas, conociendo demasiado sobre la estructura interna de la otra. Una clase accede directamente a los campos privados o métodos internos de otra, violando el principio de encapsulación.

## Ejemplo

```java
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

    // Manager manipula directamente el estado interno de Team
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
}
```

## Ejercicio

Añade gestión de miembros del equipo y límites de presupuesto por categoría.

## Problemas que encontrarás

El acoplamiento entre `Manager` y `Team` te obligará a modificar ambas clases simultáneamente, aumentando la complejidad y el riesgo de errores.
