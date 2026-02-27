# Curso Refactoring Avanzado - PHP

Ejemplos y ejercicios del Curso de Refactoring Avanzado en PHP 8.2+.

## Requisitos

- Docker y Docker Compose (recomendado)
- O PHP 8.2+ y Composer instalados localmente

## Preparación

### Opción 1: Usando Docker (Recomendado)

1. Construir y levantar el contenedor:
```bash
docker compose build
docker compose up -d
```

2. Instalar dependencias:
```bash
docker compose exec php composer install
```

3. Verificar que Xdebug está instalado:
```bash
docker compose exec php php -v
```
Deberías ver: `with Xdebug v3.3.1`

4. Acceder al contenedor para trabajar:
```bash
docker compose exec php bash
```

### Opción 2: Sin Docker

1. Instalar dependencias:
```bash
composer install
```

## Ejecutar Tests

### Con Docker
```bash
docker compose exec php composer test
```

### Sin Docker
```bash
composer test
```

## Otros Comandos Útiles

### Análisis estático con PHPStan
```bash
# Con Docker
docker compose exec php composer phpstan

# Sin Docker
composer phpstan
```

### Verificar estilo de código
```bash
# Con Docker
docker compose exec php composer cs

# Sin Docker
composer cs
```

### Corregir estilo de código automáticamente
```bash
# Con Docker
docker compose exec php composer cs:fix

# Sin Docker
composer cs:fix
```

### Generar reporte de cobertura
```bash
# Con Docker
docker compose exec php composer test:coverage

# Sin Docker
composer test:coverage
```

El reporte se generará en la carpeta `coverage/` (HTML) y `coverage/clover.xml` (formato Clover).

## Xdebug - Debugging y Code Coverage

Este proyecto incluye Xdebug 3 configurado para debugging y generación de reportes de cobertura de código.

### Características de Xdebug

- **Debug**: Debugging paso a paso con breakpoints
- **Coverage**: Generación de reportes de cobertura de código con PHPUnit
- **Puerto**: 9003 (estándar de Xdebug 3)

### Configuración de tu IDE

#### PHPStorm / IntelliJ IDEA

1. **Configurar servidor PHP**:
   - Ve a `Settings/Preferences` → `PHP` → `Servers`
   - Añade un nuevo servidor con nombre: `refactoring-php`
   - Host: `localhost`
   - Puerto: `9003`
   - Debugger: `Xdebug`
   - Marca "Use path mappings"
   - Mapea la carpeta del proyecto a `/app`

2. **Configurar Xdebug**:
   - Ve a `Settings/Preferences` → `PHP` → `Debug`
   - Puerto Xdebug: `9003`
   - Marca "Can accept external connections"

3. **Iniciar debugging**:
   - Haz clic en el icono de "Start Listening for PHP Debug Connections" (teléfono)
   - Establece breakpoints en tu código
   - Ejecuta tu script o test dentro del contenedor

#### Visual Studio Code

1. **Instalar extensión**: "PHP Debug" de Xdebug

2. **Crear configuración** (`.vscode/launch.json`):
```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Listen for Xdebug",
            "type": "php",
            "request": "launch",
            "port": 9003,
            "pathMappings": {
                "/app": "${workspaceFolder}/php"
            }
        }
    ]
}
```

3. **Iniciar debugging**:
   - Presiona F5 o ve a Run → Start Debugging
   - Establece breakpoints
   - Ejecuta tu código dentro del contenedor

### Activar Xdebug para una sesión

Xdebug está configurado en modo `trigger`, lo que significa que necesitas activarlo explícitamente:

#### Para Code Coverage (PHPUnit)

La cobertura ya está configurada en el script `composer test:coverage` con `XDEBUG_MODE=coverage`, así que solo ejecuta:
```bash
docker compose exec php composer test:coverage
```

#### Para Debugging en tiempo real

##### Opción 1: Variable de entorno (Recomendado)
```bash
docker compose exec php bash
XDEBUG_MODE=debug php script.php
```

##### Opción 2: Trigger para sesión específica
```bash
docker compose exec php bash
XDEBUG_TRIGGER=1 php script.php
```

### Verificar que Xdebug está funcionando

```bash
docker compose exec php php -v
```

Deberías ver algo como:
```
PHP 8.2.x (cli) (built: ...)
    with Xdebug v3.3.1, Copyright (c) 2002-2023, by Derick Rethans
```

### Logs de Xdebug

Si tienes problemas de conexión, revisa los logs:
```bash
docker compose exec php cat /tmp/xdebug.log
```

### Desactivar Xdebug temporalmente

Si Xdebug ralentiza la ejecución y no lo necesitas:

```bash
docker compose exec php php -dzend_extension= script.php
```

### Code Coverage con Xdebug

La cobertura de código funciona automáticamente con PHPUnit:

```bash
# Generar reporte HTML y texto
docker compose exec php composer test:coverage

# Ver reporte HTML
open coverage/html/index.html  # macOS
xdg-open coverage/html/index.html  # Linux
```

El reporte incluirá:
- **HTML**: Reporte visual navegable en `coverage/html/`
- **Texto**: Resumen en la consola
- **Clover XML**: Para integración con CI/CD en `coverage/clover.xml`

### Troubleshooting Xdebug

#### Error: "No code coverage driver available"

Esto significa que Xdebug no está activo o no detecta el modo coverage. **Soluciones**:

