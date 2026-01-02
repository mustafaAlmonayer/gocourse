package intermediate

import "fmt"

func pointers() {

	var ptr *int

	// pointer to pointer
	// var ptr **int

	var a int = 10

	fmt.Println(a)
	fmt.Println(ptr)
	// fmt.Println(*ptr) // dereferencing a nil pointer causes a runtime panic
	modifyValue(&a)

	fmt.Println(a)

}
func modifyValue(value *int) {
	*value = 20
}
