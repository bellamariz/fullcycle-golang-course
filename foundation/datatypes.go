package foundation

import "fmt"

// auxiliary functions

func showType(t interface{}) {
	fmt.Printf("Variable type is %T and value is %v\n", t, t)
}

// test programs

func TypeAssertion() {
	// type of empty interface accepts any data type
	// since golang needs to know the data type, using type assertion guarantees the integrity of the value

	var num interface{} = 10
	var text interface{} = "Dear reader,"

	showType(num)
	showType(text)

	// golang does not know the data type
	println(text)
	// output: (0x10962c0,0x10c6340)

	// therefore we use type assertion
	println(text.(string))
	// output: Dear reader,

	res, ok := num.(int)
	fmt.Printf("Result: %v - Ok: %v\n", res, ok)
}
