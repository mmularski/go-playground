package main

import "fmt"

// Task 1: Basic embedding
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	Salary     float64
	Department string
}

// Task 2: Method promotion
func (p Person) Introduce() string {
	// TODO: Return introduction string
	return ""
}

// Task 3: Method overriding
func (e Employee) Introduce() string {
	// TODO: Return employee-specific introduction
	return ""
}

// Task 4: Multiple embedding
type Address struct {
	Street  string
	City    string
	Country string
}

type Contact struct {
	Email string
	Phone string
}

type Customer struct {
	Person
	Address
	Contact
}

// Task 5: Interface embedding
type Reader interface {
	Read() string
}

type Writer interface {
	Write(string)
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct {
	content string
}

func (f *File) Read() string {
	// TODO: Return file content
	return ""
}

func (f *File) Write(data string) {
	// TODO: Write data to file
}

// Task 6: Practical application
type Vehicle struct {
	Brand string
	Model string
	Year  int
}

func (v Vehicle) GetInfo() string {
	// TODO: Return vehicle info
	return ""
}

type Car struct {
	Vehicle
	Doors int
}

func (c Car) GetInfo() string {
	// TODO: Return car info with doors
	return ""
}

type Motorcycle struct {
	Vehicle
	EngineSize int
}

func (m Motorcycle) GetInfo() string {
	// TODO: Return motorcycle info with engine size
	return ""
}

func main() {
	// TODO: Demonstrate your embedding here
}