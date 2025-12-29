package basics

import (
	"fmt"
	"slices"
)

func slicesEx() {
	//var sliceName []type
	// var numbers []int
	// var numbers1 = []int{1, 2, 3}
	// numbers2 := []int{9, 8, 7}

	// slice := make([]int, 5)
	a := [5]int{1, 2, 3, 4, 5}
	slice2 := a[1:4]

	fmt.Println(slice2)

	slice2 = append(slice2, 5, 6, 7)

	fmt.Println(slice2)

	sliceCopy := make([]int, len(slice2))
	copy(sliceCopy, slice2)

	fmt.Println(sliceCopy)

	var nilSlice []int
	fmt.Println(nilSlice)

	for i, v := range slice2 {
		fmt.Println(i, v)
	}

	// fmt.Println(slice2[3])

	// slice2[3] = 100
	// fmt.Println(slice2[3])
	// not allowed
	// fmt.Println(slice2 == sliceCopy)

	if slices.Equal(slice2, sliceCopy) {
		fmt.Println("Slices are equal")
	} else {
		fmt.Println("Slices are not equal")
	}

	twoD := make([][]int, 3)

	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
	twoD[0] = append(twoD[0], 100)
	fmt.Println(twoD)

	// slice[low:high]
	slice3 := slice2[2:4]
	fmt.Println(slice3)
	fmt.Println(len(slice3), cap(slice3))
}
