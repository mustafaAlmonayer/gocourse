package intermediate

import (
	"flag"
	"fmt"
	"os"
)

func commandLine() {

	fmt.Println("Command:", os.Args[0])
	fmt.Println("Arg1:", os.Args[1])

	var name string
	var age int
	var male bool

	flag.StringVar(&name, "name", "", "Name of the user")
	flag.IntVar(&age, "age", 0, "Age of the user")
	flag.BoolVar(&male, "male", false, "Is male")

	flag.Parse()

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(male)

}
