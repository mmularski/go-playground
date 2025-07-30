package main

// Task 1: Basic interfaces
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	// TODO: Calculate circle area
	return 0.0
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	// TODO: Calculate rectangle area
	return 0.0
}

func PrintArea(s Shape) {
	// TODO: Print area of any shape
}

// Task 2: Interface composition
type Readable interface {
	Read() string
}

type Writable interface {
	Write(string)
}

type ReadWritable interface {
	Readable
	Writable
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

// Task 3: Empty interface
func PrintAnything(v interface{}) {
	// TODO: Handle different types with type switch
}

// Task 4: Multiple interfaces
type Stringer interface {
	String() string
}

type Comparable interface {
	Compare(other Comparable) int
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	// TODO: Return formatted person string
	return ""
}

func (p Person) Compare(other Comparable) int {
	// TODO: Compare persons by age
	return 0
}

// Task 5: Interface best practices
type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (cl ConsoleLogger) Log(message string) {
	// TODO: Log message to console
}

// Task 6: Practical application
type Drawable interface {
	Draw()
}

type Canvas struct {
	shapes []Drawable
}

func (c *Canvas) AddShape(shape Drawable) {
	// TODO: Add shape to canvas
}

func (c Canvas) DrawAll() {
	// TODO: Draw all shapes in canvas
}

type Square struct{}

func (s Square) Draw() {
	// TODO: Draw a square
}

func main() {
	// TODO: Demonstrate your interfaces here
}
