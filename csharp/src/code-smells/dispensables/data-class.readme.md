# Data Class

Clase de datos.

## Definición

Una clase de datos es aquella que solo contiene campos y métodos para acceder a ellos (getters/setters), sin poseer lógica de negocio propia. Esto suele derivar en modelos de dominio anémicos, donde el comportamiento está disperso en otros servicios o clases que manipulan estos datos.

## Ejemplo

`UserRecord` es una clase que solo almacena datos, mientras que la validación y la lógica de creación residen en `UserService`.

```typescript
export class UserRecord {
  constructor(
    public id: string,
    public name: string,
    public email: string,
    public createdAt: Date,
  ) {
  }
}

class UserService {
  createUser(name: string, email: string): UserRecord {
    if (!email.includes('@')) {
      throw new Error('Invalid email')
    }

    return new UserRecord(uuidv4(), name, email, new Date())
  }

  updateUserEmail(user: UserRecord, newEmail: string): void {
    if (!newEmail.includes('@')) {
      throw new Error('Invalid email')
    }
    user.email = newEmail
  }
}

class UserReportGenerator {
  generateUserSummary(user: UserRecord): string {
    return `User ${user.name} (${user.email}) created on ${user.createdAt.toLocaleDateString()}`
  }
}

// Example usage orchestrating behavior in separate services rather than the data class itself
export function demoDataClass(): string {
  const service = new UserService()
  const report = new UserReportGenerator()
  const user = service.createUser('Lina', 'lina@example.com')
  service.updateUserEmail(user, 'lina+news@example.com')
  return report.generateUserSummary(user)
}
```

## Ejercicio

Implementa reglas de dominio adicionales, como requerir verificación de email o restringir el registro a ciertos dominios (ej. `company.com`).

## Problemas que encontrarás

Tendrás que modificar múltiples servicios y lugares que manipulan `UserRecord`. Esto demuestra cómo separar el comportamiento de los datos provoca que cambios simples se dispersen ampliamente por el código (Shotgun Surgery).
