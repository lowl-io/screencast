package main

import (
	"fmt"
	"math"
)

func main()  {

	defaultFl := 1.2

	var fl32 float32 = 5.5
	var fl64 float64 = 12.3

	fmt.Printf("defaultFl = %T %v\n", defaultFl, defaultFl)
	fmt.Println("fl32 =", fl32)
	fmt.Println("fl64 =", fl64)


	fmt.Println("MAX float32        =", math.MaxFloat32)
	fmt.Println("MIN float32        =", math.SmallestNonzeroFloat32)

	fmt.Println("MAX float64        =", math.MaxFloat64)
	fmt.Println("MIN float64        =", math.SmallestNonzeroFloat64)

}
