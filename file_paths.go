package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	relativePath := "./data/file.txt"
	absolutePath := "/home/user/docs/file.txt"

	joinedPath := filepath.Join("downloads", "file.zip")

	fmt.Println("Joined Path:", joinedPath)

	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println(normalizedPath)

	dir, file := filepath.Split("/home/user/docs/file.txt")
	fmt.Println(dir, file)

	fmt.Println(filepath.Base("/home/user/docs/file.txt"))

	fmt.Println(filepath.Abs(relativePath))
	fmt.Println(filepath.Abs(absolutePath))

	fmt.Println(filepath.IsAbs(relativePath))
	fmt.Println(filepath.IsAbs(absolutePath))

	fmt.Println(filepath.Ext(file))
	fmt.Println(strings.TrimSuffix(file, filepath.Ext(file)))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/c", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	absPath, err := filepath.Abs(relativePath)
}
