package intermediate

import (
	"fmt"
	"strconv"
)

func numberParsing() {

	numString := "123456"
	num, err := strconv.Atoi(numString)
	if err != nil {
		panic(err)
	}
	fmt.Println(num)

	num1, err := strconv.ParseInt(numString, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(num1)

	floatString := "3.14"

	num2, err := strconv.ParseFloat(floatString, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(num2)

	binaryString := "1010"
	num3, err := strconv.ParseInt(binaryString, 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(num3)

}
