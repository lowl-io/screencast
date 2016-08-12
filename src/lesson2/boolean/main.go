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
	fmt.Println("fifth  = ", fifth)
	fmt.Println()

	fmt.Println("!true  = ", !true)
	fmt.Println("!false = ", !false)
	fmt.Println("true && true   = ", true && true)
	fmt.Println("true && false  = ", true && false)
	fmt.Println("false && false = ", false && false)
	fmt.Println()

	fmt.Println("true || true   = ", true || true)
	fmt.Println("true || false  = ", true || false)
	fmt.Println("false || false = ", false || false)
	fmt.Println()

	fmt.Println("2 < 3  = ", 2 < 3)
	fmt.Println("2 > 3  = ", 2 > 3)
	fmt.Println("3 < 3  = ", 3 < 3)
	fmt.Println("3 <= 3 = ", 3 <= 3)
	fmt.Println("3 > 3  = ", 3 > 3)
	fmt.Println("3 >= 3 = ", 3 >= 3)
	fmt.Println("2 == 3 = ", 2 == 3)
	fmt.Println("3 == 3 = ", 3 == 3)
	fmt.Println("2 != 3 = ", 2 != 3)
	fmt.Println("3 != 3 = ", 3 != 3)
}
