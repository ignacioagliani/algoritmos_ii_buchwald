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
	"io"
	"math/rand/v2"
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

/*
Ejercicio 7.12. Funciones que reciben funciones.
a) Escribir una funcion llamada map, que reciba una función y una lista y devuelva la lista
que resulta de aplicar la función recibida a cada uno de los elementos de la lista recibida.
*/
/* Función de Ejemplo */
func ElevarCuadrado (n int) int {
	return n * n
}

func Map(funcion func(int) int,arreglo []int) []int {
	mapeado := []int{}
	for _,elemento := range arreglo {
		mapeado = append(mapeado,funcion(elemento))
	}
	return mapeado
}

/*
b) Escribir una función llamada filter, que reciba una función y una lista y devuelva una
lista con los elementos de la lista recibida para los cuales la función recibida devuelve un
valor verdadero.
*/
func Filter(funcion func(int) bool, arreglo []int) []int {
	arreglo_filtrado := []int{}
	for _,elemento := range arreglo {
		if funcion(elemento) {
			arreglo_filtrado = append(arreglo_filtrado,elemento)
		}
	}
	return arreglo_filtrado
}

/*
Ejercicio 11.1. Escribir una función, llamada head que reciba un archivo y un número N e
imprima las primeras N líneas del archivo.
*/
func Head(ruta_archivo string, N int) {
	archivo,err := os.Open(ruta_archivo)
	if err != nil {
		fmt.Println("No se pudo abrir el archivo")
		return
	}
	defer archivo.Close()

	buffer := bufio.NewScanner(archivo)
	lineas_leidas := 0
	for N > 0 && lineas_leidas < N && buffer.Scan() {
		fmt.Printf("%s\n",buffer.Text())
		lineas_leidas += 1
	}
	err = buffer.Err()
	if err != nil {
		fmt.Println(err)
	}
}

/*
Ejercicio 11.2. Escribir una función, llamada cp, que copie todo el contenido de un archivo (sea
de texto o binario) a otro, de modo que quede exactamente igual.
*/
func Cp(ruta_archivo_origen,ruta_archivo_destino string) {
	destino,err1 := os.Create(ruta_archivo_destino)
	if err1 != nil {
		fmt.Println("Error al abrir el archivo de destino!")
		return
	}
	defer destino.Close()

	origen,err2 := os.Open(ruta_archivo_origen)
	if err2 != nil {
		fmt.Println("Error al abrir el archivo de origen")
		return
	}
	defer origen.Close()

	_,err3 := io.Copy(destino,origen)
	if err3 != nil {
		fmt.Println("No se pudieron copiar los bytes")
		return
	}
	/*
	Si se quiere solo copiar texto normal, se hace así (quitando el bloque io.Copy()):
	linea := bufio.NewScanner(origen)
	datawriter := bufio.NewWriter(destino)
	for linea.Scan() {
		_,err4 := datawriter.WriteString(linea.Text() + "\n")
		if err4 != nil {
			fmt.Println("Error al Escribir")
			return
		}
	}
	datawriter.Flush()
	*/
}

/*
Ejercicio 11.3. Escribir una función, llamada wc, que dado un archivo de texto, lo procese e
imprima por pantalla cuántas líneas, cuantas palabras y cuántos caracteres contiene el archivo.
*/
func Wc(ruta_archivo string) {
	lineas := 0
	palabras := 0
	caracteres := 0
	archivo,err := os.Open(ruta_archivo)
	if err != nil {
		fmt.Println("No se pudo abrir el archivo")
		return
	}
	defer archivo.Close()

	linea := bufio.NewScanner(archivo)
	for linea.Scan() {
		lineas += 1
		palabras_separadas := strings.Split(linea.Text()," ")
		for _,palabra := range palabras_separadas {
			palabras += 1
			for range palabra {
				caracteres += 1
			}
		}
	}
	fmt.Printf("Cantidad Lineas: %d\nCantidad Palabras: %d\nCantidad Letras: %d\n",lineas,palabras,caracteres)
}

