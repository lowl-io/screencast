package main

import (
	m "screencast/lesson9/myMath"
	"fmt"
)

func main() {
	a := 5
	b := 2

	fmt.Println("Add:", m.Add(a, b))
	fmt.Println("Sub:", m.Sub(a, b))
	fmt.Println("Mul:", m.Mul(a, b))
	fmt.Println("Div:", m.Div(a, b))

	x := 4.0

	fmt.Println("Sqrt:", m.Sqrt(x))
}
