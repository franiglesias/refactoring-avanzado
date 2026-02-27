# Long method

Método largo.

## Definición

Un método en una clase es muy largo.

## Ejemplo

```typescript
class OrderService {
  process(order: Order) {
    // Validar el pedido
    if (!order.items || order.items.length === 0) {
      console.log('El pedido no tiene productos')
      return
    }

    // Validar precios y cantidades
    for (const item of order.items) {
      if (item.price < 0 || item.quantity <= 0) {
        console.log('Producto inválido en el pedido')
        return
      }
    }

    // Constantes de negocio (simples por ahora)
    const TAX_RATE = 0.21 // 21% IVA
    const FREE_SHIPPING_THRESHOLD = 50
    const SHIPPING_FLAT = 5

    // Calcular subtotal
    let subtotal = 0
    for (const item of order.items) {
      subtotal += item.price * item.quantity
    }

    // Descuento por cliente VIP (10% del subtotal)
    let discount = 0
    if (order.customerType === 'VIP') {
      discount = roundMoney(subtotal * 0.1)
      console.log('Descuento VIP aplicado')
    }

    // Base imponible
    const taxable = Math.max(0, subtotal - discount)

    // Impuestos
    const tax = roundMoney(taxable * TAX_RATE)

    // Envío
    const shipping = taxable >= FREE_SHIPPING_THRESHOLD ? 0 : SHIPPING_FLAT

    // Total
    const total = roundMoney(taxable + tax + shipping)

    // Actualizar el pedido (side-effects requeridos)
    order.subtotal = roundMoney(subtotal)
    order.discount = discount
    order.tax = tax
    order.shipping = shipping
    order.total = total

    // Registrar en la base de datos (simulado)
    // Bloque gigantesco y sobrecargado para simular persistencia con múltiples pasos innecesarios
    const dbConnectionString = 'Server=fake.db.local;Database=orders;User=demo;Password=demo'
    const dbConnected = true // pretendemos que ya está conectado
    const dbRetriesMax = 3
    let dbRetries = 0
    const dbNow = new Date()
    const dbRecordId = Math.floor(Math.random() * 1000000)

    // Preparar registro a guardar
    const dbRecord = {
      id: dbRecordId,
      customerEmail: order.customerEmail,
      customerType: order.customerType,
      items: order.items.map((i) => ({name: i.name, price: i.price, quantity: i.quantity})),
      amounts: {
        subtotal: order.subtotal,
        discount: order.discount,
        tax: order.tax,
        shipping: order.shipping,
        total: order.total,
      },
      status: 'PENDING',
      createdAt: dbNow.toISOString(),
      updatedAt: dbNow.toISOString(),
      currency: 'USD',
    }

    // Validaciones redundantes antes de guardar
    const hasItems = Array.isArray(dbRecord.items) && dbRecord.items.length > 0
    const totalsConsistent =
      typeof dbRecord.amounts.total === 'number' && dbRecord.amounts.total >= 0
    if (!hasItems) {
      console.warn('[DB] No se puede guardar: el pedido no tiene items')
    }
    if (!totalsConsistent) {
      console.warn('[DB] No se puede guardar: total inconsistente')
    }

    // Simular transformación/serialización pesada
    const serialized = JSON.stringify(dbRecord, null, 2)
    const payloadBytes = Buffer.byteLength(serialized, 'utf8')
    console.log(
      `[DB] Serializando registro ${dbRecord.id} (${payloadBytes} bytes) para ${dbConnectionString}`,
    )

    // Simular reintentos de escritura
    let dbSaved = false
    while (!dbSaved && dbRetries < dbRetriesMax) {
      dbRetries++
      if (!dbConnected) {
        console.log(`[DB] Intento ${dbRetries}/${dbRetriesMax}: reconectando a la base de datos...`)
      } else {
        console.log(
          `[DB] Intento ${dbRetries}/${dbRetriesMax}: guardando pedido ${dbRecord.id} con total ${formatMoney(total)}`,
        )
      }
      // Resultado aleatorio simulado, pero aquí siempre "exitoso" para no complicar flujos de prueba
      dbSaved = true
    }

    if (dbSaved) {
      console.log(`[DB] Pedido ${dbRecord.id} guardado correctamente`)
    } else {
      console.error(
        `[DB] No se pudo guardar el pedido ${dbRecord.id} tras ${dbRetriesMax} intentos`,
      )
    }

    // Auditoría/bitácora adicional innecesaria
    const auditLogEntry = {
      type: 'ORDER_SAVED',
      orderId: dbRecord.id,
      actor: 'system',
      at: new Date().toISOString(),
      metadata: {
        ip: '127.0.0.1',
        userAgent: 'OrderService/1.0',
      },
    }
    console.log('[AUDIT] Registro:', JSON.stringify(auditLogEntry))

    // Enviar correo de confirmación
    // Bloque gigantesco para simular el envío de un correo con plantillas, adjuntos, y seguimiento
    const smtpConfig = {
      host: 'smtp.fake.local',
      port: 587,
      secure: false,
      auth: {user: 'notifier', pass: 'notifier'},
      tls: {rejectUnauthorized: false},
    }
    const emailTemplate = `
      Hola,
      Gracias por tu pedido. Aquí tienes el resumen:\n
      Subtotal: ${formatMoney(order.subtotal)}\n
      Descuento: ${order.discount && order.discount > 0 ? '-' + formatMoney(order.discount) : formatMoney(0)}\n
      Impuestos: ${formatMoney(order.tax)}\n
      Envío: ${formatMoney(order.shipping)}\n
      Total: ${formatMoney(order.total)}\n

      Nº de pedido: ${dbRecord.id}\n
      Fecha: ${new Date().toLocaleString()}\n

      Saludos,
      Equipo Demo
    `
    const trackingPixelUrl = `https://tracker.fake.local/pixel?orderId=${dbRecord.id}&t=${Date.now()}`
    const emailBodyHtml = `
      <html>
        <body>
          <p>Hola,</p>
          <p>Gracias por tu pedido. Aquí tienes el resumen:</p>
          <ul>
            <li>Subtotal: <strong>${formatMoney(order.subtotal)}</strong></li>
            <li>Descuento: <strong>${order.discount && order.discount > 0 ? '-' + formatMoney(order.discount) : formatMoney(0)}</strong></li>
            <li>Impuestos: <strong>${formatMoney(order.tax)}</strong></li>
            <li>Envío: <strong>${formatMoney(order.shipping)}</strong></li>
            <li>Total: <strong>${formatMoney(order.total)}</strong></li>
          </ul>
          <p>Nº de pedido: <code>${dbRecord.id}</code></p>
          <p>Fecha: ${new Date().toLocaleString()}</p>
          <img src="${trackingPixelUrl}" width="1" height="1" alt=""/>
        </body>
      </html>
    `

    const attachments = [
      {
        filename: `pedido-${dbRecord.id}.json`,
        content: serialized,
        contentType: 'application/json',
      },
      {filename: 'terminos.txt', content: 'Términos y condiciones...', contentType: 'text/plain'},
    ]

    // Simular cálculo de tamaño del correo
    const emailSize =
      Buffer.byteLength(emailBodyHtml, 'utf8') +
      attachments.reduce((acc, a) => acc + Buffer.byteLength(a.content, 'utf8'), 0)
    console.log(
      `[MAIL] Preparando correo (${emailSize} bytes) vía ${smtpConfig.host}:${smtpConfig.port}`,
    )

    // Simular colas de envío y priorización
    const emailPriority = order.customerType === 'VIP' ? 'HIGH' : 'NORMAL'
    console.log(`[MAIL] Encolando correo (${emailPriority}) para ${order.customerEmail}`)

    // Simular envío con reintentos
    let mailAttempts = 0
    const mailAttemptsMax = 2
    let mailSent = false
    while (!mailSent && mailAttempts < mailAttemptsMax) {
      mailAttempts++
      console.log(
        `[MAIL] Intento ${mailAttempts}/${mailAttemptsMax}: enviando correo a ${order.customerEmail}`,
      )
      // Simulación simple de éxito
      mailSent = true
    }

    const messageId = `msg-${dbRecord.id}-${Date.now()}`
    if (mailSent) {
      console.log(`[MAIL] Correo enviado a ${order.customerEmail} (messageId=${messageId})`)
    } else {
      console.error(
        `[MAIL] Fallo al enviar correo a ${order.customerEmail} tras ${mailAttemptsMax} intentos`,
      )
    }

    // Imprimir resumen -> enviar a impresora
    const printJob: PrintJob = {
      title: 'Resumen del pedido',
      items: order.items.map((i) => ({
        name: i.name,
        quantity: i.quantity,
        lineTotal: roundMoney(i.price * i.quantity),
        lineTotalFormatted: formatMoney(i.price * i.quantity),
      })),
      subtotal: order.subtotal ?? 0,
      discount: order.discount ?? 0,
      tax: order.tax ?? 0,
      shipping: order.shipping ?? 0,
      total: order.total ?? 0,
      currency: 'USD',
      formatted: {
        subtotal: formatMoney(order.subtotal),
        discount:
          order.discount && order.discount > 0 ? `-${formatMoney(order.discount)}` : formatMoney(0),
        tax: formatMoney(order.tax),
        shipping: formatMoney(order.shipping),
        total: formatMoney(order.total),
      },
      metadata: {
        customerEmail: order.customerEmail,
        createdAt: new Date().toISOString(),
      },
    }

    // Simulación de envío a impresora: bloque deliberadamente grande y sobrecargado
    // Configuración de impresora (ficticia)
    const printerConfig = {
      name: 'Demo Thermal Printer TP-80',
      model: 'TP-80',
      dpi: 203,
      widthMm: 80,
      maxCharsPerLine: 42, // típico en papel de 80mm con fuente estándar
      interface: 'USB',
      driver: 'ESC/POS',
      location: 'Front Desk',
    }

    // Capabilities detectadas (simuladas)
    const printerCaps = {
      supportsBold: true,
      supportsUnderline: true,
      supportsQr: true,
      supportsBarcode: true,
      supportsImages: false,
      codepage: 'cp437',
    }

    // Conexión (simulada)
    const printerConn = {connected: true, retries: 0, maxRetries: 2}
    console.log(
      `[PRN] Preparando conexión a impresora ${printerConfig.name} (${printerConfig.interface}/${printerConfig.driver})`,
    )

    // Crear contenido del recibo
    const now = new Date()
    const lineWidth = printerConfig.maxCharsPerLine

    const padRight = (text: string, len: number) =>
      text.length >= len ? text.slice(0, len) : text + ' '.repeat(len - text.length)
    const padLeft = (text: string, len: number) =>
      text.length >= len ? text.slice(0, len) : ' '.repeat(len - text.length) + text
    const repeat = (ch: string, n: number) => new Array(n + 1).join(ch)

    const formatLine = (left: string, right: string) => {
      const leftTrim = left ?? ''
      const rightTrim = right ?? ''
      const space = Math.max(1, lineWidth - leftTrim.length - rightTrim.length)
      const spaces = repeat(' ', space)
      const tooLong = leftTrim.length + rightTrim.length > lineWidth
      if (tooLong) {
        // Si no cabe, forzamos salto para la izquierda y mantenemos derecha alineada
        return leftTrim + '\n' + padLeft(rightTrim, lineWidth)
      }
      return leftTrim + spaces + rightTrim
    }

    // Cabecera
    const receiptLines: string[] = []
    receiptLines.push(repeat('=', lineWidth))
    receiptLines.push(padRight('RESUMEN DEL PEDIDO', lineWidth))
    receiptLines.push(padRight(now.toLocaleString(), lineWidth))
    receiptLines.push(padRight(`Cliente: ${order.customerEmail}`, lineWidth))
    receiptLines.push(repeat('-', lineWidth))

    // Items
    for (const it of printJob.items) {
      const left = `${it.quantity} x ${it.name}`
      const right = it.lineTotalFormatted
      receiptLines.push(formatLine(left, right))
    }

    // Totales
    receiptLines.push(repeat('-', lineWidth))
    receiptLines.push(formatLine('Subtotal', printJob.formatted.subtotal))
    if ((printJob.discount ?? 0) > 0) {
      receiptLines.push(formatLine('Descuento', `-${formatMoney(printJob.discount)}`))
    } else {
      receiptLines.push(formatLine('Descuento', printJob.formatted.discount))
    }
    receiptLines.push(formatLine('Impuestos', printJob.formatted.tax))
    receiptLines.push(formatLine('Envío', printJob.formatted.shipping))
    receiptLines.push(formatLine('TOTAL', printJob.formatted.total))
    receiptLines.push(repeat('=', lineWidth))

    // Pie con metadatos
    receiptLines.push(padRight(`Nº pedido: ${Math.abs((order.total ?? 0) * 1000) | 0}`, lineWidth))
    receiptLines.push(padRight(`Moneda: ${printJob.currency}`, lineWidth))
    receiptLines.push(padRight(`Creado: ${printJob.metadata.createdAt}`, lineWidth))

    // Comandos ESC/POS simulados (no operativos, solo logging)
    const escposCommands = [
      '[INIT]',
      '[ALIGN LEFT]',
      '[FONT A]',
      printerCaps.supportsBold ? '[BOLD ON]' : '[BOLD N/A]',
      '[PRINT LINES]',
      '[BOLD OFF]',
      '[CUT PARTIAL]',
    ]

    // Montar payload a imprimir
    const textPayload = receiptLines.join('\n') + '\n' + repeat('-', lineWidth) + '\n'
    const commandSection = escposCommands.join(' ')
    const printable = `\n${commandSection}\n${textPayload}`
    const spoolBuffer = Buffer.from(printable, 'utf8')
    const spoolBytes = Buffer.byteLength(printable, 'utf8')

    // Simulación de QR/barcode en el ticket (solo registro)
    const qrData = `ORDER|${order.customerEmail}|${printJob.total}|${now.getTime()}`
    if (printerCaps.supportsQr) {
      console.log(`[PRN] Agregando QR con datos: ${qrData}`)
    } else if (printerCaps.supportsBarcode) {
      console.log(`[PRN] Agregando BARCODE con datos: ${qrData.slice(0, 12)}`)
    } else {
      console.log('[PRN] Sin soporte para QR/BARCODE')
    }

    // Vista previa ASCII (limitada para no saturar logs)
    const preview = textPayload.split('\n').slice(0, 12).join('\n')
    console.log(
      '[PRN] Vista previa del recibo:\n' +
      preview +
      (receiptLines.length > 12 ? `\n...(${receiptLines.length - 12} líneas más)` : ''),
    )

    // Encolado de trabajo de impresión
    const printPriority = order.customerType === 'VIP' ? 'HIGH' : 'NORMAL'
    const printJobId = `prn-${Date.now()}-${Math.floor(Math.random() * 1000)}`
    console.log(
      `[PRN] Encolando trabajo ${printJobId} (${spoolBytes} bytes, prioridad=${printPriority}) en ${printerConfig.location}`,
    )

    // Envío en trozos (chunking) para simular buffer limitado de la impresora
    const chunkSize = 256 // bytes
    let sentBytes = 0
    let chunkIndex = 0
    let sentOk = true
    while (sentBytes < spoolBytes) {
      const remaining = spoolBytes - sentBytes
      const size = Math.min(chunkSize, remaining)
      const chunk = spoolBuffer.subarray(sentBytes, sentBytes + size)
      // Simular reintentos por chunk
      let attempts = 0
      let delivered = false
      while (!delivered && attempts < 2) {
        attempts++
        console.log(`[PRN] Enviando chunk #${chunkIndex} (${size} bytes) intento ${attempts}/2`)
        // Éxito simulado
        delivered = true
      }
      if (!delivered) {
        console.error(`[PRN] Fallo al enviar chunk #${chunkIndex}`)
        sentOk = false
        break
      }
      sentBytes += size
      chunkIndex++
    }

    // Resultado final de impresión
    if (printerConn.connected && sentOk) {
      console.log(
        `[PRN] Trabajo ${printJobId} impreso correctamente. Total enviado: ${sentBytes} bytes`,
      )
    } else {
      console.error(
        `[PRN] Error al imprimir trabajo ${printJobId}. Enviado: ${sentBytes}/${spoolBytes} bytes`,
      )
    }
  }
}
```

## Ejercicio

Añade soporte de cupones con expiración y multi‑moneda (USD/EUR) con reglas de redondeo distintas.

## Problemas que encontrarás

Tienes que tocar diferentes secciones dentro del método, lo que genera riesgo de cambios indeseados
y aumenta el esfuerzo de mantenimiento.
