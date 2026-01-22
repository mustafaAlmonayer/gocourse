package intermediate

import (
	"flag"
	"fmt"
	"os"
)

func cliSubCommands() {

	subcommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subcommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	firstFlag := subcommand1.Bool("processing", false, "Command processing status")
	secondFlag := subcommand1.Int("bytes", 1024, "Byte length for result")
	thirdFlag := subcommand2.String("language", "Go", "Enter your language")

	if len(os.Args) < 2 {
		panic("You need more args")
	}

	switch os.Args[1] {
	case "firstSub":
		subcommand1.Parse(os.Args[2:])
		fmt.Println(*firstFlag)
		fmt.Println(*secondFlag)
	case "secondSub":
		subcommand2.Parse(os.Args[2:])
		fmt.Println(*thirdFlag)
	default:
		panic("No commands")
	}

}
