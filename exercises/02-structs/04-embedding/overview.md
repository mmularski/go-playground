# Struct Embedding - Overview

## What is Embedding?

Embedding allows you to include one struct type within another, promoting composition over inheritance.

## Basic Embedding

```go
type Animal struct {
    Name string
    Age  int
}

type Dog struct {
    Animal  // Embedded struct
    Breed   string
}
```

## Accessing Embedded Fields

```go
dog := Dog{
    Animal: Animal{Name: "Buddy", Age: 3},
    Breed:  "Golden Retriever",
}

// Access embedded fields directly
fmt.Println(dog.Name)  // "Buddy"
fmt.Println(dog.Age)   // 3
fmt.Println(dog.Breed) // "Golden Retriever"
```

## Method Promotion

```go
func (a Animal) Speak() string {
    return fmt.Sprintf("%s says hello", a.Name)
}

// Dog automatically gets the Speak method
dog := Dog{Animal: Animal{Name: "Buddy"}, Breed: "Golden"}
fmt.Println(dog.Speak()) // "Buddy says hello"
```

## Method Overriding

```go
func (d Dog) Speak() string {
    return fmt.Sprintf("%s barks", d.Name)
}

// Dog's Speak method overrides Animal's Speak method
```

## Multiple Embedding

```go
type Walker interface {
    Walk() string
}

type Swimmer interface {
    Swim() string
}

type Duck struct {
    Animal  // Embedded struct
    Walker  // Embedded interface
    Swimmer // Embedded interface
}
```

## Interface Embedding

```go
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
```

## Best Practices

1. **Composition over inheritance** - Use embedding to compose behavior
2. **Keep it simple** - Don't embed too many types
3. **Be explicit** - Use field names when needed to avoid conflicts
4. **Override carefully** - Understand method promotion rules

## Next Steps

Now that you understand embedding, try the tasks in `tasks.md`!