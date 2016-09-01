package main

import (
	"fmt"
)

func main() {
	// Объявление массива (1 вариант)
	var x [5]int
	fmt.Println(x)

	// Объявление массива (2 вариант)
	y := [5]float64{1.2, 2.3, 0.4, 32.3, 32.5}
	fmt.Println(y)

	// Нахождение среднего значения
	marks := [5]float64{2.0, 4.0, 5.0, 3.0, 4.0}
	sum := 0.0
	for i := 0; i < len(marks); i++ {
		sum += marks[i]
	}

	average := sum / float64(len(marks))
	fmt.Println("Срдний балл:", average)

	// Создание среза с использованием
	x := make([]float64, 5, 10)
	fmt.Println(x)

	// Создание среза с использованием slice
	arr := [5]int{1, 2, 3, 4, 5}
	y := arr[2:4]
	fmt.Println(y)

	// append()
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5, 6, 7)
	fmt.Println(slice2)

	//copy()
	slice_1 := []int{1, 2, 3, 4, 5, 6}
	slice_2 := make([]int, 2)
	copy(slice_2, slice_1)
	fmt.Println(slice_1, slice_2)

	var x map[string]int
	fmt.Println(x)

	// Создание и нициализации карты
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)

	delete(n, "foo")
	fmt.Println(n)
}