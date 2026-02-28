package calisthenics_exercises

import "fmt"

type C struct {
	u string
	p string
	s string
	e string
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
