package intermediate

import "fmt"

func fmtPackage() {

	// Printing Functions
	fmt.Print("Hello ")
	fmt.Print("World! ")
	fmt.Print(12, 356)
	fmt.Print("\n")

	fmt.Println("Hello ")
	fmt.Println("World! ")
	fmt.Println(12, 356)

	name := "John"
	age := 25

	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Binary: %b, hex: %X\n", age, age)

	// Formating Functions
	s := fmt.Sprint("Hello ", "World!", 123, 456)
	fmt.Print(s)

	s = fmt.Sprintln("Hello ", "World!", 123, 456)
	fmt.Print(s)

	s = fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(s)

	// Scanning Functions
	var nm string
	var ag int

	// normal scan
	fmt.Print("Emter your name & age: ")
	fmt.Scan(&nm, &ag)
	fmt.Printf("Your name %s,  Your Age: %d\n", nm, ag)

	// Scan new line
	fmt.Print("Emter your name & age: ")
	fmt.Scanln(&nm, &ag)
	fmt.Printf("Your name %s,  Your Age: %d\n", nm, ag)

	// Scan with certain format
	fmt.Print("Emter your name & age: ")
	fmt.Scanf("%s %d", &nm, &ag)
	fmt.Printf("Your name %s,  Your Age: %d\n", nm, ag)

	err := checkAge(15)
	if err != nil {
		fmt.Println("You can't drive", err)
	}

}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young to drive", age)
	}
	return nil
}
