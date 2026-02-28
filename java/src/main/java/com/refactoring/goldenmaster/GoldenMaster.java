package com.refactoring.goldenmaster;

import com.refactoring.refactoring.Order;
import com.refactoring.refactoring.OrderGenerator;
import com.refactoring.refactoring.ReceiptPrinter;

import java.util.Random;

/**
 * Ejercicio: Golden Master Testing
 *
 * El objetivo es crear una prueba de Golden Master para caracterizar el comportamiento
 * del código legado (ReceiptPrinter) y establecer una red de seguridad para refactorizar.
 *
 * Pasos:
 * 1. Identificar fuentes de no determinismo (Random, Date)
 * 2. Introducir costuras (seams) para controlar el comportamiento
 * 3. Generar un conjunto amplio de entradas
 * 4. Capturar la salida maestra
 * 5. Escribir pruebas que detecten cambios en la salida
 *
 * Ver README.md para instrucciones detalladas.
 */
public class GoldenMaster {

    public static void main(String[] args) {
        // Ejemplo de uso con comportamiento determinista para testing
        ReceiptPrinterForTest printer = new ReceiptPrinterForTest();
        Order order = OrderGenerator.generateOrder("ORD-001", "Ana", 3, 2);
        String receipt = printer.print(order);
        System.out.println(receipt);
    }

    /**
     * Subclase para testing que elimina el no determinismo.
     * Usa un Random con semilla fija y una fecha fija.
     */
    public static class ReceiptPrinterForTest extends ReceiptPrinter {
        public ReceiptPrinterForTest() {
            super(new Random(12345L)); // Semilla fija para reproducibilidad
        }

        @Override
        protected java.util.Date getCurrentDate() {
            // Fecha fija: 2022-01-01
            return new java.util.Date(1640995200000L); // 2022-01-01 00:00:00 UTC
        }

        @Override
        protected double discount() {
            // Sin descuento para simplificar los tests
            return 0.0;
        }
    }
}
