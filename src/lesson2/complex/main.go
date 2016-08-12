package main

import (
	"fmt"
)

func main() {

	someVar1 := 15.2 + 2.7i
	someVar2 := complex64(98.3 + 21.4i)
	someVar3 := complex(12.1, 3.4)

	fmt.Printf("someVar1 = %T %v\n", someVar1, someVar1);
	fmt.Printf("someVar2 = %T %v\n", someVar2, someVar2);
	fmt.Printf("someVar3 = %T %v\n", someVar3, someVar3);
}
