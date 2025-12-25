package main

import (
	"fmt"
	practica "practica_lenguaje_go/funciones"
)

func main() {
	m1 := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	m2 := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	ms,e := practica.SumarMatrices(m1,m2)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Print(ms)
}
