package main

import (
	//"fmt"
	practica "practica_lenguaje_go/funciones"
)

func main() {
	libros := practica.ArregloLibroEjemplo()
	practica.BuscarPorAutor(libros,"DI IORIO Y LUCERO, MARIA EUGENIA")
}
