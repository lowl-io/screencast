package main

import "fmt"

func main() {

	//r := Rectangle{Width: 2, Height: 3}
	//
	//fmt.Println("Perim:", r.perim())

	r := Rectangle{Width: 3, Height: 5}
	c := Circle{Radius: 6}

	fmt.Println("Rectangle:")
	measure(r)

	fmt.Println("Circle:")
	measure(c)
}
