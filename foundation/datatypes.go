package foundation

import "fmt"

type MyNumber int

// using an interface as a constraint for generics

type number interface {
	~int | float64
}

// auxiliary functions

func showType(t interface{}) {
	fmt.Printf("Variable type is %T and value is %v\n", t, t)
}

func sumMap[T int | float64](m map[string]T) T {
	var sum T

	for _, v := range m {
		sum += v
	}

	return sum
}

func sumMapWithConstraint[T number](m map[string]T) T {
	var sum T

	for _, v := range m {
		sum += v
	}

	return sum
}

func equal[T comparable](a, b T) bool {
	return a == b
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

// generics are usable with functions and methods

func Generics() {
	intMap := map[string]int{"Stranger Things": 1000, "Sabrina": 500, "Riverdale": 200}
	floatMap := map[string]float64{"The Boys": 1000.5, "Good Omens": 500.9, "Jack Ryan": 200.2}
	newTypeMap := map[string]MyNumber{"The Boys": 1000, "Good Omens": 500, "Jack Ryan": 200}

	// using generics
	fmt.Printf("Sum map of integers: %v\n", sumMap(intMap))
	fmt.Printf("Sum map of floats: %v\n", sumMap(floatMap))

	// using constraints (interface + generics)
	fmt.Printf("Sum map of integers with constraint: %v\n", sumMapWithConstraint(intMap))
	fmt.Printf("Sum map of floats with constraint: %v\n", sumMapWithConstraint(floatMap))

	// when using an user-defined data type
	// must add a ~ to the constraint for golang to understand the distinction
	fmt.Printf("Sum map of MyNumber with constraint: %v\n", sumMapWithConstraint(newTypeMap))

	// comparable is a golang constraint that verifies whether generic values can be compared
	fmt.Printf("Are they equal? %v\n", equal(10, 10))
}
