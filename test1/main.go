package main

import "fmt"

// Global defiitions/declarations needs to be long form
var intglobal int32 = 8899

var intglobal2 int32

func main() {

	intglobal2 = 7766

	// Shorthand notation only valid within function.
	charvar := "Barry was here"

	intvar := 11223344
	intvar2 := 11223344
	var i32var int32 = 55667788

	fmt.Printf("%s %s %d %d", charvar, " ", intvar, i32var)

	if intvar == intvar2 {
		fmt.Println("Ints are identical")
	}
	sliceTest := []string{"one", "two", "three"}
	fmt.Println(sliceTest)

	// Short form map
	testMap := map[string]string{"one": "1", "two": "2", "three": "3"}

	fmt.Println(testMap)

	// remove map entry
	delete(testMap, "two")

	fmt.Println(testMap)
}
