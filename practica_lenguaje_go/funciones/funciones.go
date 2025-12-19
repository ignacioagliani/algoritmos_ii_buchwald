package funciones
/*
Aclaración: todos los ejercicios que sean de crear un programa, los voy a crear como función :)
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math"
)

/*
Ejercicio 1.1. Escribir una función que reciba dos números y devuelva su producto.
*/

func Multiplicar(x,y int) int {
	return x * y
}

/*
Ejercicio 1.2. Utilizando la función del ejercicio anterior, escribir un programa que
pida al usuario dos números, y luego muestre el producto.
*/
func PedirYMultiplicar() {
	buffer1 := bufio.NewScanner(os.Stdin)
	buffer1.Scan()
	input1 := buffer1.Text()
	numero1,err1 := strconv.Atoi(input1)
	if err1 != nil {
		fmt.Printf("%s es un núnmero inválido\n",input1)
		return
	}
	buffer2 := bufio.NewScanner(os.Stdin)
	buffer2.Scan()
	input2 := buffer2.Text()
	numero2,err2 := strconv.Atoi(input2)
	if err2 != nil {
		fmt.Printf("%s es un número inválido\n",input2)
		return
	}
	var multiplicacion int = Multiplicar(numero1,numero2)
	fmt.Printf("Resultado: %d\n",multiplicacion)
}

/*
Ejercicio 1.5. Escribir una función que, dado un número entero 𝑛, permita calcular su factorial.
*/

func CalcularFactorialIterativo(n int) int {
	var factorial int = 1
	for i:=1; i <= n; i++ {
		factorial *= i
	}
	return factorial
}

func CalcularFactorialRecursivo(n int) int {
	if n == 0 {
		return 1
	}
	return n * CalcularFactorialRecursivo(n - 1)
}

/*
Ejercicio 2.1. Escribir una función que reciba una cantidad de pesos, una tasa de interés
y un número de años y devuelva el monto final a obtener. La fórmula a utilizar es:

𝐶𝑛 = 𝐶 × (1 + (𝑥/100))^𝑛

Donde 𝐶 es el capital inicial, 𝑥 es la tasa de interés y 𝑛 es el número de años a calcular.
*/
func CalcularInteres(capital,tasa_int,años float64) float64 {
	return capital * math.Pow(((1 + (tasa_int/100))),años)
}

/*
Ejercicio 2.2. Utilizando la función del ejercicio anterior, escribir un programa que le
pregunte al usuario la cantidad de pesos inicial, la tasa de interés y el número de años
y muestre el monto final a obtener.
*/
func CalcularInteresUsuario() {
	fmt.Printf("Capital: $")
	buffer1 := bufio.NewScanner(os.Stdin)
	buffer1.Scan()
	input1 := buffer1.Text()
	capital,err1 := strconv.ParseFloat(input1,64) // convierte a float64
	if err1 != nil || capital < 0 {
		fmt.Printf("$%s no es un capital válido\n",input1)
		return
	}
	fmt.Printf("Tasa de Interés: ")
	buffer2 := bufio.NewScanner(os.Stdin)
	buffer2.Scan()
	input2 := buffer2.Text()
	tasa_interes,err2 := strconv.ParseFloat(input2,64)
	if err2 != nil || tasa_interes < 0 {
		fmt.Printf("%s es una tasa de interés inválida\n",input2)
		return
	}
	fmt.Printf("Tiempo (en Años): ")
	buffer3 := bufio.NewScanner(os.Stdin)
	buffer3.Scan()
	input3 := buffer3.Text()
	años,err3 := strconv.ParseFloat(input3,64)
	if err3 != nil || años < 0 {
		fmt.Printf("%s Tiempo Inválido\n",input3)
		return
	}
	var retorno float64 = CalcularInteres(capital,tasa_interes,años)
	fmt.Printf("Retorno: %f\n",retorno)
}

/*
Ejercicio 2.6. Escribir una función que imprima todos los números pares entre dos números
que se le pidan al usuario
*/
func ImprimirParesUsuario() {
	fmt.Printf("Número 1: ")
	buffer1 := bufio.NewScanner(os.Stdin)
	buffer1.Scan()
	input1 := buffer1.Text()
	numero1,err1 := strconv.Atoi(input1)
	if err1 != nil {
		fmt.Printf("%s no es un número válido\n",input1)
		return
	}
	fmt.Printf("Número 2: ")
	buffer2 := bufio.NewScanner(os.Stdin)
	buffer2.Scan()
	input2 := buffer2.Text()
	numero2,err2 := strconv.Atoi(input2)
	if err2 != nil {
		fmt.Printf("%s no es unn npumero vaĺido\n",input2)
		return
	}
	if numero1 > numero2 {
		numero1,numero2 = numero2,numero1
	}
	for i:=numero1; i <= numero2; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}
}
