package intermediate

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

func xmlExp() {

	person := Person44{Name: "John", Age: 30, City: "London", Email: "email@example.com"}

	data, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(data))

	xmlStr := `
		<person>
			<name>Jane</name>
			<age>25</age>
		</person>
	`

	person44 := Person44{}

	err = xml.Unmarshal([]byte(xmlStr), &person44)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(person44)

}
