export class Level2 {
  constructor(private value: number) {
  }

  getValue(): number {
    return this.value
  }
}

export class Level1 {
  constructor(private next: Level2) {
  }

  getNext(): Level2 {
    return this.next
  }
}

export class Root {
  constructor(private next: Level1) {
  }

  getNext(): Level1 {
    return this.next
  }
}

export function readDeep(root: Root): number {
  return root.getNext().getNext().getValue()
}
