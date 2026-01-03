package intermediate

import "fmt"

func formatingVerbs() {

	i := 111_505.5
	str := "Hello World!"

	fmt.Printf("%v\n", i)   // print default formating
	fmt.Printf("%#v\n", i)  // print as it will be in Go source code
	fmt.Printf("%T\n", i)   // print type
	fmt.Printf("%v%%\n", i) // print value and %

	fmt.Printf("%v\n", str)  // print default formating
	fmt.Printf("%#v\n", str) // print as it will be in Go source code
	fmt.Printf("%T\n", str)  // print type

	// Integer Formatting
	n := 18
	fmt.Printf("%b\n", n)   // binary
	fmt.Printf("%d\n", n)   // base 10
	fmt.Printf("%+d\n", n)  // base 1- with sign allways
	fmt.Printf("%o\n", n)   // base 8
	fmt.Printf("%O\n", n)   // base 8 with Leading 0o
	fmt.Printf("%x\n", n)   // hex lowercase
	fmt.Printf("%X\n", n)   // hex uppercase
	fmt.Printf("%#x\n", n)  // hex lowercase with leading 0x
	fmt.Printf("%#X\n", n)  // hex uppercase with leading 0x
	fmt.Printf("%4d\n", n)  // pad left 4
	fmt.Printf("%-4d\n", n) // pad right 4
	fmt.Printf("%04d\n", n) // pad with 0 left

	// Strings Formating
	txt := "World"
	fmt.Printf("%s\n", txt)   // default fromating
	fmt.Printf("%q\n", txt)   // print as Go source
	fmt.Printf("%8s\n", txt)  // pad left 8
	fmt.Printf("%-8s\n", txt) // pad right 8
	fmt.Printf("%x\n", txt)   // hex value
	fmt.Printf("% x\n", txt)  // hex value with spaces

	// Boolean Formating
	t := true
	f := false
	fmt.Printf("%v\n", t) // default fromating
	fmt.Printf("%t\n", t) // print true or false text
	fmt.Printf("%t\n", f) // print true or false text

	// Float Formating
	flt := 12.011
	fmt.Printf("%v\n", flt)    // default fromating
	fmt.Printf("%e\n", flt)    // print in e formating
	fmt.Printf("%f\n", flt)    // print as go source
	fmt.Printf("%.2f\n", flt)  // only two decimal
	fmt.Printf("%6.2f\n", flt) // min of total len 6 with. with two decimal
	fmt.Printf("%g\n", flt)    // exponent as needed

}
