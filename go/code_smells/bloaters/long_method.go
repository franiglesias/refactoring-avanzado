package bloaters

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

// Code smell: Long Method [Método largo].
// El método Process tiene cientos de líneas realizando múltiples operaciones:
// validación, cálculo de precios, persistencia en BD, envío de correos, impresión de tickets.

// Ejercicio: Añade soporte de cupones con expiración y multi-moneda (USD/EUR) con reglas de redondeo distintas.

// Tienes que tocar diferentes secciones dentro del método, lo que genera riesgo de cambios indeseados
// y aumenta el esfuerzo de mantenimiento.

type OrderItem struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type Order struct {
	CustomerEmail string
	CustomerType  string // "NORMAL" o "VIP"
	Items         []OrderItem
	Subtotal      float64
	Discount      float64
	Tax           float64
	Shipping      float64
	Total         float64
}

type PrintJob struct {
	Title    string
	Items    []PrintItem
	Subtotal float64
	Discount float64
	Tax      float64
	Shipping float64
	Total    float64
	Currency string
	Formatted struct {
		Subtotal string
		Discount string
		Tax      string
		Shipping string
		Total    string
	}
	Metadata struct {
		CustomerEmail string
		CreatedAt     string
	}
}

type PrintItem struct {
	Name               string
	Quantity           int
	LineTotal          float64
	LineTotalFormatted string
}

type OrderService struct{}

