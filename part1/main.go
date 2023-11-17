package part1

import (
	"fmt"

	"github.com/bellamariz/fullcycle-golang-course/part1/foundation"
)

func separator(s string) {
	fmt.Printf("\n-------------------------------------------------")
	fmt.Printf("\n-----%s-----", s)
	fmt.Printf("\n-------------------------------------------------\n")
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

	separator("Loops and Conditionals")
	foundation.Loops()
	foundation.Conditionals()
}

func main() {
	callFoundation()
}
