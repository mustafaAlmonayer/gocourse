package intermediate

import (
	"os"
)

func writingFiles() {
	file, err := os.Create("output.txt")

	if err != nil {
		panic(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	date := []byte("Hello World!\n\n\n\n\n")
	_, err = file.Write(date)
	if err != nil {
		panic(err)
	}

	file2, err := os.Create("write_string.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := file2.Close()
		if err != nil {
			panic(err)
		}
	}()

	_, err = file2.WriteString("Hello Go!\n")
	if err != nil {
		panic(err)
	}
}
