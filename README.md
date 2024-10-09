# documento-unico-de-identidad  

[![en](https://img.shields.io/badge/lang-en-red.svg)](https://github.com/Franceskynov/documento-unico-de-identidad/blob/main/README.md)
[![es](https://img.shields.io/badge/lang-es-yellow.svg)](https://github.com/Franceskynov/documento-unico-de-identidad/blob/main/README.es.md)

A Go (golang) simple library for validate Unique Identity Document (DUI) in El Salvador


## Installation

First, we need to get the library into your project, using the next command:

```bash
go get github.com/Franceskynov/documento-unico-de-identidad
```

## Usage

To get started with the library, import the current package.

```go

	import (
	    dui "github.com/Franceskynov/documento-unico-de-identidad"
    )

```

Then you can use the library for something like

```go

    document := "00016297-5"
    err := dui.Validate(document)
	
	if err == nil {
		fmt.Printf("Current document is : %s \n and has a valid format", document)
	} else {
		fmt.Println("Please enter a valid document")
	}
```
