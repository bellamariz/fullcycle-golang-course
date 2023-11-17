// FullCycle: Go Expert - Part 1 (Foundation)
package foundation

import (
	"fmt"

	"github.com/bellamariz/fullcycle-golang-course/foundation/basics"
	"github.com/bellamariz/fullcycle-golang-course/foundation/datatypes"
	"github.com/bellamariz/fullcycle-golang-course/foundation/loop_conditional"
	"github.com/bellamariz/fullcycle-golang-course/foundation/modules"
	"github.com/bellamariz/fullcycle-golang-course/foundation/structs"
)

func separator(s string) {
	fmt.Printf("\n#################################################")
	fmt.Printf("\n##### %s #####", s)
	fmt.Printf("\n#################################################\n\n")
}

func CallFoundation() {
	separator("Basics")
	basics.Types()
	basics.Arrays()
	basics.Slices()
	basics.Maps()
	basics.Functions()

	separator("Structs")
	structs.Structs()
	structs.Pointers()

	separator("Data Types")
	datatypes.TypeAssertion()
	datatypes.Generics()

	separator("Packages and Modules")
	modules.Modules()

	separator("Loops and Conditionals")
	loop_conditional.Loops()
	loop_conditional.Conditionals()
}
