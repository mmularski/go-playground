package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Task 1: Define User struct
type User struct {
	// TODO: Add fields for user data
}

// Task 2: Define Server struct
type Server struct {
	// TODO: Add fields for storing users and managing IDs
}

// Task 3: Create NewServer function
func NewServer() *Server {
	// TODO: Initialize and return a new Server instance
	return nil
}

// Task 4: Implement handleGetUsers
func (s *Server) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	// TODO: Return all users as JSON
}

// Task 5: Implement handleGetUser
func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Get a specific user by ID
}

// Task 6: Implement handleCreateUser
func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Create a new user
}

// Task 7: Implement handleUpdateUser
func (s *Server) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Update an existing user
}

// Task 8: Implement handleDeleteUser
func (s *Server) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Delete a user
}

// Task 9: Implement extractUserID helper function
func extractUserID(r *http.Request) (int, error) {
	// TODO: Extract user ID from URL path
	return 0, nil
}

// Task 10: Implement error handling
func writeError(w http.ResponseWriter, status int, message string) {
	// TODO: Create a proper error response
}

// Task 11: Implement logging middleware
func loggingMiddleware(next http.Handler) http.Handler {
	// TODO: Create middleware that logs request details
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement logging logic
		next.ServeHTTP(w, r)
	})
}

// Task 12: Implement CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	// TODO: Create CORS middleware
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement CORS logic
		next.ServeHTTP(w, r)
	})
}

// Task 13: Set up routing in main function
func main() {
	// TODO: Set up the REST API server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}