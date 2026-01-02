package intermediate

import (
	"fmt"
)

func recursion() {
	fmt.Println("Factorial of 5 is:", factorial(5))
	fmt.Println("Sum of digits of 1234 is:", someOfDigits(1234))
}

func factorial(n int) int {
	fmt.Println("N is:", n)
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func someOfDigits(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + someOfDigits(n/10)
}
