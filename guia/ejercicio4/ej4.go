package main

import (
	"fmt"
	"math"
	"guia/ejercicio4/composicion"
)

func f(x float64) float64 {
	// f(x) = x^2
	return x*x
}

func g(x float64) float64 {
	// g(x) = sin(x)
	return math.Sin(x)
}

func main() {
	var x float64 = 10
	comp := composicion.CrearComposicion()
	comp.AgregarFuncion(f)
	comp.AgregarFuncion(g)
	resultado := comp.Aplicar(x)
	fmt.Println(resultado)
}
