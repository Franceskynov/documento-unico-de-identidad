# documento-unico-de-identidad  
[![en](https://img.shields.io/badge/lang-en-red.svg)](https://github.com/Franceskynov/documento-unico-de-identidad/blob/main/README.md)
[![es](https://img.shields.io/badge/lang-es-yellow.svg)](https://github.com/Franceskynov/documento-unico-de-identidad/blob/main/README.es.md)


Es una simple librería escrita en Go (golang) para validar los DUI (Documento Unico de Identidad)


## Instalación

En primer lugar, necesitamos agregar la librería en tu proyecto, para ello se requiere de ejecutar el siguiente comando:

```bash
go get github.com/Franceskynov/documento-unico-de-identidad
```

## Uso

Para hacer uso de esta librería es necesario importarla y usar un identificador para facilitar su uso:

```go

	import (
	    dui "github.com/Franceskynov/documento-unico-de-identidad"
    )

```

Una vez que la librería esta disponible en nuestro proyecto podremos hacer uso de esta, como se muestra a continuación:

```go

    document := "00016297-5"
    err := dui.Validate(document)
	
	if err == nil {
		fmt.Printf("El documento ingresado es: %s \n y tiene un formato valido", document)
	} else {
		fmt.Println("Por favor ingrese un documento valido")
	}
```
