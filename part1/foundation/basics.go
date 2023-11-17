package foundation

import (
	"errors"
	"fmt"
)

// defining constants
const (
	batman = "Bruce Wayne"
	robin  = "Richard Grayson"
	bats   = batman + " and " + robin
)

// create new type called ID (of type int)
type ID int

// defining global variables
var city string = "Gotham City"
var person ID = 1234567890

// variables and functions that start with a capital letter are visible to other packages for importing
// those with non-capital letters are local to the current package

func Types() {
	fmt.Printf("Value: %v - Type: %T\n", person, person)
	fmt.Printf("Value: %v - Type: %T\n", city, city)
}

// array: unmutable size, ordered
func Arrays() {
	// declaring array
	var myArray [3]int

	// adding elements by indexing
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30

	// we iterate over arrays, slices and maps using a for loop
	for i, v := range myArray {
		fmt.Printf("Index: %d - Value: %d\n", i, v)
	}
}

// slices: mutable size, ordered
func Slices() {
	// slice could be initialized statically
	// mySlice := []int{10, 20, 30, 40, 50}

	// or we can use the make command
	mySlice := make([]int, 0, 6) // length: 0, capacity: 6

	// append command is used to add elements to the end of the slice
	mySlice = append(mySlice, 10, 20, 30, 40, 50)

	fmt.Printf("Length: %d - Capacity: %d - Slice: %v\n", len(mySlice), cap(mySlice), mySlice)

	// slices are accessed by indexing and can be re-sliced into smaller slices
	fmt.Printf("Length: %d - Capacity: %d - Slice: %v \n", len(mySlice[:2]), cap(mySlice), mySlice[:2])
	fmt.Printf("Length: %d - Capacity: %d - Slice: %v \n", len(mySlice[2:]), cap(mySlice), mySlice[2:])

	// when we hit max capacity, to append new values, Golang doubles the original size (6 -> 12)
	mySlice = append(mySlice, 60, 70)
	fmt.Printf("Length: %d - Capacity: %d - Slice: %v\n", len(mySlice), cap(mySlice), mySlice)
}

// maps: hash tables, unordered
func Maps() {
	// map could be initialized statically
	myMap := map[ID]string{111: "Freddie", 222: "Roger", 333: "Deaky", 444: "Brian"}

	// or we can use the make command
	// myMap := make(map[ID]string)

	// insert or update an element in map
	myMap[0] = "Mary"

	// retrieve an element from map
	mary := myMap[0]
	fmt.Println(mary)

	// delete an element from map
	delete(myMap, 0)

	for id, member := range myMap {
		fmt.Printf("Band member %s has ID %d\n", member, id)
	}
}

// functions: generic
func Sum(a, b int) (int, error) {
	if a+b >= 50 {
		return a + b, errors.New("sum bigger than 50")
	}

	return a + b, nil
}

// function: multiple unknown parameters
func SumMany(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

// function: anonymous functions/function-within-function
func FuncWithinFunc() func() int {
	newTotal := func() int {
		return SumMany(1, 2, 3, 4, 5)
	}

	return newTotal
}

func Functions() {
	sum, err := Sum(20, 30)

	// checking if function returned an error is default for Golang
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Sum: %d\n", sum)

	total := SumMany(0, 1, 1, 2, 3, 5, 8)
	fmt.Printf("Total: %d\n", total)

	newTotal := FuncWithinFunc()
	fmt.Printf("Total: %d\n", newTotal())
}
