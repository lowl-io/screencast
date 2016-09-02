package main

import "fmt"

func main() {

	// Оригинальная перменная не изменится
	x := 5
	zero(x)
	fmt.Println("Without pointer:", x) // x = 5

	// Оригинальная прерменная изменится
	zeroPointer(&x)
	fmt.Println("Pointer:", x) // x = 0

	// Пример функции new()
	xPtr := new(int)
	one(xPtr)
	fmt.Println("NEW:", *xPtr) // x = 1

}

func one(xPtr *int) {
	*xPtr = 1
}

// Аргумент копируется в функцию,
// не изменяя оригинальную перменную
func zero(x int) {
	x = 0
}

// Оригинальная перменная x и xPtr
// указывают на одно и тоже место в памяти,
// Теперь мы можем изменить оригинальную перменную
func zeroPointer(xPtr *int) {
	*xPtr = 0
}


