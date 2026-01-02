package intermediate

import "fmt"

func closures() {
	// sum := adder()
	// sum()
	// sum()
	// sum()
	// sum()
	// sum2 := adder()
	// sum2()
	// sum2()
	// sum2()

	subtraction := func() func(int) int {
		countDowm := 99
		return func(x int) int {
			countDowm -= x
			return countDowm
		}
	}()
	fmt.Println(subtraction(1))
	fmt.Println(subtraction(2))
	fmt.Println(subtraction(3))

}

func adder() func() int {
	i := 0
	fmt.Println("Previous Value:", i)
	return func() int {
		i++
		fmt.Println("Current Value:", i)
		return i
	}
}
