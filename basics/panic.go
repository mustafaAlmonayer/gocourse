package basics

import "fmt"

func panicEx() {
	// Panic is a built-in function that stops the ordinary flow of control
	// and begins panicking. When the function F calls panic, execution of F
	// stops, any deferred functions in F are executed normally, and then
	// F returns to its caller. To the caller, F then behaves like a call to
	// panic. This continues up the stack until all functions in the current
	// goroutine have returned, at which point the program crashes with a
	// log message that includes the value passed to panic.

	// valid input
	process2(10)

	// invalid input to demonstrate panic
	process2(-5)

	// Handling panic using recover
}

func process2(input int) {

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	if input < 0 {
		defer fmt.Println("Before Panic")
		panic("negative input not allowed")
	}
	fmt.Println("Processing:", input)
}
