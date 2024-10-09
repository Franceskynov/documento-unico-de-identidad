package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	dui "github.com/Franceskynov/documento-unico-de-identidad"
)

func main() {

	fmt.Println("Por favor ingrese un DUI")
	document := Reader() 
	err := dui.Validate(document)
	
	if err == nil {
		fmt.Printf("El documento ingresado es: %s \n y tiene un formato valido", document)
	} else {
		fmt.Println("Por favor ingrese un documento valido")
	}
}

func Reader() string {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(line)
}