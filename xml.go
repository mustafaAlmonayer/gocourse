package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person44 struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	City    string   `xml:"city"`
	Email   string   `xml:"email"`
}

func main() {

	person := Person44{Name: "John", Age: 30, City: "London", Email: "email@example.com"}

	data, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(data))

}
