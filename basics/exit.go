package basics

import (
	"fmt"
	"os"
)

func exit() {
	fmt.Println("Starting the application...")
	os.Exit(0)
	fmt.Println("This line will not be executed.")
}
