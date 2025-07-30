# Functions - Tasks

## Task 1: Basic Functions
Create and call basic functions.

**Requirements:**
- Create a function that takes no parameters and returns nothing
- Create a function that takes parameters and returns a value
- Create a function that takes multiple parameters
- Call all functions from main()

## Task 2: Multiple Returns
Create functions that return multiple values.

**Requirements:**
- Create a function that returns two values
- Create a function that returns three values
- Use the blank identifier (_) to ignore some return values
- Demonstrate different ways to handle multiple returns

## Task 3: Named Returns
Create functions with named return values.

**Requirements:**
- Create a function with named return values
- Create a function that uses naked return
- Create a function that modifies named return values
- Demonstrate when to use named returns

## Task 4: Variadic Functions
Create functions that accept variable number of arguments.

**Requirements:**
- Create a function that sums variable number of integers
- Create a function that concatenates variable number of strings
- Create a function that finds maximum of variable number of values
- Pass slices to variadic functions using ...

## Task 5: Function Types
Work with function types and function variables.

**Requirements:**
- Create a function type definition
- Assign functions to variables
- Pass functions as parameters
- Return functions from functions

## Task 6: Anonymous Functions and Closures
Create anonymous functions and closures.

**Requirements:**
- Create anonymous functions
- Create closures that capture variables
- Use anonymous functions as arguments
- Demonstrate closure behavior

## Task 7: Recursion
Create recursive functions.

**Requirements:**
- Create a recursive factorial function
- Create a recursive Fibonacci function
- Create a recursive function to calculate sum of array
- Understand base cases and recursive cases

---

## How to complete these tasks:

1. Navigate to the `student/` directory
2. Create your code in `main.go`
3. Run your code: `go run main.go`

## Example structure for student/main.go:
```go
package main

import "fmt"

func main() {
    // Task 1: Basic Functions
    greet()
    result := add(5, 3)
    fmt.Printf("5 + 3 = %d\n", result)

    // Task 2: Multiple Returns
    quotient, remainder := divide(17, 5)
    fmt.Printf("17 / 5 = %d remainder %d\n", quotient, remainder)

    // Task 3: Named Returns
    min, max := findMinMax([]int{3, 1, 4, 1, 5, 9, 2, 6})
    fmt.Printf("Min: %d, Max: %d\n", min, max)

    // Task 4: Variadic Functions
    sum := sumAll(1, 2, 3, 4, 5)
    fmt.Printf("Sum: %d\n", sum)

    // Task 5: Function Types
    numbers := []int{1, 2, 3, 4, 5}
    doubled := apply(numbers, func(x int) int { return x * 2 })
    fmt.Printf("Doubled: %v\n", doubled)

    // Task 6: Anonymous Functions and Closures
    counter := createCounter()
    fmt.Printf("Counter: %d\n", counter())
    fmt.Printf("Counter: %d\n", counter())

    // Task 7: Recursion
    fact := factorial(5)
    fmt.Printf("5! = %d\n", fact)
}

func greet() {
    fmt.Println("Hello from a function!")
}

func add(a, b int) int {
    return a + b
}

func divide(a, b int) (int, int) {
    return a / b, a % b
}

func findMinMax(numbers []int) (min, max int) {
    if len(numbers) == 0 {
        return
    }
    min, max = numbers[0], numbers[0]
    for _, num := range numbers {
        if num < min {
            min = num
        }
        if num > max {
            max = num
        }
    }
    return
}

func sumAll(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

type IntFunction func(int) int

func apply(numbers []int, fn IntFunction) []int {
    result := make([]int, len(numbers))
    for i, num := range numbers {
        result[i] = fn(num)
    }
    return result
}

func createCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}
```

## Testing Your Code

Run your program:
```bash
cd student/
go run main.go
```

You should see output demonstrating all function concepts.