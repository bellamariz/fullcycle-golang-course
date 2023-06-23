package foundation

import "fmt"

// creating new types
type ID int

// global variables
var number ID = 12

func Types() {
	fmt.Printf("Value: %v - Type: %T\n", number, number)
}

// array: unmutable size, ordered
func Arrays() {
	var myArray [3]int

	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30

	for i, v := range myArray {
		fmt.Printf("Index: %d - Value: %d\n", i, v)
	}
}

// slices: mutable size, ordered
func Slices() {
	mySlice := make([]int, 0, 6) // length: 0, capacity: 6

	mySlice = append(mySlice, 10, 20, 30, 40, 50)

	fmt.Printf("Length: %d - Capacity: %d - Slice: %v\n", len(mySlice), cap(mySlice), mySlice)
	fmt.Printf("Length: %d - Capacity: %d - Slice: %v \n", len(mySlice[:2]), cap(mySlice), mySlice[:2])
	fmt.Printf("Length: %d - Capacity: %d - Slice: %v \n", len(mySlice[2:]), cap(mySlice), mySlice[2:])

	mySlice = append(mySlice, 60, 70)

	// when we hit max capacity, to append new values, Golang doubles the original size (6 -> 12)
	fmt.Printf("Length: %d - Capacity: %d - Slice: %v\n", len(mySlice), cap(mySlice), mySlice)
}

// maps: hash tables, unordered
func Maps() {
	// myMap := make(map[ID]string)
	myMap := map[ID]string{111: "Freddie", 222: "Roger", 333: "Deaky", 444: "Brian"}

	myMap[0] = "Mary"
	delete(myMap, 0)

	for id, member := range myMap {
		fmt.Printf("Band member %s has ID %d\n", member, id)
	}
}
