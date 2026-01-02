package fraccion

type Fraccion interface {
	Sumar(otra Fraccion) Fraccion
	Multiplicar(otra Fraccion) Fraccion
	ParteEntera() int
	Representacion() string
}
