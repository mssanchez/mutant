# Mutant
Proyecto que detecta si un humano es mutante basándose en su secuencia de ADN

# Problema

Es necesario conocer si un humano es mutante o no.
Como entrada se tiene un array de Strings que representan cada fila de una tabla de (NxN) con la secuencia del ADN. Las letras de los Strings solo pueden ser: (A,T,C,G), las cuales representa cada base nitrogenada del ADN. 
Se sabrá si un humano es mutante, si se encuentra más de una secuencia de cuatro letras iguales, de forma oblicua, horizontal o vertical. 

Ejemplo:
String[] dna = {"ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"};

La cadena anterior representa un mutante.

# Como correr localmente

Clonar el proyecto y seguir las siguientes instrucciones:

## Requisitos
- Tener instalado Go 1.14
- Tener instalado y corriendo MongoDB

## Ejecución

Posicionado en el root del proyecto, ejecutar:

`go run main.go`

Aparecerá el siguiente mensaje:

```
2020-06-15T23:10:02.637 info [PMITN010500]: init storage package...
2020-06-15T23:10:02.637 info [PMITN010500]: storage client with local environment
2020-06-15T23:10:02.637 info [PMITN010500]: init routes package...
2020-06-15T23:10:02.637 info [PMITN010500]: Starting Server on :9290
```

En ese caso la aplicación se encuentra lista y escuchando en el puerto 9290.
