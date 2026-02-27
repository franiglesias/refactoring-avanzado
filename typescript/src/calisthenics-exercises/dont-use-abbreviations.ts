export class C {
  u: string
  p: string
  s: string
  e: 'dev' | 'prod'

  constructor(u: string, p: string, s: string, e: 'dev' | 'prod') {
    this.u = u
    this.p = p
    this.s = s
    this.e = e
  }

  cnx(): string {
    return `${this.u}:${this.p}@${this.s}/${this.e}`
  }
}
