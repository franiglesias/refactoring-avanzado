package calisthenics_exercises

// Regla 9: No más de un punto por línea
// Los objetos deben ser tratados como cajas negras y no depender de nuestro
// conocimiento previo de su estructura interna.

// Ejercicio: Elimina el acoplamiento a la estructura interna de los objetos.

type Address struct {
	street string
	city   string
}

func (a *Address) GetCity() string {
	return a.city
}

type Customer struct {
	name    string
	address *Address
}

func (c *Customer) GetAddress() *Address {
	return c.address
}

type Order struct {
	customer *Customer
}

func (o *Order) GetCustomer() *Customer {
	return o.customer
}

// Ejemplo de uso con violación de la regla:
// Para seleccionar los pedidos por localidad de destino y asignar la ruta de transporte

func GetDestination(order *Order) string {
	return order.GetCustomer().GetAddress().GetCity()
}
