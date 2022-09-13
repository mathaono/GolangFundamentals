package loops

import "fmt"

func Loops() {
	fmt.Println("----------------")

	nomes := []string{"Matheus", "Carlos", "Marly"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	fmt.Println("----------------")

	pessoas := map[string]string{
		"nome":      "matheus",
		"sobrenome": "aono",
	}

	for chave, valor := range pessoas {
		fmt.Println(chave, valor)
	}
}
