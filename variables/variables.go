package variables

import "fmt"

func Variable() {
	var var1 string = "Declaração explícita de variável"
	var2 := "Declaração implícita de variável"

	fmt.Println(var1)
	fmt.Println(var2)
}

type Numbers struct {
	num1 int64
	num2 int64
}
