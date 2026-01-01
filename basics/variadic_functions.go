package basics

import (
	"fmt"
)

func variadicFunctions() {

	// ... Ellipsis
	fmt.Println("Sum of 1,2,3,4,5:", sum(1, 2, 3, 4, 5))
	fmt.Println("Sum of 10,20,30:", sum(10, 20, 30))
}

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}
