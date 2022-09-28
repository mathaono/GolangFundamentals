package competition

import (
	"fmt"
	"time"
)

//Concorrência != Paralelismo

func Write(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
