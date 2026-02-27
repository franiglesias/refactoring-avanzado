# Long Parameter List

Lista larga de parámetros.

## Definición

Una función recibe más de tres o cuatro parámetros.

## Ejemplo

```typescript
class ReportGenerator {
  generateReport(
    title: string,
    startDate: Date,
    endDate: Date,
    includeCharts: boolean,
    includeSummary: boolean,
    authorName: string,
    authorEmail: string,
  ) {
    console.log(`Generando reporte: ${title}`)
    console.log(`Desde ${startDate.toDateString()} hasta ${endDate.toDateString()}`)
    console.log(`Autor: ${authorName} (${authorEmail})`)
    if (includeCharts) console.log('Incluyendo gráficos...')
    if (includeSummary) console.log('Incluyendo resumen...')
    console.log('Reporte generado exitosamente.')
  }
}

export function demoLongParameterList(): void {
  const gen = new ReportGenerator()
  gen.generateReport(
    'Ventas Q1',
    new Date('2025-01-01'),
    new Date('2025-03-31'),
    true,
    false,
    'Pat Smith',
    'pat@example.com',
  )
}
```

## Ejercicio

Añade dos opciones más (p. ej., locale y pageSize) al reporte.

## Problemas que encontrarás

Con más de tres parámetros es difícil recordar con exactitud cuáles son, el orden o el tipo de cada
uno. Añadir parámetros no hace más que aumentar la dificultad de uso y mantenimiento.
