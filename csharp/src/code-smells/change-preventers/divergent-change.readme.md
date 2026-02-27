# Divergent change

## Definición

Una clase tiene múltiples razones para cambiar, lo que normalmente indica que se ocupa de muchas responsabilidades que se deberían separar en clases especialistas más pequeñas.

## Ejemplo

ProfileManager maneja validación, persistencia, exportación y envío de emails—múltiples razones para cambiar concentradas en una sola clase.

```typescript
export type User = { id: string; name: string; email: string }

export class ProfileManager {
  private store = new Map<string, User>()

  register(user: User): void {
    if (!user.name.trim()) throw new Error('invalid name')
    if (!user.email.includes('@')) throw new Error('invalid email')
    this.store.set(user.id, user)
  }

  updateEmail(id: string, newEmail: string): void {
    if (!newEmail.includes('@')) throw new Error('invalid email')
    const u = this.store.get(id)
    if (!u) throw new Error('not found')
    this.store.set(id, {...u, email: newEmail})
  }

  exportAsJson(): string {
    return JSON.stringify(Array.from(this.store.values()))
  }

  exportAsCsv(): string {
    const rows = [
      'id,name,email',
      ...Array.from(this.store.values()).map((u) => `${u.id},${u.name},${u.email}`),
    ]
    return rows.join('\n')
  }

  sendWelcomeEmail(user: User): string {
    return `Welcome ${user.name}! Sent to ${user.email}`
  }
}

export function demoDivergentChange(pm: ProfileManager, u: User): string {
  pm.register(u)
  pm.updateEmail(u.id, u.email)
  return pm.exportAsJson()
}

```

## Ejercicio

Añade un número de teléfono con validación, inclúyelo en las exportaciones y envía un SMS.

## Problemas que encontrarás

Tocarás validación, almacenamiento, exportAsJson/Csv y mensajería en un solo lugar, demostrando cómo un cambio fuerza ediciones en responsabilidades no relacionadas.