/*
Ejercicio 11.4. Escribir una función, llamada grep, que reciba una cadena y un archivo de texto,
e imprima las líneas del archivo que contienen la cadena recibida.
*/
func Grep(cadena,ruta_archivo string) {
	archivo,err := os.Open(ruta_archivo)
	if err != nil {
		fmt.Println("No se pudo abrir el archivo")
		return
	}
	defer archivo.Close()

	buffer := bufio.NewScanner(archivo)
	for buffer.Scan() {
		if strings.Contains(buffer.Text(),cadena) {
			fmt.Printf("%s\n",buffer.Text())
		}
	}
}

/*
Ejercicio 15.1. Escribir una función recursiva que reciba un número positivo 𝑛 y devuelva la
cantidad de dígitos que tiene.
*/
func DevolverCantidadDigitos(n int) int {
	if n / 10 == 0 {
		return 1
	}
	return 1 + DevolverCantidadDigitos(n / 10)
}
/*
Ejercicio 15.2. Escribir una función recursiva que simule el siguiente experimento: Se tiene una
rata en una jaula con 3 caminos, entre los cuales elige al azar (cada uno tiene la misma probabilidad),
si elige el 1 luego de 3 minutos vuelve a la jaula, si elige el 2 luego de 5 minutos
vuelve a la jaula, en el caso de elegir el 3 luego de 7 minutos sale de la jaula. La rata no aprende,
siempre elige entre los 3 caminos con la misma probabilidad, pero quiere su libertad, por lo que
recorrerá los caminos hasta salir de la jaula.
La función debe devolver el tiempo que tarda la rata en salir de la jaula.
*/
func Experimento() int {
	eleccion_rata := rand.IntN(3) + 1 // Rango entre [0,3) a diferencia del randint de Python. Asi que le sumo 1.
	switch eleccion_rata {
	case 1:
		return 3 + Experimento()
	case 2:
		return 5 + Experimento()
	default:
		return 7
	}
}

/*
Ejercicio 15.3. Escribir una función recursiva que reciba 2 enteros n y b y devuelva True si n es
potencia de b.
Ejemplos:
es_potencia(8, 2) -> True
es_potencia(64, 4) -> True
es_potencia(70, 10) -> False
*/
func EsPotencia(n,b int) bool {
	if n < 0 || b < 0 {
		fmt.Println("No se admiten número negativos")
		return false
	}
	if n == 1 {
		return true
	}
	if b == 0 {
		return n == 0
	}
	if b == 1 {
		return n == 1
	}
	return _EsPotencia(n,b,b)
}

func _EsPotencia(n,b,j int) bool {
	if b == n {
		return true
	} else if b > n {
		return false
	} else {
		return _EsPotencia(n,b*j,j)
	}
}

/*
Ejercicio 15.5. Escribir dos funciones mutualmente recursivas par(n) e impar(n)
que determinen la paridad del numero natural dado, conociendo solo que:
• 1 es impar.
• Si un número es impar, su antecesor es par; y viceversa.
*/
func Par(n uint) bool {
	if n == 1 {
		return false
	}
	if n == 2 || n == 0 {
		return true
	}
	return Impar(n-2)
}

func Impar(n uint) bool {
	if n == 1 {
		return true
	}
	if n == 2 || n == 0 {
		return false
	}
	return Par(n-2)
}

/*
Ejercicio 15.6. Escribir una función recursiva que calcule recursivamente el n-ésimo número
triangular (el número 1 + 2 + 3 + ⋯ + 𝑛).
*/
func CalcularNumeroTriangular(n uint) uint {
	if n == 0 {
		return 0
	}
	return n + CalcularNumeroTriangular(n - 1)
}

/*
Ejercicio 15.8. Escribir una funcion recursiva que encuentre el mayor elemento de una lista.
*/
func MayorElemento(arreglo []int) []int {
	if len(arreglo) == 1 {
		return arreglo
	}
	if arreglo[0] > arreglo[1] {
		arreglo = append(arreglo[:1], arreglo[2:]...)
		return MayorElemento(arreglo)
	}
	return MayorElemento(arreglo[1:])
}

