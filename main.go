package main

import (
	"fmt"
	"fullcycle-golang-course/main/foundation"
)

func callFoundation() {
	foundation.Types()
	foundation.Arrays()
	foundation.Slices()
	foundation.Maps()

	sum, err := foundation.Sum(20, 30)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Sum: %d\n", sum)

	total := foundation.SumMany(0, 1, 1, 2, 3, 5, 8)
	fmt.Printf("Total: %d\n", total)

	newTotal := func() int {
		return foundation.SumMany(1, 2, 3, 4, 5)
	}
	fmt.Printf("Total: %d\n", newTotal())

	foundation.Structs()
	foundation.Pointers()
	foundation.TypeAssertion()
	foundation.Generics()
}

func main() {
	// callFoundation()

	foundation.Generics()
}
