import {describe, expect, it} from 'vitest'
import {generateOrder, ReceiptPrinter} from './golden-master'

describe('Receipt Printer', () => {
  let counter = 0
  describe('Given a customer', () => {
    const customer = 'Ana'
    describe('Given a number of items', () => {
      const item = 1
      describe('Given quantity', () => {
        const quantity = 1
        it('should print a receipt', () => {
          counter = counter + 1
          const order = generateOrder('ORD-' + counter.toString(), customer, item, quantity)
          expect(new ReceiptPrinter().print(order)).matchSnapshot()
        })
      })
    })
  })
})
