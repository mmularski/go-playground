package main

// Task 1: Basic structs
type Person struct {
	Name  string
	Age   int
	Email string
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Book struct {
	Title  string
	Author string
	Year   int
}

// Task 4: Nested structs
type Address struct {
	City    string
	Country string
}

type PersonWithAddress struct {
	Person  Person
	Address Address
}

func main() {
	// Task 1: Create and use structs
	// TODO: Create and use the Person, Rectangle, and Book structs

	// Task 2: Struct fields and tags
	// TODO: Work with struct fields and JSON tags

	// Task 3: Different initialization methods
	// TODO: Try different ways to initialize structs

	// Task 4: Nested structs
	// TODO: Work with nested structs

	// Task 5: Anonymous structs
	// TODO: Create and use anonymous structs

	// Task 6: Struct methods
	// TODO: Add methods to your structs
}
