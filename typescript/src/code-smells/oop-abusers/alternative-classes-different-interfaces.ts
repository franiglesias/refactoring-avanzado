export class TextLogger {
  log(message: string): void {
    console.log(`[text] ${message}`)
  }
}

export class MessageWriter {
  write(entry: string): void {
    console.log(`[text] ${entry}`)
  }
}

export function useAltClasses(choice: 'logger' | 'writer', msg: string): void {
  if (choice === 'logger') {
    new TextLogger().log(msg)
  } else {
    new MessageWriter().write(msg)
  }
}
