package couplers

// Code smell: Message Chains [Cadenas de mensajes].
// La navegación profunda por grafos de objetos acopla a los clientes a la estructura
// de los intermediarios y conduce a código frágil.

// Ejercicio: Inserta un nuevo Level entre Root y Level1, o reubica getValue.

// Observa cómo cada cliente que usa root.GetNext().GetNext().GetValue() debe cambiar,
// revelando cómo las cadenas de mensajes vuelven costosas refactorizaciones simples.

type Level2 struct {
	value int
}

func NewLevel2(value int) *Level2 {
	return &Level2{value: value}
}

func (l *Level2) GetValue() int {
	return l.value
}

type Level1 struct {
	next *Level2
}

func NewLevel1(next *Level2) *Level1 {
	return &Level1{next: next}
}

func (l *Level1) GetNext() *Level2 {
	return l.next
}

type Root struct {
	next *Level1
}

func NewRoot(next *Level1) *Root {
	return &Root{next: next}
}

func (r *Root) GetNext() *Level1 {
	return r.next
}

func ReadDeep(root *Root) int {
	return root.GetNext().GetNext().GetValue()
}
