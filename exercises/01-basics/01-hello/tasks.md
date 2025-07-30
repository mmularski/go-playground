# Hello World - Tasks

## Task 1: Basic Hello World
Create a program that prints "Hello, World!" to the console.

**Requirements:**
- Use `fmt.Println()` to print the message
- The message should be exactly "Hello, World!"

**File to create:** `student/main.go`

## Task 2: Personalized Greeting
Create a function that takes a name as a parameter and prints a personalized greeting.

**Requirements:**
- Create a function called `greet(name string)`
- The function should print "Hello, [name]!"
- Test it with different names

## Task 3: Multiple Greetings
Create a program that prints multiple greetings in different ways.

**Requirements:**
- Use `fmt.Print()` to print without a newline
- Use `fmt.Printf()` to print with formatting
- Use `fmt.Println()` to print with a newline
- Print at least 3 different greetings

## Task 4: Formatted Output
Create a program that demonstrates different formatting options.

**Requirements:**
- Print a number with `%d`
- Print a string with `%s`
- Print a float with `%.2f`
- Print multiple values in one statement

## Task 5: Your Own Program
Create a simple program of your choice that demonstrates what you've learned.

**Requirements:**
- Use at least 2 different print functions
- Include at least one variable
- Make it interesting and creative!

---

## How to complete these tasks:

1. Navigate to the `student/` directory
2. Create a file called `main.go`
3. Implement all tasks in the `main()` function
4. Test your code by running: `go run main.go`

## Example structure for student/main.go:
```go
package main

import "fmt"

func main() {
    // Task 1: Basic Hello World
    fmt.Println("Hello, World!")

    // Task 2: Personalized Greeting
    greet("Alice")
    greet("Bob")

    // Task 3: Multiple Greetings
    fmt.Print("Hello ")
    fmt.Print("from ")
    fmt.Println("Go!")

    fmt.Printf("Welcome to %s!\n", "Go Programming")

    // Task 4: Formatted Output
    age := 25
    name := "John"
    height := 1.75
    fmt.Printf("Name: %s, Age: %d, Height: %.2fm\n", name, age, height)

    // Task 5: Your own program
    // ... your creative code here
}

func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
```

## Testing Your Code

Run your program:
```bash
cd student/
go run main.go
```

You should see output showing:
- "Hello, World!" message
- Personalized greetings for different names
- Multiple greetings using different print functions
- Formatted output with various data types
- Your creative program
```