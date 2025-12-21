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
	"strings"
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

/*
Ejercicio 3.4. Área de polígonos
a) Escribir una función que reciba las coordenadas de un vector en ℝ^3 (x,y,z)
y devuelva la norma del vector, dada por ‖(𝑥, 𝑦, 𝑧)‖ = √(𝑥^2 + 𝑦^2 + 𝑧^2).
Ejemplo: norma(3, 2, -4) → 5.3851
*/
func CalcularNorma(vector [3]float64) float64 {
	norma := math.Pow(math.Pow(vector[0],2)+math.Pow(vector[1],2)+math.Pow(vector[2],2),0.5)
	return norma
}

/*
b) Escribir una función que reciba las coordenadas de dos vectores en ℝ^3
(x1,y1,z1,x2,y2,z2) y devuelva las coordenadas del vector diferencia (debe
devolver 3 valores numéricos).
Ejemplo: diferencia(8, 7, -3, 5, 3, 2) → (3, 4, -5)
*/
func DevolverVectorDiferencia(x1,y1,z1,x2,y2,z2 int) [3]int {
	componente_x := x1 - x2
	componente_y := y1 - y2
	componente_z := z1 - z2
	vector_diferencia := [3]int{componente_x,componente_y,componente_z}
	return vector_diferencia
}

/*
Ejercicio 4.1. Escribir dos funciones que resuelvan los siguientes problemas:
a) Dado un número entero 𝑛, indicar si es par o no.
*/
func EsPar(n int) bool {
	if n % 2 == 0 {
		return true
	}
	return false
}
/*
b) Dado un número entero 𝑛, indicar si es primo o no.
*/
func EsPrimo(n int) bool {
	switch {
	case n <= 1:
		return false
	case n == 2:
		return true
	default:
		for i := 2;i <= int(math.Sqrt(float64(n))); i++ { // Es suficiente probar hasta la raiz cuadrada
			if n % i == 0 {
				return false
			}
		}
		return true
	}
}

/*
Ejercicio 4.2. Escribir una implementación propia de la función abs, que devuelva el valor
absoluto de cualquier valor que reciba.
*/
func DevolverValorAbsoluto(n float64) float64 {
	if n < 0 {
		return (-1) * n
	}
	return n
}

/*
Ejercicio 4.3. Escribir una función que reciba por parámetro una dimensión 𝑛, e imprima
la matriz identidad correspondiente a esa dimensión.
*/
func DevolverMatrizIdentidad(n int) [][]int {
	matriz_identidad := make([][]int, n)
	for i := 0; i < n; i++ {
		fila := make([]int,n)
		fila[i] = 1
		matriz_identidad[i] = fila
	}
	return matriz_identidad
}

/*
Ejercicio 5.1. Escribir un programa que permita al usuario ingresar un conjunto de notas,
preguntando a cada paso si desea ingresar más notas, e imprimiendo el promedio correspondiente.
*/
func CalcularPromedioNotas() float64 {
	var suma_notas float64 = 0
	var contador int = 0
	buffer := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("¿Querés seguir ingresando notas? S/N")
		buffer.Scan()
		opcion := buffer.Text()
		switch strings.ToLower(opcion) {
		case "s":
			fmt.Printf("Nota: ")
			buffer.Scan()
			nota_str := buffer.Text()
			nota,err := strconv.ParseFloat(nota_str,64)
			if err != nil {
				fmt.Printf("%s es una nota inválida!\n",nota_str)
				continue
			}
			suma_notas += nota
			contador += 1
		case "n":
			if contador == 0 {
				fmt.Println("No se cargaron notas")
				return -1
			}
			return suma_notas / float64(contador)
		default:
			fmt.Println("Opción Inválida")
			return -1
		}
	}
}
