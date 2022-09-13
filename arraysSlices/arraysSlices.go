package arraysslices

import "fmt"

type Pessoa struct {
	Nome string
	Cpf  string
}

func Arrays(p Pessoa) {

	slice := []Pessoa{p}
	fmt.Println(slice)
}

//Arrays Internos

func InternalSlices() {

	slice := make([]int, 10, 20)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
