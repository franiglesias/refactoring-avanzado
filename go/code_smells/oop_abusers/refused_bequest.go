package oop_abusers

import "fmt"

// Code smell: Refused Bequest [Herencia rechazada].
// ReadOnlyController implementa la interfaz Controller pero no puede (o no debe)
// implementar comportamientos que modifiquen el estado, dejando métodos vacíos.

// Ejercicio: Añade un método de ciclo de vida Pause a la interfaz Controller
// y haz que Start y Stop sean obligatorios con lógica real.

// ReadOnlyController se verá forzado a implementar métodos que no tienen sentido
// para su propósito, lo que te obligará a lanzar excepciones o dejar implementaciones
// vacías que violan el Principio de Sustitución de Liskov.

type Controller interface {
	Start()
	Stop()
}

type Resettable interface {
	Reset()
}

type BaseController struct{}

func (b *BaseController) Start() {
	fmt.Println("starting")
}

func (b *BaseController) Stop() {
	fmt.Println("stopping")
}

func (b *BaseController) Reset() {
	fmt.Println("resetting")
}

type ReadOnlyController struct{}

func (r *ReadOnlyController) Start() {
	// Método vacío - refused bequest
}

func (r *ReadOnlyController) Stop() {
	// Método vacío - refused bequest
}

func DemoRefusedBequest(readonly bool) {
	var controller Controller
	if readonly {
		controller = &ReadOnlyController{}
	} else {
		controller = &BaseController{}
	}
	controller.Start()
	controller.Stop()
}