/*
Ejercicio 15.9. Escribir una función recursiva para replicar los elementos de una lista una
cantidad n de veces. Por ejemplo:
replicar([1, 3, 3, 7], 2) -> ([1, 1, 3, 3, 3, 3, 7, 7])
*/
func Replicar(arreglo []int, n uint) []int {
	return replicar(arreglo,n,[]int{})
}

func replicar(arreglo []int, n uint, arreglo_replicado []int) []int {
	if len(arreglo) == 0 || n == 0 {
		return arreglo_replicado
	}
	for range n {
		arreglo_replicado = append(arreglo_replicado,arreglo[0])
	}
	return replicar(arreglo[1:],n,arreglo_replicado)
}

/*
Ejercicio 15.10. Escribir una funcion recursiva que dada una cadena determine si en la misma
hay más letras A o letras E.
*/
func masAoE (cadena string) int {
	if len(cadena) == 0 {
		return 0
	}
	cantidad_A_contra_E := 0
	if string(cadena[0]) == "A" {
		cantidad_A_contra_E = 1
	}
	if string(cadena[0]) == "E" {
		cantidad_A_contra_E = -1
	}
	return cantidad_A_contra_E + masAoE(cadena[1:])
}

func MasAoE(cadena string) {
	cantidad_A := masAoE(cadena)
	if cantidad_A > 0 {
		fmt.Println("Más A")
	} else if cantidad_A < 0 {
		fmt.Println("Más E")
	} else {
		fmt.Println("Misma cantidad")
	}
}

/* Ejercicios de Structs (No son de la guía de Essaya) */

/*
1.	Definí un struct Persona { Nombre string; Edad int } y escribí:
•	EsMayor() que devuelva true si Edad >= 18
•	ImprimirInformacion() para imprimir “Nombre (Edad)”.
*/
// Esta es Pública
type Persona struct {
	Nombre string
	Edad int
}

// Ya hay una función EsMayor antes, así que le cambio el nombre
func (persona Persona) EsMayorStruct() bool {
	return persona.Edad >= 18
}

func (persona Persona) ImprimirInformacion() {
	fmt.Printf("%s (%d)\n",persona.Nombre,persona.Edad)
}

/*
2.	struct Rectangulo { Ancho, Alto float64 }:
•	método Area()
•	método Perimetro()
•	función que reciba un slice de rectángulos y devuelva el de mayor área.
*/

// Esta es Pública
type Rectangulo struct {
	Ancho float64
	Alto float64
}

func (rec Rectangulo) Area() float64 {
	return rec.Alto * rec.Ancho
}

func (rec Rectangulo) Perimetro() float64 {
	return 2 * rec.Alto + 2 * rec.Ancho
}

/*
3.	struct CuentaBancaria { Titular string; Saldo float64 }:
•	métodos Depositar(monto float64) error (no permitir saldo negativo).
•	Extraer(monto float64). error (no permitir saldo negativo)
•	método Transferir(a *CuentaBancaria, monto float64) error.
*/

// Esta es privada
type CuentaBancaria struct {
	titular string
	saldo float64
}

func CrearCuenta(nombre string, monto float64) (CuentaBancaria, error) {
	if monto < 0 {
		return CuentaBancaria{}, fmt.Errorf("monto inválido")
	}
	return CuentaBancaria{titular: nombre, saldo: monto}, nil
}

func (cuenta *CuentaBancaria) Depositar(monto float64) error {
	if monto < 0 {
		return fmt.Errorf("ERROR: no se admiten números negativos")
	}
	cuenta.saldo += monto
	return nil
}

func (cuenta *CuentaBancaria) Extraer(monto float64) error {
	if monto < 0 || monto > cuenta.saldo {
		return fmt.Errorf("ERROR: monto inválido")
	}
	cuenta.saldo -= monto
	return nil
}

func (cuenta *CuentaBancaria) Transferir(a *CuentaBancaria, monto float64) error {
	if monto < 0 {
		return fmt.Errorf("ERROR: monto inválido")
	}
	if a == nil {
		return fmt.Errorf("La cuenta de destino no existe.")
	}
	extraccion := cuenta.Extraer(monto)
	if extraccion != nil {
		return extraccion
	}
	deposito := a.Depositar(monto)
	if deposito != nil {
		cuenta.saldo += monto
		return deposito
	}
	return nil
}

