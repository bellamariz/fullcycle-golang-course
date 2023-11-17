package loop_conditional

import (
	"errors"
	"fmt"
)

func myError(isError bool) error {
	if isError {
		return errors.New("my error")
	}

	return nil
}

func Loops() {
	// default loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// loop for iterating over data (arrays, slices, maps)
	albums := []string{"speak now", "reputation", "lover", "midnights", "folklore"}
	for i, v := range albums {
		fmt.Printf("index: %v, album: %v\n", i, v)
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

func Conditionals() {
	// default conditional: if and else
	iphone := 9000.00
	samsung := 5000.00

	// and: a && b
	// or: a || b
	if iphone > samsung {
		fmt.Println("iPhones are more expensive...")
	} else if iphone == samsung {
		fmt.Println("They are both expensive...")
	} else {
		fmt.Println("Samsung phones are more expensive...")
	}

	// the closest to in-line if Golang has, but it is not very intuitive to read
	if err := myError(true); err != nil {
		fmt.Println("An error has occured!")
	}

	// switch case
	hawkingsCharacter := "Steve Harrington"
	switch hawkingsCharacter {
	case "Vecna":
		fmt.Println("You are dead!")
	case "Eleven":
		fmt.Println("You are saved!")
	case "Demagorgon":
		fmt.Println("You are also dead...")
	case "Nancy Wheeler":
		fmt.Println("You are definitely saved...")
	default:
		fmt.Println("Just run...")
	}
}
