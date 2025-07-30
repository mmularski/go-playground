# Channels - Tasks

## Task 1: Basic Channel Operations
Create and use basic channels.

**Requirements:**
- Create an unbuffered channel
- Send data to the channel from a goroutine
- Receive data from the channel in main
- Demonstrate basic send/receive operations

## Task 2: Buffered Channels
Work with buffered channels.

**Requirements:**
- Create a buffered channel with capacity 3
- Send multiple values to the channel
- Receive all values from the channel
- Demonstrate the difference from unbuffered channels

## Task 3: Channel Direction
Use directional channels in functions.

**Requirements:**
- Create a function that accepts a send-only channel
- Create a function that accepts a receive-only channel
- Pass channels to these functions
- Demonstrate channel direction safety

## Task 4: Channel Closing
Work with channel closing and range loops.

**Requirements:**
- Create a goroutine that sends data and closes the channel
- Use a range loop to receive all data
- Handle channel closing properly
- Demonstrate the comma ok idiom

## Task 5: Select Statement
Use select to handle multiple channels.

**Requirements:**
- Create multiple channels
- Use select to handle different channel operations
- Include a default case
- Include a timeout case using time.After()

## Task 6: Practical Application
Create a simple pipeline using channels.

**Requirements:**
- Create a pipeline with multiple stages
- Use channels to pass data between stages
- Implement a simple data processing pipeline
- Demonstrate channel-based communication

---

## How to complete these tasks:

1. Navigate to the `student/` directory
2. Create your code in `main.go`
3. Run your code: `go run main.go`

## Example structure for student/main.go:
```go
package main

import (
    "fmt"
    "time"
)

// Task 1: Basic channel operations
func basicChannel() {
    ch := make(chan int)

    go func() {
        ch <- 42
    }()

    value := <-ch
    fmt.Printf("Received: %d\n", value)
}

// Task 2: Buffered channels
func bufferedChannel() {
    ch := make(chan int, 3)

    ch <- 1
    ch <- 2
    ch <- 3

    fmt.Printf("Received: %d\n", <-ch)
    fmt.Printf("Received: %d\n", <-ch)
    fmt.Printf("Received: %d\n", <-ch)
}

// Task 3: Channel direction
func sendOnly(ch chan<- int) {
    for i := 1; i <= 5; i++ {
        ch <- i
    }
    close(ch)
}

func receiveOnly(ch <-chan int) {
    for value := range ch {
        fmt.Printf("Received: %d\n", value)
    }
}

// Task 4: Channel closing
func channelClosing() {
    ch := make(chan int)

    go func() {
        for i := 1; i <= 5; i++ {
            ch <- i
            time.Sleep(time.Millisecond * 100)
        }
        close(ch)
    }()

    for value := range ch {
        fmt.Printf("Received: %d\n", value)
    }
}

// Task 5: Select statement
func selectExample() {
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
        case <-time.After(time.Second * 3):
            fmt.Println("Timeout")
        }
    }
}

// Task 6: Practical application
func pipeline() {
    // Stage 1: Generate numbers
    numbers := make(chan int)
    go func() {
        for i := 1; i <= 10; i++ {
            numbers <- i
        }
        close(numbers)
    }()

    // Stage 2: Square numbers
    squares := make(chan int)
    go func() {
        for n := range numbers {
            squares <- n * n
        }
        close(squares)
    }()

    // Stage 3: Print results
    for square := range squares {
        fmt.Printf("Square: %d\n", square)
    }
}

func main() {
    fmt.Println("=== Task 1: Basic Channel Operations ===")
    basicChannel()

    fmt.Println("\n=== Task 2: Buffered Channels ===")
    bufferedChannel()

    fmt.Println("\n=== Task 3: Channel Direction ===")
    ch := make(chan int)
    go sendOnly(ch)
    receiveOnly(ch)

    fmt.Println("\n=== Task 4: Channel Closing ===")
    channelClosing()

    fmt.Println("\n=== Task 5: Select Statement ===")
    selectExample()

    fmt.Println("\n=== Task 6: Pipeline ===")
    pipeline()
}
```

## Testing Your Code

Run your program:
```bash
cd student/
go run main.go
```

You should see output demonstrating all channel concepts.