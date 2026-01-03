package intermediate

import (
	"fmt"
	"unicode/utf8"
)

func stringRunes() {
	message := "Hello, Go"
	rawMessage := `Hello, \nGo` // discard escape sequences
	englishAndChineseMessage := "Hello, 世界"

	fmt.Println(message)
	fmt.Println(rawMessage)

	// A string is an array of unicode code characters called runes
	// a rune is an alias for int32 that represents a unicode character

	fmt.Println("Length of message variable is: ", len(message))
	fmt.Println("Length of rawMessage variable is: ", len(rawMessage))
	fmt.Println("Length of englishAndChineseMessage variable is: ", len(englishAndChineseMessage))

	// len() return back the byte array length and not the string length in characters

	fmt.Println("First char in message", message[0]) // ASCII
	fmt.Println("First char in rawMessage", rawMessage[0])
	fmt.Println("First char in englishAndChineseMessage", englishAndChineseMessage[0])

	fmt.Println("message variable")
	for i, char := range message {
		fmt.Println(i, char)
	}
	fmt.Println("rawMessage variable")

	for i, char := range rawMessage {
		fmt.Println(i, char)
	}
	fmt.Println("englishAndChineseMessage variable")
	for i, char := range englishAndChineseMessage {
		fmt.Println(i, char)
	}

	fmt.Println("Rune count for message variable", utf8.RuneCountInString(message))
	fmt.Println("Rune count for rawMessage variable", utf8.RuneCountInString(rawMessage))
	fmt.Println("Rune count for message englishAndChineseMessage", utf8.RuneCountInString(englishAndChineseMessage))

	// Define a Rune
	var ch rune = 'a'
	fmt.Printf("%c\n", ch)
	jps := '日'
	fmt.Printf("%c\n", jps)

	cstr := string(ch)

	fmt.Printf("%T\n", cstr)
}
