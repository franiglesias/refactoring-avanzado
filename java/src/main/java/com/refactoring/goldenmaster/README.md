# Ejercicio: Crear una prueba de Golden Master para código legado (ReceiptPrinter)

## Objetivo

Caracterizar el comportamiento actual del código legado y establecer una red de seguridad para refactorizar con confianza.

## Contexto

La clase `ReceiptPrinter` en el paquete `com.refactoring.refactoring` es código legado que:
- Calcula totales con impuestos variables según categoría
- Aplica descuentos aleatorios ("descuento de la suerte")
- Formatea recibos con fecha actual
- Tiene comportamiento no determinista (Random, Date)

**No debes modificar `ReceiptPrinter` al inicio del ejercicio**. Primero crearás el Golden Master.

## Pasos recomendados (haz commits entre pasos)

### 1. Identificar fuentes de no determinismo

Las fuentes que rompen cualquier intento de escribir tests deterministas:
- `Random` para el descuento aleatorio
- `Date.now()` para la fecha actual del recibo

### 2. Introducir costuras (SEAMS) mínimas sin cambiar el comportamiento

Ya hemos preparado algunas costuras en `ReceiptPrinter`:
- Constructor que acepta un `Random` para testing
- Método protegido `discount()` que puedes sobrescribir
- Método protegido `getCurrentDate()` que puedes sobrescribir

**Estrategia**: Extiende `ReceiptPrinter` y sobreescribe estos métodos para controlar el comportamiento.

Ejemplo en `GoldenMaster.java`:

```java
public static class ReceiptPrinterForTest extends ReceiptPrinter {
    public ReceiptPrinterForTest() {
        super(new Random(12345L)); // Semilla fija
    }

    @Override
    protected Date getCurrentDate() {
        return new Date(1640995200000L); // Fecha fija
    }

    @Override
    protected double discount() {
        return 0.0; // Sin descuento aleatorio
    }
}
```

### 3. Generar un conjunto amplio y estable de entradas

Usa `OrderGenerator.generateOrder()` para crear pedidos con diferentes combinaciones:
- Diferentes clientes
- Diferentes cantidades de items (1-5)
- Diferentes cantidades por item (1-4)

Ejemplo:

```java
@Test
public void testGoldenMaster() {
    ReceiptPrinterForTest printer = new ReceiptPrinterForTest();

    // Generar múltiples casos
    for (String customer : new String[]{"Ana", "Luis", "Mar"}) {
        for (int numItems = 1; numItems <= 3; numItems++) {
            for (int qty = 1; qty <= 2; qty++) {
                String orderId = String.format("ORD-%s-%d-%d", customer, numItems, qty);
                Order order = OrderGenerator.generateOrder(orderId, customer, numItems, qty);
                String receipt = printer.print(order);

                // Aquí capturarías la salida o compararías con el golden master
                // Por ejemplo, usando approval tests o archivos snapshot
            }
        }
    }
}
```

### 4. Capturar la salida maestra

Opciones:
- **Approval Testing**: Usa una librería como ApprovalTests.Java
- **Snapshot Testing**: Guarda las salidas en archivos y compáralas
- **Assertions manuales**: Captura las salidas y verifica que no cambien

### 5. Escribir la prueba usando la clase derivada

La prueba debe:
- Usar `ReceiptPrinterForTest` para comportamiento determinista
- Generar múltiples casos de entrada
- Capturar y verificar la salida
- Fallar si la salida cambia

### 6. Refactorizar con seguridad

Una vez que tienes el Golden Master:
- Puedes refactorizar `ReceiptPrinter` con confianza
- Los tests detectarán cualquier cambio de comportamiento
- Gradualmente puedes invertir dependencias (inyectar Clock, Random)

### 7. Introducir tests unitarios/integración

Eventualmente reemplaza el Golden Master con tests más específicos:
- Tests unitarios para cálculos individuales
- Tests de integración para flujos completos
- Esto permite añadir nueva funcionalidad con TDD

## Criterios de aceptación

- La prueba de Golden Master captura múltiples casos de entrada
- La prueba falla ante cambios en la salida
- Las fuentes de no determinismo están controladas mediante Seams
- El comportamiento original de `ReceiptPrinter` sigue disponible (compatibilidad)
- Puedes refactorizar con confianza sabiendo que los tests te protegen

## Recursos

- Clases necesarias ya están en `com.refactoring.refactoring`:
  - `Order`, `OrderItem`, `OrderGenerator`, `ReceiptPrinter`
- Ejemplo de uso en `GoldenMaster.java`

## Ejecución

```bash
# Compilar
javac -d out src/main/java/com/refactoring/**/*.java

# Ejecutar ejemplo
java -cp out com.refactoring.goldenmaster.GoldenMaster

# Para tests, usa tu framework favorito (JUnit, TestNG)
```

## Notas

- El código legado ya tiene algunas costuras preparadas
- No modifiques `ReceiptPrinter` hasta tener el Golden Master
- El ejercicio simula una situación real: código sin tests que necesitas refactorizar
