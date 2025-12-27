package basics

import (
	"fmt"
)

func swithSt() {
	fruit := "apple"
	switch fruit {
	case "apple":
		fmt.Println("It's an apple!")
	case "orange":
		fmt.Println("It's an orange!")
	case "banana":
		fmt.Println("It's a banana!")
	default:
	}

	day := "Monday"
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	number := 15
	switch {
	case number < 10:
		fmt.Println("The number is less than 10.")
	case number >= 10 && number <= 20:
		fmt.Println("The number is between 10 and 20.")
	default:
		fmt.Println("The number is greater than 20.")
	}

	num := 2
	switch {
	case num > 1:
		fmt.Println("One")
		fallthrough
	case num == 2:
		fmt.Println("Two")
	default:
		fmt.Println("Other number")
	}

	checkTyepe(10)
	checkTyepe(10.13)
	checkTyepe("dsdaw")
	checkTyepe(true)

}

func checkTyepe(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("x is an integer")
	case float64:
		fmt.Println("x is a float")
	case string:
		fmt.Println("x is a string")
	case bool:
		fmt.Println("x is a boolean")
	default:
		fmt.Println("Unknown type")
	}
}
