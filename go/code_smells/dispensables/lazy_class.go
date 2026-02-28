package dispensables

import "fmt"

// Code smell: Lazy Class [Clase perezosa].
// La clase ShippingLabelBuilder solo tiene un método que realiza una concatenación
// de strings simple, algo que podría resolverse con una función pura.

// Ejercicio: Reescribe el código para eliminar la necesidad de la clase ShippingLabelBuilder.

// Mantener una estructura de clase para una lógica tan simple te obliga a instanciar objetos
// innecesariamente y añade capas de abstracción que dificultan la legibilidad del código
// sin ofrecer beneficios a cambio.

type Address struct {
	Name  string
	Line1 string
	City  *string
}

type ShippingLabelBuilder struct{}

func (s *ShippingLabelBuilder) Build(a Address) string {
	cityStr := ""
	if a.City != nil {
		cityStr = ", " + *a.City
	}
	return fmt.Sprintf("%s — %s%s", a.Name, a.Line1, cityStr)
}

func PrintShippingLabel() {
	city := "New York"
	address := Address{
		Name:  "John Doe",
		Line1: "123 Main St",
		City:  &city,
	}

	labelBuilder := &ShippingLabelBuilder{}
	label := labelBuilder.Build(address)
	fmt.Println(label)
}
