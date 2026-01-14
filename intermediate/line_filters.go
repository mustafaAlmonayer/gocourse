package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func lineFilters() {

	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	keyword := "important"

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			updatedLine := strings.ReplaceAll(line, keyword, "necessary")
			fmt.Println(updatedLine)
		}
	}

}
