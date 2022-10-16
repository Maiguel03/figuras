package figuras

type Cuadrado struct {
	Lado float64
}

func (cua *Cuadrado) calcularArea() float64 {
	return cua.Lado * cua.Lado
}

func (cua *Cuadrado) calcularPerimetro() float64 {
	return 4 * cua.Lado
}
