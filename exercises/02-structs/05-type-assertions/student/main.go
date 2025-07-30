package main

import "fmt"

// Task 1: Basic type assertions
func processBasicType(v interface{}) {
	// TODO: Use type assertion to check if v is a string
	// TODO: Use type assertion to check if v is an int
	// TODO: Handle failed assertions gracefully
}

// Task 2: Type switches
func describeType(v interface{}) {
	// TODO: Use type switch to handle different types
	// TODO: Handle string, int, bool, float64, and default cases
}

// Task 3: Interface type assertions
type Stringer interface {
	String() string
}

type Person struct {
	Name string
}

func (p Person) String() string {
	return p.Name
}

func checkStringer(v interface{}) {
	// TODO: Check if v implements Stringer interface
	// TODO: Call String() method if it does
}

// Task 4: Safe type assertions
func safeString(v interface{}) (string, bool) {
	// TODO: Safely assert v to string
	// TODO: Return string value and success status
	return "", false
}

func safeInt(v interface{}) (int, bool) {
	// TODO: Safely assert v to int
	// TODO: Return int value and success status
	return 0, false
}

// Task 5: Complex type assertions
func processSlice(v interface{}) {
	// TODO: Check if v is a slice of integers
	// TODO: Process the slice if it is
}

func processMap(v interface{}) {
	// TODO: Check if v is a map[string]int
	// TODO: Process the map if it is
}

// Task 6: Practical application
type DataProcessor struct{}

func (dp DataProcessor) Process(data interface{}) string {
	// TODO: Use type switch to process different data types
	// TODO: Return appropriate string representation
	return ""
}

func main() {
	// TODO: Demonstrate your type assertions here
}