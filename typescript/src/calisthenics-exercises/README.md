# Calistenia

Un conjunto de nueve reglas propuestas por Jeff Bay que, aplicadas a nuestro código, pueden mejorar su calidad y aproximarlo a un mejor diseño.

Se basan en fijarnos en ciertas características aplicables a cualquier código.

Si observamos que el código incumple alguna de esas reglas, intentamos arreglarlo para que se ajuste a ella.

Las reglas son:

1. No usar abreviaturas
2. No usar ELSE
3. Un solo nivel de indentación
4. Encapsular primitivos
5. Encapsular colecciones
6. No usar getters y setters
7. Mantener las unidades de código pequeñas
8. Máximo de dos variables de instancia por clase
9. No más de un punto por línea

## 1. No usar abreviaturas

Los identificadores abreviados oscurecen la intención y hacen que sea más costoso comprender los conceptos que maneja el código y como se utiliza.

Buscamos expandir los nombres para transmitir la intención.

Aceptación: No quedan abreviaturas ambiguas, el código se lee casi como si fuera prosa.

Refactor:

* Renombrar identificadores.

[Ejercicio](dont-use-abbreviations.ts)

## 2. No usar ELSE

Las cláusulas ELSE ocultan reglas importantes y dificultan la comprensión del flujo del código.

Buscamos reducir la cantidad de ELSE para mejorar la legibilidad y comprensión del código.

Aceptación: El código no tiene ELSE, facilitando la comprensión y lectura.

Refactor:

* Aplicar retorno temprano.
* Reemplazar con cláusulas de guarda.

[Ejercicio](dont-use-else.ts)

## 3. Un solo nivel de indentación

Múltiples niveles de indentación en un bloque de código revelan una mezcla de nivel de abstracción.

Buscamos reducir la cantidad de niveles de indentación para mejorar la legibilidad y comprensión del código.

Aceptación: El código se organiza en bloques con un solo nivel de indentación, facilitando la comprensión y lectura.

Refactor:

* Reemplazar condicionales anidadas con cláusulas de guarda.
* Descomponer condicionales.

[Ejercicio](one-indentation-level.ts)

## 4. Empaquetar primitivos

Usar tipos primitivos del lenguaje para representar conceptos del dominio suele generar duplicidad de código y requerir técnicas de programación defensiva, ya que no podemos confiar en su validez o integridad.

Buscamos introducir objetos que representen y protejan las invariantes de los conceptos del dominio, lo que nos llevará a un código más simple y limpio.

Refactor:

* Reemplazar valor con objeto.

[Ejercicio](no-primitives.ts)

## 5. Colecciones de primera clase

Los tipos colectivos (arrays, maps, sets, etc.) son una forma de representar colecciones de objetos, pero al igual que ocurre con los tipos primitivos, no nos protegen de problemas de validez o integridad de las invariantes del dominio.

Por tanto, buscamos encapsular estos tipos colectivos en objetos que representen conceptos del dominio y protejan sus invariantes.

Refactor:

* Envolver colecciones en objetos.

[Ejercicio](first-class-collections.ts)

## 6. No usar getters y setters (ni propiedades públicas)

Exponer la estructura interna de los objetos genera acoplamiento y dificulta la evolución del código.

Buscamos evitar el uso de getters y setters, o de propiedades públicas, de modo que evitemos el acoplamiento a la estructura interna y tengamos flexibilidad para cambiar el código en el futuro.

Refactor:

* Encapsular campos.
* Introducir Doble Despacho.

[Ejercicio](no-getters-or-setters.ts)

## 7. Mantener las unidades de código pequeñas

Cuando más grande es una unidad de código, más probable es que esté haciéndose cargo de varias responsabilidades y tomando muchas decisiones, lo que aumenta la complejidad ciclomática y el riesgo de introducir errores.

Buscamos descomponer las unidades grandes en otras más pequeñas, cada una de ellas con una única responsabilidad bien definida.

[Ejercicio](small-entities.ts)

## 8. Máximo de dos variables de instancia por clase

Cuando una clase tiene muchas variables de instancia es posible que estas puedan agruparse representando conceptos. A menudo, la clase gestiona uno o dos conceptos del dominio, pero necesitamos varios valores primitivos para representar cada uno de ellos.

Buscamos agrupar estas variables de instancia en value objects que representen conceptos del dominio.

Refactor:

* Introducir value objects.

[Ejercicio](not-more-than-2-instance-variables.ts)

## 9. No más de un punto por línea

Los objetos deben ser tratados como cajas negras y no depender de nuestro conocimiento previo de su estructura interna. Si usamos elementos internos de un objeto, resultará muy difícil romper este acoplamiento en el futuro, limitando la evolución del propio objeto y del resto del código.

Por tanto, buscamos interactuar con objetos que el código conoce directamente.

Refactor:

* Esconder la delegación (con cuidado).
* Extraer o mover métodos.

[Ejercicio](only-one-dot-per-line.ts)
