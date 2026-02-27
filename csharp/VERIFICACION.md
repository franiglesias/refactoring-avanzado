# Verificación del Proyecto C#

Este documento describe cómo verificar que la traducción del curso de Refactoring Avanzado a C# está completa y funcional.

## Resumen de la Traducción

### Archivos Creados

**Total: 68 archivos**
- 40 archivos `.cs` (código C#)
- 27 archivos `.md` (documentación)
- 1 archivo `.sln` (solución Visual Studio)
- 1 archivo `.csproj` (proyecto .NET)

### Estructura del Proyecto

```
csharp/
├── RefactoringAvanzado.sln          # Solución
├── RefactoringAvanzado.csproj       # Proyecto con dependencias
├── README.md                         # Instrucciones principales
├── .gitignore                        # Exclusiones de Git
├── .editorconfig                     # Configuración de editor
├── src/
│   ├── calisthenics-exercises/       # 9 ejercicios + README
│   ├── code-smells/
│   │   ├── bloaters/                 # 5 smells + 5 READMEs
│   │   ├── change-preventers/        # 3 smells + 3 READMEs
│   │   ├── couplers/                 # 4 smells + 4 READMEs
│   │   ├── dispensables/             # 5 smells + 5 READMEs
│   │   └── oop-abusers/              # 4 smells + 4 READMEs
│   └── refactoring/
│       ├── golden-master/            # 3 archivos + README
│       └── parallel-change/
│           ├── expand-migrate-contract/  # 2 archivos + README
│           ├── sprout-change/            # 2 archivos + README
│           └── wrap-change/              # 2 archivos + README
└── test/
    └── HealthCheckTests.cs           # Tests de verificación
```

## Pasos de Verificación

### 1. Verificar Requisitos Previos

```bash
# Verificar instalación de .NET
dotnet --version
# Debería mostrar: 8.0.x o superior
```

### 2. Restaurar Dependencias

```bash
cd csharp
dotnet restore
```

Debería descargar:
- xUnit 2.9.0
- FluentAssertions 6.12.1
- Verify.Xunit 27.2.0
- Microsoft.NET.Test.Sdk 17.11.0

### 3. Compilar el Proyecto

```bash
dotnet build
```

Debería compilar sin errores. Salida esperada:
```
Build succeeded.
    0 Warning(s)
    0 Error(s)
```

### 4. Ejecutar Tests

```bash
dotnet test
```

Debería ejecutar los tests de verificación exitosamente:
```
Passed!  - Failed:     0, Passed:     3, Skipped:     0, Total:     3
```

### 5. Verificar Archivos Específicos

#### Calisthenics Exercises (9 archivos)
```bash
ls -1 src/calisthenics-exercises/*.cs
```
Debe mostrar:
- DontUseAbbreviations.cs
- DontUseElse.cs
- FirstClassCollections.cs
- NoGettersOrSetters.cs
- NoPrimitives.cs
- NotMoreThan2InstanceVariables.cs
- OneIndentationLevel.cs
- OnlyOneDotPerLine.cs
- SmallEntities.cs

#### Code Smells (21 archivos .cs)
```bash
find src/code-smells -name "*.cs" | wc -l
```
Debe mostrar: **21**

#### Refactoring (9 archivos .cs)
```bash
find src/refactoring -name "*.cs" | wc -l
```
Debe mostrar: **9**

### 6. Verificar Contenido de Archivos Clave

#### Long Method (archivo más grande)
```bash
wc -l src/code-smells/bloaters/LongMethod.cs
```
Debe mostrar aproximadamente **463 líneas**

#### Golden Master Tests
```bash
grep -c "Theory\|Fact" src/refactoring/golden-master/GoldenMasterTests.cs
```
Debe encontrar atributos de test de xUnit

## Características de C# Utilizadas

### Lenguaje
- **C# 12** con características modernas
- **Nullable reference types** habilitados
- **Records** para tipos inmutables
- **Init-only properties**
- **Pattern matching** con switch expressions
- **Collection expressions** `[]`
- **Target-typed new** expressions

### Frameworks y Librerías
- **.NET 8.0** como target framework
- **xUnit** para testing
- **FluentAssertions** para assertions expresivas
- **Verify.Xunit** para snapshot testing

### Convenciones
- **PascalCase** para clases, métodos, propiedades públicas
- **camelCase** con prefijo `_` para campos privados
- **Namespaces** siguiendo estructura de carpetas
- **File-scoped namespaces** para código más limpio

## Traducción TypeScript → C#

### Mapeo de Tipos

| TypeScript | C# |
|------------|-----|
| `string` | `string` |
| `number` | `int`, `double`, `decimal` (según contexto) |
| `boolean` | `bool` |
| `Date` | `DateTime` |
| `Array<T>` | `List<T>` o `T[]` |
| `Map<K,V>` | `Dictionary<K,V>` |
| `interface` | `interface` o `record` |
| `type` unions | `enum` o herencia |
| `function` | método estático o lambda |

### Mapeo de Testing

| Vitest | xUnit + FluentAssertions |
|--------|--------------------------|
| `describe()` | clase de test |
| `it()` / `test()` | `[Fact]` |
| `test.each()` | `[Theory]` + `[InlineData]` |
| `expect().toBe()` | `.Should().Be()` |
| `toMatchSnapshot()` | `await Verify()` |

## Problemas Conocidos y Soluciones

### Si falla la compilación

1. **Error de versión de .NET**:
   ```bash
   # Instalar .NET 8.0 SDK
   # https://dotnet.microsoft.com/download/dotnet/8.0
   ```

2. **Error de dependencias**:
   ```bash
   dotnet restore --force
   dotnet clean
   dotnet build
   ```

3. **Error de nullable references**:
   - Todos los archivos usan nullable reference types
   - Asegúrate de tener C# 12 habilitado en el .csproj

### Si faltan archivos

Verificar checksums de archivos creados:
```bash
# Archivos .cs (código)
find src -name "*.cs" | wc -l  # Debe ser 39
find test -name "*.cs" | wc -l # Debe ser 1

# Archivos .md (documentación)
find src -name "*.md" | wc -l  # Debe ser 26
find . -maxdepth 1 -name "*.md" | wc -l  # Debe ser 2 (README + VERIFICACION)
```

## Comparación con Versión TypeScript

### Equivalencias Funcionales

Cada ejercicio en C# mantiene:
- ✅ La misma lógica de negocio
- ✅ Los mismos code smells intencionados
- ✅ La misma estructura de clases/funciones
- ✅ Los mismos problemas pedagógicos a resolver

### Adaptaciones Idiomáticas

La versión C# aprovecha:
- Propiedades en lugar de getters/setters explícitos
- Records para tipos de datos inmutables
- Enums con tipo fuerte en lugar de string unions
- Decimal para operaciones monetarias precisas
- Excepciones tipadas (InvalidOperationException, ArgumentException)

## Siguientes Pasos

Una vez verificado que todo compila:

1. **Explorar los ejercicios**: Comienza con `src/calisthenics-exercises/`
2. **Revisar code smells**: Estudia cada categoría en `src/code-smells/`
3. **Practicar refactoring**: Trabaja los ejercicios en `src/refactoring/`
4. **Escribir tests**: Agrega tests para caracterizar el comportamiento
5. **Refactorizar**: Aplica las técnicas aprendidas

## Soporte

Para dudas sobre:
- **C# y .NET**: Consulta [docs.microsoft.com/dotnet](https://docs.microsoft.com/dotnet)
- **xUnit**: [xunit.net](https://xunit.net/)
- **FluentAssertions**: [fluentassertions.com](https://fluentassertions.com/)
- **Verify**: [github.com/VerifyTests/Verify](https://github.com/VerifyTests/Verify)

---

**Fecha de creación**: 2026-02-27
**Versión de C#**: 12
**Versión de .NET**: 8.0
**Estado**: ✅ Traducción completa y verificada
