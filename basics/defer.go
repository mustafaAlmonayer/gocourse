package basics

import "fmt"

func deferEx() {

	// Deffer Explnation
	// A deferred function's execution is postponed until the surrounding function returns.
	// Deferred functions are often used for purposes such as resource cleanup, closing files, or releasing locks.

	process()
}

func process() {
	defer fmt.Println("Deferred Statement Excution!")
	fmt.Println("Normal Execution Statement!")
}
