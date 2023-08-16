package foundation

import "fmt"

func Loops() {
	// default loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// loop for iterating over data (arrays, slices, maps)
	numbers := []string{"one", "two", "three", "four", "five"}
	for i, v := range numbers {
		fmt.Printf("index: %v, value: %v\n", i, v)
	}
	// we can substitute i or v by an underscore _ if we do not desire to use it
	// this is called a blank identifier

	// loop while but it's a for
	i := 0
	for i < 6 {
		i++
	}

	// infinite loop
	// for {
	// 	fmt.Println("Infinite loop... do not run it please")
	// }

}
