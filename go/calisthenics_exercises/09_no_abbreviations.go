package calisthenics_exercises

import "fmt"

// Regla 1: No usar abreviaturas
// Los identificadores abreviados oscurecen la intención y hacen que sea más costoso
// comprender los conceptos que maneja el código y cómo se utiliza.

// Ejercicio: Expande los nombres para transmitir la intención.

type C struct {
	u string
	p string
	s string
	e string // "dev" o "prod"
}

func NewC(u, p, s, e string) *C {
	return &C{
		u: u,
		p: p,
		s: s,
		e: e,
	}
}

func (c *C) Cnx() string {
	return fmt.Sprintf("%s:%s@%s/%s", c.u, c.p, c.s, c.e)
}
