package main

import (
	//"fmt"
	practica "practica_lenguaje_go/funciones"
)

func main() {
	cta1,err1 := practica.CrearCuenta("Ignacio",100)
	cta2,err2 := practica.CrearCuenta("Otro",100)
	if err1 != nil || err2 != nil {
		return
	}
	cta1.ConsultarSaldo()
	cta2.ConsultarSaldo()
	cta1.Transferir(&cta2,5)
	cta1.ConsultarSaldo()
	cta2.ConsultarSaldo()
}
