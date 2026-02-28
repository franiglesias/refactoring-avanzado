# Long Parameter List

Lista larga de parámetros.

## Definición

Una función recibe más de tres o cuatro parámetros.

## Ejemplo

```php
<?php

declare(strict_types=1);

namespace CodeSmells\Bloaters;

use DateTime;

class ReportGenerator
{
    public function generateReport(
        string $title,
        DateTime $startDate,
        DateTime $endDate,
        bool $includeCharts,
        bool $includeSummary,
        string $authorName,
        string $authorEmail
    ): void {
        echo "Generando reporte: {$title}\n";
        echo "Desde {$startDate->format('Y-m-d')} hasta {$endDate->format('Y-m-d')}\n";
        echo "Autor: {$authorName} ({$authorEmail})\n";
        if ($includeCharts) {
            echo "Incluyendo gráficos...\n";
        }
        if ($includeSummary) {
            echo "Incluyendo resumen...\n";
        }
        echo "Reporte generado exitosamente.\n";
    }
}

function demoLongParameterList(): void
{
    $gen = new ReportGenerator();
    $gen->generateReport(
        'Ventas Q1',
        new DateTime('2025-01-01'),
        new DateTime('2025-03-31'),
        true,
        false,
        'Pat Smith',
        'pat@example.com'
    );
}
```

## Ejercicio

Añade dos opciones más (p. ej., locale y pageSize) al reporte.

## Problemas que encontrarás

Con más de tres parámetros es difícil recordar con exactitud cuáles son, el orden o el tipo de cada uno. Añadir parámetros no hace más que aumentar la dificultad de uso y mantenimiento.
