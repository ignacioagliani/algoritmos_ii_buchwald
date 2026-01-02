package numeroComplejo

type NumeroComplejo interface {
	Multiplicar(otro NumeroComplejo) NumeroComplejo 
	Sumar(otro NumeroComplejo) NumeroComplejo
	ParteReal() float64
	ParteImaginaria() float64
	Modulo() float64
	Angulo() float64
	Ver()
}
