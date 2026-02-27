class Address {
  private readonly street: string;
  private readonly city: string;

  constructor(street: string, city: string) {
    this.street = street;
    this.city = city;
  }

  getCity(): string {
    return this.city
  }
}

class Customer {
  private readonly name: string;
  private readonly address: Address;

  constructor(name: string, address: Address) {
    this.name = name;
    this.address = address;
  }

  getAddress(): Address {
    return this.address
  }
}

class Order {
  private readonly customer: Customer;

  constructor(customer: Customer) {
    this.customer = customer;
  }

  getCustomer(): Customer {
    return this.customer
  }
}

// Ejemplo de uso:

const order = new Order(new Customer('John Doe', new Address('Elm Street', 'Madrid')))
// Por ejemplo, para seleccionar los pedidos por localidad de destino y asignar la ruta de transporte
const destination = order.getCustomer().getAddress().getCity()

