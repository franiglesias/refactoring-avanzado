package com.refactoring.codesmells.bloaters;

import java.util.*;

/**
 * Code Smell: Long Method [Método largo]
 * Este método process() maneja muchas responsabilidades:
 * validación, cálculos, persistencia, envío de email e impresión.
 * Es difícil de leer, mantener y probar.
 */
public class LongMethod {

    public static void main(String[] args) {
        OrderService service = new OrderService();
        Order order = new Order(
            "cliente@example.com",
            CustomerType.VIP,
            Arrays.asList(
                new OrderItem("Producto A", 50.0, 2),
                new OrderItem("Producto B", 30.0, 1)
            )
        );
        service.process(order);
    }

    public enum CustomerType {
        NORMAL, VIP
    }

    public static class OrderItem {
        String name;
        double price;
        int quantity;

        public OrderItem(String name, double price, int quantity) {
            this.name = name;
            this.price = price;
            this.quantity = quantity;
        }
    }

    public static class Order {
        String customerEmail;
        CustomerType customerType;
        List<OrderItem> items;
        Double subtotal;
        Double discount;
        Double tax;
        Double shipping;
        Double total;

        public Order(String customerEmail, CustomerType customerType, List<OrderItem> items) {
            this.customerEmail = customerEmail;
            this.customerType = customerType;
            this.items = items;
        }
    }

    public static class OrderService {
        public void process(Order order) {
            // Validar el pedido
            if (order.items == null || order.items.isEmpty()) {
                System.out.println("El pedido no tiene productos");
                return;
            }

            // Validar precios y cantidades
            for (OrderItem item : order.items) {
                if (item.price < 0 || item.quantity <= 0) {
                    System.out.println("Producto inválido en el pedido");
                    return;
                }
            }

            // Constantes de negocio
            final double TAX_RATE = 0.21; // 21% IVA
            final double FREE_SHIPPING_THRESHOLD = 50;
            final double SHIPPING_FLAT = 5;

            // Calcular subtotal
            double subtotal = 0;
            for (OrderItem item : order.items) {
                subtotal += item.price * item.quantity;
            }

            // Descuento por cliente VIP (10% del subtotal)
            double discount = 0;
            if (order.customerType == CustomerType.VIP) {
                discount = roundMoney(subtotal * 0.1);
                System.out.println("Descuento VIP aplicado");
            }

            // Base imponible
            double taxable = Math.max(0, subtotal - discount);

            // Impuestos
            double tax = roundMoney(taxable * TAX_RATE);

            // Envío
            double shipping = taxable >= FREE_SHIPPING_THRESHOLD ? 0 : SHIPPING_FLAT;

            // Total
            double total = roundMoney(taxable + tax + shipping);

            // Actualizar el pedido
            order.subtotal = roundMoney(subtotal);
            order.discount = discount;
            order.tax = tax;
            order.shipping = shipping;
            order.total = total;

            // Simular guardado en base de datos
            String dbConnectionString = "Server=fake.db.local;Database=orders;User=demo";
            int dbRecordId = new Random().nextInt(1000000);
            System.out.println("[DB] Serializando registro " + dbRecordId + " para " + dbConnectionString);
            System.out.println("[DB] Pedido " + dbRecordId + " guardado correctamente");

            // Auditoría
            System.out.println("[AUDIT] Registro: ORDER_SAVED orderId=" + dbRecordId);

            // Enviar correo de confirmación
            String smtpHost = "smtp.fake.local";
            int smtpPort = 587;
            String emailBody = String.format(
                "Hola,\nGracias por tu pedido. Aquí tienes el resumen:\n" +
                "Subtotal: %s\n" +
                "Descuento: %s\n" +
                "Impuestos: %s\n" +
                "Envío: %s\n" +
                "Total: %s\n" +
                "Nº de pedido: %d\n" +
                "Fecha: %s\n",
                formatMoney(order.subtotal),
                order.discount > 0 ? "-" + formatMoney(order.discount) : formatMoney(0.0),
                formatMoney(order.tax),
                formatMoney(order.shipping),
                formatMoney(order.total),
                dbRecordId,
                new Date()
            );

            System.out.println("[MAIL] Preparando correo vía " + smtpHost + ":" + smtpPort);
            String emailPriority = order.customerType == CustomerType.VIP ? "HIGH" : "NORMAL";
            System.out.println("[MAIL] Encolando correo (" + emailPriority + ") para " + order.customerEmail);
            System.out.println("[MAIL] Correo enviado a " + order.customerEmail);

            // Imprimir resumen
            String printerName = "Demo Thermal Printer TP-80";
            int lineWidth = 42;
            System.out.println("[PRN] Preparando conexión a impresora " + printerName);

            StringBuilder receipt = new StringBuilder();
            receipt.append("=".repeat(lineWidth)).append("\n");
            receipt.append("RESUMEN DEL PEDIDO\n");
            receipt.append(new Date()).append("\n");
            receipt.append("Cliente: ").append(order.customerEmail).append("\n");
            receipt.append("-".repeat(lineWidth)).append("\n");

            for (OrderItem item : order.items) {
                double lineTotal = roundMoney(item.price * item.quantity);
                receipt.append(String.format("%d x %s %s\n",
                    item.quantity, item.name, formatMoney(lineTotal)));
            }

            receipt.append("-".repeat(lineWidth)).append("\n");
            receipt.append(String.format("Subtotal %s\n", formatMoney(order.subtotal)));
            if (order.discount > 0) {
                receipt.append(String.format("Descuento -%s\n", formatMoney(order.discount)));
            }
            receipt.append(String.format("Impuestos %s\n", formatMoney(order.tax)));
            receipt.append(String.format("Envío %s\n", formatMoney(order.shipping)));
            receipt.append(String.format("TOTAL %s\n", formatMoney(order.total)));
            receipt.append("=".repeat(lineWidth)).append("\n");

            System.out.println("[PRN] Vista previa del recibo:\n" + receipt.toString());
            System.out.println("[PRN] Trabajo impreso correctamente");
        }

        private double roundMoney(double n) {
            return Math.round(n * 100.0) / 100.0;
        }

        private String formatMoney(double n) {
            return String.format("$%.2f", n);
        }
    }
}
