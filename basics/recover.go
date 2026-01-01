package basics

import "fmt"

func recoverEx() {
	// Recover from panic explanation
	// In Go, the recover function is used to regain control of a panicking goroutine.
	// It can only be called within a deferred function. When a panic occurs,
	// the deferred function can call recover to retrieve the value passed to panic
	// and stop the panic from propagating further.
	process3()
}

func process3() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in process3:", r.(string))
		}
	}()
	fmt.Println("Calling process3")
	panic("Something went wrong in process3")
}
