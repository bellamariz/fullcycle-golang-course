package modules

import (
	"fmt"

	"github.com/bellamariz/fullcycle-golang-course/foundation/mypackage"
)

// without importing the 'mypackage' package, this module this cannot find the function MySum()
// we create the go.mod file and import this external package accordingly
func Modules() {
	s := mypackage.MySum(10.5, 20.5)
	fmt.Printf("Resultado: %v\n", s)
}
