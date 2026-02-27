import {v4 as uuidv4} from 'uuid'

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

export function demoDataClass(): string {
  const service = new UserService()
  const report = new UserReportGenerator()
  const user = service.createUser('Lina', 'lina@example.com')
  service.updateUserEmail(user, 'lina+news@example.com')
  return report.generateUserSummary(user)
}
