package main

import (
	"fmt"
	"fullcycle-golang-course/main/foundation"
)

func separator(s string) {
	fmt.Printf("\n-------------------------------------------------\n")
	fmt.Printf("--------------------%s--------------------\n", s)
	fmt.Printf("-------------------------------------------------\n")
}

func callFoundation() {
	separator("Basics")
	foundation.Types()
	foundation.Arrays()
	foundation.Slices()
	foundation.Maps()
	foundation.Functions()

	separator("Structs")
	foundation.Structs()
	foundation.Pointers()

	separator("Data Types")
	foundation.TypeAssertion()
	foundation.Generics()
}

func main() {
	callFoundation()
}
