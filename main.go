package main

import (
	"fmt"

	"github.com/bellamariz/fullcycle-golang-course/foundation"
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

	separator("Packages and Modules")
	foundation.Modules()
}

func main() {
	callFoundation()
}
