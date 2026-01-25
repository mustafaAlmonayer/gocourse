package main

import "fmt"

func main() {
	var a int = 32
	b := int32(a)
	c := float64(b)

	e := 3.14
	f := int(e)
	fmt.Println(f, c)

	// Type(value)
	g := "Hello"
	h := []byte(g)
	fmt.Println(h)
}
