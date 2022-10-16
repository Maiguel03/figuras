package figuras

import "fmt"

type Operacion interface {
	calcularArea() float64
	calcularPerimetro() float64
}

func CalcularOperaciones(area Operacion) {
	fmt.Println("Medida: ", area)
	fmt.Println("Area: ", area.calcularArea())
	fmt.Println("Perimetro: ", area.calcularPerimetro())
}
