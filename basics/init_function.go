package basics

import "fmt"

// init is called before the main function
// its a special function in Go
// used for initialization tasks
// such as setting up configurations or resources
func init() {
	// Initialization code here
	fmt.Println("Initialization function called")
}

// Multiple init functions can be defined in the same package
func init() {
	// Initialization code here
	fmt.Println("Initialization function2 called")
}

func init() {
	// Initialization code here
	fmt.Println("Initialization function3 called")
}

func initFunctions() {

	fmt.Println("Main function called")

}
