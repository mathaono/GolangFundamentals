package pointers

import "fmt"

func Pointers(v1 int64) {
	var1 := v1
	var2 := var1

	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	var3 := v1
	var pointer *int64
	pointer = &var3
	fmt.Println(var3, *pointer)

	var3++

	fmt.Println(var3, pointer)
	fmt.Println(var3, *pointer)
}
