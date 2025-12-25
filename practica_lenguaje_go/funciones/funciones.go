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
/*
Ejercicio 5.3. Manejo de contraseñas
a) Escribir un programa que contenga una contraseña inventada, que le pregunte al usuario
la contraseña, y no le permita continuar hasta que la haya ingresado correctamente.
*/
func PedirContraseña() {
	var contraseña string = "abc123"
	buffer := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Contraseña: ")
		buffer.Scan()
		contraseña_usuario := buffer.Text()
		if contraseña_usuario == contraseña {
			fmt.Println("Acceso Concedido!")
			return
		}
	}
}

/*
b) Modificar el programa anterior para que solamente permita una cantidad fija de intentos.
*/
func PedirContraseñaLimitado(intentos int) {
	contador_intentos := 0
	contraseña := "aaa"
	buffer := bufio.NewScanner(os.Stdin)
	for {
		if contador_intentos < intentos {
			fmt.Print("Contraseña: ")
			buffer.Scan()
			contraseña_usuario := buffer.Text()
			if contraseña_usuario == contraseña {
				fmt.Println("Acceso concedido!")
				return
			}
			contador_intentos += 1
		} else {
			fmt.Println("No se pudo acceder. Se acabaron los intentos!")
			return
		}
	}
}

/*
Ejercicio 5.10. Escribir una función que reciba un número natural e imprima todos los números
primos que hay hasta ese número.
*/
func ImprimirRangoPrimos(n int) {
	for i := range n+1 {
		if EsPrimo(i) {
			fmt.Println(i)
		}
	}
}

/*
Ejercicio 6.1. Escribir funciones que dada una cadena de caracteres:
a) Imprima los dos primeros caracteres.
*/
func ImprimirPrimerosDosCaracteres(cadena string) {
	if len(cadena) < 2 {
		fmt.Println(cadena)
	} else {
		fmt.Println(cadena[0:2])
	}
}

/*
b) Imprima los tres últimos caracteres.
*/
func ImprimirUltimosTresCaracteres(cadena string) {
	if len(cadena) < 3 {
		fmt.Println(cadena)
	} else {
		fmt.Println(cadena[len(cadena) - 3:])
	}
}

/*
c) Imprima dicha cadena cada dos caracteres. Ej.: 'recta' debería imprimir 'rca'
*/
func ImprimirCadaDos(cadena string) {
	for i := 0; i < len(cadena); i += 2 {
		fmt.Printf("%c",cadena[i])
	}
	fmt.Print("\n")
}

/*
d) Dicha cadena en sentido inverso. Ej.: 'hola mundo!' debe imprimir '!odnum aloh'
*/
func ImprimirInverso(cadena string) {
	indice := len(cadena) - 1
	for indice >= 0 {
		fmt.Printf("%c",cadena[indice])
		indice -= 1
	}
	fmt.Print("\n")
}

/*
e) Imprima la cadena en un sentido y en sentido inverso. Ej: 'reflejo' imprime
'reflejoojelfer'.
*/
func ImprimirReflejo(cadena string) {
	fmt.Print(cadena)
	ImprimirInverso(cadena)
}

/*
Ejercicio 6.5. Escribir una función que dada una cadena de caracteres, devuelva:
a) La primera letra de cada palabra. Por ejemplo, si recibe 'Universal Serial Bus' debe
devolver 'USB'.
*/
func DevolverIniciales(cadena string) string {
	var iniciales string
	iniciales += string(cadena[0])
	for i := 0; i < len(cadena); i++ {
		if string(cadena[i]) == " " {
			iniciales += string(cadena[i+1])
		}
	}
	return iniciales
}

/*
b) Dicha cadena con la primera letra de cada palabra en mayúsculas. Por ejemplo, si recibe
'república argentina' debe devolver 'República Argentina'.
*/
func Capitalizar(cadena string) string {
	cadena = strings.TrimSpace(cadena)
	var capitalizada string
	capitalizada += strings.ToUpper(string(cadena[0]))
	for i := 1; i < len(cadena); i++ {
		if string(cadena[i]) == " " {
			capitalizada += " "
			capitalizada += strings.ToUpper(string(cadena[i+1]))
			i += 2
		}
		capitalizada += string(cadena[i])
	}
	return capitalizada
}

/*
Ejercicio 6.7. Escribir funciones que dadas dos cadenas de caracteres:
a) Indique si la segunda cadena es una subcadena de la primera. Por ejemplo, 'cadena'
es una subcadena de 'subcadena'.
*/
func EsSubcadena(cadena,subcadena string) bool {
	if strings.Contains(cadena,subcadena) {
		return true
	}
	return false
}

