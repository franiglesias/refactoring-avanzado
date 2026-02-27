export class User {
  constructor(
    public readonly id: string,
    public readonly name: string,
    public readonly email: string,
    public readonly firstName?: string,
    public readonly lastName?: string,
  ) {
  }
}

// --- Consumers of User.name ---

export function formatGreeting(user: User): string {
  return `Hello, ${user.name}!`
}

export function formatEmailHeader(user: User): string {
  return `From: ${user.name} <${user.email}>`
}

export function formatDisplayName(user: User): string {
  return `${user.name} (${user.id})`
}

export function buildUserSummary(users: User[]): string {
  return users.map((u) => `- ${u.name}`).join('\n')
}