func (cuenta CuentaBancaria) ConsultarSaldo() {
	fmt.Printf("Titular: %s\nSaldo: $%f\n",cuenta.titular,cuenta.saldo)
}

/*
4.	struct Producto { ID int; Nombre string; Precio float64 }:
•	función AplicarDescuento(p Producto, porcentaje float64) Producto
•	función que devuelva el producto más caro de un slice.
*/
type Producto struct {
	id int
	nombre string
	precio float64
}

func AplicarDescuento(p Producto, porcentaje float64) (Producto,error) {
	if porcentaje > 100 || porcentaje < 0 {
		return p,fmt.Errorf("Porcentaje Inválido")
	}
	descuento := p.precio - (p.precio*(porcentaje/100))
	p.precio = descuento
	return p,nil
}

func MasCaro(productos []Producto) (Producto,error) {
	if len(productos) == 0 {
		return Producto{},fmt.Errorf("Lista vacía")
	}
	var indice int = 0
	var mayor_precio float64 = productos[0].precio
	for i,prod := range productos {
		if prod.precio > mayor_precio {
			indice = i
			mayor_precio = prod.precio
		}
	}
	return productos[indice],nil
}

/*
5.	struct Punto { X, Y float64 } y struct Segmento { A, B Punto }:
•	método Longitud() para el segmento
•	método Mover(dx, dy float64) que traslade ambos puntos.
*/
type Punto struct {
	X,Y float64
}
type Segmento struct {
	A,B Punto
}

func (s Segmento) Longitud() float64 {
	A := s.A
	B := s.B
	len_x := A.X - B.X
	len_y := A.Y - B.Y
	return math.Sqrt(len_x*len_x + len_y*len_y)
}

func (s *Segmento) Mover(dx,dy float64) {
	s.A.X += dx
	s.B.X += dx
	s.A.Y += dy
	s.B.Y += dy
}

/*
6.	struct Alumno { Legajo int; Nombre string; Notas []float64 }:
•	método Promedio() (float64, error) (error si no tiene notas)
•	función que filtre alumnos con promedio >= 7.
*/
type Alumno struct {
	legajo int
	nombre string
	notas []float64
}
func (alumno Alumno) Promedio() (float64,error) {
	if len(alumno.notas) == 0 {
		return -1,fmt.Errorf("El alumno no cuenta con notas cargadas")
	}
	var suma_notas float64 = 0
	for _,nota := range alumno.notas {
		suma_notas += nota
	}
	return suma_notas/float64(len(alumno.notas)),nil
}

func FiltrarPromedio(arreglo_alumnos []Alumno) []Alumno {
	alumnos_filtrados := []Alumno{}
	for _,alumno := range arreglo_alumnos {
		promedio,err := alumno.Promedio()
		if err == nil {
			if promedio >= 7 {
				alumnos_filtrados = append(alumnos_filtrados,alumno)
			}
		}
	}
	return alumnos_filtrados
}

func CrearAlumnos() []Alumno { //Necesaria para poder usarse porque los campos del struct son privados
	yo := Alumno{111111,"Ignacio",[]float64{10,10,10,9,9,8}}
	otro1 := Alumno{111111,"Otro1",[]float64{}}
	otro2 := Alumno{111111,"Otro12",[]float64{2,5,6}}
	otro3 := Alumno{111111,"Otro3",[]float64{7,7,7}}
	return []Alumno{yo,otro1,otro2,otro3}
}

/*
7.	struct Pedido { Numero int; Items []Item } con struct Item { Nombre string; Cantidad int; PrecioUnit float64 }:
•	método Total() float64
•	función que agregue un item (si existe, sumar cantidad).
*/
type Pedido struct {
	numero int
	items []Items
}

type Items struct {
	nombre string
	cantidad int
	precioUnitario float64
}

func (pedido Pedido) Total() float64 {
	if len(pedido.items) == 0 {
		return 0
	}
	var total float64 = 0
	for _,item := range pedido.items {
		total += item.precioUnitario * float64(item.cantidad)
	}
	return total
}