func (s *OrderService) Process(order *Order) {
	// Validar el pedido
	if order.Items == nil || len(order.Items) == 0 {
		fmt.Println("El pedido no tiene productos")
		return
	}

	// Validar precios y cantidades
	for _, item := range order.Items {
		if item.Price < 0 || item.Quantity <= 0 {
			fmt.Println("Producto inválido en el pedido")
			return
		}
	}

	// Constantes de negocio (simples por ahora)
	const TAX_RATE = 0.21           // 21% IVA
	const FREE_SHIPPING_THRESHOLD = 50.0
	const SHIPPING_FLAT = 5.0

	// Calcular subtotal
	subtotal := 0.0
	for _, item := range order.Items {
		subtotal += item.Price * float64(item.Quantity)
	}

	// Descuento por cliente VIP (10% del subtotal)
	discount := 0.0
	if order.CustomerType == "VIP" {
		discount = roundMoney(subtotal * 0.1)
		fmt.Println("Descuento VIP aplicado")
	}

	// Base imponible
	taxable := math.Max(0, subtotal-discount)

	// Impuestos
	tax := roundMoney(taxable * TAX_RATE)

	// Envío
	shipping := 0.0
	if taxable >= FREE_SHIPPING_THRESHOLD {
		shipping = 0
	} else {
		shipping = SHIPPING_FLAT
	}

	// Total
	total := roundMoney(taxable + tax + shipping)

	// Actualizar el pedido (side-effects requeridos)
	order.Subtotal = roundMoney(subtotal)
	order.Discount = discount
	order.Tax = tax
	order.Shipping = shipping
	order.Total = total

	// Registrar en la base de datos (simulado)
	// Bloque gigantesco y sobrecargado para simular persistencia con múltiples pasos innecesarios
	dbConnectionString := "Server=fake.db.local;Database=orders;User=demo;Password=demo"
	dbConnected := true // pretendemos que ya está conectado
	const dbRetriesMax = 3
	dbRetries := 0
	dbNow := time.Now()
	dbRecordId := rand.Intn(1000000)

	// Preparar registro a guardar
	type DBAmounts struct {
		Subtotal float64 `json:"subtotal"`
		Discount float64 `json:"discount"`
		Tax      float64 `json:"tax"`
		Shipping float64 `json:"shipping"`
		Total    float64 `json:"total"`
	}

	type DBItem struct {
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		Quantity int     `json:"quantity"`
	}

	dbItems := make([]DBItem, len(order.Items))
	for i, item := range order.Items {
		dbItems[i] = DBItem{Name: item.Name, Price: item.Price, Quantity: item.Quantity}
	}

	dbRecord := map[string]interface{}{
		"id":            dbRecordId,
		"customerEmail": order.CustomerEmail,
		"customerType":  order.CustomerType,
		"items":         dbItems,
		"amounts": DBAmounts{
			Subtotal: order.Subtotal,
			Discount: order.Discount,
			Tax:      order.Tax,
			Shipping: order.Shipping,
			Total:    order.Total,
		},
		"status":    "PENDING",
		"createdAt": dbNow.Format(time.RFC3339),
		"updatedAt": dbNow.Format(time.RFC3339),
		"currency":  "USD",
	}

	// Validaciones redundantes antes de guardar
	hasItems := len(dbItems) > 0
	totalsConsistent := order.Total >= 0
	if !hasItems {
		fmt.Println("[DB] No se puede guardar: el pedido no tiene items")
	}
	if !totalsConsistent {
		fmt.Println("[DB] No se puede guardar: total inconsistente")
	}

	// Simular transformación/serialización pesada
	serialized, _ := json.MarshalIndent(dbRecord, "", "  ")
	payloadBytes := len(serialized)
	fmt.Printf("[DB] Serializando registro %d (%d bytes) para %s\n", dbRecordId, payloadBytes, dbConnectionString)

	// Simular reintentos de escritura
	dbSaved := false
	for !dbSaved && dbRetries < dbRetriesMax {
		dbRetries++
		if !dbConnected {
			fmt.Printf("[DB] Intento %d/%d: reconectando a la base de datos...\n", dbRetries, dbRetriesMax)
		} else {
			fmt.Printf("[DB] Intento %d/%d: guardando pedido %d con total %s\n", dbRetries, dbRetriesMax, dbRecordId, formatMoney(total))
		}
		// Resultado aleatorio simulado, pero aquí siempre "exitoso" para no complicar flujos de prueba
		dbSaved = true
	}

	if dbSaved {
		fmt.Printf("[DB] Pedido %d guardado correctamente\n", dbRecordId)
	} else {
		fmt.Printf("[DB] No se pudo guardar el pedido %d tras %d intentos\n", dbRecordId, dbRetriesMax)
	}

	// Auditoría/bitácora adicional innecesaria
	auditLogEntry := map[string]interface{}{
		"type":    "ORDER_SAVED",
		"orderId": dbRecordId,
		"actor":   "system",
		"at":      time.Now().Format(time.RFC3339),
		"metadata": map[string]string{
			"ip":        "127.0.0.1",
			"userAgent": "OrderService/1.0",
		},
	}
	auditJSON, _ := json.Marshal(auditLogEntry)
	fmt.Printf("[AUDIT] Registro: %s\n", string(auditJSON))

	// Enviar correo de confirmación
	// Bloque gigantesco para simular el envío de un correo con plantillas, adjuntos, y seguimiento
	smtpConfig := map[string]interface{}{
		"host":   "smtp.fake.local",
		"port":   587,
		"secure": false,
		"auth": map[string]string{
			"user": "notifier",
			"pass": "notifier",
		},
		"tls": map[string]bool{
			"rejectUnauthorized": false,
		},
	}

	discountStr := formatMoney(0)
	if order.Discount > 0 {
		discountStr = "-" + formatMoney(order.Discount)
	}

	emailTemplate := fmt.Sprintf(`
      Hola,
      Gracias por tu pedido. Aquí tienes el resumen:

      Subtotal: %s
      Descuento: %s
      Impuestos: %s
      Envío: %s
      Total: %s

      Nº de pedido: %d
      Fecha: %s

      Saludos,
      Equipo Demo
    `, formatMoney(order.Subtotal), discountStr, formatMoney(order.Tax), formatMoney(order.Shipping), formatMoney(order.Total), dbRecordId, time.Now().Format(time.RFC1123))

	trackingPixelUrl := fmt.Sprintf("https://tracker.fake.local/pixel?orderId=%d&t=%d", dbRecordId, time.Now().UnixMilli())

	emailBodyHtml := fmt.Sprintf(`
      <html>
        <body>
          <p>Hola,</p>
          <p>Gracias por tu pedido. Aquí tienes el resumen:</p>
          <ul>
            <li>Subtotal: <strong>%s</strong></li>
            <li>Descuento: <strong>%s</strong></li>
            <li>Impuestos: <strong>%s</strong></li>
            <li>Envío: <strong>%s</strong></li>
            <li>Total: <strong>%s</strong></li>
          </ul>
          <p>Nº de pedido: <code>%d</code></p>
          <p>Fecha: %s</p>
          <img src="%s" width="1" height="1" alt=""/>
        </body>
      </html>
    `, formatMoney(order.Subtotal), discountStr, formatMoney(order.Tax), formatMoney(order.Shipping), formatMoney(order.Total), dbRecordId, time.Now().Format(time.RFC1123), trackingPixelUrl)

	type Attachment struct {
		Filename    string
		Content     string
		ContentType string
	}

	attachments := []Attachment{
		{
			Filename:    fmt.Sprintf("pedido-%d.json", dbRecordId),
			Content:     string(serialized),
			ContentType: "application/json",
		},
		{
			Filename:    "terminos.txt",
			Content:     "Términos y condiciones...",
			ContentType: "text/plain",
		},
	}

	// Simular cálculo de tamaño del correo
	emailSize := len(emailBodyHtml)
	for _, a := range attachments {
		emailSize += len(a.Content)
	}
	fmt.Printf("[MAIL] Preparando correo (%d bytes) vía %s:%d\n", emailSize, smtpConfig["host"], smtpConfig["port"])

	// Simular colas de envío y priorización
	emailPriority := "NORMAL"
	if order.CustomerType == "VIP" {
		emailPriority = "HIGH"
	}
	fmt.Printf("[MAIL] Encolando correo (%s) para %s\n", emailPriority, order.CustomerEmail)

	// Simular envío con reintentos
	mailAttempts := 0
	const mailAttemptsMax = 2
	mailSent := false
	for !mailSent && mailAttempts < mailAttemptsMax {
		mailAttempts++
		fmt.Printf("[MAIL] Intento %d/%d: enviando correo a %s\n", mailAttempts, mailAttemptsMax, order.CustomerEmail)
		// Simulación simple de éxito
		mailSent = true
	}

	messageId := fmt.Sprintf("msg-%d-%d", dbRecordId, time.Now().UnixMilli())
	if mailSent {
		fmt.Printf("[MAIL] Correo enviado a %s (messageId=%s)\n", order.CustomerEmail, messageId)
	} else {
		fmt.Printf("[MAIL] Fallo al enviar correo a %s tras %d intentos\n", order.CustomerEmail, mailAttemptsMax)
	}

	// Imprimir resumen -> enviar a impresora
	printJob := PrintJob{
		Title:    "Resumen del pedido",
		Items:    make([]PrintItem, len(order.Items)),
		Subtotal: order.Subtotal,
		Discount: order.Discount,
		Tax:      order.Tax,
		Shipping: order.Shipping,
		Total:    order.Total,
		Currency: "USD",
	}

	for i, item := range order.Items {
		lineTotal := roundMoney(item.Price * float64(item.Quantity))
		printJob.Items[i] = PrintItem{
			Name:               item.Name,
			Quantity:           item.Quantity,
			LineTotal:          lineTotal,
			LineTotalFormatted: formatMoney(lineTotal),
		}
	}

	printJob.Formatted.Subtotal = formatMoney(order.Subtotal)
	if order.Discount > 0 {
		printJob.Formatted.Discount = "-" + formatMoney(order.Discount)
	} else {
		printJob.Formatted.Discount = formatMoney(0)
	}
	printJob.Formatted.Tax = formatMoney(order.Tax)
	printJob.Formatted.Shipping = formatMoney(order.Shipping)
	printJob.Formatted.Total = formatMoney(order.Total)
	printJob.Metadata.CustomerEmail = order.CustomerEmail
	printJob.Metadata.CreatedAt = time.Now().Format(time.RFC3339)

	// Simulación de envío a impresora: bloque deliberadamente grande y sobrecargado
	// Configuración de impresora (ficticia)
	printerConfig := map[string]interface{}{
		"name":            "Demo Thermal Printer TP-80",
		"model":           "TP-80",
		"dpi":             203,
		"widthMm":         80,
		"maxCharsPerLine": 42, // típico en papel de 80mm con fuente estándar
		"interface":       "USB",
		"driver":          "ESC/POS",
		"location":        "Front Desk",
	}

	// Capabilities detectadas (simuladas)
	printerCaps := map[string]interface{}{
		"supportsBold":      true,
		"supportsUnderline": true,
		"supportsQr":        true,
		"supportsBarcode":   true,
		"supportsImages":    false,
		"codepage":          "cp437",
	}

	// Conexión (simulada)
	printerConn := map[string]interface{}{
		"connected":  true,
		"retries":    0,
		"maxRetries": 2,
	}
	fmt.Printf("[PRN] Preparando conexión a impresora %s (%s/%s)\n", printerConfig["name"], printerConfig["interface"], printerConfig["driver"])

	// Crear contenido del recibo
	now := time.Now()
	lineWidth := printerConfig["maxCharsPerLine"].(int)

	padRight := func(text string, length int) string {
		if len(text) >= length {
			return text[:length]
		}
		return text + strings.Repeat(" ", length-len(text))
	}

	padLeft := func(text string, length int) string {
		if len(text) >= length {
			return text[:length]
		}
		return strings.Repeat(" ", length-len(text)) + text
	}

	repeat := func(ch string, n int) string {
		return strings.Repeat(ch, n)
	}

	formatLine := func(left, right string) string {
		space := lineWidth - len(left) - len(right)
		if space < 1 {
			space = 1
		}
		tooLong := len(left)+len(right) > lineWidth
		if tooLong {
			// Si no cabe, forzamos salto para la izquierda y mantenemos derecha alineada
			return left + "\n" + padLeft(right, lineWidth)
		}
		return left + strings.Repeat(" ", space) + right
	}

	// Cabecera
	var receiptLines []string
	receiptLines = append(receiptLines, repeat("=", lineWidth))
	receiptLines = append(receiptLines, padRight("RESUMEN DEL PEDIDO", lineWidth))
	receiptLines = append(receiptLines, padRight(now.Format(time.RFC1123), lineWidth))
	receiptLines = append(receiptLines, padRight(fmt.Sprintf("Cliente: %s", order.CustomerEmail), lineWidth))
	receiptLines = append(receiptLines, repeat("-", lineWidth))

	// Items
	for _, it := range printJob.Items {
		left := fmt.Sprintf("%d x %s", it.Quantity, it.Name)
		right := it.LineTotalFormatted
		receiptLines = append(receiptLines, formatLine(left, right))
	}

	// Totales
	receiptLines = append(receiptLines, repeat("-", lineWidth))
	receiptLines = append(receiptLines, formatLine("Subtotal", printJob.Formatted.Subtotal))
	if printJob.Discount > 0 {
		receiptLines = append(receiptLines, formatLine("Descuento", "-"+formatMoney(printJob.Discount)))
	} else {
		receiptLines = append(receiptLines, formatLine("Descuento", printJob.Formatted.Discount))
	}
	receiptLines = append(receiptLines, formatLine("Impuestos", printJob.Formatted.Tax))
	receiptLines = append(receiptLines, formatLine("Envío", printJob.Formatted.Shipping))
	receiptLines = append(receiptLines, formatLine("TOTAL", printJob.Formatted.Total))
	receiptLines = append(receiptLines, repeat("=", lineWidth))

	// Pie con metadatos
	receiptLines = append(receiptLines, padRight(fmt.Sprintf("Nº pedido: %d", int(math.Abs(order.Total*1000))), lineWidth))
	receiptLines = append(receiptLines, padRight(fmt.Sprintf("Moneda: %s", printJob.Currency), lineWidth))
	receiptLines = append(receiptLines, padRight(fmt.Sprintf("Creado: %s", printJob.Metadata.CreatedAt), lineWidth))

	// Comandos ESC/POS simulados (no operativos, solo logging)
	escposCommands := []string{
		"[INIT]",
		"[ALIGN LEFT]",
		"[FONT A]",
	}
	if printerCaps["supportsBold"].(bool) {
		escposCommands = append(escposCommands, "[BOLD ON]")
	} else {
		escposCommands = append(escposCommands, "[BOLD N/A]")
	}
	escposCommands = append(escposCommands, "[PRINT LINES]", "[BOLD OFF]", "[CUT PARTIAL]")

	// Montar payload a imprimir
	textPayload := strings.Join(receiptLines, "\n") + "\n" + repeat("-", lineWidth) + "\n"
	commandSection := strings.Join(escposCommands, " ")
	printable := "\n" + commandSection + "\n" + textPayload
	spoolBytes := len(printable)

	// Simulación de QR/barcode en el ticket (solo registro)
	qrData := fmt.Sprintf("ORDER|%s|%.2f|%d", order.CustomerEmail, printJob.Total, now.UnixMilli())
	if printerCaps["supportsQr"].(bool) {
		fmt.Printf("[PRN] Agregando QR con datos: %s\n", qrData)
	} else if printerCaps["supportsBarcode"].(bool) {
		fmt.Printf("[PRN] Agregando BARCODE con datos: %s\n", qrData[:min(len(qrData), 12)])
	} else {
		fmt.Println("[PRN] Sin soporte para QR/BARCODE")
	}

	// Vista previa ASCII (limitada para no saturar logs)
	previewLines := receiptLines
	if len(previewLines) > 12 {
		previewLines = previewLines[:12]
	}
	preview := strings.Join(previewLines, "\n")
	if len(receiptLines) > 12 {
		fmt.Printf("[PRN] Vista previa del recibo:\n%s\n...(%d líneas más)\n", preview, len(receiptLines)-12)
	} else {
		fmt.Printf("[PRN] Vista previa del recibo:\n%s\n", preview)
	}

	// Encolado de trabajo de impresión
	printPriority := "NORMAL"
	if order.CustomerType == "VIP" {
		printPriority = "HIGH"
	}
	printJobId := fmt.Sprintf("prn-%d-%d", time.Now().UnixMilli(), rand.Intn(1000))
	fmt.Printf("[PRN] Encolando trabajo %s (%d bytes, prioridad=%s) en %s\n", printJobId, spoolBytes, printPriority, printerConfig["location"])

	// Envío en trozos (chunking) para simular buffer limitado de la impresora
	const chunkSize = 256 // bytes
	sentBytes := 0
	chunkIndex := 0
	sentOk := true
	printableBytes := []byte(printable)
	for sentBytes < spoolBytes {
		remaining := spoolBytes - sentBytes
		size := chunkSize
		if remaining < chunkSize {
			size = remaining
		}
		// Simular reintentos por chunk
		attempts := 0
		delivered := false
		for !delivered && attempts < 2 {
			attempts++
			fmt.Printf("[PRN] Enviando chunk #%d (%d bytes) intento %d/2\n", chunkIndex, size, attempts)
			// Éxito simulado
			delivered = true
		}
		if !delivered {
			fmt.Printf("[PRN] Fallo al enviar chunk #%d\n", chunkIndex)
			sentOk = false
			break
		}
		sentBytes += size
		chunkIndex++
	}

	// Resultado final de impresión
	if printerConn["connected"].(bool) && sentOk {
		fmt.Printf("[PRN] Trabajo %s impreso correctamente. Total enviado: %d bytes\n", printJobId, sentBytes)
	} else {
		fmt.Printf("[PRN] Error al imprimir trabajo %s. Enviado: %d/%d bytes\n", printJobId, sentBytes, spoolBytes)
	}
}

func roundMoney(n float64) float64 {
	return math.Round(n*100) / 100
}

func formatMoney(n float64) string {
	return fmt.Sprintf("$%.2f", n)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
