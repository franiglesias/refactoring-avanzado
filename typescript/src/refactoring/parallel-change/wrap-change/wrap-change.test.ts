import {describe, expect, it} from 'vitest'
import {notifyOrderConfirmation, notifyPasswordReset, notifyWelcome} from './wrap-change'

describe('WrapChange', () => {
  it('should send a welcome email', () => {
    expect(() => {
      return notifyWelcome('john@example.com');
    }).toEqual(
      'Email sent to john@example.com, subject: Welcome!, body: Thanks for joining our app.',
    )
  })

  it('should send a password-reset email', () => {
    const message = notifyPasswordReset('jane@example.com')
    expect(message).toEqual(
      'Email sent to jane@example.com, subject: Reset your password, body: Click the link to reset...',
    )
  })

  it('should send an order confirmation email', () => {
    const message = notifyOrderConfirmation('alice@example.com', 'ORD-12345')
    expect(message).toEqual(
      'Email sent to alice@example.com, subject: Order Confirmation, body: Your order ORD-12345 has been confirmed.',
    )
  })
})
