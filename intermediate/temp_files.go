package intermediate

import (
	"fmt"
	"os"
)

func tempFile() {
	tempFileName := "temporaryFile"
	tempFile, err := os.CreateTemp("", tempFileName)

	checkError(err)

	fmt.Println("Temporal file created", tempFile.Name())

	defer os.Remove(tempFile.Name())
	defer tempFile.Close()
}
