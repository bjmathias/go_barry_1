package main

import "fmt"

// Global defiitions/declarations needs to be long form
var intglobal int32 = 8899

var intglobal2 int32

// Function for closure example
func addUp() func(int, int) int {
	return func(x int, y int) int {
		return x + y
	}
}

func main() {

	intglobal2 = 7766

	// Shorthand notation only valid within function.
	charVar := "Barry was here"

	intVar := 11223344
	intVar2 := 11223344
	var i32Var int32 = 55667788

	fmt.Printf("%s %s %d %d", charVar, " ", intVar, i32Var)

	if intVar == intVar2 {
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

	// Ranges
	intArray := []int{9999, 8888, 7777, 6666, 5555, 4444, 3333, 2222, 1111}

	for i, id := range intArray {
		fmt.Printf("index=%d id=%d\n", i, id)
	}

	// Index not used
	for _, id := range intArray {
		fmt.Printf("id=%d\n", id)
	}

	// Map iteration (can just use k for keys alone)
	for k, v := range testMap {
		fmt.Printf("key=%s Value=%s\n", k, v)
	}

	// Pointers
	aaa := 77
	ppp := &aaa

	fmt.Printf("aaa=%d type=%T, ppp = %x type=%T", aaa, aaa, ppp, ppp)

	// Closure uses anonymous function above
	addFunc := addUp()

	for _, ii := range intArray {
		fmt.Printf("addup result=%d\n", addFunc(ii, 2))
	}

	// Structures
	type Person struct {
		firstName   string
		lastName    string
		nationality string
		age         int
	}

	personTest := Person{
		firstName:   "Barry",
		lastName:    "Golang",
		nationality: "Martian",
		age:         26}

	personTest.age = 32

	// Simpler form definition
	personTest2 := Person{
		"Harry", "Bolang", "Venetian", 43}

	personTest.age = 32

	fmt.Println(personTest)
	fmt.Println(personTest2)

}
