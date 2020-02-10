package goArea

import (
	"math"
)

//Pi proporcao numerica definida pela relaçao entre
//o perimetro de uma circuferencia e ao seu diametro
const Pi = 3.1413

//Circulo Responsavel por calcular a rea da circuferencia
func Circulo(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Retangulo calcular a area de um retangulo
func Retangulo(base, altura float64) float64 {
	return base * altura
}

// nao é visivel
func _Triangulo(base, altura float64) float64 {
	return (base * altura) / 2
}
