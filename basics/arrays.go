package basics

import "fmt"

func arrays() {
	//var arrayName [size]dataType
	var numbers [5]int
	fmt.Println("Numbers:", numbers)

	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5

	fruits := [4]string{"Apple", "Banana", "Orange", "Mango"}
	fmt.Println(fruits)
	fmt.Println("Third fruit:", fruits[2])

	orginalArray := [3]int{1, 2, 3}
	copiedArray := orginalArray
	copiedArray[0] = 100
	fmt.Println("Original array:", orginalArray)
	fmt.Println("Copied array:", copiedArray)

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("numbers[%d] = %d\n", i, numbers[i])
	}

	for i, v := range fruits {
		fmt.Printf("fruits[%d] = %s\n", i, v)
	}

	a, b := someFunction()
	fmt.Println("Values from someFunction:", a, b)

	fmt.Println("Length of numbers array:", len(numbers))

	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}
	fmt.Println("Are array1 and array2 equal?", array1 == array2)

	// 2d array
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("2D Array (Matrix):", matrix)

	orginalArray2 := [3]int{1, 2, 3}
	var copiedArray2 *[3]int = &orginalArray2
	copiedArray2[0] = 100
	fmt.Println("Original array:", orginalArray2)
	fmt.Println("Copied array:", copiedArray2)
	fmt.Println("Copied array:", *copiedArray2)
}

func someFunction() (int, int) {
	return 1, 2
}
