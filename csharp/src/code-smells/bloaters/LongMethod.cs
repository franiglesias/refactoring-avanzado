using System.Text;

namespace RefactoringAvanzado.CodeSmells.Bloaters;

public class OrderService
{
    public void Process(Order order)
    {
        // Validar el pedido
        if (order.Items == null || order.Items.Count == 0)
        {
            Console.WriteLine("El pedido no tiene productos");
            return;
        }

        // Validar precios y cantidades
        foreach (var item in order.Items)
        {
            if (item.Price < 0 || item.Quantity <= 0)
            {
                Console.WriteLine("Producto inválido en el pedido");
                return;
            }
        }

        // Constantes de negocio (simples por ahora)
        const double TAX_RATE = 0.21; // 21% IVA
        const double FREE_SHIPPING_THRESHOLD = 50;
        const double SHIPPING_FLAT = 5;

        // Calcular subtotal
        double subtotal = 0;
        foreach (var item in order.Items)
        {
            subtotal += item.Price * item.Quantity;
        }

        // Descuento por cliente VIP (10% del subtotal)
        double discount = 0;
        if (order.CustomerType == "VIP")
        {
            discount = RoundMoney(subtotal * 0.1);
            Console.WriteLine("Descuento VIP aplicado");
        }

        // Base imponible
        double taxable = Math.Max(0, subtotal - discount);

        // Impuestos
        double tax = RoundMoney(taxable * TAX_RATE);

        // Envío
        double shipping = taxable >= FREE_SHIPPING_THRESHOLD ? 0 : SHIPPING_FLAT;

        // Total
        double total = RoundMoney(taxable + tax + shipping);

        // Actualizar el pedido (side-effects requeridos)
        order.Subtotal = RoundMoney(subtotal);
        order.Discount = discount;
        order.Tax = tax;
        order.Shipping = shipping;
        order.Total = total;

        // Registrar en la base de datos (simulado)
        // Bloque gigantesco y sobrecargado para simular persistencia con múltiples pasos innecesarios
        string dbConnectionString = "Server=fake.db.local;Database=orders;User=demo;Password=demo";
        bool dbConnected = true; // pretendemos que ya está conectado
        int dbRetriesMax = 3;
        int dbRetries = 0;
        DateTime dbNow = DateTime.Now;
        int dbRecordId = new Random().Next(1000000);

        // Preparar registro a guardar
        var dbRecord = new
        {
            id = dbRecordId,
            customerEmail = order.CustomerEmail,
            customerType = order.CustomerType,
            items = order.Items.Select(i => new { name = i.Name, price = i.Price, quantity = i.Quantity }).ToList(),
            amounts = new
            {
                subtotal = order.Subtotal,
                discount = order.Discount,
                tax = order.Tax,
                shipping = order.Shipping,
                total = order.Total
            },
            status = "PENDING",
            createdAt = dbNow.ToString("o"),
            updatedAt = dbNow.ToString("o"),
            currency = "USD"
        };

        // Validaciones redundantes antes de guardar
        bool hasItems = dbRecord.items != null && dbRecord.items.Count > 0;
        bool totalsConsistent = dbRecord.amounts.total >= 0;
        if (!hasItems)
        {
            Console.WriteLine("[DB] No se puede guardar: el pedido no tiene items");
        }
        if (!totalsConsistent)
        {
            Console.WriteLine("[DB] No se puede guardar: total inconsistente");
        }

        // Simular transformación/serialización pesada
        string serialized = System.Text.Json.JsonSerializer.Serialize(dbRecord, new System.Text.Json.JsonSerializerOptions { WriteIndented = true });
        int payloadBytes = Encoding.UTF8.GetByteCount(serialized);
        Console.WriteLine($"[DB] Serializando registro {dbRecord.id} ({payloadBytes} bytes) para {dbConnectionString}");

        // Simular reintentos de escritura
        bool dbSaved = false;
        while (!dbSaved && dbRetries < dbRetriesMax)
        {
            dbRetries++;
            if (!dbConnected)
            {
                Console.WriteLine($"[DB] Intento {dbRetries}/{dbRetriesMax}: reconectando a la base de datos...");
            }
            else
            {
                Console.WriteLine($"[DB] Intento {dbRetries}/{dbRetriesMax}: guardando pedido {dbRecord.id} con total {FormatMoney(total)}");
            }
            // Resultado aleatorio simulado, pero aquí siempre "exitoso" para no complicar flujos de prueba
            dbSaved = true;
        }

        if (dbSaved)
        {
            Console.WriteLine($"[DB] Pedido {dbRecord.id} guardado correctamente");
        }
        else
        {
            Console.WriteLine($"[DB] No se pudo guardar el pedido {dbRecord.id} tras {dbRetriesMax} intentos");
        }

        // Auditoría/bitácora adicional innecesaria
        var auditLogEntry = new
        {
            type = "ORDER_SAVED",
            orderId = dbRecord.id,
            actor = "system",
            at = DateTime.Now.ToString("o"),
            metadata = new
            {
                ip = "127.0.0.1",
                userAgent = "OrderService/1.0"
            }
        };
        Console.WriteLine($"[AUDIT] Registro: {System.Text.Json.JsonSerializer.Serialize(auditLogEntry)}");

        // Enviar correo de confirmación
        // Bloque gigantesco para simular el envío de un correo con plantillas, adjuntos, y seguimiento
        var smtpConfig = new
        {
            host = "smtp.fake.local",
            port = 587,
            secure = false,
            auth = new { user = "notifier", pass = "notifier" },
            tls = new { rejectUnauthorized = false }
        };
        string emailTemplate = $@"
      Hola,
      Gracias por tu pedido. Aquí tienes el resumen:

      Subtotal: {FormatMoney(order.Subtotal)}

      Descuento: {(order.Discount > 0 ? "-" + FormatMoney(order.Discount) : FormatMoney(0))}

      Impuestos: {FormatMoney(order.Tax)}

      Envío: {FormatMoney(order.Shipping)}

      Total: {FormatMoney(order.Total)}


      Nº de pedido: {dbRecord.id}

      Fecha: {DateTime.Now}


      Saludos,
      Equipo Demo
    ";
        string trackingPixelUrl = $"https://tracker.fake.local/pixel?orderId={dbRecord.id}&t={DateTimeOffset.Now.ToUnixTimeMilliseconds()}";
        string emailBodyHtml = $@"
      <html>
        <body>
          <p>Hola,</p>
          <p>Gracias por tu pedido. Aquí tienes el resumen:</p>
          <ul>
            <li>Subtotal: <strong>{FormatMoney(order.Subtotal)}</strong></li>
            <li>Descuento: <strong>{(order.Discount > 0 ? "-" + FormatMoney(order.Discount) : FormatMoney(0))}</strong></li>
            <li>Impuestos: <strong>{FormatMoney(order.Tax)}</strong></li>
            <li>Envío: <strong>{FormatMoney(order.Shipping)}</strong></li>
            <li>Total: <strong>{FormatMoney(order.Total)}</strong></li>
          </ul>
          <p>Nº de pedido: <code>{dbRecord.id}</code></p>
          <p>Fecha: {DateTime.Now}</p>
          <img src=""{trackingPixelUrl}"" width=""1"" height=""1"" alt=""""/>
        </body>
      </html>
    ";

        var attachments = new[]
        {
            new { filename = $"pedido-{dbRecord.id}.json", content = serialized, contentType = "application/json" },
            new { filename = "terminos.txt", content = "Términos y condiciones...", contentType = "text/plain" }
        };

        // Simular cálculo de tamaño del correo
        int emailSize = Encoding.UTF8.GetByteCount(emailBodyHtml) +
            attachments.Sum(a => Encoding.UTF8.GetByteCount(a.content));
        Console.WriteLine($"[MAIL] Preparando correo ({emailSize} bytes) vía {smtpConfig.host}:{smtpConfig.port}");

        // Simular colas de envío y priorización
        string emailPriority = order.CustomerType == "VIP" ? "HIGH" : "NORMAL";
        Console.WriteLine($"[MAIL] Encolando correo ({emailPriority}) para {order.CustomerEmail}");

        // Simular envío con reintentos
        int mailAttempts = 0;
        int mailAttemptsMax = 2;
        bool mailSent = false;
        while (!mailSent && mailAttempts < mailAttemptsMax)
        {
            mailAttempts++;
            Console.WriteLine($"[MAIL] Intento {mailAttempts}/{mailAttemptsMax}: enviando correo a {order.CustomerEmail}");
            // Simulación simple de éxito
            mailSent = true;
        }

        string messageId = $"msg-{dbRecord.id}-{DateTimeOffset.Now.ToUnixTimeMilliseconds()}";
        if (mailSent)
        {
            Console.WriteLine($"[MAIL] Correo enviado a {order.CustomerEmail} (messageId={messageId})");
        }
        else
        {
            Console.WriteLine($"[MAIL] Fallo al enviar correo a {order.CustomerEmail} tras {mailAttemptsMax} intentos");
        }

        // Imprimir resumen -> enviar a impresora
        var printJob = new PrintJob
        {
            Title = "Resumen del pedido",
            Items = order.Items.Select(i => new PrintJobItem
            {
                Name = i.Name,
                Quantity = i.Quantity,
                LineTotal = RoundMoney(i.Price * i.Quantity),
                LineTotalFormatted = FormatMoney(i.Price * i.Quantity)
            }).ToList(),
            Subtotal = order.Subtotal ?? 0,
            Discount = order.Discount ?? 0,
            Tax = order.Tax ?? 0,
            Shipping = order.Shipping ?? 0,
            Total = order.Total ?? 0,
            Currency = "USD",
            Formatted = new PrintJobFormatted
            {
                Subtotal = FormatMoney(order.Subtotal),
                Discount = order.Discount > 0 ? $"-{FormatMoney(order.Discount)}" : FormatMoney(0),
                Tax = FormatMoney(order.Tax),
                Shipping = FormatMoney(order.Shipping),
                Total = FormatMoney(order.Total)
            },
            Metadata = new PrintJobMetadata
            {
                CustomerEmail = order.CustomerEmail,
                CreatedAt = DateTime.Now.ToString("o")
            }
        };

        // Simulación de envío a impresora: bloque deliberadamente grande y sobrecargado
        // Configuración de impresora (ficticia)
        var printerConfig = new
        {
            name = "Demo Thermal Printer TP-80",
            model = "TP-80",
            dpi = 203,
            widthMm = 80,
            maxCharsPerLine = 42, // típico en papel de 80mm con fuente estándar
            interfaceType = "USB",
            driver = "ESC/POS",
            location = "Front Desk"
        };

        // Capabilities detectadas (simuladas)
        var printerCaps = new
        {
            supportsBold = true,
            supportsUnderline = true,
            supportsQr = true,
            supportsBarcode = true,
            supportsImages = false,
            codepage = "cp437"
        };

        // Conexión (simulada)
        var printerConn = new { connected = true, retries = 0, maxRetries = 2 };
        Console.WriteLine($"[PRN] Preparando conexión a impresora {printerConfig.name} ({printerConfig.interfaceType}/{printerConfig.driver})");

        // Crear contenido del recibo
        DateTime now = DateTime.Now;
        int lineWidth = printerConfig.maxCharsPerLine;

        string PadRight(string text, int len) =>
            text.Length >= len ? text.Substring(0, len) : text + new string(' ', len - text.Length);
        string PadLeft(string text, int len) =>
            text.Length >= len ? text.Substring(0, len) : new string(' ', len - text.Length) + text;
        string Repeat(char ch, int n) => new string(ch, n);

        string FormatLine(string left, string right)
        {
            string leftTrim = left ?? "";
            string rightTrim = right ?? "";
            int space = Math.Max(1, lineWidth - leftTrim.Length - rightTrim.Length);
            string spaces = Repeat(' ', space);
            bool tooLong = leftTrim.Length + rightTrim.Length > lineWidth;
            if (tooLong)
            {
                // Si no cabe, forzamos salto para la izquierda y mantenemos derecha alineada
                return leftTrim + "\n" + PadLeft(rightTrim, lineWidth);
            }
            return leftTrim + spaces + rightTrim;
        }

        // Cabecera
        var receiptLines = new List<string>();
        receiptLines.Add(Repeat('=', lineWidth));
        receiptLines.Add(PadRight("RESUMEN DEL PEDIDO", lineWidth));
        receiptLines.Add(PadRight(now.ToString(), lineWidth));
        receiptLines.Add(PadRight($"Cliente: {order.CustomerEmail}", lineWidth));
        receiptLines.Add(Repeat('-', lineWidth));

        // Items
        foreach (var it in printJob.Items)
        {
            string left = $"{it.Quantity} x {it.Name}";
            string right = it.LineTotalFormatted;
            receiptLines.Add(FormatLine(left, right));
        }

        // Totales
        receiptLines.Add(Repeat('-', lineWidth));
        receiptLines.Add(FormatLine("Subtotal", printJob.Formatted.Subtotal));
        if (printJob.Discount > 0)
        {
            receiptLines.Add(FormatLine("Descuento", $"-{FormatMoney(printJob.Discount)}"));
        }
        else
        {
            receiptLines.Add(FormatLine("Descuento", printJob.Formatted.Discount));
        }
        receiptLines.Add(FormatLine("Impuestos", printJob.Formatted.Tax));
        receiptLines.Add(FormatLine("Envío", printJob.Formatted.Shipping));
        receiptLines.Add(FormatLine("TOTAL", printJob.Formatted.Total));
        receiptLines.Add(Repeat('=', lineWidth));

        // Pie con metadatos
        receiptLines.Add(PadRight($"Nº pedido: {Math.Abs((int)((order.Total ?? 0) * 1000))}", lineWidth));
        receiptLines.Add(PadRight($"Moneda: {printJob.Currency}", lineWidth));
        receiptLines.Add(PadRight($"Creado: {printJob.Metadata.CreatedAt}", lineWidth));

        // Comandos ESC/POS simulados (no operativos, solo logging)
        var escposCommands = new[]
        {
            "[INIT]",
            "[ALIGN LEFT]",
            "[FONT A]",
            printerCaps.supportsBold ? "[BOLD ON]" : "[BOLD N/A]",
            "[PRINT LINES]",
            "[BOLD OFF]",
            "[CUT PARTIAL]"
        };

        // Montar payload a imprimir
        string textPayload = string.Join("\n", receiptLines) + "\n" + Repeat('-', lineWidth) + "\n";
        string commandSection = string.Join(" ", escposCommands);
        string printable = $"\n{commandSection}\n{textPayload}";
        byte[] spoolBuffer = Encoding.UTF8.GetBytes(printable);
        int spoolBytes = spoolBuffer.Length;

        // Simulación de QR/barcode en el ticket (solo registro)
        string qrData = $"ORDER|{order.CustomerEmail}|{printJob.Total}|{now.Ticks}";
        if (printerCaps.supportsQr)
        {
            Console.WriteLine($"[PRN] Agregando QR con datos: {qrData}");
        }
        else if (printerCaps.supportsBarcode)
        {
            Console.WriteLine($"[PRN] Agregando BARCODE con datos: {qrData.Substring(0, Math.Min(12, qrData.Length))}");
        }
        else
        {
            Console.WriteLine("[PRN] Sin soporte para QR/BARCODE");
        }

        // Vista previa ASCII (limitada para no saturar logs)
        string preview = string.Join("\n", receiptLines.Take(12));
        Console.WriteLine("[PRN] Vista previa del recibo:\n" +
            preview +
            (receiptLines.Count > 12 ? $"\n...({receiptLines.Count - 12} líneas más)" : ""));

        // Encolado de trabajo de impresión
        string printPriority = order.CustomerType == "VIP" ? "HIGH" : "NORMAL";
        string printJobId = $"prn-{DateTimeOffset.Now.ToUnixTimeMilliseconds()}-{new Random().Next(1000)}";
        Console.WriteLine($"[PRN] Encolando trabajo {printJobId} ({spoolBytes} bytes, prioridad={printPriority}) en {printerConfig.location}");

        // Envío en trozos (chunking) para simular buffer limitado de la impresora
        int chunkSize = 256; // bytes
        int sentBytes = 0;
        int chunkIndex = 0;
        bool sentOk = true;
        while (sentBytes < spoolBytes)
        {
            int remaining = spoolBytes - sentBytes;
            int size = Math.Min(chunkSize, remaining);
            byte[] chunk = spoolBuffer.Skip(sentBytes).Take(size).ToArray();
            // Simular reintentos por chunk
            int attempts = 0;
            bool delivered = false;
            while (!delivered && attempts < 2)
            {
                attempts++;
                Console.WriteLine($"[PRN] Enviando chunk #{chunkIndex} ({size} bytes) intento {attempts}/2");
                // Éxito simulado
                delivered = true;
            }
            if (!delivered)
            {
                Console.WriteLine($"[PRN] Fallo al enviar chunk #{chunkIndex}");
                sentOk = false;
                break;
            }
            sentBytes += size;
            chunkIndex++;
        }

        // Resultado final de impresión
        if (printerConn.connected && sentOk)
        {
            Console.WriteLine($"[PRN] Trabajo {printJobId} impreso correctamente. Total enviado: {sentBytes} bytes");
        }
        else
        {
            Console.WriteLine($"[PRN] Error al imprimir trabajo {printJobId}. Enviado: {sentBytes}/{spoolBytes} bytes");
        }
    }

    private static double RoundMoney(double n)
    {
        return Math.Round(n * 100) / 100;
    }

    private static string FormatMoney(double? n)
    {
        double v = n ?? 0;
        return $"${v:F2}";
    }
}

public class Order
{
    public string CustomerEmail { get; set; } = "";
    public string CustomerType { get; set; } = "NORMAL";
    public List<OrderItem> Items { get; set; } = new();
    public double? Subtotal { get; set; }
    public double? Discount { get; set; }
    public double? Tax { get; set; }
    public double? Shipping { get; set; }
    public double? Total { get; set; }
}

public class OrderItem
{
    public string Name { get; set; } = "";
    public double Price { get; set; }
    public int Quantity { get; set; }
}

public class PrintJob
{
    public string Title { get; set; } = "";
    public List<PrintJobItem> Items { get; set; } = new();
    public double Subtotal { get; set; }
    public double Discount { get; set; }
    public double Tax { get; set; }
    public double Shipping { get; set; }
    public double Total { get; set; }
    public string Currency { get; set; } = "";
    public PrintJobFormatted Formatted { get; set; } = new();
    public PrintJobMetadata Metadata { get; set; } = new();
}

public class PrintJobItem
{
    public string Name { get; set; } = "";
    public int Quantity { get; set; }
    public double LineTotal { get; set; }
    public string LineTotalFormatted { get; set; } = "";
}

public class PrintJobFormatted
{
    public string Subtotal { get; set; } = "";
    public string Discount { get; set; } = "";
    public string Tax { get; set; } = "";
    public string Shipping { get; set; } = "";
    public string Total { get; set; } = "";
}

public class PrintJobMetadata
{
    public string CustomerEmail { get; set; } = "";
    public string CreatedAt { get; set; } = "";
}
