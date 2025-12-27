package basics

import "fmt"

func forLoop() {
	// Over Range
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	// Over Collection
	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, number := range numbers {
		fmt.Println("Index:", index, "Number:", number)
	}

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	// new syntax
	for i := range 10 {
		fmt.Println(i)
	}
}
