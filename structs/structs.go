package structs

import "fmt"

type Usuario struct {
	Nome  string
	Idade int64
}

func Structs(user1, user2 string, idade1, idade2 int64) {
	usr1 := Usuario{user1, idade1}
	fmt.Println(usr1)

	var usr2 = Usuario{}
	usr2.Nome = user2
	usr2.Idade = idade2
	fmt.Println(usr2)
}
