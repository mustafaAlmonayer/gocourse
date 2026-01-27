package advanced

import (
	"fmt"
	"time"
)

func channels() {
	// make
	greeting := make(chan string)

	greetString := "Hello"
	go func() {
		greeting <- greetString // blocking
		fmt.Println("Hi")
		greeting <- greetString
	}()
	go func() {
		for val := range greeting {
			println(val)
		}
	}()
	time.Sleep(1 * time.Second)
}
