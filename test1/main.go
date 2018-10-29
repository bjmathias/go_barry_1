package main

import (
	"bytes"
	"fmt"
	"math"
	"net/http"
	"strconv"
)

// Global defiitions/declarations needs to be long form
var intglobal int32 = 8899

var intglobal2 int32

// Function for closure example
func addUp() func(int, int) int {
	return func(x int, y int) int {
		return x + y
	}
}

// Person Structure example
type Person struct {
	firstName   string
	lastName    string
	nationality string
	age         int
}

func (per Person) sayHello() string {

	return "Hello " + per.firstName
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
	personTest2.age++

	fmt.Println(personTest)
	fmt.Println(personTest2)

	fmt.Printf("%s\n", personTest.sayHello())

	// String length
	var stringvar string = "abcdefg"
	fmt.Printf("Length of string = %d\n", len(stringvar))

	// Boolean expression result (should be true due to last element)
	extst := (true && false) || (false && true) || !(false && false)
	fmt.Printf("Expression result is : %t\n", extst)

	// Basic increment (should use x++)
	x := 5
	x += 1
	fmt.Printf("Incremented x = %d\n", x)

	// Basic commmand line input and math package use
	fmt.Print("Enter a number to be squared: ")
	var input float64
	fmt.Scanf("%f", &input)
	inputsq := math.Pow(input, 2)
	fmt.Printf("Input sqaured = %f\n", inputsq)

	i := 5

	// Switch statement, note no break required.
	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}

	// Loop using mudulus operator
	for i = 1; i <= 100; i++ {
		if (i % 3) == 0 {
			fmt.Printf("Number %d is divisible by 3\n", i)
		}
	}

	// FizzBuzz experimenting with a string buffer.
	var buffer bytes.Buffer
	for i = 1; i <= 100; i++ {
		if (i % 3) == 0 {
			buffer.WriteString("Fizz")
		}
		if (i % 5) == 0 {
			buffer.WriteString("Buzz")
		}
		if buffer.Len() == 0 {
			buffer.WriteString(strconv.Itoa(i))
		}
		fmt.Println(buffer.String())
		// Clear buffer for next iteration.
		buffer.Reset()
	}

	// Basic slices & arrays, a string slice with 10 elements
	strArray := make([]string, 10)

	// Underlying array with cpacity of 50 elements
	strArray2 := make([]string, 10, 50)

	strArray[0] = "Hello "
	strArray2[0] = "Barry"

	strArray3 := append(strArray, strArray2[0])

	fmt.Println(strArray3)

	// Prints c,d,e as the range excludes the end index (5)
	zArray := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(zArray[2:5])

	// Basic map creation
	m := make(map[string]int)
	m["key"] = 7878
	fmt.Println(m["key"])

	// Checking for map lookup fail
	delete(m, "key")
	val, ok := m["key"]
	if !ok {
		fmt.Println("Array is empty")
	} else {
		fmt.Println(val)
	}

	// Basic web server test
	//httpListen()
}

// httpTestFunc handler for an http /test request
func httpTestFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Server response")
	// Echo the url
	fmt.Fprintln(w, "Request URL: "+r.URL.RequestURI()+"\n")
	// Echo the url (GET) parameter map
	fmt.Fprintln(w, r.URL.Query())
	// Get and echo specific request parameter
	fmt.Fprintln(w, "Origin="+r.URL.Query().Get("origin"))

}

// httpListen starts a basic http server on designated port
func httpListen() {

	httpPort := 8840
	fmt.Printf("Starting http server on port %d.......\n", httpPort)
	http.HandleFunc("/test", httpTestFunc)
	http.ListenAndServe(":"+strconv.Itoa(httpPort), nil)
	return
}
