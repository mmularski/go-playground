# HTTP Package - Overview

## What is the HTTP Package?

The `net/http` package provides HTTP client and server implementations. It's the foundation for building web applications and making HTTP requests in Go.

## HTTP Client

### Basic GET Request
```go
import (
    "fmt"
    "io"
    "net/http"
)

func basicGetRequest() {
    resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
    if err != nil {
        log.Printf("Error: %v", err)
        return
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Printf("Error reading body: %v", err)
        return
    }

    fmt.Printf("Status: %s\n", resp.Status)
    fmt.Printf("Response: %s\n", string(body))
}
```

### POST Request with JSON
```go
import (
    "bytes"
    "encoding/json"
    "net/http"
)

func postWithJSON() {
    data := map[string]interface{}{
        "title":  "Test Post",
        "body":   "This is a test post",
        "userId": 1,
    }

    jsonData, err := json.Marshal(data)
    if err != nil {
        log.Printf("Error marshaling JSON: %v", err)
        return
    }

    resp, err := http.Post("https://jsonplaceholder.typicode.com/posts",
        "application/json",
        bytes.NewBuffer(jsonData))
    if err != nil {
        log.Printf("Error: %v", err)
        return
    }
    defer resp.Body.Close()

    // Read response...
}
```

## Custom HTTP Client

```go
import (
    "net/http"
    "time"
)

func customClient() {
    client := &http.Client{
        Timeout: 10 * time.Second,
    }

    req, err := http.NewRequest("GET", "https://httpbin.org/delay/2", nil)
    if err != nil {
        log.Printf("Error creating request: %v", err)
        return
    }

    // Add headers
    req.Header.Add("User-Agent", "Go-Client/1.0")
    req.Header.Add("Accept", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        log.Printf("Error: %v", err)
        return
    }
    defer resp.Body.Close()

    fmt.Printf("Status: %s\n", resp.Status)
    fmt.Printf("Headers: %v\n", resp.Header)
}
```

## HTTP Server

### Basic Server Setup
```go
import (
    "encoding/json"
    "net/http"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    switch r.Method {
    case "GET":
        // Return all users
        json.NewEncoder(w).Encode(users)
    case "POST":
        // Create new user
        var newUser User
        if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        // Add user logic...
        json.NewEncoder(w).Encode(newUser)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func startServer() {
    http.HandleFunc("/users", handleUsers)
    http.HandleFunc("/health", handleHealth)

    fmt.Println("Starting server on :8080")
    http.ListenAndServe(":8080", nil)
}
```

### Health Check Endpoint
```go
func handleHealth(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "status": "healthy",
        "time":   time.Now().Format(time.RFC3339),
    })
}
```

## Middleware Pattern

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

func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

## File Server

```go
func serveStaticFiles() {
    // Serve static files from ./static directory
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
}
```

## REST API Patterns

### CRUD Operations
```go
func handleUserCRUD(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    switch r.Method {
    case "GET":
        // Read - return user(s)
        id := r.URL.Query().Get("id")
        if id != "" {
            // Get specific user
        } else {
            // Get all users
        }
    case "POST":
        // Create - add new user
        var user User
        json.NewDecoder(r.Body).Decode(&user)
        // Add to database...
    case "PUT":
        // Update - modify existing user
        var user User
        json.NewDecoder(r.Body).Decode(&user)
        // Update in database...
    case "DELETE":
        // Delete - remove user
        id := r.URL.Query().Get("id")
        // Delete from database...
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
```

## Error Handling

```go
func handleWithErrors(w http.ResponseWriter, r *http.Request) {
    // Validate request
    if r.Method != "POST" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Parse JSON
    var data map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Process data...
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
```

## Best Practices

1. **Always close response bodies** - Use `defer resp.Body.Close()`
2. **Set timeouts** - Use custom client with timeout
3. **Handle errors properly** - Check all error returns
4. **Set appropriate headers** - Content-Type, CORS, etc.
5. **Use middleware** - For logging, auth, CORS
6. **Validate input** - Check request data
7. **Return proper status codes** - Use appropriate HTTP status codes

## Common Status Codes

- `200 OK` - Success
- `201 Created` - Resource created
- `400 Bad Request` - Invalid request
- `401 Unauthorized` - Authentication required
- `403 Forbidden` - Access denied
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server error

## Example Program

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    users := []User{
        {ID: 1, Name: "Alice", Age: 25},
        {ID: 2, Name: "Bob", Age: 30},
    }

    json.NewEncoder(w).Encode(users)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "status": "healthy",
        "time":   time.Now().Format(time.RFC3339),
    })
}

func main() {
    http.HandleFunc("/users", handleUsers)
    http.HandleFunc("/health", handleHealth)

    fmt.Println("Server starting on :8080")
    http.ListenAndServe(":8080", nil)
}
```

## Next Steps

Now that you understand HTTP, try the tasks in `tasks.md`!
