package main

import "fmt"

func main() {

	var first, second bool
	var third bool = true
	fourth := !third
	var fifth = true

	fmt.Println("first  = ", first)
	fmt.Println("second = ", second)
	fmt.Println("third  = ", third)
	fmt.Println("fourth = ", fourth)
	fmt.Println("fifth  = ", fifth, "\n")

	fmt.Println("!true  = ", !true)
	fmt.Println("!false = ", !false, "\n")
	fmt.Println("true && true   = ", true && true)
	fmt.Println("true && false  = ", true && false)
	fmt.Println("false && false = ", false && false, "\n")

	fmt.Println("true || true   = ", true || true)
	fmt.Println("true || false  = ", true || false)
	fmt.Println("false || false = ", false || false, "\n")

	fmt.Println("2 < 3  = ", 2 < 3)        // true
	fmt.Println("2 > 3  = ", 2 > 3)        // false
	fmt.Println("3 < 3  = ", 3 < 3)        // false
	fmt.Println("3 <= 3 = ", 3 <= 3)       // true
	fmt.Println("3 > 3  = ", 3 > 3)        // false
	fmt.Println("3 >= 3 = ", 3 >= 3)       // true
	fmt.Println("2 == 3 = ", 2 == 3)       // false
	fmt.Println("3 == 3 = ", 3 == 3)       // true
	fmt.Println("2 != 3 = ", 2 != 3)       // true
	fmt.Println("3 != 3 = ", 3 != 3, "\n") // false
}
