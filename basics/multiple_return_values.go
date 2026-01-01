package basics

import "errors"

func multyReturnVal() {
	a, b := 10, 3
	quotient, remainder := divide(a, b)
	println("Quotient:", quotient)
	println("Remainder:", remainder)

	quotient2, remainder2 := divideShortSyntax(a, b)
	println("Quotient:", quotient2)
	println("Remainder:", remainder2)

	result, err := compare(a, b)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Comparison Result:", result)
	}
}

func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func divideShortSyntax(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if a < b {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Unable to compare which is greater")
	}
}
