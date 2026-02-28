# Middleman

Intermediario.

## Definición

Ocurre cuando una clase realiza una única acción: delegar el trabajo a otra clase. Si una clase existe solo como un "pasamanos" hacia otro objeto, es posible que estemos ante una capa de abstracción innecesaria que oscurece al colaborador real.

## Ejemplo

`Shop` hace poco más que delegar a `Catalog`, añadiendo una capa innecesaria que oculta dónde ocurre realmente la lógica.

```go
type Catalog struct {
	items map[string]string
}

func NewCatalog() *Catalog {
	return &Catalog{
		items: make(map[string]string),
	}
}

func (c *Catalog) Add(id, name string) {
	c.items[id] = name
}

func (c *Catalog) Find(id string) (string, bool) {
	name, ok := c.items[id]
	return name, ok
}

func (c *Catalog) List() []string {
	result := make([]string, 0, len(c.items))
	for _, name := range c.items {
		result = append(result, name)
	}
	return result
}

type Shop struct {
	catalog *Catalog
}

func NewShop(catalog *Catalog) *Shop {
	return &Shop{catalog: catalog}
}

func (s *Shop) Add(id, name string) {
	s.catalog.Add(id, name)
}

func (s *Shop) Find(id string) (string, bool) {
	return s.catalog.Find(id)
}

func (s *Shop) List() []string {
	return s.catalog.List()
}
```

## Ejercicio

Añade una funcionalidad `searchByPrefix` en `Catalog` y propágala a través de `Shop`.

## Problemas que encontrarás

Añadirás métodos a `Shop` que solo pasan a través hacia `Catalog`, fomentando la duplicación accidental y ocultando dónde vive el comportamiento real cuando necesites cambiarlo después.
