package intermediate

import (
	"fmt"
	"os"
	"strings"
)

func envs() {
	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println(user)
	fmt.Println(home)

	err := os.Setenv("FRUIT", "apple")
	if err != nil {
		panic(err)
	}

	for _, e := range os.Environ() {
		fmt.Println(strings.SplitN(e, "=", 2))
	}

	err = os.Unsetenv("FRUIT")
	if err != nil {
		panic(err)
	}

	fmt.Println(os.Getenv("FRUIT"))
}
