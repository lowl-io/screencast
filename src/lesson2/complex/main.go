package main

import (
	"fmt"
	"reflect"
)

func main() {

	someVar1 := 15.2 + 2.7i
	someVar2 := complex64(98.3 + 21.4i)
	someVar3 := complex(12.1, 3.4)

	fmt.Println("someVar1 =", reflect.TypeOf(someVar1), someVar1);
	fmt.Println("someVar2 =", reflect.TypeOf(someVar2), someVar2);
	fmt.Println("someVar3 =", reflect.TypeOf(someVar3), someVar3);
}
