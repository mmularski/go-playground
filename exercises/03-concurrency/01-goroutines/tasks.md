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


## Testing Your Code

Run your program:
```bash
cd student/
go run main.go
```

You should see output demonstrating all goroutine concepts.