/*
b) Devuelva la que sea anterior en orden alfábetico. Por ejemplo, si recibe 'kde' y 'gnome'
debe devolver 'gnome'.
*/
func DevolverMenor(cadena1,cadena2 string) string {
	if cadena1 < cadena2 {
		return cadena1
	}
	return cadena2
}

/* Todo lo que sea tuplas/listas lo cambio por slices! */
/*
Ejercicio 7.1. Escribir una función que reciba una tupla de elementos e indique si se encuentran
ordenados de menor a mayor o no.
*/
func EstaOrdenadoMenorMayor(arreglo []int) bool {
	if len(arreglo) < 2 {
		return true
	}
	var i int = arreglo[0]
	for _,elemento := range arreglo {
		if elemento < i {
			return false
		}
		i = elemento
	}
	return true
}

/*
Ejercicio 7.2. Dominó.
a) Escribir una función que indique si dos fichas de dominó encajan o no. Las fichas son
recibidas en dos tuplas, por ejemplo: (3,4) y (5,4)
*/
func EncajaDomino(ficha1,ficha2 []int) bool {
	if len(ficha1) != 2 || len(ficha2) != 2 {
		return false
	}
	if ficha1[0] == ficha2[0] || ficha1[0] == ficha2[1] || ficha1[1] == ficha2[0] || ficha1[1] == ficha2[1] {
		return true
	}
	return false
}

/*
Ejercicio 7.3. Campaña electoral
a) Escribir una función que reciba una tupla con nombres, y para cada nombre imprima
el mensaje Estimado <nombre>, vote por mí.
*/
func ImprimirMensajeElectoral(nombres []string) {
	for _,nombre := range nombres {
		fmt.Printf("Estimado %s, vote por mí\n",nombre)
	}
}

/*
b) Escribir una función que reciba una tupla con nombres, una posición de origen p y una
cantidad n, e imprima el mensaje anterior para los n nombres que se encuentran a partir
de la posición p.
*/
func ImprimirAlgunosMensajes(nombres []string, p int, n int) {

	if p < 0 || p >= len(nombres) || p + n > len(nombres){
		return
	}
	seleccionados := nombres[p:p + n]
	for _,nombre := range seleccionados {
		fmt.Printf("Estimado %s, vote por mí\n",nombre)
	}
}

/*
c) Modificar las funciones anteriores para que tengan en cuenta el género del destinatario,
para ello, deberán recibir una tupla de tuplas, conteniendo el nombre y el género.
*/
func ImprimirMensajesDiscriminadoGenero(personas [][]string) {
	for _,info := range personas {
		genero := strings.ToUpper(strings.TrimSpace(info[1]))
		if genero == "H" {
			fmt.Printf("Estimado %s, vote por mí\n",info[0])
		} else if genero == "M" {
			fmt.Printf("Estimada %s, vote por mí\n",info[0])
		} else {
			fmt.Println("Sexo Inexistente")
		}
	}
}

/*
Ejercicio 7.4. Vectores
a) Escribir una función que reciba dos vectores y devuelva su producto escalar.
*/
func CalcularProductoEscalar(vector1, vector2 []int) (int,error) {
	len_vector1 := len(vector1)
	len_vector2 := len(vector2)
	if len_vector1 != len_vector2 {
		return 0,fmt.Errorf("Error: los vectores no poseen el mismo tamaño.")
	}
	producto_escalar := 0
	for i := 0; i < len_vector1; i++ {
		producto_escalar += vector1[i] * vector2[i]
	}
	return producto_escalar,nil
}

/*
b) Escribir una función que reciba dos vectores y devuelva si son o no ortogonales.
*/
func SonOrtogonales(vector1,vector2 []int) bool {
	prod_escalar,err := CalcularProductoEscalar(vector1,vector2)
	if prod_escalar != 0 || err != nil {
		return false
	}
	return true
}

/*
c) Escribir una función que reciba dos vectores y devuelva si son paralelos o no.
*/
func SonParalelos(vector1,vector2 []int) bool {
	len_vector1 := len(vector1)
	len_vector2 := len(vector2)
	if len_vector1 != len_vector2 {
		return false
	}
	for i := 0; i < len_vector1; i++ {
		if vector1[0] > vector2[0] {
			if vector1[i] % vector2[i] != 0 {
				return false
			}
		} else {
			if vector2[i] % vector1[i] != 0 {
				return false
			}
		}
	}
	return true
}

