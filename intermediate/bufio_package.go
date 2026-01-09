package intermediate

import (
	"bufio"
	"fmt"
	"os"
)

func bufioEx() {

	// reader := bufio.NewReader(strings.NewReader("Hello, bufio package!\n"))
	// data := make([]byte, 21)

	// n, err := reader.Read(data)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(n, string(data))

	// str, err := reader.ReadString('\n')
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(str)

	writer := bufio.NewWriter(os.Stdout)

	data := []byte("Hello bufio package!\n")

	n, err := writer.Write(data)
	if err != nil {
		panic(err)
	}
	err = writer.Flush()
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

	str := "This is a string.\n"
	n, err = writer.WriteString(str)
	if err != nil {
		panic(err)
	}
	err = writer.Flush()
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

}
