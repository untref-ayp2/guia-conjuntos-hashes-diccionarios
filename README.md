# Guía: Conjuntos, Hash Tables y Diccionarios

## Conjuntos

1. Implementar la operaciones `Union`, `Intersection`, `Difference`, `SymmetricDifference`, `Equal`, `Subset` y `Superset` para el conjunto de enteros dado como ejemplo ([`set/intset.go`](set/intset.go)).
2. Implementar un conjunto genérico que permita almacenar cualquier tipo de dato ([`set/set.go`](set/set.go)).
3. Implementar un conjunto basado en arreglos dinámicos. Implementar tantos las operaciones sobre elementos como las operaciones entre conjuntos.
4. Implementar un conjunto basado en listas enlazadas. Implementar tantos las operaciones sobre elementos como las operaciones entre conjuntos ([`set/listset.go`](set/listset.go)).
5. Comparar el orden de las operaciones entre las implementaciones de arreglos y listas enlazadas.

## _Hashe Tables_

1. Modificar la tabla de _hash_ abierto para que las claves puedan ser de distintos tipos (usar el paquete [`maphash`](https://pkg.go.dev/hash/maphash) de Go).

2. Implementar una tabla de _hash_ cerrada. Para ello se debe implementar una lista enlazada que almacene los elementos en cada posición del arreglo. Cuando se produce una colisión, el nuevo elemento se agrega a la lista en la posición correspondiente. La tabla debe tener los mismos métodos que la tabla de _hash_ abierta: `Put`, `Get`, `Remove`, `Keys`, `Values`, `Size`, `IsEmpty`, `Clear` y `String`. Las claves deben ser de cualquier tipo.

3. Escribir casos de pruebas que cubran todas las operaciones de los puntos anteriores.

## Diccionarios

1. Implementar un diccionario sobre una tabla de hashing genérica, como la vista en Tablas de Hashing.

2. Analizar el orden de todas las operaciones.

3. Escriba una función que reciba un texto en castellano y un diccionario, y devuelva la oración traducida palabra por palabra según el diccionario. Si una palabra no se encuentra en el diccionario, deberá sustituirla por "error" en la cadena resultante.

   > Para parsear el string pueden usar: `func Split(str, sep string) []string`

4. Escriba una función que reciba una cadena de caracteres y calcule la frecuencia de aparición de cada caracter.

   Por ejemplo:

   ```go
   Frecuencia("hola como estas...")
   ```

   debe devolver

   ```text
   {a: 2, c: 1, e: 1, h: 1, l: 1, m: 1, o: 3, s: 2, t: 1}
   ```

   > Recuerden que al recorrer una cadena deben castear cada valor a string ya que GO representa los caracteres como enteros. No se deben incluir los blancos en el conjunto. Por lo tanto se recomienda usar [`stringUtils`](https://pkg.go.dev/github.com/agrison/go-commons-lang/stringUtils) para chequear si un caracter es un blanco (espacio, tabulador, salto de línea, etc.)

5. Escriba una función que reciba 2 slices y devuelva una nueva lista que es el resultado de la interseccion de los 2 slices anteriores, la complejidad del metodo debe ser _O(n)_

6. Tenemos que ayudar a los docentes a preparar la información solicitada por el departamento de alumnos.

   Cuando toman asistencia anotan los presentes por día de clase, pero ahora desde alumnos se les pide que envíen la información de asistencia por alumno.

   Por ejemplo, si la lista que se recibe es:

   ```json
   {"Mie 10": ["Ana", "Pedro"], "Vie 12":, ["Ana", "Luz"], "Mie 17":, ["Luz", "Pedro"]}
   ```

   el resultado debe ser:

   ```json
   {
     "Ana": ["Mie 10", "Vie 12"],
     "Luz": ["Vie 12", "Mie 17"],
     "Pedro": ["Mie 10", "Mie 17"]
   }
   ```
