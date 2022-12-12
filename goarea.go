package goarea

import "math"

// O Pi é uma proporção númerica definida pela relação entre
// o perimetro de uma circuferência e seu diâmetro
const Pi float64 = 3.1416

// Circ é responsável por cálcular a circuferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}


// React é responsável de calcular a área de um retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//não é visível
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
