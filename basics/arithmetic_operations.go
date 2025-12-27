package basics

import "fmt"

func arithmetic_operations() {
	var a, b = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Subtraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division:", result)

	result = a % b
	fmt.Println("Modulus:", result)

	const p float64 = 22 / 7.0
	fmt.Println("Constant p (22/7):", p)
}
