export function transfer(
  amount: number,
  fromIban: string,
  toIban: string,
  currency: string,
): string {
  if (!fromIban || !toIban || !currency) {
    throw new Error('Missing data')
  }
  if (amount <= 0) {
    throw new Error('Invalid amount')
  }
  if (fromIban.length < 24) {
    throw new Error('Invalid IBAN From')
  } else if (toIban.length < 24) {
    throw new Error('Invalid IBAN To')
  }
  if (fromIban === toIban) {
    throw new Error('Same account')
  }
  // simular una transferencia
  return `${amount} ${currency} from ${fromIban} to ${toIban}`
}
