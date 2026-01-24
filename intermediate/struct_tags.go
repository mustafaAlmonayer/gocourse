package intermediate

import (
	"encoding/json"
	"fmt"
	"log"
)

// - to always omit
// omitempty to omit when zero value is presint
// default is alwaus include
type Person22 struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
	Age       int    `json:"age"`
}

func structTags() {
	person := Person22{FirstName: "Jane", LastName: "Doe", Age: 50}

	date, err := json.Marshal(person)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(date))
}