func AgregarItem(item Items, pedido *Pedido) {
	for i,_ := range pedido.items {
		if pedido.items[i].nombre == item.nombre {
			pedido.items[i].cantidad += item.cantidad
			return
		}
	}
	pedido.items = append(pedido.items,item)
}

func PedidoFicticio () Pedido {
	return Pedido{1,[]Items{Items{"Cheeseburger",2,10000},Items{"Papa Fritas",2,7000}}}
}

func ItemaAgregar() Items {
	return Items{"Coca-Cola",2,4000}
}

/*
8.	struct Libro { ISBN string; Titulo string; Autor string; Stock int }:
•	método Prestar() error (si no hay stock, error)
•	método Devolver() (incrementa stock)
•	función para buscar libros por autor.
*/
type Libro struct {
	isbn string
	titulo string
	autor string
	stock int
}

func (libro *Libro) Prestar() error {
	if libro.stock == 0 {
		return fmt.Errorf("No hay stock. No se puede prestar.")
	}
	libro.stock -= 1
	return nil
}

func (libro *Libro) Devolver(){
	libro.stock += 1
}

func BuscarPorAutor(arreglo_libros []Libro, autor string) []Libro {
	libros_autor := []Libro{}
	for i,_ := range arreglo_libros {
		if arreglo_libros[i].autor == autor {
			libros_autor = append(libros_autor,arreglo_libros[i])
		}
	}
	return libros_autor
}

func LibroEjemplo() Libro {
	return Libro{"1515145","Análisis Vectorial","Claudio Pita Ruiz",0}
}
func ArregloLibroEjemplo() []Libro {
	return []Libro{
		Libro{"1515145","Análisis Vectorial","Claudio Pita Ruiz",10},
		Libro{"1515146","Espacios métricos homogéneos de Lie-Banach","DI IORIO Y LUCERO, MARIA EUGENIA",20},
		Libro{"1515147","Álgebra Lineal","Claudio Pita Ruiz",15},
	}
}

/*
9.	struct Fecha { Dia, Mes, Anio int }:
•	método EsValida() bool
•	método AntesQue(otra Fecha) bool
*/
type Fecha struct {
	dia, mes, anio int
}

func (fecha Fecha) EsValida() bool {
	if fecha.dia < 1 || fecha.dia > 31 {
		return false
	}
	if fecha.mes < 1 || fecha.mes > 12 {
		return false
	}
	if fecha.mes == 2 {
		if (fecha.anio % 400 == 0) || (fecha.anio % 4 == 0 && fecha.anio % 100 != 0) {
			if fecha.dia > 29 {
				return false
			}
			return true
		}
		if fecha.dia > 28 {
			return false
		}
		return true
	}
	if fecha.mes == 4 || fecha.mes == 6 || fecha.mes == 9 || fecha.mes == 11 {
		if fecha.dia > 30 {
			return false
		}
	}
	return true
}

func (fecha Fecha) AntesQue(otra Fecha) bool {
	if !fecha.EsValida() || !otra.EsValida() {
		return false
	}
	if fecha.anio > otra.anio {
		return false
	}
	if fecha.anio == otra.anio {
		if fecha.mes == otra.mes {
			if fecha.dia >= otra.dia {
				return false
			}
		} else if fecha.mes > otra.mes {
			return false
		}
	}
	return true
}

/*
10.	struct Vector { Datos []int }:
•	método Norma() float64
•	método ProductoEscalar(otro Vector) (int, error) (error si longitudes distintas)
*/
type Vector struct {
	datos []int
}

func (v Vector) Norma() float64 {
	if len(v.datos) == 0 {
		return 0
	}
	suma_componentes_cuadrado := 0
	for _,componente := v.datos {
		suma_componentes_cuadrado += componente * componente
	}
	return math.Sqrt(float64(suma_componentes_cuadrado))
}

func (v Vector) ProductoEscalar(otro Vector) (int, error) {
	if len(v.datos) != len(otro.datos) {
		return 0,fmt.Errorf("Error: no coincide la longitud")
	}
	producto_escalar := 0
	for i,_ := range v.datos {
		producto_escalar += v.datos[i] * otro.datos[i]
	}
	return producto_escalar,nil
}
