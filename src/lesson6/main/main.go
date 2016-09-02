package main

import "fmt"

func main() {

	marks := [5]float64{1.2, 2.5, 5.0, 2.3, 4.3}

	result := average(marks)

	fmt.Println("Среднее значение:", result)

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	fmt.Println("Factorial: ", fact(4))
}

// Вычисление среднего значения
func average(marks [5]float64) float64 {
	sum := 0.0
	for i := 0; i < len(marks); i++ {
		sum += marks[i]
	}

	return sum / float64(len(marks))
}

// Замыкание
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func fact(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * fact(x-1)
}