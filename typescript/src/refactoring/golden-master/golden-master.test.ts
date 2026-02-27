import {describe, expect, it} from 'vitest'
import {generateOrder, ReceiptPrinter} from './golden-master'

class ReceiptPrinterWithoutDiscountForTest extends ReceiptPrinter {
  protected getCurrentDate() {
    return new Date(2022, 1, 1)
  }

  protected discount() {
    return 0;
  }
}

class ReceiptPrinterWithDiscountForTest extends ReceiptPrinter {
  protected getCurrentDate() {
    return new Date(2022, 1, 1)
  }

  protected discount() {
    return 0.05;
  }
}


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
          const pedido = generateOrder('ORD-' + counter.toString(), customer, item, quantity)
          expect(new ReceiptPrinterWithoutDiscountForTest().print(pedido)).matchSnapshot()
        })
      })
    })
  })
})
