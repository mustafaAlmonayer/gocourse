package main

import (
	"fmt"
	"time"
)

func main() {

	// Current Local Time
	fmt.Println(time.Now())

	specificTime := time.Date(2024, time.July, 30, 12, 0, 0, 0, time.UTC)
	fmt.Println(specificTime)

	parsedTime, _ := time.Parse("2006-01-02", "2020-05-01")
	fmt.Println(parsedTime)
}
