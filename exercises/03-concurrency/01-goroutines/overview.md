# Goroutines - Overview

## What are Goroutines?

Goroutines are lightweight threads managed by the Go runtime. They allow you to run functions concurrently without the overhead of operating system threads.

## Basic Goroutine

### Starting a Goroutine
```go
func sayHello(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

// Start a goroutine
go sayHello("Alice")
```

### Multiple Goroutines
```go
go sayHello("Alice")
go sayHello("Bob")
go sayHello("Charlie")

// Note: Main function might exit before goroutines complete
time.Sleep(1 * time.Second) // Wait for goroutines
```

## Synchronization with WaitGroup

Use `sync.WaitGroup` to wait for goroutines to complete:

```go
func processItem(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Signal completion when function exits

    fmt.Printf("Processing item %d...\n", id)
    time.Sleep(500 * time.Millisecond) // Simulate work
    fmt.Printf("Item %d completed\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go processItem(i, &wg)
    }

    wg.Wait() // Wait for all goroutines to complete
    fmt.Println("All items processed!")
}
```

## Shared Data Safety

### Unsafe Concurrent Access
```go
func unsafeCounter(wg *sync.WaitGroup, counter *int) {
    defer wg.Done()

    for i := 0; i < 1000; i++ {
        *counter++ // Race condition!
    }
}

// This can lead to incorrect results due to race conditions
```

### Safe Concurrent Access with Mutex
```go
func safeCounter(wg *sync.WaitGroup, counter *int, mu *sync.Mutex) {
    defer wg.Done()

    for i := 0; i < 1000; i++ {
        mu.Lock()
        *counter++
        mu.Unlock()
    }
}

func main() {
    var counter int
    var mu sync.Mutex
    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go safeCounter(&wg, &counter, &mu)
    }

    wg.Wait()
    fmt.Printf("Counter: %d\n", counter)
}
```

## Goroutines with Channels

### Basic Channel Communication
```go
func sendNumbers(ch chan int, count int) {
    for i := 1; i <= count; i++ {
        ch <- i // Send to channel
        time.Sleep(100 * time.Millisecond)
    }
    close(ch) // Close channel when done
}

func receiveNumbers(ch chan int, wg *sync.WaitGroup) {
    defer wg.Done()

    for num := range ch { // Receive until channel closed
        fmt.Printf("Received: %d\n", num)
    }
}

func main() {
    ch := make(chan int)
    var wg sync.WaitGroup

    wg.Add(1)
    go sendNumbers(ch, 5)
    go receiveNumbers(ch, &wg)

    wg.Wait()
}
```

### Buffered Channels
```go
func producer(ch chan int, count int) {
    for i := 1; i <= count; i++ {
        ch <- i
        fmt.Printf("Produced: %d\n", i)
        time.Sleep(200 * time.Millisecond)
    }
    close(ch)
}

func consumer(ch chan int, wg *sync.WaitGroup) {
    defer wg.Done()

    for num := range ch {
        fmt.Printf("Consumed: %d\n", num)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    bufferedCh := make(chan int, 3) // Buffer of 3
    var wg sync.WaitGroup

    wg.Add(1)
    go producer(bufferedCh, 5)
    go consumer(bufferedCh, &wg)

    wg.Wait()
}
```

## Worker Pool Pattern

```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(500 * time.Millisecond) // Simulate work
        results <- job * 2 // Send result
    }
}

func main() {
    numJobs := 10
    numWorkers := 3

    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    // Start workers
    for i := 1; i <= numWorkers; i++ {
        go worker(i, jobs, results)
    }

    // Send jobs
    for i := 1; i <= numJobs; i++ {
        jobs <- i
    }
    close(jobs)

    // Collect results
    for i := 1; i <= numJobs; i++ {
        result := <-results
        fmt.Printf("Result: %d\n", result)
    }
}
```

## Timeout with Select

```go
func longRunningTask() string {
    time.Sleep(3 * time.Second)
    return "Task completed"
}

func taskWithTimeout(timeout time.Duration) {
    done := make(chan string, 1)

    go func() {
        result := longRunningTask()
        done <- result
    }()

    select {
    case result := <-done:
        fmt.Printf("Task completed: %s\n", result)
    case <-time.After(timeout):
        fmt.Printf("Task timed out after %v\n", timeout)
    }
}
```

## Context for Cancellation

```go
func contextWorker(ctx context.Context, name string) {
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("%s: Context cancelled\n", name)
            return
        default:
            fmt.Printf("%s: Working...\n", name)
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    go contextWorker(ctx, "Worker1")
    go contextWorker(ctx, "Worker2")

    time.Sleep(3 * time.Second)
}
```

## Best Practices

1. **Always use WaitGroup** to wait for goroutines to complete
2. **Use mutexes** when multiple goroutines access shared data
3. **Use channels** for communication between goroutines
4. **Use context** for cancellation and timeouts
5. **Avoid goroutine leaks** by ensuring goroutines always exit
6. **Use buffered channels** when you know the number of messages

## Example Program

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(500 * time.Millisecond)
        results <- job * 2
    }
}

func main() {
    jobs := make(chan int, 5)
    results := make(chan int, 5)

    // Start workers
    for i := 1; i <= 3; i++ {
        go worker(i, jobs, results)
    }

    // Send jobs
    for i := 1; i <= 5; i++ {
        jobs <- i
    }
    close(jobs)

    // Collect results
    for i := 1; i <= 5; i++ {
        result := <-results
        fmt.Printf("Result: %d\n", result)
    }
}
```

## Next Steps

Now that you understand goroutines, try the tasks in `tasks.md`!
