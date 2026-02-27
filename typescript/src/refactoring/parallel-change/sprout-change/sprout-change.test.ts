import {describe, expect, it} from 'vitest'
import {calculateTotal, CartItem, Region} from './sprout-change'

describe('Sprout Change', () => {
  function executeSubject(cart: CartItem[], region: Region) {
    return calculateTotal(cart, region)
  }

  const cart: CartItem[] = [
    {id: 'p1', price: 10, qty: 2, category: 'general'},
    {id: 'b1', price: 20, qty: 1, category: 'books'},
    {id: 'f1', price: 15, qty: 4, category: 'food'},
  ]

  it('should calculate the total for EU', () => {
    expect(executeSubject(cart, 'EU')).toEqual(104)
  })

  it('should calculate the total for US', () => {
    expect(executeSubject(cart, 'US')).toEqual(107)
  })
})