1. **Reconstruir el contenedor** (si es la primera vez o tras cambios):
   ```bash
   docker compose down
   docker compose build --no-cache
   docker compose up -d
   docker compose exec php composer install
   ```

2. **Verificar que Xdebug está instalado**:
   ```bash
   docker compose exec php php -v
   ```
   Debe mostrar: `with Xdebug v3.3.1`

3. **Verificar que Xdebug está en modo coverage**:
   ```bash
   docker compose exec php php -i | grep xdebug.mode
   ```
   Debe mostrar: `develop,debug,coverage`

4. **Usar la variable de entorno explícitamente**:
   ```bash
   docker compose exec php bash -c "XDEBUG_MODE=coverage composer test:coverage"
   ```

#### El debugger no se conecta

1. **Verificar que el puerto está expuesto**:
   ```bash
   docker compose ps
   ```
   Debe mostrar: `0.0.0.0:9003->9003/tcp`

2. **Verificar logs de Xdebug**:
   ```bash
   docker compose exec php cat /tmp/xdebug.log
   ```

3. **Verificar firewall**: Asegúrate de que el puerto 9003 no está bloqueado

## Contenido

### Mantenimiento diario de código: Calistenia

Un conjunto de reglas para escribir código nuevo o evaluar código existente y modificarlo para acercarlo a un mejor diseño.

Ejercicios disponibles en `src/calisthenics-exercises/`:
- `OneIndentationLevel.php` - Un solo nivel de indentación
- `DontUseElse.php` - No usar else
- `OnlyOneDotPerLine.php` - Solo un punto por línea (Ley de Demeter)
- `NoPrimitives.php` - No usar primitivos
- `NotMoreThanTwoInstanceVariables.php` - No más de dos variables de instancia
- `NoGettersOrSetters.php` - No getters ni setters
- `FirstClassCollections.php` - Colecciones de primera clase
- `SmallEntities.php` - Entidades pequeñas
- `DontUseAbbreviations.php` - No usar abreviaturas

### Code Smells

En estos ejercicios de Code Smells se presenta cada _smell_ con un ejemplo de código y se propone un ejercicio.

Cada ejercicio presenta una dificultad debida al _code smell_, que deberías abordar primero con un refactor para reducir el coste de cambio.

Sugerencias para realizar los ejercicios:

1. Introduce tests para caracterizar el comportamiento actual del código
2. Intenta resolver el ejercicio sin refactorizar primero.
3. Realiza un refactor para reducir el coste del cambio.
4. Completa el ejercicio tras el refactor.

#### Bloaters

Code smells en los que se complica el cambio en el código por producir unidades demasiado grandes o por obligarnos a introducir mucho código auxiliar.

- `DataClump.php` - Grupos de datos que aparecen juntos repetidamente
- `PrimitiveObsession.php` - Uso excesivo de tipos primitivos
- `LongParameterList.php` - Lista larga de parámetros

#### Change Preventers

Code smells que hacen que cualquier cambio sea costoso e incluso arriesgado al obligarnos a intervenir en muchos lugares del código a la vez.

- `DivergentChange.php` - Cambios divergentes
- `ShotgunSurgery.php` - Cirugía de escopeta
- `ParallelInheritanceHierarchy.php` - Jerarquías de herencia paralelas

#### Couplers

Code smells en los que cambios en una unidad fuerzan cambios en otra que tiene un acoplamiento muy fuerte.

- `FeatureEnvy.php` - Envidia de características
- `InappropriateIntimacy.php` - Intimidad inapropiada
- `MessageChains.php` - Cadenas de mensajes
- `Middleman.php` - Intermediario innecesario

#### Dispensables

Code smells debidos a código innecesario, que introduce ruido dificultando la inteligibilidad del código.

- `Comments.php` - Comentarios excesivos
- `DataClass.php` - Clase de datos
- `DeadCode.php` - Código muerto
- `DuplicatedCode.php` - Código duplicado
- `LazyClass.php` - Clase perezosa

#### OOP Abusers

Code smells debido a la aplicación inadecuada de la orientación a objetos.

- `AlternativeClassesDifferentInterfaces.php` - Clases alternativas con interfaces diferentes
- `RefusedBequest.php` - Herencia rechazada
- `SwitchStatements.php` - Uso excesivo de switch
- `TemporalInstanceVariables.php` - Variables de instancia temporales

## Estructura del Proyecto

```
php/
├── src/
│   ├── calisthenics-exercises/    # Ejercicios de calistenia
│   ├── code-smells/                # Ejemplos de code smells
│   │   ├── bloaters/
│   │   ├── change-preventers/
│   │   ├── couplers/
│   │   ├── dispensables/
│   │   └── oop-abusers/
│   └── refactoring/                # Técnicas de refactoring
├── tests/                          # Tests unitarios
├── composer.json                   # Dependencias PHP
├── phpunit.xml                     # Configuración PHPUnit
├── phpstan.neon                    # Configuración PHPStan
├── Dockerfile                      # Imagen Docker
└── docker-compose.yml             # Orquestación Docker
```

## Detener el Entorno Docker

```bash
docker compose down
```

## Notas

- Este proyecto usa PHP 8.2+ con características modernas como readonly properties y tipos de unión
- Se incluye PHPUnit para testing
- PHPStan está configurado en nivel 8 para análisis estático estricto
- PHP_CodeSniffer está configurado con el estándar PSR-12
