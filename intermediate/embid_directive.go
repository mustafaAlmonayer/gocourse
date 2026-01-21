package intermediate

import (
	"embed"
	"fmt"
)

// //go:embed example.txt
var content string

// //go:embed example
var exampleFolder embed.FS

func embidDirective() {
	fmt.Println(content)
	cont, err := exampleFolder.ReadFile("example/example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(cont))
}
