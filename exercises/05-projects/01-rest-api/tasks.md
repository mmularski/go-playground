# REST API Project - Tasks

## Task 1: Basic HTTP Server Setup

Create a basic HTTP server that can handle different routes and HTTP methods.

**Requirements:**
- Set up a server that listens on port 8080
- Handle basic routing for `/api/users`
- Implement proper HTTP method handling (GET, POST, PUT, DELETE)
- Return appropriate HTTP status codes

## Task 2: User Management API

Implement a complete CRUD (Create, Read, Update, Delete) API for users.

**Requirements:**
- Create a `User` struct with fields: ID, Name, Email, Created
- Implement GET `/api/users` to list all users
- Implement GET `/api/users/{id}` to get a specific user
- Implement POST `/api/users` to create a new user
- Implement PUT `/api/users/{id}` to update a user
- Implement DELETE `/api/users/{id}` to delete a user

## Task 3: JSON Handling

Implement proper JSON serialization and deserialization.

**Requirements:**
- Use struct tags for JSON field mapping
- Handle JSON encoding/decoding with proper error handling
- Implement custom JSON marshaling for date formatting
- Add input validation for required fields

## Task 4: Error Handling

Implement comprehensive error handling throughout the API.

**Requirements:**
- Create custom error types for different scenarios
- Return appropriate HTTP status codes for different errors
- Implement proper error messages in JSON format
- Handle edge cases (invalid IDs, missing fields, etc.)

## Task 5: URL Routing and Parameters

Implement advanced routing with path and query parameters.

**Requirements:**
- Extract user ID from URL path parameters
- Handle query parameters for filtering and pagination
- Implement proper URL parsing and validation
- Add support for optional query parameters

## Task 6: Middleware Implementation

Add middleware for logging, CORS, and authentication.

**Requirements:**
- Implement logging middleware to track requests
- Add CORS middleware for cross-origin requests
- Create authentication middleware (basic implementation)
- Chain multiple middleware functions

## Task 7: Data Storage

Implement in-memory data storage with proper concurrency handling.

**Requirements:**
- Use a thread-safe data structure for storing users
- Implement proper locking mechanisms
- Handle concurrent access to shared data
- Add data persistence (optional: save to file)

## Task 8: API Documentation

Create comprehensive API documentation.

**Requirements:**
- Document all endpoints with examples
- Include request/response schemas
- Add error code documentation
- Create usage examples

## Task 9: Testing

Write comprehensive tests for the API.

**Requirements:**
- Test all HTTP endpoints
- Test error scenarios
- Test concurrent access
- Use `httptest` package for testing

## Task 10: Advanced Features

Implement advanced API features.

**Requirements:**
- Add pagination support
- Implement search functionality
- Add sorting options
- Implement rate limiting (basic)

## Bonus Challenges

### Challenge 1: Database Integration
- Replace in-memory storage with a real database
- Use SQLite or PostgreSQL
- Implement proper database transactions

### Challenge 2: Authentication
- Implement JWT token authentication
- Add user registration and login endpoints
- Implement password hashing

### Challenge 3: File Upload
- Add support for file uploads
- Implement image processing
- Add file storage functionality

### Challenge 4: Real-time Features
- Implement WebSocket support
- Add real-time notifications
- Create a chat feature

## Testing Your API

### Using curl
```bash
# Get all users
curl http://localhost:8080/api/users

# Create a user
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com"}'

# Get a specific user
curl http://localhost:8080/api/users/1

# Update a user
curl -X PUT http://localhost:8080/api/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"John Updated","email":"john.updated@example.com"}'

# Delete a user
curl -X DELETE http://localhost:8080/api/users/1
```

### Using Postman
- Import the API endpoints into Postman
- Test different scenarios and edge cases
- Verify proper error handling

## Success Criteria

Your REST API should:
- ✅ Handle all CRUD operations for users
- ✅ Return proper HTTP status codes
- ✅ Include comprehensive error handling
- ✅ Support JSON request/response format
- ✅ Include middleware for logging and CORS
- ✅ Be properly tested with unit tests
- ✅ Handle concurrent requests safely
- ✅ Include proper documentation

## Next Steps

Start with Task 1 and work through each task systematically. Each task builds upon the previous ones, so make sure to complete them in order.