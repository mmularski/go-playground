# Context Package - Tasks

## Task 1: Basic Context
Create and use basic context.

**Requirements:**
- Create a background context
- Create a context with timeout
- Use context.Done() to check for cancellation
- Demonstrate basic context usage

## Task 2: Context with Timeout
Use context with timeout functionality.

**Requirements:**
- Create a context with timeout
- Launch a goroutine that respects the timeout
- Handle timeout cancellation
- Demonstrate timeout behavior

## Task 3: Context with Cancellation
Use context with manual cancellation.

**Requirements:**
- Create a context with cancellation
- Launch a goroutine that can be cancelled
- Cancel the context after a delay
- Handle cancellation properly

## Task 4: Context with Values
Use context to carry values.

**Requirements:**
- Create a context with values
- Pass context to functions
- Retrieve values from context
- Demonstrate context value propagation

## Task 5: Context in HTTP
Use context in HTTP request handling.

**Requirements:**
- Create a simple HTTP server
- Use request context for timeout
- Handle request cancellation
- Demonstrate context in HTTP

## Task 6: Practical Application
Create a simple API client with context.

**Requirements:**
- Create a function that simulates API calls
- Use context for timeout and cancellation
- Handle different context scenarios
- Demonstrate practical context usage

---

## How to complete these tasks:

1. Navigate to the `student/` directory
2. Create your code in `main.go`
3. Run your code: `go run main.go`

## Example structure for student/main.go:
```go
package main

import (
    "context"
    "fmt"
    "net/http"
    "time"
)

// Task 1: Basic context
func basicContext() {
    ctx := context.Background()

    select {
    case <-ctx.Done():
        fmt.Println("Context cancelled")
    default:
        fmt.Println("Context is active")
    }
}

// Task 2: Context with timeout
func contextWithTimeout() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    go func() {
        time.Sleep(3 * time.Second)
        fmt.Println("Long operation completed")
    }()

    select {
    case <-time.After(5 * time.Second):
        fmt.Println("Operation would complete")
    case <-ctx.Done():
        fmt.Println("Operation cancelled due to timeout")
    }
}

// Task 3: Context with cancellation
func contextWithCancellation() {
    ctx, cancel := context.WithCancel(context.Background())

    go func() {
        time.Sleep(2 * time.Second)
        cancel()
    }()

    select {
    case <-time.After(5 * time.Second):
        fmt.Println("Operation completed")
    case <-ctx.Done():
        fmt.Println("Operation cancelled")
    }
}

// Task 4: Context with values
func contextWithValues() {
    ctx := context.WithValue(context.Background(), "userID", "12345")
    ctx = context.WithValue(ctx, "requestID", "req-001")

    userID := ctx.Value("userID").(string)
    requestID := ctx.Value("requestID").(string)

    fmt.Printf("User ID: %s, Request ID: %s\n", userID, requestID)
}

// Task 5: Context in HTTP
func handleRequest(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    select {
    case <-time.After(5 * time.Second):
        fmt.Fprintf(w, "Request completed")
    case <-ctx.Done():
        fmt.Fprintf(w, "Request cancelled")
    }
}

func startHTTPServer() {
    http.HandleFunc("/", handleRequest)
    go http.ListenAndServe(":8080", nil)
    fmt.Println("HTTP server started on :8080")
}

// Task 6: Practical application
func apiCall(ctx context.Context, endpoint string) (string, error) {
    select {
    case <-time.After(time.Second):
        return fmt.Sprintf("Response from %s", endpoint), nil
    case <-ctx.Done():
        return "", ctx.Err()
    }
}

func practicalContext() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    endpoints := []string{"users", "posts", "comments"}

    for _, endpoint := range endpoints {
        result, err := apiCall(ctx, endpoint)
        if err != nil {
            fmt.Printf("Error calling %s: %v\n", endpoint, err)
        } else {
            fmt.Printf("Success: %s\n", result)
        }
    }
}

func main() {
    fmt.Println("=== Task 1: Basic Context ===")
    basicContext()

    fmt.Println("\n=== Task 2: Context with Timeout ===")
    contextWithTimeout()

    fmt.Println("\n=== Task 3: Context with Cancellation ===")
    contextWithCancellation()

    fmt.Println("\n=== Task 4: Context with Values ===")
    contextWithValues()

    fmt.Println("\n=== Task 5: Context in HTTP ===")
    startHTTPServer()

    fmt.Println("\n=== Task 6: Practical Application ===")
    practicalContext()

    // Keep the program running for HTTP server
    time.Sleep(time.Second * 10)
}
```

## Testing Your Code

Run your program:
```bash
cd student/
go run main.go
```

You should see output demonstrating all context concepts.