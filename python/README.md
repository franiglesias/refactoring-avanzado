# Curso Refactoring Avanzado (Python)

Ejemplos y ejercicios del Curso de Refactoring Avanzado convertidos a Python.

## Preparación

### Instalar dependencias

```bash
pip install -r requirements.txt
```

### Ejecutar tests

```bash
# Ejecutar todos los tests
pytest

# Ejecutar tests con cobertura
pytest --cov=src --cov-report=html

# Ejecutar tests específicos
pytest src/refactoring/test_golden_master.py -v
```

## Contenido

### Técnicas de Refactoring

#### Golden Master

Técnica para caracterizar el comportamiento de código legado sin tests.

- [Ejercicio Golden Master](./src/refactoring/README.md)

```bash
pytest src/refactoring/test_golden_master.py -v
```

#### Parallel Change

Técnicas para realizar cambios seguros en código en producción.

- Expand-Migrate-Contract
- Sprout Change
- Wrap Change

```bash
pytest src/refactoring/parallel_change/ -v
```

### Mantenimiento diario de código: Calistenia

Un conjunto de reglas para escribir código nuevo o evaluar código existente y modificarlo para acercarlo a un mejor diseño.

Ejercicios en [src/calisthenics_exercises/](./src/calisthenics_exercises):

1. Un nivel de indentación por método
2. No uses ELSE
3. Envuelve primitivos
4. Colecciones de primera clase
5. Un punto por línea
6. No uses abreviaciones
7. Mantén las entidades pequeñas
8. No más de 2 variables de instancia
9. Sin getters ni setters

### Code Smells

En estos ejercicios de Code Smells se presenta cada _smell_ con un ejemplo de código y se propone un ejercicio.

Cada ejercicio presenta una dificultad debida al _code smell_, que deberías abordar primero con un refactor para reducir el coste de cambio.

Sugerencias para realizar los ejercicios:

1. Introduce tests para caracterizar el comportamiento actual del código
2. Intenta resolver el ejercicio sin refactorizar primero
3. Realiza un refactor para reducir el coste del cambio
4. Completa el ejercicio tras el refactor

#### Bloaters

Code smells en los que se complica el cambio en el código por producir unidades demasiado grandes o por obligarnos a introducir mucho código auxiliar.

- [Data clump](src/code_smells/bloaters/data_clump.py) - Grupos de datos que aparecen juntos repetidamente
- [Large class](src/code_smells/bloaters/large_class.py) - Clases que hacen demasiadas cosas
- [Long method](src/code_smells/bloaters/long_method.py) - Métodos excesivamente largos (ejercicio final recomendado)
- [Long parameter list](src/code_smells/bloaters/long_parameter_list.py) - Funciones con muchos parámetros
- [Primitive obsession](src/code_smells/bloaters/primitive_obsession.py) - Uso excesivo de tipos primitivos

#### Couplers

Code smells en los que cambios en una unidad fuerzan cambios en otra que tiene un acoplamiento muy fuerte.

- [Feature envy](src/code_smells/couplers/feature_envy.py) - Métodos más interesados en otras clases

#### Dispensables

Code smells debidos a código innecesario, que introduce ruido dificultando la inteligibilidad del código.

- [Duplicated code](src/code_smells/dispensables/duplicated_code.py) - Código duplicado en múltiples lugares

#### OOP Abusers

Code smells debido a la aplicación inadecuada de la orientación a objetos.

- [Switch statements](src/code_smells/oop_abusers/switch_statements.py) - Uso de if-elif en lugar de polimorfismo

## Estructura del Proyecto

```
src/
├── refactoring/              # Técnicas de refactoring
│   ├── golden_master.py
│   ├── test_golden_master.py
│   └── parallel_change/
├── calisthenics_exercises/   # Ejercicios de Object Calisthenics
└── code_smells/              # Ejemplos de Code Smells
    ├── bloaters/
    ├── couplers/
    ├── dispensables/
    └── oop_abusers/
```

## Comandos Útiles

```bash
# Instalar dependencias
pip install -r requirements.txt

# Ejecutar todos los tests
pytest

# Ejecutar tests en modo watch (requiere pytest-watch)
pip install pytest-watch
ptw

# Ver cobertura de tests
pytest --cov=src --cov-report=term-missing

# Formatear código con black
pip install black
black src/

# Verificar tipos con mypy
pip install mypy
mypy src/
```

## Versión de Python

Este proyecto requiere Python 3.11 o superior.

## Diferencias con la versión TypeScript

- Se usa `dataclass` en lugar de tipos TypeScript
- `snake_case` en lugar de `camelCase` para nombres de variables y funciones
- `pytest` en lugar de `vitest` para tests
- Uso de type hints de Python (typing module) en lugar de tipos de TypeScript
