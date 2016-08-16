package main

import "fmt"

// Переменная globalVar глобальная и видна для всех функциий
var globalVar string = "Hello, I`am global varibale"

func main() {

	// Переменная globalVar видна только внтури main() и
	// во всех вложенных блоках, находящихся в main()
	//var globalVar string = "Hello, I`am global varibale"

	{
		//zxy := 5
		fmt.Println(globalVar)
	}

	//fmt.Println(zxy)

	// Полная форма объявления перменной
	var x string = "Hello X"
	fmt.Println(x)

	// Объявление перменной с ее последующей инициализацией
	var y string
	y = "Hello Y"
	fmt.Println(y)

	// Краткая форма объявления перменной
	z := "Hello Z"
	fmt.Println(z);

	// Форма записи объявления переменной без явного задания типа
	var a = "Hello A"
	fmt.Println(a)

	// Объявление константы
	const xyz string = "Hello XYZ"
	fmt.Println(xyz);

	// Ошибка! Значение константы не может быть изменено
	//xyz = "Hello const"

	f();

	// Объявление нескольких переменных (работает и с const)
	var (
		b = 4
		c = -3
		d = 1.2
	)

	fmt.Println(b, c, d);
}

func f() {
	fmt.Println(globalVar);
}