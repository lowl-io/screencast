package main

import "fmt"

func main() {

	// Вывод значений от 1 до 10
	for i := 1; i <= 10; i++ {
		//fmt.Print(i, " ")
	}

	// Проверка на четность числа
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println("Число:", i, "четное")
		} else {
			fmt.Println("Число:", i, "не четное")
		}
	}

	// Пример синтаксиса с else if
	i := 1
	if i == 0 {
		// do some action
	} else if i == 5 {
		// do some action
	} else if i == 4 {
		// do some action
	} else {
		// В это блок управление перейдет, если
		// все условия выше окаались false
		fmt.Println("Hello world!")
	}


	// Оператор switch и пример программы: Распорядок дня
	day := 3
	switch day {
	case 1:
		fmt.Println("Понедельник: Работа, тренировка")
	case 2:
		fmt.Println("Вторник: Работа")
	case 3:
		fmt.Println("Среда: Работа, Бассейн")
	case 4:
		fmt.Println("Четверг: Работа, Прогулка")
	case 5:
		fmt.Println("Пятница: Работа, тренировка")
	case 6:
		fmt.Println("Суббота: Природа")
	case 7:
		fmt.Println("Воскресенье: Играть весь день в ПК")
	default:
		fmt.Println("Нет такого дня недели")
	}
}
