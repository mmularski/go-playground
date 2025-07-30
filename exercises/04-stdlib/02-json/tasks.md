# JSON - Tasks

## Task 1: Basic JSON Marshaling
Create and marshal JSON data.

**Requirements:**
- Create a struct with JSON tags
- Marshal struct to JSON
- Handle marshaling errors
- Print formatted JSON

## Task 2: JSON Unmarshaling
Parse JSON data into structs.

**Requirements:**
- Create JSON string data
- Unmarshal JSON into struct
- Handle unmarshaling errors
- Access unmarshaled data

## Task 3: Nested JSON Structures
Work with complex JSON structures.

**Requirements:**
- Create nested structs
- Marshal/unmarshal nested data
- Handle complex JSON hierarchies
- Demonstrate nested structure access

## Task 4: JSON Arrays and Slices
Work with JSON arrays.

**Requirements:**
- Marshal slice of structs to JSON
- Unmarshal JSON array into slice
- Handle array processing
- Demonstrate slice operations

## Task 5: JSON Maps
Work with JSON objects as maps.

**Requirements:**
- Create map[string]interface{}
- Marshal/unmarshal map data
- Handle dynamic JSON structures
- Demonstrate map operations

## Task 6: Custom JSON Marshaling
Implement custom marshaling.

**Requirements:**
- Create custom types
- Implement MarshalJSON method
- Implement UnmarshalJSON method
- Handle custom JSON formatting

---

## How to complete these tasks:

1. Navigate to the `student/` directory
2. Create your code in `main.go`
3. Run your code: `go run main.go`

## Example structure for student/main.go:
```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// Task 1: Basic JSON marshaling
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
    City string `json:"city,omitempty"`
}

func basicJSONMarshaling() {
    person := Person{Name: "John Doe", Age: 30, City: "New York"}

    data, err := json.Marshal(person)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Marshaled JSON: %s\n", string(data))

    // Pretty print
    prettyData, _ := json.MarshalIndent(person, "", "  ")
    fmt.Printf("Pretty JSON:\n%s\n", string(prettyData))
}

// Task 2: JSON unmarshaling
func jsonUnmarshaling() {
    jsonData := `{"name":"Jane Smith","age":25,"city":"Los Angeles"}`

    var person Person
    err := json.Unmarshal([]byte(jsonData), &person)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Unmarshaled person: %+v\n", person)
}

// Task 3: Nested JSON structures
type Address struct {
    Street  string `json:"street"`
    City    string `json:"city"`
    Country string `json:"country"`
}

type Employee struct {
    ID      int     `json:"id"`
    Name    string  `json:"name"`
    Address Address `json:"address"`
}

func nestedJSONStructures() {
    employee := Employee{
        ID:   1,
        Name: "Alice Johnson",
        Address: Address{
            Street:  "123 Main St",
            City:    "Boston",
            Country: "USA",
        },
    }

    data, _ := json.MarshalIndent(employee, "", "  ")
    fmt.Printf("Nested JSON:\n%s\n", string(data))

    // Unmarshal nested JSON
    jsonData := `{
        "id": 2,
        "name": "Bob Wilson",
        "address": {
            "street": "456 Oak Ave",
            "city": "Chicago",
            "country": "USA"
        }
    }`

    var newEmployee Employee
    json.Unmarshal([]byte(jsonData), &newEmployee)
    fmt.Printf("Unmarshaled employee: %+v\n", newEmployee)
}

// Task 4: JSON arrays and slices
type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

func jsonArraysAndSlices() {
    products := []Product{
        {ID: 1, Name: "Laptop", Price: 999.99},
        {ID: 2, Name: "Mouse", Price: 29.99},
        {ID: 3, Name: "Keyboard", Price: 89.99},
    }

    data, _ := json.MarshalIndent(products, "", "  ")
    fmt.Printf("Products JSON:\n%s\n", string(data))

    // Unmarshal JSON array
    jsonData := `[
        {"id": 4, "name": "Monitor", "price": 299.99},
        {"id": 5, "name": "Headphones", "price": 149.99}
    ]`

    var newProducts []Product
    json.Unmarshal([]byte(jsonData), &newProducts)
    fmt.Printf("Unmarshaled products: %+v\n", newProducts)
}

// Task 5: JSON maps
func jsonMaps() {
    config := map[string]interface{}{
        "database": map[string]string{
            "host": "localhost",
            "port": "5432",
            "name": "mydb",
        },
        "api": map[string]interface{}{
            "timeout": 30,
            "retries": 3,
            "enabled": true,
        },
    }

    data, _ := json.MarshalIndent(config, "", "  ")
    fmt.Printf("Config JSON:\n%s\n", string(data))

    // Unmarshal into map
    jsonData := `{
        "server": {
            "host": "0.0.0.0",
            "port": 8080
        },
        "logging": {
            "level": "info",
            "file": "app.log"
        }
    }`

    var settings map[string]interface{}
    json.Unmarshal([]byte(jsonData), &settings)
    fmt.Printf("Unmarshaled settings: %+v\n", settings)
}

// Task 6: Custom JSON marshaling
type CustomTime time.Time

func (ct CustomTime) MarshalJSON() ([]byte, error) {
    t := time.Time(ct)
    return json.Marshal(t.Format("2006-01-02"))
}

func (ct *CustomTime) UnmarshalJSON(data []byte) error {
    var s string
    if err := json.Unmarshal(data, &s); err != nil {
        return err
    }
    t, err := time.Parse("2006-01-02", s)
    if err != nil {
        return err
    }
    *ct = CustomTime(t)
    return nil
}

type Event struct {
    ID   int        `json:"id"`
    Name string     `json:"name"`
    Date CustomTime `json:"date"`
}

func customJSONMarshaling() {
    event := Event{
        ID:   1,
        Name: "Go Conference",
        Date: CustomTime(time.Now()),
    }

    data, _ := json.MarshalIndent(event, "", "  ")
    fmt.Printf("Event with custom time:\n%s\n", string(data))

    // Unmarshal custom time
    jsonData := `{"id": 2, "name": "Workshop", "date": "2024-01-15"}`
    var newEvent Event
    json.Unmarshal([]byte(jsonData), &newEvent)
    fmt.Printf("Unmarshaled event: %+v\n", newEvent)
}

func main() {
    fmt.Println("=== Task 1: Basic JSON Marshaling ===")
    basicJSONMarshaling()

    fmt.Println("\n=== Task 2: JSON Unmarshaling ===")
    jsonUnmarshaling()

    fmt.Println("\n=== Task 3: Nested JSON Structures ===")
    nestedJSONStructures()

    fmt.Println("\n=== Task 4: JSON Arrays and Slices ===")
    jsonArraysAndSlices()

    fmt.Println("\n=== Task 5: JSON Maps ===")
    jsonMaps()

    fmt.Println("\n=== Task 6: Custom JSON Marshaling ===")
    customJSONMarshaling()
}
```

## Testing Your Code

Run your program:
```bash
cd student/
go run main.go
```

You should see output demonstrating all JSON concepts.