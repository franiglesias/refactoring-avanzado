<?php

declare(strict_types=1);

namespace RefactoringAvanzado\CodeSmells\Bloaters;

use DateTime;

class OrderService
{
    public function process(Order $order): void
    {
        // Validar el pedido
        if (!$order->items || count($order->items) === 0) {
            echo "El pedido no tiene productos\n";
            return;
        }

        // Validar precios y cantidades
        foreach ($order->items as $item) {
            if ($item['price'] < 0 || $item['quantity'] <= 0) {
                echo "Producto inválido en el pedido\n";
                return;
            }
        }

        // Constantes de negocio (simples por ahora)
        $TAX_RATE = 0.21; // 21% IVA
        $FREE_SHIPPING_THRESHOLD = 50;
        $SHIPPING_FLAT = 5;

        // Calcular subtotal
        $subtotal = 0;
        foreach ($order->items as $item) {
            $subtotal += $item['price'] * $item['quantity'];
        }

        // Descuento por cliente VIP (10% del subtotal)
        $discount = 0;
        if ($order->customerType === 'VIP') {
            $discount = $this->roundMoney($subtotal * 0.1);
            echo "Descuento VIP aplicado\n";
        }

        // Base imponible
        $taxable = max(0, $subtotal - $discount);

        // Impuestos
        $tax = $this->roundMoney($taxable * $TAX_RATE);

        // Envío
        $shipping = $taxable >= $FREE_SHIPPING_THRESHOLD ? 0 : $SHIPPING_FLAT;

        // Total
        $total = $this->roundMoney($taxable + $tax + $shipping);

        // Actualizar el pedido (side-effects requeridos)
        $order->subtotal = $this->roundMoney($subtotal);
        $order->discount = $discount;
        $order->tax = $tax;
        $order->shipping = $shipping;
        $order->total = $total;

        // Registrar en la base de datos (simulado)
        // Bloque gigantesco y sobrecargado para simular persistencia con múltiples pasos innecesarios
        $dbConnectionString = 'Server=fake.db.local;Database=orders;User=demo;Password=demo';
        $dbConnected = true; // pretendemos que ya está conectado
        $dbRetriesMax = 3;
        $dbRetries = 0;
        $dbNow = new DateTime();
        $dbRecordId = random_int(1, 1000000);

        // Preparar registro a guardar
        $dbRecord = [
            'id' => $dbRecordId,
            'customerEmail' => $order->customerEmail,
            'customerType' => $order->customerType,
            'items' => array_map(fn($i) => [
                'name' => $i['name'],
                'price' => $i['price'],
                'quantity' => $i['quantity']
            ], $order->items),
            'amounts' => [
                'subtotal' => $order->subtotal,
                'discount' => $order->discount,
                'tax' => $order->tax,
                'shipping' => $order->shipping,
                'total' => $order->total,
            ],
            'status' => 'PENDING',
            'createdAt' => $dbNow->format('c'),
            'updatedAt' => $dbNow->format('c'),
            'currency' => 'USD',
        ];

        // Validaciones redundantes antes de guardar
        $hasItems = is_array($dbRecord['items']) && count($dbRecord['items']) > 0;
        $totalsConsistent = is_numeric($dbRecord['amounts']['total']) && $dbRecord['amounts']['total'] >= 0;

        if (!$hasItems) {
            echo "[DB] No se puede guardar: el pedido no tiene items\n";
        }
        if (!$totalsConsistent) {
            echo "[DB] No se puede guardar: total inconsistente\n";
        }

        // Simular transformación/serialización pesada
        $serialized = json_encode($dbRecord, JSON_PRETTY_PRINT);
        $payloadBytes = strlen($serialized);
        echo "[DB] Serializando registro {$dbRecord['id']} ({$payloadBytes} bytes) para {$dbConnectionString}\n";

        // Simular reintentos de escritura
        $dbSaved = false;
        while (!$dbSaved && $dbRetries < $dbRetriesMax) {
            $dbRetries++;
            if (!$dbConnected) {
                echo "[DB] Intento {$dbRetries}/{$dbRetriesMax}: reconectando a la base de datos...\n";
            } else {
                echo "[DB] Intento {$dbRetries}/{$dbRetriesMax}: guardando pedido {$dbRecord['id']} con total {$this->formatMoney($total)}\n";
            }
            // Resultado aleatorio simulado, pero aquí siempre "exitoso" para no complicar flujos de prueba
            $dbSaved = true;
        }

        if ($dbSaved) {
            echo "[DB] Pedido {$dbRecord['id']} guardado correctamente\n";
        } else {
            echo "[DB] No se pudo guardar el pedido {$dbRecord['id']} tras {$dbRetriesMax} intentos\n";
        }

        // Auditoría/bitácora adicional innecesaria
        $auditLogEntry = [
            'type' => 'ORDER_SAVED',
            'orderId' => $dbRecord['id'],
            'actor' => 'system',
            'at' => (new DateTime())->format('c'),
            'metadata' => [
                'ip' => '127.0.0.1',
                'userAgent' => 'OrderService/1.0',
            ],
        ];
        echo "[AUDIT] Registro: " . json_encode($auditLogEntry) . "\n";

        // Enviar correo de confirmación
        // Bloque gigantesco para simular el envío de un correo con plantillas, adjuntos, y seguimiento
        $smtpConfig = [
            'host' => 'smtp.fake.local',
            'port' => 587,
            'secure' => false,
            'auth' => ['user' => 'notifier', 'pass' => 'notifier'],
            'tls' => ['rejectUnauthorized' => false],
        ];

        $emailTemplate = "
      Hola,
      Gracias por tu pedido. Aquí tienes el resumen:\n
      Subtotal: {$this->formatMoney($order->subtotal)}\n
      Descuento: " . ($order->discount && $order->discount > 0 ? '-' . $this->formatMoney($order->discount) : $this->formatMoney(0)) . "\n
      Impuestos: {$this->formatMoney($order->tax)}\n
      Envío: {$this->formatMoney($order->shipping)}\n
      Total: {$this->formatMoney($order->total)}\n

      Nº de pedido: {$dbRecord['id']}\n
      Fecha: " . (new DateTime())->format('Y-m-d H:i:s') . "\n

      Saludos,
      Equipo Demo
    ";

        $trackingPixelUrl = "https://tracker.fake.local/pixel?orderId={$dbRecord['id']}&t=" . time();
        $emailBodyHtml = "
      <html>
        <body>
          <p>Hola,</p>
          <p>Gracias por tu pedido. Aquí tienes el resumen:</p>
          <ul>
            <li>Subtotal: <strong>{$this->formatMoney($order->subtotal)}</strong></li>
            <li>Descuento: <strong>" . ($order->discount && $order->discount > 0 ? '-' . $this->formatMoney($order->discount) : $this->formatMoney(0)) . "</strong></li>
            <li>Impuestos: <strong>{$this->formatMoney($order->tax)}</strong></li>
            <li>Envío: <strong>{$this->formatMoney($order->shipping)}</strong></li>
            <li>Total: <strong>{$this->formatMoney($order->total)}</strong></li>
          </ul>
          <p>Nº de pedido: <code>{$dbRecord['id']}</code></p>
          <p>Fecha: " . (new DateTime())->format('Y-m-d H:i:s') . "</p>
          <img src=\"{$trackingPixelUrl}\" width=\"1\" height=\"1\" alt=\"\"/>
        </body>
      </html>
    ";

        $attachments = [
            [
                'filename' => "pedido-{$dbRecord['id']}.json",
                'content' => $serialized,
                'contentType' => 'application/json',
            ],
            [
                'filename' => 'terminos.txt',
                'content' => 'Términos y condiciones...',
                'contentType' => 'text/plain'
            ],
        ];

        // Simular cálculo de tamaño del correo
        $emailSize = strlen($emailBodyHtml) + array_reduce(
            $attachments,
            fn($acc, $a) => $acc + strlen($a['content']),
            0
        );
        echo "[MAIL] Preparando correo ({$emailSize} bytes) vía {$smtpConfig['host']}:{$smtpConfig['port']}\n";

        // Simular colas de envío y priorización
        $emailPriority = $order->customerType === 'VIP' ? 'HIGH' : 'NORMAL';
        echo "[MAIL] Encolando correo ({$emailPriority}) para {$order->customerEmail}\n";

        // Simular envío con reintentos
        $mailAttempts = 0;
        $mailAttemptsMax = 2;
        $mailSent = false;
        while (!$mailSent && $mailAttempts < $mailAttemptsMax) {
            $mailAttempts++;
            echo "[MAIL] Intento {$mailAttempts}/{$mailAttemptsMax}: enviando correo a {$order->customerEmail}\n";
            // Simulación simple de éxito
            $mailSent = true;
        }

        $messageId = "msg-{$dbRecord['id']}-" . time();
        if ($mailSent) {
            echo "[MAIL] Correo enviado a {$order->customerEmail} (messageId={$messageId})\n";
        } else {
            echo "[MAIL] Fallo al enviar correo a {$order->customerEmail} tras {$mailAttemptsMax} intentos\n";
        }

        // Imprimir resumen -> enviar a impresora
        $printJob = [
            'title' => 'Resumen del pedido',
            'items' => array_map(fn($i) => [
                'name' => $i['name'],
                'quantity' => $i['quantity'],
                'lineTotal' => $this->roundMoney($i['price'] * $i['quantity']),
                'lineTotalFormatted' => $this->formatMoney($i['price'] * $i['quantity']),
            ], $order->items),
            'subtotal' => $order->subtotal ?? 0,
            'discount' => $order->discount ?? 0,
            'tax' => $order->tax ?? 0,
            'shipping' => $order->shipping ?? 0,
            'total' => $order->total ?? 0,
            'currency' => 'USD',
            'formatted' => [
                'subtotal' => $this->formatMoney($order->subtotal),
                'discount' => $order->discount && $order->discount > 0 ? '-' . $this->formatMoney($order->discount) : $this->formatMoney(0),
                'tax' => $this->formatMoney($order->tax),
                'shipping' => $this->formatMoney($order->shipping),
                'total' => $this->formatMoney($order->total),
            ],
            'metadata' => [
                'customerEmail' => $order->customerEmail,
                'createdAt' => (new DateTime())->format('c'),
            ],
        ];

        // Simulación de envío a impresora: bloque deliberadamente grande y sobrecargado
        // Configuración de impresora (ficticia)
        $printerConfig = [
            'name' => 'Demo Thermal Printer TP-80',
            'model' => 'TP-80',
            'dpi' => 203,
            'widthMm' => 80,
            'maxCharsPerLine' => 42, // típico en papel de 80mm con fuente estándar
            'interface' => 'USB',
            'driver' => 'ESC/POS',
            'location' => 'Front Desk',
        ];

        // Capabilities detectadas (simuladas)
        $printerCaps = [
            'supportsBold' => true,
            'supportsUnderline' => true,
            'supportsQr' => true,
            'supportsBarcode' => true,
            'supportsImages' => false,
            'codepage' => 'cp437',
        ];

        // Conexión (simulada)
        $printerConn = ['connected' => true, 'retries' => 0, 'maxRetries' => 2];
        echo "[PRN] Preparando conexión a impresora {$printerConfig['name']} ({$printerConfig['interface']}/{$printerConfig['driver']})\n";

        // Crear contenido del recibo
        $now = new DateTime();
        $lineWidth = $printerConfig['maxCharsPerLine'];

        $padRight = function (string $text, int $len): string {
            return strlen($text) >= $len ? substr($text, 0, $len) : $text . str_repeat(' ', $len - strlen($text));
        };

        $padLeft = function (string $text, int $len): string {
            return strlen($text) >= $len ? substr($text, 0, $len) : str_repeat(' ', $len - strlen($text)) . $text;
        };

        $repeat = fn(string $ch, int $n): string => str_repeat($ch, $n);

        $formatLine = function (string $left, string $right) use ($lineWidth, $padLeft, $repeat): string {
            $leftTrim = $left ?? '';
            $rightTrim = $right ?? '';
            $space = max(1, $lineWidth - strlen($leftTrim) - strlen($rightTrim));
            $spaces = $repeat(' ', $space);
            $tooLong = strlen($leftTrim) + strlen($rightTrim) > $lineWidth;
            if ($tooLong) {
                // Si no cabe, forzamos salto para la izquierda y mantenemos derecha alineada
                return $leftTrim . "\n" . $padLeft($rightTrim, $lineWidth);
            }
            return $leftTrim . $spaces . $rightTrim;
        };

        // Cabecera
        $receiptLines = [];
        $receiptLines[] = $repeat('=', $lineWidth);
        $receiptLines[] = $padRight('RESUMEN DEL PEDIDO', $lineWidth);
        $receiptLines[] = $padRight($now->format('Y-m-d H:i:s'), $lineWidth);
        $receiptLines[] = $padRight("Cliente: {$order->customerEmail}", $lineWidth);
        $receiptLines[] = $repeat('-', $lineWidth);

        // Items
        foreach ($printJob['items'] as $it) {
            $left = "{$it['quantity']} x {$it['name']}";
            $right = $it['lineTotalFormatted'];
            $receiptLines[] = $formatLine($left, $right);
        }

        // Totales
        $receiptLines[] = $repeat('-', $lineWidth);
        $receiptLines[] = $formatLine('Subtotal', $printJob['formatted']['subtotal']);
        if (($printJob['discount'] ?? 0) > 0) {
            $receiptLines[] = $formatLine('Descuento', '-' . $this->formatMoney($printJob['discount']));
        } else {
            $receiptLines[] = $formatLine('Descuento', $printJob['formatted']['discount']);
        }
        $receiptLines[] = $formatLine('Impuestos', $printJob['formatted']['tax']);
        $receiptLines[] = $formatLine('Envío', $printJob['formatted']['shipping']);
        $receiptLines[] = $formatLine('TOTAL', $printJob['formatted']['total']);
        $receiptLines[] = $repeat('=', $lineWidth);

        // Pie con metadatos
        $receiptLines[] = $padRight("Nº pedido: " . abs((int)(($order->total ?? 0) * 1000)), $lineWidth);
        $receiptLines[] = $padRight("Moneda: {$printJob['currency']}", $lineWidth);
        $receiptLines[] = $padRight("Creado: {$printJob['metadata']['createdAt']}", $lineWidth);

        // Comandos ESC/POS simulados (no operativos, solo logging)
        $escposCommands = [
            '[INIT]',
            '[ALIGN LEFT]',
            '[FONT A]',
            $printerCaps['supportsBold'] ? '[BOLD ON]' : '[BOLD N/A]',
            '[PRINT LINES]',
            '[BOLD OFF]',
            '[CUT PARTIAL]',
        ];

        // Montar payload a imprimir
        $textPayload = implode("\n", $receiptLines) . "\n" . $repeat('-', $lineWidth) . "\n";
        $commandSection = implode(' ', $escposCommands);
        $printable = "\n{$commandSection}\n{$textPayload}";
        $spoolBytes = strlen($printable);

        // Simulación de QR/barcode en el ticket (solo registro)
        $qrData = "ORDER|{$order->customerEmail}|{$printJob['total']}|{$now->getTimestamp()}";
        if ($printerCaps['supportsQr']) {
            echo "[PRN] Agregando QR con datos: {$qrData}\n";
        } elseif ($printerCaps['supportsBarcode']) {
            echo "[PRN] Agregando BARCODE con datos: " . substr($qrData, 0, 12) . "\n";
        } else {
            echo "[PRN] Sin soporte para QR/BARCODE\n";
        }

        // Vista previa ASCII (limitada para no saturar logs)
        $previewLines = array_slice(explode("\n", $textPayload), 0, 12);
        $preview = implode("\n", $previewLines);
        echo "[PRN] Vista previa del recibo:\n" . $preview .
            (count($receiptLines) > 12 ? "\n...(" . (count($receiptLines) - 12) . " líneas más)" : "") . "\n";

        // Encolado de trabajo de impresión
        $printPriority = $order->customerType === 'VIP' ? 'HIGH' : 'NORMAL';
        $printJobId = 'prn-' . time() . '-' . random_int(1, 1000);
        echo "[PRN] Encolando trabajo {$printJobId} ({$spoolBytes} bytes, prioridad={$printPriority}) en {$printerConfig['location']}\n";

        // Envío en trozos (chunking) para simular buffer limitado de la impresora
        $chunkSize = 256; // bytes
        $sentBytes = 0;
        $chunkIndex = 0;
        $sentOk = true;
        while ($sentBytes < $spoolBytes) {
            $remaining = $spoolBytes - $sentBytes;
            $size = min($chunkSize, $remaining);
            // Simular reintentos por chunk
            $attempts = 0;
            $delivered = false;
            while (!$delivered && $attempts < 2) {
                $attempts++;
                echo "[PRN] Enviando chunk #{$chunkIndex} ({$size} bytes) intento {$attempts}/2\n";
                // Éxito simulado
                $delivered = true;
            }
            if (!$delivered) {
                echo "[PRN] Fallo al enviar chunk #{$chunkIndex}\n";
                $sentOk = false;
                break;
            }
            $sentBytes += $size;
            $chunkIndex++;
        }

        // Resultado final de impresión
        if ($printerConn['connected'] && $sentOk) {
            echo "[PRN] Trabajo {$printJobId} impreso correctamente. Total enviado: {$sentBytes} bytes\n";
        } else {
            echo "[PRN] Error al imprimir trabajo {$printJobId}. Enviado: {$sentBytes}/{$spoolBytes} bytes\n";
        }
    }

    private function roundMoney(float $n): float
    {
        return round($n * 100) / 100;
    }

    private function formatMoney(?float $n): string
    {
        $v = is_numeric($n) ? $n : 0;
        return '$' . number_format($v, 2);
    }
}

class Order
{
    public string $customerEmail;
    /** @var 'NORMAL'|'VIP' */
    public string $customerType;
    /** @var array<array{name: string, price: float, quantity: int}> */
    public array $items;
    public ?float $subtotal = null;
    public ?float $discount = null;
    public ?float $tax = null;
    public ?float $shipping = null;
    public ?float $total = null;
}
