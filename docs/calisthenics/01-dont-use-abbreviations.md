# No Usar Abreviaturas

Regla 1 de Object Calisthenics

## Definición

Los identificadores (nombres de variables, métodos, clases, parámetros, etc.) no deben usar abreviaturas. Todos los nombres deben escribirse completos para transmitir claramente su intención y significado.

## Descripción

Una **abreviatura** es cualquier acortamiento de una palabra o frase que sacrifica claridad por brevedad. En el pasado, cuando los editores tenían limitaciones de espacio en pantalla y el autocompletado no existía, las abreviaturas podían tener sentido. Hoy en día, son obstáculos innecesarios para la comprensión.

El problema con las abreviaturas es que su significado es ambiguo y depende del contexto:

- `usr` podría ser "user", "usuario" o "user service repository"
- `calc` podría ser "calculate", "calculator" o "calculation"
- `msg` podría ser "message", "messaging" o "messenger"
- `addr` podría ser "address" o "addresser"

Los nombres completos eliminan esta ambigüedad y hacen que el código sea **autoexplicativo**. El código debería leerse como prosa, no como un crucigrama que hay que descifrar.

Además, las abreviaturas suelen ser **inconsistentes** dentro del mismo código base: un desarrollador usa `msg`, otro usa `message`, y un tercero usa `msj`. Esta inconsistencia aumenta la carga cognitiva y dificulta la búsqueda de código.

## Síntomas

- Identificadores con menos de 4 caracteres (con excepciones como `id`, `url`, `uri`)
- Nombres que requieren consultar documentación o contexto para entender
- Inconsistencia en abreviaturas (`usr` en un lugar, `user` en otro)
- Comentarios que explican qué significa una abreviatura
- Abreviaturas "estándar" que no son universales (`qty`, `amt`, `addr`)
- Uso de consonantes sin vocales (`fltr`, `cnfg`, `mgr`)
- Abreviaturas de dominio que no son universales en la empresa
- Nombres que no se pueden pronunciar en voz alta
- Dificultad para buscar código (buscar "message" no encuentra "msg")

## Ejemplo

### Antes (Violación)

```pseudocode
class UsrSvcRepo {
  method getPwdByUsr(usrId) {
    usr = fetchUsrFromDb(usrId)
    if (usr.isPremiumUsr()) {
      return usr.encPwd
    }
    return usr.pwd
  }

  method updUsrAddr(usrId, newAddr) {
    usr = fetchUsrFromDb(usrId)
    usr.addr = newAddr
    saveUsrToDb(usr)
  }

  method calcUsrDisc(usrId, amt) {
    usr = fetchUsrFromDb(usrId)
    discPct = getDiscPctByUsrLvl(usr.lvl)
    return amt * discPct / 100
  }
}
```

### Después (Cumplimiento)

```pseudocode
class UserServiceRepository {
  method getPasswordByUser(userId) {
    user = fetchUserFromDatabase(userId)
    if (user.isPremiumUser()) {
      return user.encryptedPassword
    }
    return user.password
  }

  method updateUserAddress(userId, newAddress) {
    user = fetchUserFromDatabase(userId)
    user.address = newAddress
    saveUserToDatabase(user)
  }

  method calculateUserDiscount(userId, amount) {
    user = fetchUserFromDatabase(userId)
    discountPercentage = getDiscountPercentageByUserLevel(user.level)
    return amount * discountPercentage / 100
  }
}
```

**Diferencias clave**:
- `UsrSvcRepo` → `UserServiceRepository`: Claridad total sobre el propósito
- `getPwdByUsr` → `getPasswordByUser`: Sin ambigüedad
- `encPwd` → `encryptedPassword`: Especifica que está encriptada
- `updUsrAddr` → `updateUserAddress`: Acción y objeto claros
- `calcUsrDisc` → `calculateUserDiscount`: Verbo completo, sustantivo completo
- `discPct` → `discountPercentage`: Unidad explícita
- `amt` → `amount`: Evita confusión con otras abreviaturas

## Ejercicio

**Tarea**: En el código proporcionado en tu lenguaje, expande todas las abreviaturas para que el código sea completamente autoexplicativo. El código debe leerse como prosa natural.

**Criterios de éxito**:
1. No quedan abreviaturas ambiguas
2. Todos los nombres transmiten claramente su intención
3. El código puede leerse en voz alta sin problemas
4. No se necesitan comentarios para explicar nombres
5. Los nombres son consistentes en todo el código

## Problemas que Encontrarás

### 1. Tentación de "es obvio por contexto"

Pensarás que `usr` es obvio en `UserService`, pero recuerda: el código se lee mucho más que se escribe, y no siempre tendrás el contexto completo.

### 2. Nombres muy largos

Al expandir todo, algunos nombres pueden parecer excesivamente largos (`getUserDiscountPercentageByMembershipLevel`). Esto es síntoma de que el código tiene otros problemas de diseño (probablemente viola Single Responsibility Principle).

