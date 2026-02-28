package oop_abusers

import (
	"fmt"
	"strings"
)

// Code smell: Temporal Instance Variables [Variables de instancia temporales].
// PizzaOrder utiliza variables de instancia que solo son válidas entre la llamada
// a Start() y Place().

// Ejercicio: Añade una validación para que no se pueda llamar a Place()
// si no se ha añadido al menos un ingrediente.

// Te darás cuenta de que el objeto es una "máquina de estados" frágil.
// Si un cliente olvida llamar a Start() o intenta llamar a AddTopping() fuera de orden,
// el sistema puede fallar silenciosamente o requerir comprobaciones constantes
// de nulidad en cada método.

type PizzaOrder struct {
	size     *string
	toppings []string
	address  *string
}

func NewPizzaOrder() *PizzaOrder {
	return &PizzaOrder{
		toppings: []string{},
	}
}

func (p *PizzaOrder) Start(size string) {
	p.size = &size
	p.toppings = []string{}
	p.address = nil
}

func (p *PizzaOrder) AddTopping(topping string) {
	if p.size == nil {
		return
	}
	p.toppings = append(p.toppings, topping)
}

func (p *PizzaOrder) SetDeliveryAddress(address string) {
	p.address = &address
}

func (p *PizzaOrder) Place() string {
	sizeStr := "?"
	if p.size != nil {
		sizeStr = *p.size
	}
	addressStr := "UNKNOWN"
	if p.address != nil {
		addressStr = *p.address
	}
	summary := fmt.Sprintf("Pizza %s to %s with [%s]", sizeStr, addressStr, strings.Join(p.toppings, ", "))
	p.size = nil
	p.address = nil
	p.toppings = []string{}
	return summary
}

func DemoPizzaOrder() string {
	o := NewPizzaOrder()
	o.Start("L")
	o.AddTopping("pepperoni")
	o.AddTopping("mushroom")
	o.SetDeliveryAddress("123 Main St")
	return o.Place()
}
