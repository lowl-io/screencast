package main

import (
	"fmt"
)

func main() {
	//array()
	//slice()
	mapp()
}

func array() {

	// Объявление массива (1 вариант)
	var x [5]float64

	// Объявление массива (2 вариант)
	y := [5]float64{4.3, 2.2, 3.3, 5.0, 2.1}

	fmt.Println("Array x:", x)
	fmt.Println("Array y:", y)

	// Пример вычисления среднего значения 1
	var total1 float64 = 0
	for i := 0; i < len(y); i++ {
		total1 += y[i]
	}

	fmt.Println("Средний балл:", total1 / float64(len(y)))

	// Пример вычисления среднего значения 2
	var total2 float64 = 0
	for _, value := range y {
		total2 += value
	}

	fmt.Println("Средний балл:", total2 / float64(len(y)))
}

func slice() {
	// Создание среза
	x := make([]float64, 5, 10)
	fmt.Println(x)

	// Создание среза, используя [low:high)
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr[2:4])

	// Функции срезов append() copy()
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	// copy
	slice3 := []int{1,2,3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)
}

func mapp() {
	// Объявление карты
	var x map[string]int
	fmt.Println(x)

	// Инициализация карты
	y := make(map[string]int)
	y["key"] = 10
	fmt.Println(y)
}