### 3. Abreviaturas "estándar"

Abreviaturas como `qty`, `amt`, `addr` parecen estándar, pero no lo son universalmente. `quantity`, `amount`, `address` son más claros.

### 4. Abreviaturas técnicas

Acrónimos técnicos como `HTTP`, `JSON`, `URL`, `ID` están ampliamente aceptados y deben mantenerse. La regla se aplica a abreviaturas inventadas o ambiguas.

### 5. Conflicto con convenciones de la industria

En algunos dominios hay abreviaturas establecidas (medicina, finanzas). Usa criterio: si la abreviatura es universal en tu dominio y todos en el equipo la entienden inmediatamente, puede quedar.

## Proceso de Aplicación

### 1. Identificar todas las abreviaturas

- Recorre el código buscando identificadores cortos o sin vocales
- Marca cualquier nombre que no sea inmediatamente claro
- Incluye parámetros, variables locales, campos, métodos y clases

### 2. Determinar el significado completo

- Para cada abreviatura, identifica qué palabra completa representa
- Consulta con el equipo si hay ambigüedad
- Documenta el significado antes de cambiar (para referencia)

### 3. Aplicar Rename (usando herramientas de refactoring)

- Usa la función "Rename" de tu IDE para cambios seguros
- Esto actualizará todas las referencias automáticamente
- Hazlo de manera incremental, abreviatura por abreviatura

### 4. Verificar consistencia

- Asegúrate de que el mismo concepto se nombre igual en todo el código
- Si tenías `usr` y `user`, ahora todo debe ser `user`
- Busca variaciones y unifícalas

### 5. Revisar legibilidad

- Lee el código en voz alta
- Si algo suena raro o artificial, probablemente haya un problema de diseño más profundo
- Considera si métodos/clases necesitan refactoring adicional

### 6. Actualizar tests

- Los tests también deben tener nombres completos
- Los nombres de tests son especialmente importantes porque documentan comportamiento
- `testUsrLogin` → `testUserLogin` o mejor: `shouldAuthenticateUserWithValidCredentials`

## Técnicas de Refactoring Aplicables

- **Rename Variable**: Cambiar nombre de variable local
- **Rename Field**: Cambiar nombre de campo/propiedad
- **Rename Method**: Cambiar nombre de método/función
- **Rename Class**: Cambiar nombre de clase/tipo
- **Rename Parameter**: Cambiar nombre de parámetro
- **Extract Variable**: Si expandir el nombre revela complejidad, extraer a variable con nombre descriptivo

## Beneficios

### 1. Código Autoexplicativo

El código se lee como prosa, reduciendo la necesidad de comentarios y documentación externa.

### 2. Onboarding Más Rápido

Nuevos desarrolladores pueden entender el código sin necesidad de aprender un diccionario de abreviaturas del proyecto.

### 3. Búsqueda Más Efectiva

Buscar "user" encuentra todas las ocurrencias, no solo las que usaron esa abreviatura particular.

### 4. Menos Errores

No hay confusión sobre qué significa cada abreviatura, reduciendo malentendidos que llevan a bugs.

### 5. Mejor Autocompletado

Los IDEs modernos pueden sugerir nombres completos más efectivamente que abreviaturas ambiguas.

### 6. Code Review Más Fácil

Los revisores pueden entender el código sin necesitar contexto adicional sobre abreviaturas internas.

### 7. Mantenibilidad a Largo Plazo

El código sigue siendo comprensible años después, incluso cuando los autores originales ya no están.

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/calisthenics-exercises/dont-use-abbreviations.ts)
- [Go](../../go/calisthenics_exercises/09_no_abbreviations.go)
- [Java](../../java/src/main/java/com/refactoring/calisthenics/NoAbbreviations.java)
- [PHP](../../php/src/calisthenics-exercises/DontUseAbbreviations.php)
- [Python](../../python/src/calisthenics_exercises/dont_use_abbreviations.py)
- [C#](../../csharp/src/calisthenics-exercises/DontUseAbbreviations.cs)

## Referencias en Español

- [Naming things](https://franiglesias.github.io/naming-things/) - Guía completa sobre cómo nombrar elementos en programación
- [Object Calisthenics para adelgazar las clases](https://franiglesias.github.io/calistenics-and-small-classes/) - Incluye discusión sobre nombres expresivos

## Referencias

- **"The ThoughtWorks Anthology"** - Jeff Bay - Capítulo original sobre Object Calisthenics
- **"Clean Code"** - Robert C. Martin - Capítulo 2: "Meaningful Names"
- [Object Calisthenics - William Durand](https://williamdurand.fr/2013/06/03/object-calisthenics/) - Regla #1
- [Naming Things in Code](https://www.martinfowler.com/bliki/TwoHardThings.html) - Martin Fowler sobre la dificultad de nombrar
- **"Code Complete"** - Steve McConnell - Capítulo sobre convenciones de nombres
