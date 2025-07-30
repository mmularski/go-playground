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


## Testing Your Code

Run your program:
```bash
cd student/
go run main.go
```

You should see output demonstrating all context concepts.