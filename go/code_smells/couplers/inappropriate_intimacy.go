package couplers

import "fmt"

// Code smell: Inappropriate Intimacy [Intimidad inapropiada].
// Team y Manager exponen y modifican el estado interno del otro,
// creando un acoplamiento fuerte y diseños frágiles.

// Ejercicio: Añade una traza de auditoría cuando cambien los presupuestos y aplica reglas de presupuesto mínimo.

// Como Team y Manager tocan libremente los campos del otro, tendrás que esparcir comprobaciones
// y registros en muchos lugares, aumentando el acoplamiento y las regresiones.

type Budget struct {
	Amount float64
}

func NewBudget(amount float64) (*Budget, error) {
	if amount < 900 {
		return nil, fmt.Errorf("Budget too low")
	}
	if amount > 3000 {
		return nil, fmt.Errorf("Budget too high")
	}
	return &Budget{Amount: amount}, nil
}

func (b *Budget) Raise(amount float64) error {
	if b.Amount+amount > 3000 {
		return fmt.Errorf("Budget exceeded")
	}
	b.Amount += amount
	return nil
}

type Team struct {
	name    string
	budget  *Budget
	manager *Manager
}

func NewTeam(name string, budget float64) (*Team, error) {
	b, err := NewBudget(budget)
	if err != nil {
		return nil, err
	}
	return &Team{
		name:   name,
		budget: b,
	}, nil
}

func (t *Team) AssignManager(m *Manager) {
	t.manager = m
	m.AssignTeam(t)
}

func (t *Team) RaiseBudget(amount float64) error {
	newBudget, err := NewBudget(t.budget.Amount + amount)
	if err != nil {
		return err
	}
	t.budget = newBudget
	return nil
}

func (t *Team) Rename(newName string) {
	t.name = newName
}

type Manager struct {
	name string
	team *Team
}

func NewManager(name string) *Manager {
	return &Manager{name: name}
}

func (m *Manager) AssignTeam(t *Team) error {
	if m.team != nil {
		return fmt.Errorf("Team already assigned")
	}
	m.team = t
	return nil
}

func (m *Manager) RaiseTeamBudget(amount float64) error {
	if m.team != nil {
		return m.team.RaiseBudget(amount)
	}
	return nil
}

func (m *Manager) RenameTeam(newName string) {
	if m.team != nil {
		m.team.Rename(newName)
	}
}

func DemoInappropriateIntimacy() (*Team, error) {
	t, err := NewTeam("Core", 1000)
	if err != nil {
		return nil, err
	}
	m := NewManager("Alice")
	frontendTeam, _ := NewTeam("Frontend", 2000)
	m.AssignTeam(frontendTeam)
	t.AssignManager(m)
	m.RaiseTeamBudget(200)
	m.RenameTeam("Platform")
	return t, nil
}