/*
d) Escribir una función que reciba un vector y devuelva su norma.
*/
func Norma(vector []int) (float64,error) {
	len_vector := len(vector)
	if len_vector == 0 {
		return 0,fmt.Errorf("El vector no tiene componentes")
	}
	suma_componentes_cuadrado := 0
	for i := 0; i < len_vector; i++ {
		suma_componentes_cuadrado += vector[i] * vector[i]
	}
	norma := math.Sqrt(float64(suma_componentes_cuadrado))
	return norma,nil
}

/*
Ejercicio 7.5. Dada una lista de números enteros, escribir una función que:
a) Devuelva una lista con todos los que sean primos.
*/
func DevolverArregloPrimos(arreglo_numeros []int) []int {
	arreglo_primos := []int{}
	for _,numero := range arreglo_numeros {
		if EsPrimo(numero) {
			arreglo_primos = append(arreglo_primos,numero)
		}
	}
	return arreglo_primos
}

/*
b) Devuelva la sumatoria y el promedio de los valores.
*/
func DevolverSumaYPromedio(arreglo_numeros []int) (int,float64,error) {
	if len(arreglo_numeros) == 0 {
		return 0,0,fmt.Errorf("Error: el arreglo está vacío!")
	}
	var suma float64 = 0
	var contador float64 = 0
	for _,numero := range arreglo_numeros {
		suma += float64(numero)
		contador += 1
	}
	return int(suma),suma/contador,nil
}

/*
c) Devuelva una lista con el factorial de cada uno de esos números.
*/
func DevolverArregloFactorial(arreglo []int) []int {
	arreglo_factorial := []int{}
	for _,numero := range arreglo{
		arreglo_factorial = append(arreglo_factorial,CalcularFactorialRecursivo(numero))
	}
	return arreglo_factorial
}

/*
Ejercicio 7.6. Dada una lista de números enteros y un entero k, escribir una función que:
a) Devuelva tres listas, una con los menores, otra con los mayores y otra con los iguales a k.
*/
func DevolverListasMenoresMayoresIguales(arreglo []int, k int) ([]int,[]int,[]int) {
	menores := []int{}
	mayores := []int{}
	iguales := []int{}
	for _,numero := range arreglo {
		if numero < k {
			menores = append(menores,numero)
		} else if numero > k {
			mayores = append(mayores,numero)
		} else {
			iguales = append(iguales,numero)
		}
	}
	return menores,mayores,iguales
}

/*
b) Devuelva una lista con aquellos que son múltiplos de k.
*/
func DevolverListaMultiplos(arreglo []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	multiplos := []int{}
	for _,numero := range arreglo {
		if numero % k == 0 {
			multiplos = append(multiplos,numero)
		}
	}
	return multiplos
}

/*
Ejercicio 7.7. Escribir una función que reciba una lista de tuplas (Apellido, Nombre, Inicial_
segundo_nombre) y devuelva una lista de cadenas donde cada una contenga primero el
nombre, luego la inicial con un punto, y luego el apellido.
*/
func DevolverArregloNombresConFormato(arreglo [][]string) []string {
	arreglo_nombres_con_formato := []string{}
	for _,tupla := range arreglo {
		if len(tupla) < 3 {
			continue
		}
		nombre_formateado := tupla[1] + " " + tupla[2] + ". " + tupla[0]
		arreglo_nombres_con_formato = append(arreglo_nombres_con_formato,nombre_formateado)
	}
	return arreglo_nombres_con_formato
}

/*
Ejercicio 7.8. Inversión de listas
a) Realizar una función que, dada una lista, devuelva una nueva lista cuyo contenido sea
igual a la original pero invertida. Así, dada la lista ['Di', 'buen', 'día', 'a', 'papa'],
deberá devolver ['papa', 'a', 'día', 'buen', 'Di'].
*/
func DevolverArregloInvertido(arreglo []string) []string {
	arreglo_invertido := []string{}
	i := len(arreglo) - 1
	for i >= 0{
		arreglo_invertido = append(arreglo_invertido,arreglo[i])
		i--
	}
	return arreglo_invertido
}
/*
Ejercicio 7.10. Matrices.
a) Escribir una función que reciba dos matrices y devuelva la suma.
*/
func SumarMatrices(matriz1,matriz2 [][]int) ([][]int,error) {
	if len(matriz1) != len(matriz2) || len(matriz1[0]) != len(matriz2[0]){
		return [][]int{},fmt.Errorf("Error: Tamaños Distintos!")
	}
	matriz_suma := [][]int{}
	for i := 0; i < len(matriz1); i++ {
		fila := []int{}
		for j := 0; j < len(matriz1[0]); j++ {
			fila = append(fila,matriz1[i][j] + matriz2[i][j])
		}
		matriz_suma = append(matriz_suma,fila)
	}
	return matriz_suma,nil
}
