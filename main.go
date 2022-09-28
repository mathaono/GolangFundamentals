package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	jsons "github.com/GolangFundamentals/JSONS"
	arraysslices "github.com/GolangFundamentals/arraysSlices"
	"github.com/GolangFundamentals/assistant"
	"github.com/GolangFundamentals/functions"
	"github.com/GolangFundamentals/genericInheritance"
	"github.com/GolangFundamentals/httpProtocol"
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

	/* maps.Maps()

	loops.Loops()

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		waitgroup.Write("teste 1")
		waitGroup.Done()
	}()

	go func() {
		waitgroup.Write("teste 2")
		waitGroup.Done()
	}()

	waitGroup.Wait()

	channel := make(chan string)
	channels.Write("teste 3", channel) */

	fmt.Println("-----------------------------------")

	user := jsons.Usuario{Nome: "Matheus", Idade: 26}
	fmt.Println(user)

	userJSON, error := json.Marshal(user)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(userJSON)
	fmt.Println(bytes.NewBuffer(userJSON))

	fmt.Println("-----------------------------------")

	userMap := map[string]string{
		"nome":  "Matheus",
		"idade": "26",
	}
	fmt.Println(userMap)

	userJSON2, error := json.Marshal(userMap)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(userJSON2)
	fmt.Println(bytes.NewBuffer(userJSON2))

	fmt.Println("-----------------------------------")

	userInJSON := `{"nome": "Matheus", "idade": 26}`
	fmt.Println(userInJSON)

	var userUnmarchal jsons.Usuario
	fmt.Println(userUnmarchal)

	error = json.Unmarshal([]byte(userInJSON), &userUnmarchal)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(userUnmarchal)

	fmt.Println("-----------------------------------")

	userInMap := `{"nome": "Carlos", "idade": "63"}`
	fmt.Println(userInMap)

	var userUnmarchal2 = make(map[string]string)
	fmt.Println(userUnmarchal2)

	error = json.Unmarshal([]byte(userInMap), &userUnmarchal2)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(userUnmarchal2)

	httpProtocol.UpServer()
}
