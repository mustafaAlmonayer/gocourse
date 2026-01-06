package intermediate

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func stringFunctions() {

	str := "Hello Go!"

	fmt.Println(len(str))

	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result)
	fmt.Println(str[0])

	fmt.Println(str[1:5])

	// String Conversion
	num := 18
	str3 := strconv.Itoa(num)
	fmt.Println(len(str3))
	fmt.Println(str3)

	// String splitting
	fruits := "Apple, Orange, Banana"
	parts := strings.Split(fruits, ",")

	for _, part := range parts {
		fmt.Println(strings.Trim(part, " "))
	}

	countries := []string{"Germany", "France", "Italy"}

	joined := strings.Join(countries, ", ")
	fmt.Println(joined)

	// Contains
	fmt.Println(strings.Contains(str, "G"))

	// Replace

	replaced := strings.ReplaceAll(str, "Go", "World")
	fmt.Println(replaced)

	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.ToUpper(str))

	// Regexp
	re := regexp.MustCompile(`\d+`)
	fmt.Println(re.FindAllString("Hello, 123 Go 11!", -1))

	// Unicode
	str6 := "Hello, 世界"
	fmt.Println(utf8.RuneCountInString(str6))

	// Builder

	builder := strings.Builder{}

	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World!")

	fmt.Println(builder.String())

}
