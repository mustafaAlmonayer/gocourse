package main

import (
	"os"
	"text/template"
)

func main() {
	tmpl := template.Must(template.New("example").Parse("Welcome, {{.name}}! how are you doing?"))

	data := map[string]any{
		"name": "John",
	}

	if err := tmpl.Execute(os.Stdout, data); err != nil {
		panic(err)
	}
}
