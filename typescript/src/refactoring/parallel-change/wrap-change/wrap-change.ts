// Dependencia externa que NO podemos modificar (simulada)
// En el mundo real, esto podría ser una librería de terceros o un servicio legacy
export class LegacyEmailService {
  sendEmail(to: string, subject: string, body: string): string {
    // Implementación rígida y limitada
    // Problema: no valida emails, no soporta retry, no tiene logging, etc.
    return `Email sent to ${to}, subject: ${subject}, body: ${body}`
  }
}

// Código de nuestra aplicación que usa directamente el servicio legacy
const legacyService = new LegacyEmailService()

export function notifyWelcome(userEmail: string): string {
  return legacyService.sendEmail(userEmail, 'Welcome!', 'Thanks for joining our app.')
}

export function notifyPasswordReset(userEmail: string): string {
  return legacyService.sendEmail(userEmail, 'Reset your password', 'Click the link to reset...')
}

export function notifyOrderConfirmation(userEmail: string, orderId: string): string {
  return legacyService.sendEmail(
    userEmail,
    'Order Confirmation',
    `Your order ${orderId} has been confirmed.`
  )
}
