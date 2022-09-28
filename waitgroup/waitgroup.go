package waitgroup

import (
	"fmt"
	"time"
)

//Concorrência != Paralelismo

func Write(texto string) {
	for i := 0; i < 1; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
