package fraccion

import (
	"fmt"
)

type fraccion struct {
	numerador int
	denominador int
}

func CrearFraccion(numerador,denominador int) fraccion {
	if denominador == 0 {
		panic("ERROR: denominador es cero!")
	}
	return fraccion{numerador,denominador}
}

func (f fraccion) Sumar(otro Fraccion) Fraccion {
	ot := otro.(fraccion)
	num := (f.numerador * ot.denominador) + (f.denominador * ot.numerador)
	den := f.denominador * ot.denominador
	return CrearFraccion(num, den)
}

func (f fraccion) Multiplicar(otro Fraccion) Fraccion {
	ot := otro.(fraccion)
	num := f.numerador * ot.numerador
	den := f.denominador * ot.denominador
	return CrearFraccion(num, den)
}

func (f fraccion) ParteEntera() int {
	return f.numerador / f.denominador
}

func (f fraccion) Representacion() string {
	return fmt.Sprintf("%d/%d", f.numerador, f.denominador)
}
