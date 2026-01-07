package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func textTemplates() {
	tmpl := template.Must(template.New("example").Parse("Welcome, {{.name}}! how are you doing?\n"))

	data := map[string]any{
		"name": "John",
	}

	if err := tmpl.Execute(os.Stdout, data); err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	if err := tmpl.Execute(os.Stdin, map[string]string{"name": strings.TrimSpace(name)}); err != nil {
		panic(err)
	}
}
