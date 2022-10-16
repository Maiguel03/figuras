package figuras

import "math"

type Circulo struct {
	Radio float64
}

func (cir *Circulo) calcularArea() float64 {
	return math.Pi * (cir.Radio * cir.Radio)
}

func (cir *Circulo) calcularPerimetro() float64 {
	return 2 * math.Pi * cir.Radio
}
