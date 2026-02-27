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
