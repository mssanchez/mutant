# Mutant
Proyecto que detecta si un humano es mutante basándose en su secuencia de ADN.

Escrito en [Go].

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
2020-06-27T20:42:39.882 info [PMITN010500]: init storage package...
2020-06-27T20:42:39.882 info [PMITN010500]: storage client with local environment
2020-06-27T20:42:39.883 info [PMITN010500]: init routes package...
2020-06-27T20:42:39.883 info [PMITN010500]: Starting Server on :8080
```

En ese caso la aplicación se encuentra lista y escuchando en el puerto 8080.

[Go]:https://golang.org/

# Tests

## Como ejecutar tests

Sobre el root del proyecto ejecutar:
```
go test ./...
```

## Coverage

Para evaluar el coverage del proyecto, posicionado en el root del proyecto ejecutar:
```
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```
Lo cual generará el siguiente reporte:
```
mutant/pkg/app/app.go:24:		NewApplication		100.0%
mutant/pkg/app/app.go:44:		RunServer		0.0%
mutant/pkg/app/app.go:57:		Shutdown		0.0%
mutant/pkg/config/config.go:36:		NewConfiguration	100.0%
mutant/pkg/config/config.go:49:		readApplicationConfig	88.2%
mutant/pkg/config/config.go:82:		readDatastorePassword	83.3%
mutant/pkg/errors/error.go:34:		New			100.0%
mutant/pkg/errors/error.go:39:		Newf			100.0%
mutant/pkg/errors/error.go:44:		Wrapf			100.0%
mutant/pkg/errors/error.go:49:		New			100.0%
mutant/pkg/errors/error.go:54:		Newf			100.0%
mutant/pkg/errors/error.go:59:		Wrapf			100.0%
mutant/pkg/errors/error.go:73:		AddSingleContext	100.0%
mutant/pkg/errors/error.go:84:		Context			100.0%
mutant/pkg/errors/error.go:89:		Error			100.0%
mutant/pkg/errors/error.go:94:		Type			100.0%
mutant/pkg/mutant/mutant.go:13:		NewMutant		100.0%
mutant/pkg/mutant/mutant.go:28:		IsMutant		100.0%
mutant/pkg/mutant/mutant.go:68:		isVerticalMutant	100.0%
mutant/pkg/mutant/mutant.go:86:		isObliqueMutant		100.0%
mutant/pkg/mutant/mutant.go:105:	isHorizontalMutant	100.0%
mutant/pkg/mutant/mutant.go:120:	validateSequence	100.0%
mutant/pkg/mutant/validator.go:12:	Validate		100.0%
mutant/pkg/mutant/validator.go:22:	Validate		100.0%
mutant/pkg/routes/healthCheck.go:9:	addHealthCheckRoutes	100.0%
mutant/pkg/routes/healthCheck.go:14:	healthCheck		100.0%
mutant/pkg/routes/mutant.go:17:		addMutantHandler	100.0%
mutant/pkg/routes/mutant.go:26:		Mutant			100.0%
mutant/pkg/routes/routes.go:18:		NewHTTPRoutes		0.0%
mutant/pkg/routes/routes.go:29:		AddAllHTTPRoutes	0.0%
mutant/pkg/routes/stats.go:16:		addStatsHandler		100.0%
mutant/pkg/routes/stats.go:25:		stats			100.0%
mutant/pkg/stats/stats.go:11:		NewStats		100.0%
mutant/pkg/stats/stats.go:26:		MutantStats		100.0%
mutant/pkg/stats/stats.go:44:		safeDivideFloat		100.0%
total:					(statements)		91.0%
```
El reporte y un HTML con los resultados se pueden encontrar en el directorio [coverage](coverage).

# Ejemplos

### Verificar si una cadena de ADN es mutante
```
curl -v -X POST 'https://positive-apex-280419.ue.r.appspot.com/mutant' \
--header 'Content-Type: application/json' \
--data-raw '{
  "dna": [
    "ATGCGA",
    "CAGTGC",
    "TTATGT",
    "AGAAGG",
    "CCCCTA",
    "TCACTG"
  ]
}'
```

### Visualizar estadísticas de cadenas de ADN
```
curl -v -X GET 'https://positive-apex-280419.ue.r.appspot.com/stats'
```

# Deploy

### Requisitos
- Instalar el SDK de Google Cloud siguiendo las siguientes [instrucciones](https://cloud.google.com/sdk/install?hl=es)
- Crear un archivo de nombre 'password' en el directorio /pkg/config/support/dev con las credenciales para conectarse a la base de datos.
Deberá tener el siguiente formato:
    ```
    password=some-value
    ```

### Ejecución

Ejecutar el script de [deploy](deploy.sh) parado sobre el root del proyecto.
El script pedirá confirmación antes de deployar.

Una vez finalizado, se muestran los logs de la aplicación.

## Análisis de codigo

Se utilizó la herramienta [golint](https://github.com/golang/lint).

Para ejecutar el análisis, sobre el root del proyecto ejecutar:
```
golint ./...
```