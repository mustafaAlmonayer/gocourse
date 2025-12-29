package basics

import "fmt"

func rangeEx() {
	message := "Hello Wolrd!"

	for i, v := range message {
		fmt.Println(i, v)
		fmt.Printf("Index %d, Rune: %c\n", i, v)
	}
}
