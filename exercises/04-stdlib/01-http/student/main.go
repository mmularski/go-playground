package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Task 1: Basic HTTP server
func basicHTTPServer() {
	// TODO: Create a simple HTTP server
}

// Task 2: HTTP client
func httpClient() {
	// TODO: Make HTTP requests
}

// Task 3: HTTP handlers
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Handle different HTTP methods
}

// Task 4: Middleware
func loggingMiddleware(next http.Handler) http.Handler {
	// TODO: Create logging middleware
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Log request details
	})
}

func authMiddleware(next http.Handler) http.Handler {
	// TODO: Create authentication middleware
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Check authentication
	})
}

// Task 5: File server
func fileServer() {
	// TODO: Serve static files
}

// Task 6: REST API
var users = []User{
	{ID: 1, Name: "John Doe", Email: "john@example.com"},
	{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
}

func restAPI() {
	// TODO: Create CRUD operations
}

func main() {
	// Task 1: Basic HTTP server
	go func() {
		fmt.Println("Starting basic HTTP server...")
		basicHTTPServer()
	}()

	// Wait a bit for server to start
	time.Sleep(time.Second)

	// Task 2: HTTP client
	fmt.Println("\n=== Task 2: HTTP Client ===")
	httpClient()

	// Task 3: HTTP handlers
	fmt.Println("\n=== Task 3: HTTP Handlers ===")
	http.HandleFunc("/user", userHandler)

	// Task 4: Middleware
	fmt.Println("\n=== Task 4: Middleware ===")
	// TODO: Apply middleware to routes

	// Task 5: File server
	fmt.Println("\n=== Task 5: File Server ===")
	fileServer()

	// Task 6: REST API
	fmt.Println("\n=== Task 6: REST API ===")
	restAPI()

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
