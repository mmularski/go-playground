# JSON - Overview

## What is JSON?

JSON (JavaScript Object Notation) is a lightweight data interchange format. Go provides excellent support for JSON through the `encoding/json` package.

## Basic JSON Marshaling

```go
import "encoding/json"

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

person := Person{Name: "John", Age: 30}
data, err := json.Marshal(person)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(data)) // {"name":"John","age":30}
```

## JSON Unmarshaling

```go
jsonData := `{"name":"Jane","age":25}`
var person Person
err := json.Unmarshal([]byte(jsonData), &person)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
```

## JSON Tags

```go
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password string `json:"-"` // Exclude from JSON
    Email    string `json:"email,omitempty"` // Omit if empty
}
```

## Nested Structures

```go
type Address struct {
    Street  string `json:"street"`
    City    string `json:"city"`
    Country string `json:"country"`
}

type Employee struct {
    Name    string  `json:"name"`
    Age     int     `json:"age"`
    Address Address `json:"address"`
}
```

## Arrays and Slices

```go
type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

products := []Product{
    {ID: 1, Name: "Laptop", Price: 999.99},
    {ID: 2, Name: "Mouse", Price: 29.99},
}

data, _ := json.Marshal(products)
fmt.Println(string(data))
```

## Maps

```go
config := map[string]interface{}{
    "database": map[string]string{
        "host": "localhost",
        "port": "5432",
    },
    "api": map[string]interface{}{
        "timeout": 30,
        "retries": 3,
    },
}

data, _ := json.MarshalIndent(config, "", "  ")
fmt.Println(string(data))
```

## Custom Marshaling

```go
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
```

## Streaming JSON

```go
// Encoder for streaming
encoder := json.NewEncoder(os.Stdout)
encoder.SetIndent("", "  ")

for _, product := range products {
    encoder.Encode(product)
}

// Decoder for streaming
decoder := json.NewDecoder(strings.NewReader(jsonData))
for decoder.More() {
    var product Product
    if err := decoder.Decode(&product); err != nil {
        break
    }
    fmt.Printf("Product: %+v\n", product)
}
```

## Best Practices

1. **Use struct tags** for field naming and control
2. **Handle errors** from Marshal/Unmarshal
3. **Use omitempty** for optional fields
4. **Use pointers** for nullable fields
5. **Validate JSON** before processing

## Next Steps

Now that you understand JSON, try the tasks in `tasks.md`!