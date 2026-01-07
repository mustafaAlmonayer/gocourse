package intermediate

import (
	"fmt"
	"regexp"
)

func regex() {

	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	email1 := "user@email.com"
	email2 := "invalid_email"
	fmt.Println(re.MatchString(email1))
	fmt.Println(re.MatchString(email2))

	// groups
	re = regexp.MustCompile(`(\d){4}-(\d{2})-(\d{2})`)

	data := "2024-07-30"

	fmt.Println(re.FindStringSubmatch(data))
}
