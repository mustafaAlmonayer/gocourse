package intermediate

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func directories() {
	err := os.Mkdir("subdir", 0755)
	checkError(err)

	// defer os.RemoveAll("subdir")

	os.WriteFile("subdir/file", []byte{}, 0755)
	checkError(os.MkdirAll("subdir/parent/child", 0755))
	checkError(os.MkdirAll("subdir/parent/child1", 0755))
	checkError(os.MkdirAll("subdir/parent/child2", 0755))
	checkError(os.MkdirAll("subdir/parent/child3", 0755))
	os.WriteFile("subdir/parent/file", []byte{}, 0755)
	os.WriteFile("subdir/parent/child/file", []byte{}, 0755)

	result, err := os.ReadDir("subdir/parent")
	checkError(err)
	fmt.Println(result)

	for _, entry := range result {
		fmt.Println(entry)
	}

	fmt.Println("Reading...")

	checkError(os.Chdir("subdir/parent/child"))

	result, err = os.ReadDir(".")

	checkError(err)
	for _, entry := range result {
		fmt.Println(entry)
	}

	checkError(os.Chdir("../../.."))

	dir, err := os.Getwd()
	checkError(err)
	fmt.Println(dir)

	pathFile := "./subdir"

	fmt.Println("Walking Dir")

	err = filepath.WalkDir(
		pathFile,
		func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path)
			return nil
		},
	)

	checkError(err)

	err = os.RemoveAll("./subdir")
	checkError(err)
}
