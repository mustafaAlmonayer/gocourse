package advanced

import (
	"fmt"
)

func channelSynchronization() {
	done := make(chan string)
	go func() {

		for i := range 5 {
			done <- "Hello from goroutine" + fmt.Sprint(i)

		}
		close(done)

	}()

	for msg := range done {
		fmt.Println(msg)
	}
	fmt.Println("It's finished")
}
