# Select Statement - Overview

## What is Select?

The `select` statement lets a goroutine wait on multiple channel operations. It's similar to a `switch` statement but for channels.

## Basic Select Syntax

```go
select {
case <-ch1:
    // Handle ch1
case ch2 <- value:
    // Send to ch2
default:
    // Default case (non-blocking)
}
```

## Non-blocking Select

```go
ch := make(chan int)

select {
case <-ch:
    fmt.Println("Received from ch")
default:
    fmt.Println("No data available")
}
```

## Select with Multiple Channels

```go
ch1 := make(chan string)
ch2 := make(chan string)

go func() {
    time.Sleep(time.Second)
    ch1 <- "from ch1"
}()

go func() {
    time.Sleep(time.Second * 2)
    ch2 <- "from ch2"
}()

for i := 0; i < 2; i++ {
    select {
    case msg1 := <-ch1:
        fmt.Printf("Received: %s\n", msg1)
    case msg2 := <-ch2:
        fmt.Printf("Received: %s\n", msg2)
    }
}
```

## Select with Timeout

```go
ch := make(chan string)

go func() {
    time.Sleep(time.Second * 2)
    ch <- "result"
}()

select {
case result := <-ch:
    fmt.Printf("Received: %s\n", result)
case <-time.After(time.Second):
    fmt.Println("Timeout")
}
```

## Select with Default

```go
ch := make(chan int)

select {
case <-ch:
    fmt.Println("Received")
default:
    fmt.Println("No data, continuing...")
}
```

## Random Selection

When multiple cases are ready, select chooses one randomly:

```go
ch1 := make(chan string)
ch2 := make(chan string)

go func() {
    ch1 <- "from ch1"
}()

go func() {
    ch2 <- "from ch2"
}()

select {
case msg1 := <-ch1:
    fmt.Printf("Received: %s\n", msg1)
case msg2 := <-ch2:
    fmt.Printf("Received: %s\n", msg2)
}
```

## Best Practices

1. **Always include a default case** for non-blocking operations
2. **Use timeouts** to prevent indefinite blocking
3. **Handle all cases** to avoid deadlocks
4. **Use select for coordination** between multiple goroutines

## Next Steps

Now that you understand select, try the tasks in `tasks.md`!