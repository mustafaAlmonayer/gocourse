package main

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	err := os.Mkdir("subdir", 0755)
	checkError(err)

	// defer os.RemoveAll("subdir")

	os.WriteFile("subdir/file", []byte{}, 0755)
	checkError(os.MkdirAll("subdir/parent/child", 0755))
	checkError(os.MkdirAll("subdir/parent/child1", 0755))
	checkError(os.MkdirAll("subdir/parent/child2", 0755))
	checkError(os.MkdirAll("subdir/parent/child3", 0755))
	os.WriteFile("subdir/parent/file", []byte{}, 0755)

	result, err := os.ReadDir("subdir/parent")
	checkError(err)
	fmt.Println(result)

	for _, entry := range result {
		fmt.Println(entry)
	}
}
