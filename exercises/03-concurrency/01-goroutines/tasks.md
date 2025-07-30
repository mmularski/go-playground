# Goroutines - Tasks

## Task 1: Basic Goroutines
Create and run basic goroutines.

**Requirements:**
- Create a simple function that prints messages
- Launch it as a goroutine using `go`
- Demonstrate that goroutines run concurrently
- Show the difference between sequential and concurrent execution

## Task 2: Goroutines with Sleep
Use goroutines with time delays.

**Requirements:**
- Create goroutines that sleep for different durations
- Use `time.Sleep()` to simulate work
- Show how goroutines can run in parallel
- Demonstrate non-blocking execution

## Task 3: WaitGroup
Use sync.WaitGroup to wait for goroutines to complete.

**Requirements:**
- Create multiple goroutines
- Use `Add()`, `Done()`, and `Wait()` methods
- Ensure all goroutines complete before program ends
- Handle WaitGroup properly

## Task 4: Shared Data Safety
Work with shared data safely using mutexes.

**Requirements:**
- Create a shared counter variable
- Launch multiple goroutines that modify the counter
- Use `sync.Mutex` to prevent race conditions
- Demonstrate the difference between safe and unsafe access

## Task 5: Goroutine Communication
Use channels to communicate between goroutines.

**Requirements:**
- Create a goroutine that sends data through a channel
- Create a goroutine that receives data from a channel
- Use buffered and unbuffered channels
- Demonstrate basic producer-consumer pattern

## Task 6: Practical Application
Create a simple web scraper using goroutines.

**Requirements:**
- Create a function that simulates fetching a URL
- Launch multiple goroutines to fetch different URLs
- Use WaitGroup to wait for all fetches to complete
- Collect and display results from all goroutines

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
    "sync"
    "time"
)

// Task 1: Basic goroutines
func printMessage(msg string) {
    fmt.Printf("Message: %s\n", msg)
}

// Task 2: Goroutines with sleep
func delayedPrint(msg string, delay time.Duration) {
    time.Sleep(delay)
    fmt.Printf("Delayed message: %s\n", msg)
}

// Task 3: WaitGroup
func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

// Task 4: Shared data safety
type SafeCounter struct {
    mu    sync.Mutex
    count int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

func (c *SafeCounter) GetCount() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}

// Task 5: Goroutine communication
func producer(ch chan<- int) {
    for i := 1; i <= 5; i++ {
        ch <- i
        time.Sleep(time.Millisecond * 100)
    }
    close(ch)
}

func consumer(ch <-chan int) {
    for num := range ch {
        fmt.Printf("Received: %d\n", num)
    }
}

// Task 6: Practical application
func fetchURL(url string, wg *sync.WaitGroup, results chan<- string) {
    defer wg.Done()

    // Simulate network delay
    time.Sleep(time.Millisecond * 500)

    result := fmt.Sprintf("Fetched: %s", url)
    results <- result
}

func main() {
    // Task 1: Basic goroutines
    go printMessage("Hello from goroutine")
    printMessage("Hello from main")

    // Task 2: Goroutines with sleep
    go delayedPrint("First", time.Millisecond*100)
    go delayedPrint("Second", time.Millisecond*50)
    go delayedPrint("Third", time.Millisecond*200)

    // Task 3: WaitGroup
    var wg sync.WaitGroup
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }
    wg.Wait()

    // Task 4: Shared data safety
    counter := &SafeCounter{}
    for i := 0; i < 10; i++ {
        go counter.Increment()
    }
    time.Sleep(time.Second)
    fmt.Printf("Final count: %d\n", counter.GetCount())

    // Task 5: Goroutine communication
    ch := make(chan int)
    go producer(ch)
    go consumer(ch)

    // Task 6: Practical application
    urls := []string{"http://example1.com", "http://example2.com", "http://example3.com"}
    results := make(chan string, len(urls))

    var fetchWg sync.WaitGroup
    for _, url := range urls {
        fetchWg.Add(1)
        go fetchURL(url, &fetchWg, results)
    }

    go func() {
        fetchWg.Wait()
        close(results)
    }()

    for result := range results {
        fmt.Println(result)
    }
}
```

## Testing Your Code

Run your program:
```bash
cd student/
go run main.go
```

You should see output demonstrating all goroutine concepts.