package basics

import (
	"fmt"
)

func functions() {

	// Decleration of functions
	// func <name>(prameters) (return_types) {}
	fmt.Println(add(1, 2))

	//Functions Literals
	func() {
		fmt.Println("Hello Anonymous Function")
	}()

	greet := func() {
		fmt.Println("Hello from Greet Function")
	}
	greet()

	operation := add

	result := operation(3, 4)
	fmt.Println("Result:", result)
	fmt.Println(applyOpt(10, 20, func(x int, y int) int {
		return x + y
	}))

	operation2 := returnOpt(5, 15)
	fmt.Println("Operation2 Result:", operation2())
}

func add(a, b int) int {
	return a + b
}

func applyOpt(x int, y int, opt func(int, int) int) int {
	return opt(x, y)
}

func returnOpt(x int, y int) func() int {
	return func() int {
		return x + y
	}
}
