package intermediate

import "fmt"

type person struct {
	name string
	age  int
}

type employee struct {
	person
	empId  string
	salary float64
}

func (p person) introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.name, p.age)
}

// Override
func (e employee) introduce() {
	fmt.Printf("Hi, I'm %s and I earn %.2f.\n", e.empId, e.salary)
}

func structEmbedding() {

	emp := employee{
		person: person{
			name: "John",
			age:  30,
		},
		empId:  "E001",
		salary: 50000,
	}

	emp.introduce()
	fmt.Println(emp.name)
	fmt.Println(emp.age)
	fmt.Println(emp.empId)
	fmt.Println(emp.salary)

}
