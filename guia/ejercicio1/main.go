package main

import (
	"fmt"
	fraccion "guia/ejercicio1/fraccion"
)

func main() {
	fraccion1 := fraccion.CrearFraccion(1,2)
	fraccion2 := fraccion.CrearFraccion(1,3)
	fmt.Println(fraccion1.Sumar(fraccion2).Representacion())
}