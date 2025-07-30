package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Task 1: Basic context
func basicContext() {
	// TODO: Create a background context: ctx := context.Background()
	// TODO: Use context.Done() to check for cancellation:
	// select {
	// case <-ctx.Done(): fmt.Println("Context cancelled")
	// default: fmt.Println("Context is active")
	// }
}

// Task 2: Context with timeout
func contextWithTimeout() {
	// TODO: Create a context with timeout: ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// TODO: Defer cancel: defer cancel()
	// TODO: Launch a goroutine that respects the timeout:
	// go func() { time.Sleep(3 * time.Second); fmt.Println("Long operation completed") }()
	// TODO: Handle timeout cancellation:
	// select {
	// case <-time.After(5 * time.Second): fmt.Println("Operation would complete")
	// case <-ctx.Done(): fmt.Println("Operation cancelled due to timeout")
	// }
}

// Task 3: Context with cancellation
func contextWithCancellation() {
	// TODO: Create a context with cancellation: ctx, cancel := context.WithCancel(context.Background())
	// TODO: Launch a goroutine that can be cancelled:
	// go func() { time.Sleep(2 * time.Second); cancel() }()
	// TODO: Handle cancellation:
	// select {
	// case <-time.After(5 * time.Second): fmt.Println("Operation completed")
	// case <-ctx.Done(): fmt.Println("Operation cancelled")
	// }
}

// Task 4: Context with values
func contextWithValues() {
	// TODO: Create a context with values:
	// ctx := context.WithValue(context.Background(), "userID", "12345")
	// ctx = context.WithValue(ctx, "requestID", "req-001")
	// TODO: Retrieve values from context:
	// userID := ctx.Value("userID").(string)
	// requestID := ctx.Value("requestID").(string)
	// TODO: Print values: fmt.Printf("User ID: %s, Request ID: %s\n", userID, requestID)
}

// Task 5: Context in HTTP
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// TODO: Use request context for timeout: ctx := r.Context()
	// TODO: Handle request cancellation:
	// select {
	// case <-time.After(5 * time.Second): fmt.Fprintf(w, "Request completed")
	// case <-ctx.Done(): fmt.Fprintf(w, "Request cancelled")
	// }
}

func startHTTPServer() {
	// TODO: Create a simple HTTP server:
	// http.HandleFunc("/", handleRequest)
	// go http.ListenAndServe(":8080", nil)
	// fmt.Println("HTTP server started on :8080")
}

// Task 6: Practical application
func apiCall(ctx context.Context, endpoint string) (string, error) {
	// TODO: Create a function that simulates API calls:
	// select {
	// case <-time.After(time.Second): return fmt.Sprintf("Response from %s", endpoint), nil
	// case <-ctx.Done(): return "", ctx.Err()
	// }
	return "", nil
}

func practicalContext() {
	// TODO: Create a context with timeout: ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// TODO: Defer cancel: defer cancel()
	// TODO: Call apiCall with different endpoints:
	// endpoints := []string{"users", "posts", "comments"}
	// for _, endpoint := range endpoints {
	//     result, err := apiCall(ctx, endpoint)
	//     if err != nil { fmt.Printf("Error calling %s: %v\n", endpoint, err) }
	//     else { fmt.Printf("Success: %s\n", result) }
	// }
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
