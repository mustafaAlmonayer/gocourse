package intermediate

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age,omitempty"`
	Address Address `json:"address"`
}

type Employee struct {
	FullName string  `json:"full_name"`
	EmpId    string  `json:"emp_id"`
	Age      int     `json:"age"`
	Address  Address `json:"address"`
}

func jsonExp() {
	person := Person{Name: "John", Age: 30}

	data, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	person1 := Person{Name: "John"}

	data, err = json.Marshal(person1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	person2 := Person{
		Name: "Jane",
		Age:  24,
		Address: Address{
			City:  "New York",
			State: "NY",
		},
	}

	data, err = json.Marshal(person2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	jsonData :=
		`{
			"full_name": "Jenny Doe",
			"emp_id": "0009",
			"age": 30,
			"address": {
				"city": "San Jose",
				"state": "CA"
			}
		}`

	emp := Employee{}
	err = json.Unmarshal([]byte(jsonData), &emp)

	if err != nil {
		panic(err)
	}
	fmt.Println(emp)

	// Dynamic JSON
	jsonData2 :=
		`{
			"name": "John",
			"age": 30,
			"address": {
				"city": "New York",
				"state": "NY"
			}
		}`
	jsonMap := make(map[string]any)
	err = json.Unmarshal([]byte(jsonData2), &jsonMap)

	if err != nil {
		panic(err)
	}
	fmt.Println(jsonMap)
}
