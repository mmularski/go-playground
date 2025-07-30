# Collections - Overview

## What are Collections?

Collections are data structures that can hold multiple values. Go provides three main collection types: arrays, slices, and maps.

## Arrays

Arrays are fixed-size collections of elements of the same type.

### Array Declaration
```go
var numbers [5]int                    // Array of 5 integers
colors := [3]string{"red", "green", "blue"} // Array literal
scores := [...]int{95, 87, 92}       // Array with ellipsis
```

### Array Operations
```go
// Accessing elements
first := numbers[0]
last := numbers[len(numbers)-1]

// Modifying elements
numbers[0] = 10

// Length
length := len(numbers)
```

## Slices

Slices are dynamic arrays that can grow and shrink.

### Slice Declaration
```go
var slice1 []int                    // nil slice
slice2 := []int{1, 2, 3, 4, 5}     // slice literal
slice3 := make([]int, 5)            // slice with make
slice4 := make([]int, 3, 5)         // slice with length 3, capacity 5
```

### Slice Operations
```go
// Appending
slice = append(slice, 6, 7, 8)

// Slicing
first3 := slice[:3]   // First 3 elements
last3 := slice[7:]    // Last 3 elements
middle := slice[3:7]  // Middle elements

// Length and capacity
length := len(slice)
capacity := cap(slice)
```

### Slice Best Practices
```go
// Copying slices
source := []int{1, 2, 3, 4, 5}
destination := make([]int, len(source))
copy(destination, source)

// Iterating over slices
for index, value := range slice {
    fmt.Printf("Index %d: %d\n", index, value)
}
```

## Maps

Maps are key-value pairs, similar to dictionaries in other languages.

### Map Declaration
```go
var person map[string]string                    // nil map
scores := map[string]int{"Alice": 95, "Bob": 87} // map literal
grades := make(map[string]string)               // empty map
```

### Map Operations
```go
// Adding/updating
grades["Alice"] = "A"
grades["Bob"] = "B"

// Accessing
grade := grades["Alice"]

// Checking if key exists
if grade, exists := grades["David"]; exists {
    fmt.Printf("Grade: %s\n", grade)
} else {
    fmt.Println("Not found")
}

// Deleting
delete(grades, "Bob")

// Iterating
for key, value := range grades {
    fmt.Printf("%s: %s\n", key, value)
}
```

## Common Patterns

### Working with Slices
```go
// Filtering
var evenNumbers []int
for _, num := range numbers {
    if num%2 == 0 {
        evenNumbers = append(evenNumbers, num)
    }
}

// Finding maximum
max := numbers[0]
for _, num := range numbers {
    if num > max {
        max = num
    }
}
```

### Working with Maps
```go
// Counting occurrences
wordCount := make(map[string]int)
words := []string{"hello", "world", "hello", "go"}
for _, word := range words {
    wordCount[word]++
}
```

## Best Practices

1. **Use slices for dynamic collections** instead of arrays
2. **Pre-allocate slices** when you know the size
3. **Use maps for key-value lookups**
4. **Check for key existence** in maps
5. **Use range for iteration** over collections

## Example Program

```go
package main

import "fmt"

func main() {
    // Slice example
    numbers := []int{1, 2, 3, 4, 5}
    numbers = append(numbers, 6)

    for i, num := range numbers {
        fmt.Printf("Index %d: %d\n", i, num)
    }

    // Map example
    scores := map[string]int{
        "Alice": 95,
        "Bob":   87,
        "Charlie": 92,
    }

    for name, score := range scores {
        fmt.Printf("%s: %d\n", name, score)
    }
}
```

## Next Steps

Now that you understand collections, try the tasks in `tasks.md`!
