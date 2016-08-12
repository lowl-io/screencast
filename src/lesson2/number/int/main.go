package main

import (
	"fmt"
	"math"
)

func main()  {
	// Основные операции
	fmt.Println("Основные операции")

	fmt.Println("3 + 3          =", 3 + 3)
	fmt.Println("4 - 3          =", 4 - 3)
	fmt.Println("6 * 3          =", 6 * 3)
	fmt.Println("9 / 16  	    =", 9 / 16)
	fmt.Println("15 % 4  	    =", 15 % 4)

	// Операции с присваиванием
	fmt.Println("Операции с присваиванием")

	var someVar int = 5;
	fmt.Println("someVar(source) =", someVar)

	someVar--
	fmt.Println("someVar--       =", someVar)

	someVar++
	fmt.Println("someVar++       =", someVar)

	someVar += 10
	fmt.Println("someVar += 10   =", someVar)

	someVar -= 10
	fmt.Println("someVar -= 10   =", someVar)

	someVar *= 10
	fmt.Println("someVar *= 10   =", someVar)

	someVar /= 10
	fmt.Println("someVar /= 10   =", someVar)

	someVar %= 10
	fmt.Println("someVar %= 10   =", someVar)


	sizeOfInteger()
}

func sizeOfInteger() {
	fmt.Println("Пределы значений целых int")

	fmt.Printf("int8:   [%v, %v]\n", math.MinInt8,  math.MaxInt8)
	fmt.Printf("int16:  [%v, %v]\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("int32:  [%v, %v]\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("int64:  [%v, %v]\n", math.MinInt64, math.MaxInt64)

	fmt.Println("\nПределы значений целых uint")

	fmt.Printf("uint8:  [0, %v]\n", math.MaxUint8)
	fmt.Printf("uint16: [0, %v]\n", math.MaxUint16)
	fmt.Printf("uint32: [0, %v]\n", math.MaxUint32)
	fmt.Printf("uint64: [0, %d]\n", uint64(math.MaxUint64))

	fmt.Println("\nСинонимы целых типов")

	fmt.Println("byte    - int8")
	fmt.Println("rune    - int32")
	fmt.Println("int     - int32, или int64, в зависимости от ОС")
	fmt.Println("uint    - uint32, или uint64, в зависимости от ОС")
	fmt.Println("uintptr - Беззнаковое целое, пригодное для хранения значения указателя")
}