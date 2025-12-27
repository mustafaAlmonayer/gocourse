package basics

import "fmt"

func whileLoop() {
	i := 1
	for i <= 5 {
		fmt.Println("Iteration:", i)
		i++
	}

	// infinite loop example
	// 	for {
	// 		fmt.Println("This will run forever")
	// 	}

	sum := 0
	for {
		sum += 10
		if sum >= 50 {
			break
		}
	}
	fmt.Println("Final Sum:", sum)
}
