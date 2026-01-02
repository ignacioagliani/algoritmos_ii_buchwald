package numeroComplejo

import (
	"math"
)

type complejo struct {
	real float64
	imaginaria float64
}

func CrearComplejo(real,img float64) NumeroComplejo {
	return complejo{real:real,imaginaria:img}
}

// a+bi * c+di = (ac-bd)+(ad+bc)i
func (comp complejo) Multiplicar(otro NumeroComplejo) NumeroComplejo {
	complejo2 := otro.(complejo)
	parte_real := comp.real * complejo2.real - comp.imaginaria * complejo2.imaginaria
	parte_imaginaria := comp.real * complejo2.imaginaria + comp.imaginaria *complejo2.real
	return CrearComplejo(parte_real,parte_imaginaria)
}

func (comp complejo) Sumar(otro NumeroComplejo) NumeroComplejo {
	complejo2 := otro.(complejo)
	return CrearComplejo(comp.real + complejo2.real, comp.imaginaria + complejo2.imaginaria)
}

func (comp complejo) ParteReal() float64 {
	return comp.real
}

func (comp complejo) ParteImaginaria() float64 {
	return comp.imaginaria
}

func (comp complejo) Modulo() float64 {
	suma_componentes_cuadrado := comp.real * comp.real + comp.imaginaria * comp.imaginaria
	return math.Sqrt(suma_componentes_cuadrado)
}

func (comp complejo) Angulo() float64 {
	return math.Atan2(comp.imaginaria,comp.real)
}
