package main

import (
	"fmt"
	practica "practica_lenguaje_go/funciones"
)

func main() {
	var prom float64 = practica.CalcularPromedioNotas()
	fmt.Printf("%f\n",prom)
}
