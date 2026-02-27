# Ejercicio: Crear una prueba de Golden Master para el código legado (ReceiptPrinter)

## Objetivo

Caracterizar el comportamiento actual y establecer una red de seguridad para refactorizar.

## Ejecutar tests

```shell
npm run test -- src/refactoring/golden-master/golden-master.test.ts
```

## Pasos recomendados (haz commits entre pasos)

1. Identificar fuentes de no determinismo que rompen cualquier intento de escribir tests.
2. Introducir costuras (SEAMS) mínimas sin cambiar el comportamiento.
  1. Aisla las fuentes no determninistas en métodos protegidos (seams).
  2. Extiende la clase y sobreescribe las costuras para controlar su comportamiento.
3. Generar un conjunto amplio y estable de entradas. Puedes usar esta forma para que el test se ejecute con un value nuevo cada vez:

```typescript
describe.each(['Ana', 'Luis', 'Mar', 'Iván', 'Sofía'])('Given a customer %s', (customer) => {
  it('should do something', () => {
    // Aquí el test
  });
});
```
4. Capturar la salida maestra.
5. Escribir la prueba usando como SUT la clase derivada, mediante `expect(texto).toMatchSnapshot()`.
6. Refactorizar con seguridad para invertir las dependencias.
7. Introducir tests unitarios/integración para reemplazar el Golden Master y poder introducir nueva funcionalidad.

## Criterios de aceptación

- La prueba de Golden Master captura múltiples casos de entrada y falla ante cambios en la salida.
- Las fuentes de no determinismo están controladas mediante Seams o, finalmente, inversión de dependencias.
- El comportamiento original de `ReceiptPrinter` sigue disponible para el resto del código (compatibilidad), inlcuyendo el uso de comportamientos no deterministas.
