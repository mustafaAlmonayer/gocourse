package basics

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func namingConventions() {

	// Pascal Case
	// Eg. CalculateArea, UserProfile, OrderDetails
	// Structs, Interfaces, Enums

	// snake case
	// Eg. user_id, first_name, http_request
	// files

	//UPERCASE
	// Use case is for Constants
	// Eg. MAX_CONNECTIONS, DEFAULT_TIMEOUT

	// Mixed Case
	// Eg. isValid, getUserName
	// Eg. Variables, Functions

	const MAX_CONNECTIONS = 100

	var employeeID = 12345
	fmt.Println("Employee ID:", employeeID)
}
