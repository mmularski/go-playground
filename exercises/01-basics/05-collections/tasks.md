# Collections - Tasks

## Task 1: Arrays
Work with fixed-size arrays in Go.

**Requirements:**
- Create arrays of different types (int, string, bool)
- Access array elements by index
- Iterate over arrays using for loops
- Demonstrate array length and capacity
- Show array initialization methods

## Task 2: Slices
Work with dynamic slices in Go.

**Requirements:**
- Create slices using different methods
- Append elements to slices
- Slice existing slices (slicing)
- Use built-in functions: len(), cap(), append()
- Demonstrate slice capacity and length

## Task 3: Maps
Work with key-value pairs using maps.

**Requirements:**
- Create maps with different key and value types
- Add, update, and delete map entries
- Check if a key exists in a map
- Iterate over maps using range
- Demonstrate map initialization

## Task 4: Collection Operations
Perform common operations on collections.

**Requirements:**
- Find the maximum and minimum values in slices
- Count occurrences of elements
- Filter elements based on conditions
- Sort slices (both ascending and descending)
- Remove duplicates from slices

## Task 5: Nested Collections
Work with collections containing other collections.

**Requirements:**
- Create slices of slices (2D slices)
- Create maps with slice values
- Create slices with map values
- Iterate over nested collections
- Perform operations on nested structures

## Task 6: Practical Application
Create a simple inventory management system.

**Requirements:**
- Store items with properties (name, price, quantity)
- Add and remove items from inventory
- Search for items by name
- Calculate total inventory value
- Display inventory statistics

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
    // Task 1: Arrays
    numbers := [5]int{1, 2, 3, 4, 5}
    fmt.Printf("Array: %v, Length: %d\n", numbers, len(numbers))

    // Task 2: Slices
    fruits := []string{"apple", "banana"}
    fruits = append(fruits, "orange")
    fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", fruits, len(fruits), cap(fruits))

    // Task 3: Maps
    scores := map[string]int{"Alice": 95, "Bob": 87, "Charlie": 92}
    for name, score := range scores {
        fmt.Printf("%s: %d\n", name, score)
    }

    // Task 4: Collection Operations
    numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}
    max := findMax(numbers)
    fmt.Printf("Maximum: %d\n", max)

    // Task 5: Nested Collections
    matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
    for i, row := range matrix {
        fmt.Printf("Row %d: %v\n", i, row)
    }

    // Task 6: Practical Application
    inventory := createInventory()
    addItem(inventory, "Laptop", 999.99, 5)
    displayInventory(inventory)
}

func findMax(numbers []int) int {
    if len(numbers) == 0 {
        return 0
    }
    max := numbers[0]
    for _, num := range numbers {
        if num > max {
            max = num
        }
    }
    return max
}

func createInventory() map[string]Item {
    return make(map[string]Item)
}

type Item struct {
    Name     string
    Price    float64
    Quantity int
}

func addItem(inventory map[string]Item, name string, price float64, quantity int) {
    inventory[name] = Item{Name: name, Price: price, Quantity: quantity}
}

func displayInventory(inventory map[string]Item) {
    fmt.Println("Inventory:")
    for name, item := range inventory {
        fmt.Printf("  %s: $%.2f x %d\n", name, item.Price, item.Quantity)
    }
}
```

## Testing Your Code

Run your program:
```bash
cd student/
go run main.go
```

You should see output demonstrating all collection operations.