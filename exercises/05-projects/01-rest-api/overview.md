# REST API Project - Overview

## What is a REST API?

A REST API (Representational State Transfer) is a web service that follows REST principles. It uses HTTP methods to perform operations on resources.

## REST API Principles

- **Stateless** - Each request contains all information needed
- **Client-Server** - Separation of concerns
- **Cacheable** - Responses can be cached
- **Uniform Interface** - Consistent HTTP methods and status codes
- **Layered System** - Can be composed of multiple layers

## HTTP Methods

### GET - Retrieve Data
```http
GET /api/users
GET /api/users/123
```

### POST - Create Data
```http
POST /api/users
Content-Type: application/json

{
    "name": "John Doe",
    "email": "john@example.com"
}
```

### PUT - Update Data (Full Update)
```http
PUT /api/users/123
Content-Type: application/json

{
    "name": "John Doe",
    "email": "john@example.com",
    "age": 30
}
```

### PATCH - Update Data (Partial Update)
```http
PATCH /api/users/123
Content-Type: application/json

{
    "age": 31
}
```

### DELETE - Remove Data
```http
DELETE /api/users/123
```

## HTTP Status Codes

### 2xx Success
- **200 OK** - Request successful
- **201 Created** - Resource created successfully
- **204 No Content** - Request successful, no content to return

### 4xx Client Errors
- **400 Bad Request** - Invalid request
- **401 Unauthorized** - Authentication required
- **403 Forbidden** - Access denied
- **404 Not Found** - Resource not found
- **409 Conflict** - Resource conflict

### 5xx Server Errors
- **500 Internal Server Error** - Server error
- **502 Bad Gateway** - Gateway error
- **503 Service Unavailable** - Service unavailable

## Go HTTP Server Structure

### Basic HTTP Server
```go
package main

import (
    "encoding/json"
    "net/http"
    "log"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type Server struct {
    users map[int]*User
    nextID int
}

func NewServer() *Server {
    return &Server{
        users: make(map[int]*User),
        nextID: 1,
    }
}

func (s *Server) handleGetUsers(w http.ResponseWriter, r *http.Request) {
    users := make([]*User, 0, len(s.users))
    for _, user := range s.users {
        users = append(users, user)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
    // Extract user ID from URL
    // Implementation here
}

func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    user.ID = s.nextID
    s.nextID++
    s.users[user.ID] = &user

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func (s *Server) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
    // Implementation here
}

func (s *Server) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
    // Implementation here
}

func main() {
    server := NewServer()

    // Route handlers
    http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            server.handleGetUsers(w, r)
        case "POST":
            server.handleCreateUser(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    http.HandleFunc("/api/users/", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            server.handleGetUser(w, r)
        case "PUT":
            server.handleUpdateUser(w, r)
        case "DELETE":
            server.handleDeleteUser(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## URL Routing and Parameters

### Path Parameters
```go
func extractUserID(r *http.Request) (int, error) {
    // Extract ID from URL like /api/users/123
    path := strings.TrimPrefix(r.URL.Path, "/api/users/")
    return strconv.Atoi(path)
}
```

### Query Parameters
```go
func handleGetUsers(w http.ResponseWriter, r *http.Request) {
    // Handle query parameters like /api/users?limit=10&offset=0
    limit := r.URL.Query().Get("limit")
    offset := r.URL.Query().Get("offset")

    // Parse and use parameters
    // Implementation here
}
```

## Middleware Pattern

### Logging Middleware
```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        // Call the next handler
        next.ServeHTTP(w, r)

        // Log the request
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}
```

### CORS Middleware
```go
func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}
```

## Error Handling

### Custom Error Types
```go
type APIError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

func (e *APIError) Error() string {
    return e.Message
}

func writeError(w http.ResponseWriter, status int, message string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(APIError{
        Code:    status,
        Message: message,
    })
}
```

## JSON Handling

### Struct Tags
```go
type User struct {
    ID       int       `json:"id"`
    Name     string    `json:"name"`
    Email    string    `json:"email"`
    Password string    `json:"-"`  // Exclude from JSON
    Created  time.Time `json:"created_at,omitempty"`
}
```

### Custom Marshaling
```go
func (u *User) MarshalJSON() ([]byte, error) {
    type Alias User
    return json.Marshal(&struct {
        *Alias
        Created string `json:"created_at"`
    }{
        Alias:   (*Alias)(u),
        Created: u.Created.Format("2006-01-02T15:04:05Z07:00"),
    })
}
```

## Testing HTTP Handlers

### Using httptest
```go
func TestGetUsers(t *testing.T) {
    server := NewServer()

    req, err := http.NewRequest("GET", "/api/users", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(server.handleGetUsers)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    var users []User
    if err := json.NewDecoder(rr.Body).Decode(&users); err != nil {
        t.Errorf("failed to decode response: %v", err)
    }
}
```

## Project Structure

```
rest-api/
├── main.go              # Entry point
├── server/
│   ├── server.go        # Server implementation
│   ├── handlers.go      # HTTP handlers
│   └── middleware.go    # Middleware functions
├── models/
│   └── user.go          # Data models
├── storage/
│   └── memory.go        # Data storage
└── tests/
    └── handlers_test.go # Tests
```

## Best Practices

### 1. Use Proper HTTP Methods
- GET for retrieving data
- POST for creating data
- PUT for full updates
- PATCH for partial updates
- DELETE for removing data

### 2. Return Appropriate Status Codes
- 200 for successful GET requests
- 201 for successful POST requests
- 204 for successful DELETE requests
- 400 for bad requests
- 404 for not found
- 500 for server errors

### 3. Use JSON for Data Exchange
```go
w.Header().Set("Content-Type", "application/json")
```

### 4. Handle Errors Gracefully
```go
if err != nil {
    writeError(w, http.StatusInternalServerError, "Internal server error")
    return
}
```

### 5. Validate Input
```go
func validateUser(user *User) error {
    if user.Name == "" {
        return errors.New("name is required")
    }
    if user.Email == "" {
        return errors.New("email is required")
    }
    return nil
}
```

## Example Complete API

```go
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "strings"
    "time"
)

type User struct {
    ID      int       `json:"id"`
    Name    string    `json:"name"`
    Email   string    `json:"email"`
    Created time.Time `json:"created_at"`
}

type Server struct {
    users map[int]*User
    nextID int
}

func NewServer() *Server {
    return &Server{
        users: make(map[int]*User),
        nextID: 1,
    }
}

func (s *Server) handleGetUsers(w http.ResponseWriter, r *http.Request) {
    users := make([]*User, 0, len(s.users))
    for _, user := range s.users {
        users = append(users, user)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/users/"))
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    user, exists := s.users[id]
    if !exists {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    user.ID = s.nextID
    user.Created = time.Now()
    s.nextID++
    s.users[user.ID] = &user

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func main() {
    server := NewServer()

    http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            server.handleGetUsers(w, r)
        case "POST":
            server.handleCreateUser(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    http.HandleFunc("/api/users/", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            server.handleGetUser(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Next Steps

Now that you understand REST APIs, try the tasks in `tasks.md`!