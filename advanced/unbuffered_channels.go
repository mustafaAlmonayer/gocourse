package advanced

import (
	"fmt"
	"time"
)

func unbufferedChannels() {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("2 seconds passed")
	}()

	go func() {
		ch <- 1
		time.Sleep(3 * time.Second)
		fmt.Println("3 seconds passed")
	}()
	recevier := <-ch
	println(recevier)
}
