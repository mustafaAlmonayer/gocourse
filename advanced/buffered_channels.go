package advanced

import (
	"fmt"
	"time"
)

func bufferedChannels() {
	// make
	// Block on send when the buffer is full.
	// Block on receive when the buffer is empty.
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println(<-ch)
	}()
	ch <- 3
	fmt.Println("End")

}
