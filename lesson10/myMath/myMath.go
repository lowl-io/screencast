package myMath

import "math"

// Сложение a + b
func Add(a, b int) int {
	return a + b
}

// Вычитание a - b
func Sub(a, b int) int {
	return a - b
}

// Умножение a * b
func Mul(a, b int) int {
	return a * b
}

// Деление a / b
func Div(a, b int) int {
	return a / b
}

// Квадртаный корень sqrt(x)
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}