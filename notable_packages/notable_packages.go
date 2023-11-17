// FullCycle: Go Expert - Part 2 (Notable Packages)
package notable_packages

import (
	"fmt"

	"github.com/bellamariz/fullcycle-golang-course/notable_packages/files"
)

func separator(s string) {
	fmt.Printf("\n#################################################")
	fmt.Printf("\n##### %s #####", s)
	fmt.Printf("\n#################################################\n\n")
}

func CallNotablePackages() {
	separator("Manipulating Files")

	files.CreateFile()
	files.ReadFile()
}
