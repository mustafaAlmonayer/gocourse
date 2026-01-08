package intermediate

import (
	"fmt"
	"math/rand"
)

func randomNums() {

	// Seed
	val := rand.New(rand.NewSource(40))

	fmt.Println(val.Intn(6) + 5)
}
