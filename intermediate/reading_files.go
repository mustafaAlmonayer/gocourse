package intermediate

import (
	"bufio"
	"fmt"
	"os"
)

func readingFiles() {

	file, err := os.Open("output.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	// data := make([]byte, 17)
	// _, err = file.Read(data)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}
}
