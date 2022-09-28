package channels

import (
	"fmt"
	"time"
)

func Write(texto string, channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
