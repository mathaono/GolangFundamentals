package competition

import (
	"fmt"
	"time"
)

//ConcorrĂȘncia != Paralelismo

func Write(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
