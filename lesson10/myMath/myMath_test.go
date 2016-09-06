package myMath

import (
	m "screencast/lesson9/myMath"
	"testing"
)


// Тестирование функции Add(a, b)
func TestAdd(t *testing.T) {
	a := 4
	b := 6
	result := 10

	resultAdd := m.Add(a, b)

	if resultAdd != result {
		t.Error("Expcted", result, "got", resultAdd)
	}
}

// Тестирование функции Sub(a, b)
func TestSub(t *testing.T) {
	a := 6
	b := 4
	result := 2

	resultAdd := m.Sub(a, b)

	if resultAdd != result {
		t.Error("Expcted", result, "got", resultAdd)
	}
}

// Тестирование функции Mul(a, b)
func TestMul(t *testing.T) {
	a := 4
	b := 6
	result := 24

	resultAdd := m.Mul(a, b)

	if resultAdd != result {
		t.Error("Expcted", result, "got", resultAdd)
	}
}

// Тестирование функции Div(a, b)
func TestDiv(t *testing.T) {
	a := 10
	b := 5
	result := 2

	resultAdd := m.Div(a, b)

	if resultAdd != result {
		t.Error("Expcted", result, "got", resultAdd)
	}
}

// Тестирование функции Sqrt(x)
func TestSqrt(t *testing.T) {
	x := 4.0
	result := 2.0

	resultSqrt := m.Sqrt(x)

	if resultSqrt != result {
		t.Error("Expected", result, "got", resultSqrt)
	}
}