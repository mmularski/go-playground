# Channels - Overview

## What are Channels?

Channels are Go's primary mechanism for communication between goroutines. They provide a way for goroutines to send and receive data safely.

## Basic Channel Operations

```go
// Create a channel
ch := make(chan int)

// Send data to channel
ch <- 42

// Receive data from channel
value := <-ch
```

## Unbuffered Channels

```go
ch := make(chan int) // Unbuffered channel

go func() {
    ch <- 42 // Send
}()

value := <-ch // Receive
fmt.Println(value) // 42
```

## Buffered Channels

```go
ch := make(chan int, 3) // Buffered channel with capacity 3

ch <- 1
ch <- 2
ch <- 3

fmt.Println(<-ch) // 1
fmt.Println(<-ch) // 2
fmt.Println(<-ch) // 3
```

## Channel Direction

```go
// Send-only channel
func sendOnly(ch chan<- int) {
    ch <- 42
}

// Receive-only channel
func receiveOnly(ch <-chan int) {
    value := <-ch
    fmt.Println(value)
}
```

## Channel Closing

```go
ch := make(chan int)

go func() {
    for i := 1; i <= 5; i++ {
        ch <- i
    }
    close(ch) // Close the channel
}()

for value := range ch {
    fmt.Println(value)
}
```

## Select Statement

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

select {
case msg1 := <-ch1:
    fmt.Println("Received:", msg1)
case msg2 := <-ch2:
    fmt.Println("Received:", msg2)
case <-time.After(time.Second * 3):
    fmt.Println("Timeout")
}
```

## Best Practices

1. **Close from sender** - Only the sender should close a channel
2. **Check for closed channels** - Use the comma ok idiom
3. **Use buffered channels wisely** - Don't buffer unnecessarily
4. **Handle timeouts** - Use select with time.After()

## Next Steps

Now that you understand channels, try the tasks in `tasks.md`!