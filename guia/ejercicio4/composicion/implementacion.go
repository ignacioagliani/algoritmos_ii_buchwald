package composicion

type Funcion struct {
	funciones []func(float64) float64
}

func CrearComposicion() ComposicionFunciones {
	 return &Funcion{funciones: []func(float64) float64{}}
}

func (f *Funcion) AgregarFuncion(funcion func (float64) float64) {
	f.funciones = append(f.funciones,funcion)
}

func (f Funcion) Aplicar(x float64) float64 {
	i := len(f.funciones) - 1
	resultado := x
	for i >= 0 {
		resultado = f.funciones[i](resultado)
		i -= 1
	}
	return resultado
}
