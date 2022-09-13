package main

import (
	"fmt"

	arraysslices "github.com/GolangFundamentals/arraysSlices"
	"github.com/GolangFundamentals/assistant"
	"github.com/GolangFundamentals/functions"
	"github.com/GolangFundamentals/genericInheritance"
	"github.com/GolangFundamentals/loops"
	"github.com/GolangFundamentals/maps"
	"github.com/GolangFundamentals/operators"
	"github.com/GolangFundamentals/pointers"
	"github.com/GolangFundamentals/structs"
	"github.com/GolangFundamentals/variables"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("teste")
	assistant.Write()
	err := checkmail.ValidateFormat("aono.1304gmail.com.abc")
	fmt.Println(err)

	variables.Variable()

	sum := functions.Add(10, 20)
	fmt.Println(sum)

	functions.Funcs("Declarando uma variável do tipo func")
	sum, subtract := functions.Calculations(50, 25)
	fmt.Println(sum, subtract)

	fmt.Println(operators.Operators(100, 20))

	structs.Structs("Matheus", "Carlos", 26, 64)
	usr3 := structs.Usuario{}
	usr3.Nome = "Marly"
	usr3.Idade = 62
	fmt.Println(usr3)

	p1 := genericInheritance.People{Name: "Matheus", Age: 26, Height: 1.84}
	e1 := genericInheritance.Student{College: "Anhembi", Course: "ADS", People: p1}
	fmt.Println(e1)

	pointers.Pointers(10)

	pessoa1 := arraysslices.Pessoa{Nome: "Matheus", Cpf: "470.303.228-78"}
	arraysslices.Arrays(pessoa1)

	arraysslices.InternalSlices()

	maps.Maps()

	loops.Loops()
}
