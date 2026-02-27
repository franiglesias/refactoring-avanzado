# Switch Statements

Sentencias switch.

## Definición

El uso excesivo de `switch` o múltiples `if/else` basados en un código de tipo suele ser una señal de que falta polimorfismo. El problema principal es que cada vez que se añade una nueva variante (un nuevo tipo), hay que buscar y modificar todos los bloques `switch` dispersos por la aplicación.

## Ejemplo

La función `calculatePay` utiliza un `switch` para decidir cómo calcular el salario según el tipo de empleado.

```typescript
export type EmployeeKind = 'engineer' | 'manager' | 'sales'

export interface EmployeeRecord {
  kind: EmployeeKind
  base: number
  bonus?: number
  commission?: number
}

export function calculatePay(rec: EmployeeRecord): number {
  switch (rec.kind) {
    case 'engineer':
      return rec.base
    case 'manager':
      return rec.base + (rec.bonus ?? 0)
    case 'sales':
      return rec.base + (rec.commission ?? 0)
    default:
      const _exhaustive: never = rec.kind
      return _exhaustive
  }
}

export function demoSwitchStatements(): number[] {
  return [
    calculatePay({kind: 'engineer', base: 1000}),
    calculatePay({kind: 'manager', base: 1000, bonus: 200}),
    calculatePay({kind: 'sales', base: 800, commission: 500}),
  ]
}

```

## Ejercicio

Añade un nuevo tipo de empleado (`contractor`) con una regla de pago especial (ej. tarifa por horas).

## Problemas que encontrarás

Tendrás que modificar el `switch` y cualquier otro código que dependa del tipo de empleado. A medida que el sistema crece, olvidar actualizar uno de estos puntos genera bugs difíciles de detectar.
