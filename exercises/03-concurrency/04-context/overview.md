# Context Package - Overview

## What is Context?

The `context` package provides a way to carry deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.

## Basic Context Usage

```go
import "context"

// Create a background context
ctx := context.Background()

// Create a context with timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
```

## Context with Timeout

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

select {
case <-time.After(3 * time.Second):
    fmt.Println("Operation completed")
case <-ctx.Done():
    fmt.Println("Operation cancelled due to timeout")
}
```

## Context with Cancellation

```go
ctx, cancel := context.WithCancel(context.Background())

go func() {
    time.Sleep(2 * time.Second)
    cancel() // Cancel the context
}()

select {
case <-time.After(5 * time.Second):
    fmt.Println("Operation completed")
case <-ctx.Done():
    fmt.Println("Operation cancelled")
}
```

## Context with Values

```go
ctx := context.WithValue(context.Background(), "userID", "12345")

// Retrieve value
userID := ctx.Value("userID").(string)
fmt.Println("User ID:", userID)
```

## Context in HTTP Requests

```go
func handleRequest(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    select {
    case <-time.After(10 * time.Second):
        fmt.Fprintf(w, "Request completed")
    case <-ctx.Done():
        fmt.Fprintf(w, "Request cancelled")
    }
}
```

## Context with Deadline

```go
deadline := time.Now().Add(5 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()

select {
case <-time.After(10 * time.Second):
    fmt.Println("Long operation completed")
case <-ctx.Done():
    fmt.Println("Deadline exceeded")
}
```

## Best Practices

1. **Always check context.Done()** in long-running operations
2. **Pass context as first parameter** to functions
3. **Use context.WithTimeout** for timeouts
4. **Use context.WithCancel** for manual cancellation
5. **Don't store context in structs** - pass it explicitly

## Common Patterns

```go
func processWithContext(ctx context.Context, data string) error {
    select {
    case <-ctx.Done():
        return ctx.Err()
    case <-time.After(time.Second):
        // Process data
        return nil
    }
}
```

## Next Steps

Now that you understand context, try the tasks in `tasks.md`!