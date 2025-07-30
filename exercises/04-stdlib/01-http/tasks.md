# HTTP - Tasks

## Task 1: Basic HTTP Server
Create a basic HTTP server.

**Requirements:**
- Create a simple HTTP server using `net/http`
- Handle GET requests to root path "/"
- Return a simple HTML response
- Start server on port 8080

## Task 2: HTTP Client
Create an HTTP client to make requests.

**Requirements:**
- Create an HTTP client using `net/http`
- Make GET requests to external APIs
- Handle HTTP response status codes
- Parse response body

## Task 3: HTTP Handlers
Create custom HTTP handlers.

**Requirements:**
- Create custom handler functions
- Handle different HTTP methods (GET, POST)
- Parse request parameters
- Return JSON responses

## Task 4: Middleware
Implement HTTP middleware.

**Requirements:**
- Create logging middleware
- Create authentication middleware
- Chain multiple middleware functions
- Demonstrate middleware functionality

## Task 5: File Server
Create a file server.

**Requirements:**
- Serve static files from a directory
- Handle file uploads
- Serve different file types
- Implement basic file operations

## Task 6: REST API
Create a simple REST API.

**Requirements:**
- Create CRUD operations (Create, Read, Update, Delete)
- Use proper HTTP methods and status codes
- Handle JSON request/response
- Implement basic data storage

---

## How to complete these tasks:

1. Navigate to the `student/` directory
2. Create your code in `main.go`
3. Run your code: `go run main.go`

## Example structure for student/main.go:
```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
)

// Task 1: Basic HTTP server
func basicHTTPServer() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        html := `
        <html>
            <head><title>Go HTTP Server</title></head>
            <body>
                <h1>Welcome to Go HTTP Server!</h1>
                <p>This is a basic HTTP server created with Go.</p>
            </body>
        </html>
        `
        w.Header().Set("Content-Type", "text/html")
        fmt.Fprint(w, html)
    })

    fmt.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// Task 2: HTTP client
func httpClient() {
    resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
    if err != nil {
        fmt.Printf("Error making request: %v\n", err)
        return
    }
    defer resp.Body.Close()

    fmt.Printf("Status: %s\n", resp.Status)
    fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
}

// Task 3: HTTP handlers
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        user := User{ID: 1, Name: "John Doe", Email: "john@example.com"}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(user)
    case "POST":
        var user User
        if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(user)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// Task 4: Middleware
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        fmt.Printf("%s %s %v\n", r.Method, r.URL.Path, time.Since(start))
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

// Task 5: File server
func fileServer() {
    // Create a directory for static files
    os.MkdirAll("static", 0755)

    // Serve static files
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // File upload handler
    http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        file, header, err := r.FormFile("file")
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        defer file.Close()

        fmt.Printf("Uploaded file: %s\n", header.Filename)
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "File %s uploaded successfully", header.Filename)
    })
}

// Task 6: REST API
var users = []User{
    {ID: 1, Name: "John Doe", Email: "john@example.com"},
    {ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
}

func restAPI() {
    http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")

        switch r.Method {
        case "GET":
            json.NewEncoder(w).Encode(users)
        case "POST":
            var user User
            if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }
            user.ID = len(users) + 1
            users = append(users, user)
            w.WriteHeader(http.StatusCreated)
            json.NewEncoder(w).Encode(user)
        }
    })
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
    // Apply middleware to all routes
    http.Handle("/api/", loggingMiddleware(authMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Protected endpoint")
    }))))

    // Task 5: File server
    fmt.Println("\n=== Task 5: File Server ===")
    fileServer()

    // Task 6: REST API
    fmt.Println("\n=== Task 6: REST API ===")
    restAPI()

    fmt.Println("Server running on :8080")
    fmt.Println("Available endpoints:")
    fmt.Println("- GET / - Basic HTML page")
    fmt.Println("- GET/POST /user - User handler")
    fmt.Println("- GET/POST /api/users - REST API")
    fmt.Println("- POST /upload - File upload")
    fmt.Println("- GET /static/* - Static files")

    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Testing Your Code

Run your program:
```bash
cd student/
go run main.go
```

You should see output demonstrating all HTTP concepts.