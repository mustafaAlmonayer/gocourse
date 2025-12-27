package basics

// only accessible within package main
import "fmt"

var middleName = "Cane"

func variables() {
	// var age int
	// var name string = "John"
	// var name1 = "Jane"

	// count := 10
	// lastName := "Smith"
	// middleName := "Cane"

	// Default values
	// Numeric types default to 0
	// String type defaults to ""
	// Boolean type defaults to false
	// Pointers, slices, maps, functions, structs, channels, and interfaces default to nil

	// ---- Scope
	middleName := "Alexander"
	fmt.Println(middleName) // This will cause an error because firstName is not in scope here
	printName()

}

func printName() {
	firstName := "Michael"
	fmt.Println(firstName, middleName)
}
