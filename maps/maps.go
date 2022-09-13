package maps

import "fmt"

func Maps() {

	usuario := map[string]string{
		"nome":      "Matheus",
		"sobrenome": "Aono",
	}

	fmt.Println(usuario)
	fmt.Println("-------------------")

	usuario2 := map[string]map[string]string{
		"pessoa": {
			"nome":      "matheus",
			"sobrenome": "aono",
		},
		"formacao": {
			"curso":  "ads",
			"status": "formado",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "formacao")

	fmt.Println(usuario2)

	usuario2["estado"] = map[string]string{
		"UF": "SP",
	}

	fmt.Println(usuario2)
}
