package intermediate

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

type Person1 struct {
	firstName string
	lastName  string
	age       int
	// Helper Struct
	address Address
}

type Person2 struct {
	firstName string
	lastName  string
	age       int
	// Empedded Struct (Struct Colapse into this one only one . used to access)
	Address
}

type Address struct {
	city    string
	country string
}

func (p Person) printName() {
	fmt.Println(p.firstName, p.lastName)
}

func (p Person) getname() string {
	return p.firstName + " " + p.lastName
}

// Can change value
func (p *Person) incrementAge() {
	p.age++
}

// Cant change value
func (p Person) incrementAgeNoPointer() {
	p.age++
}

func structs() {

	p := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
	}

	fmt.Println(p.firstName)
	fmt.Println(p.lastName)

	p.printName()

	fmt.Println(p.getname())

	// Anonymous Structs
	user := struct {
		username string
		email    string
	}{
		username: "user",
		email:    "example@email.com",
	}
	fmt.Println(user.username)

	fmt.Println("Before", p.age)
	p.incrementAgeNoPointer()
	fmt.Println("After", p.age)

	fmt.Println("Before", p.age)
	p.incrementAge()
	fmt.Println("After", p.age)

}
