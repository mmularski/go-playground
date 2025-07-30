# Type Assertions - Overview

## What are Type Assertions?

Type assertions allow you to access the concrete value underlying an interface.

## Basic Type Assertion

```go
var i interface{} = "hello"

// Type assertion
s, ok := i.(string)
if ok {
    fmt.Println(s) // "hello"
}
```

## Type Assertion with Panic

```go
var i interface{} = "hello"

// This will panic if i is not a string
s := i.(string)
fmt.Println(s) // "hello"
```

## Type Switches

```go
func describe(i interface{}) {
    switch v := i.(type) {
    case string:
        fmt.Printf("String: %s\n", v)
    case int:
        fmt.Printf("Integer: %d\n", v)
    case bool:
        fmt.Printf("Boolean: %t\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

## Interface Type Assertions

```go
type Stringer interface {
    String() string
}

type Person struct {
    Name string
}

func (p Person) String() string {
    return p.Name
}

var i interface{} = Person{"Alice"}

// Check if i implements Stringer
if s, ok := i.(Stringer); ok {
    fmt.Println(s.String()) // "Alice"
}
```

## Multiple Type Assertions

```go
func processValue(v interface{}) {
    switch val := v.(type) {
    case int:
        fmt.Printf("Processing integer: %d\n", val)
    case string:
        fmt.Printf("Processing string: %s\n", val)
    case []int:
        fmt.Printf("Processing slice: %v\n", val)
    case map[string]int:
        fmt.Printf("Processing map: %v\n", val)
    default:
        fmt.Printf("Unknown type: %T\n", val)
    }
}
```

## Best Practices

1. **Use the comma ok idiom** - Always check if assertion succeeded
2. **Use type switches** - For multiple type checks
3. **Handle the default case** - In type switches
4. **Be specific** - Assert to the most specific type needed

## Next Steps

Now that you understand type assertions, try the tasks in `tasks.md`!