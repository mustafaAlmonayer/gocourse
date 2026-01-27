package advanced

import (
	"fmt"
	"time"
)

func goRoutines() {
	fmt.Println("Beginning program.")
	go sayHello()
	fmt.Println("After sayHello function.")
	go printNums()
	go printAlphabets()
	time.Sleep(2 * time.Second)
}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Goroutine!!")
}

func printNums() {
	for i := 0; i < 5; i++ {
		fmt.Println(i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printAlphabets() {
	for _, letter := range "abcde" {
		fmt.Println(string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}
