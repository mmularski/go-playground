# Select Statement - Tasks

## Task 1: Basic Select
Create a basic select statement.

**Requirements:**
- Create a channel
- Use select to check if data is available
- Include a default case
- Demonstrate non-blocking behavior

## Task 2: Select with Multiple Channels
Use select with multiple channels.

**Requirements:**
- Create two channels
- Launch goroutines that send to different channels
- Use select to handle both channels
- Demonstrate random selection when both are ready

## Task 3: Select with Timeout
Use select with timeout functionality.

**Requirements:**
- Create a channel that receives data after a delay
- Use time.After() for timeout
- Handle both successful receive and timeout cases
- Demonstrate timeout behavior

## Task 4: Non-blocking Operations
Use select for non-blocking channel operations.

**Requirements:**
- Create channels that may or may not have data
- Use select with default to avoid blocking
- Handle multiple channel operations
- Demonstrate graceful handling of empty channels

## Task 5: Select in Loop
Use select in a loop for continuous monitoring.

**Requirements:**
- Create multiple channels
- Use select in a loop to monitor all channels
- Handle channel closing
- Demonstrate continuous operation

## Task 6: Practical Application
Create a simple load balancer using select.

**Requirements:**
- Create multiple worker channels
- Use select to distribute work
- Handle worker availability
- Demonstrate load balancing behavior

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

// Task 1: Basic select
func basicSelect() {
    ch := make(chan int)

    select {
    case <-ch:
        fmt.Println("Received from ch")
    default:
        fmt.Println("No data available")
    }
}

// Task 2: Select with multiple channels
func multipleChannels() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(time.Millisecond * 100)
        ch1 <- "from ch1"
    }()

    go func() {
        time.Sleep(time.Millisecond * 200)
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
}

// Task 3: Select with timeout
func selectWithTimeout() {
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
}

// Task 4: Non-blocking operations
func nonBlockingSelect() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    // Try to send to ch1
    select {
    case ch1 <- 42:
        fmt.Println("Sent to ch1")
    default:
        fmt.Println("ch1 is not ready")
    }

    // Try to receive from ch2
    select {
    case value := <-ch2:
        fmt.Printf("Received from ch2: %d\n", value)
    default:
        fmt.Println("ch2 has no data")
    }
}

// Task 5: Select in loop
func selectInLoop() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {
        for i := 1; i <= 3; i++ {
            ch1 <- i
            time.Sleep(time.Millisecond * 200)
        }
        close(ch1)
    }()

    go func() {
        for i := 10; i <= 12; i++ {
            ch2 <- i
            time.Sleep(time.Millisecond * 300)
        }
        close(ch2)
    }()

    ch1Closed := false
    ch2Closed := false

    for !ch1Closed || !ch2Closed {
        select {
        case value, ok := <-ch1:
            if ok {
                fmt.Printf("Received from ch1: %d\n", value)
            } else {
                ch1Closed = true
            }
        case value, ok := <-ch2:
            if ok {
                fmt.Printf("Received from ch2: %d\n", value)
            } else {
                ch2Closed = true
            }
        }
    }
}

// Task 6: Practical application
func loadBalancer() {
    worker1 := make(chan string)
    worker2 := make(chan string)

    // Simulate workers
    go func() {
        for i := 1; i <= 3; i++ {
            time.Sleep(time.Millisecond * 100)
            worker1 <- fmt.Sprintf("Worker1 processed job %d", i)
        }
    }()

    go func() {
        for i := 1; i <= 3; i++ {
            time.Sleep(time.Millisecond * 150)
            worker2 <- fmt.Sprintf("Worker2 processed job %d", i)
        }
    }()

    // Load balancer
    for i := 0; i < 6; i++ {
        select {
        case result := <-worker1:
            fmt.Printf("Load balancer: %s\n", result)
        case result := <-worker2:
            fmt.Printf("Load balancer: %s\n", result)
        }
    }
}

func main() {
    fmt.Println("=== Task 1: Basic Select ===")
    basicSelect()

    fmt.Println("\n=== Task 2: Select with Multiple Channels ===")
    multipleChannels()

    fmt.Println("\n=== Task 3: Select with Timeout ===")
    selectWithTimeout()

    fmt.Println("\n=== Task 4: Non-blocking Operations ===")
    nonBlockingSelect()

    fmt.Println("\n=== Task 5: Select in Loop ===")
    selectInLoop()

    fmt.Println("\n=== Task 6: Load Balancer ===")
    loadBalancer()
}
```

## Testing Your Code

Run your program:
```bash
cd student/
go run main.go
```

You should see output demonstrating all select concepts.