# Database Application Project - Tasks

## Task 1: Database Connection Setup

Set up a database connection with proper configuration and connection pooling.

**Requirements:**
- Configure database connection with connection pooling
- Set up connection timeouts and retry logic
- Implement connection health checks
- Handle connection errors gracefully

## Task 2: Data Models and Structs

Define data models and structs for the application.

**Requirements:**
- Create User model with proper fields and tags
- Create Product model with relationships
- Create Order and OrderItem models
- Add JSON and database tags for proper serialization

## Task 3: Database Migrations

Implement a migration system to manage database schema changes.

**Requirements:**
- Create migrations table to track applied migrations
- Implement migration runner
- Create initial schema migrations
- Handle migration rollbacks

## Task 4: CRUD Operations

Implement basic CRUD (Create, Read, Update, Delete) operations.

**Requirements:**
- Implement Create operations with proper error handling
- Implement Read operations with filtering and pagination
- Implement Update operations with validation
- Implement Delete operations with soft delete option

## Task 5: Query Building

Create a dynamic query builder for complex queries.

**Requirements:**
- Implement WHERE clause building
- Add ORDER BY and LIMIT support
- Implement JOIN operations
- Add support for complex conditions

## Task 6: Prepared Statements

Implement prepared statements for better performance and security.

**Requirements:**
- Prepare statements for common operations
- Use parameterized queries to prevent SQL injection
- Implement statement caching
- Handle statement preparation errors

## Task 7: Transactions

Implement transaction support for complex operations.

**Requirements:**
- Implement transaction management
- Handle transaction rollbacks on errors
- Implement nested transactions
- Add transaction isolation levels

## Task 8: Connection Pooling

Implement advanced connection pooling with monitoring.

**Requirements:**
- Configure connection pool parameters
- Implement pool monitoring and statistics
- Handle pool exhaustion scenarios
- Add connection health checks

## Task 9: Error Handling

Implement comprehensive error handling for database operations.

**Requirements:**
- Handle specific database error codes
- Implement custom error types
- Add error logging and monitoring
- Provide meaningful error messages

## Task 10: Data Validation

Implement data validation before database operations.

**Requirements:**
- Validate data types and formats
- Implement business rule validation
- Add constraint checking
- Handle validation errors gracefully

## Bonus Challenges

### Challenge 1: Multi-Database Support
- Support multiple database types (PostgreSQL, MySQL, SQLite)
- Implement database abstraction layer
- Add database-specific optimizations
- Handle database-specific features

### Challenge 2: Caching Layer
- Implement Redis caching for frequently accessed data
- Add cache invalidation strategies
- Implement cache-aside pattern
- Handle cache failures gracefully

### Challenge 3: Search and Full-Text Search
- Implement basic search functionality
- Add full-text search capabilities
- Implement search result ranking
- Add search result highlighting

### Challenge 4: Data Export/Import
- Implement data export to CSV/JSON
- Add data import functionality
- Handle bulk operations
- Implement data validation for imports

## Testing Your Database Application

### Basic Usage Examples
```bash
# Start the application
go run main.go

# Test database connection
curl http://localhost:8080/health

# Create a user
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"username":"john","email":"john@example.com"}'

# Get a user
curl http://localhost:8080/users/1

# Get all users
curl http://localhost:8080/users

# Update a user
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"username":"john_updated","email":"john.updated@example.com"}'

# Delete a user
curl -X DELETE http://localhost:8080/users/1
```

### Testing Different Scenarios
```bash
# Test pagination
curl "http://localhost:8080/users?limit=5&offset=0"

# Test filtering
curl "http://localhost:8080/users?email=john@example.com"

# Test sorting
curl "http://localhost:8080/users?sort=created_at&order=desc"

# Test error handling
curl http://localhost:8080/users/999  # Should return 404
```

## Success Criteria

Your database application should:
- ✅ Connect to database with proper connection pooling
- ✅ Implement comprehensive CRUD operations
- ✅ Handle database migrations automatically
- ✅ Use prepared statements for security
- ✅ Support transactions for complex operations
- ✅ Include proper error handling and validation
- ✅ Provide RESTful API endpoints
- ✅ Include comprehensive logging and monitoring
- ✅ Handle concurrent access safely
- ✅ Be well-documented with usage examples

## Project Structure

```
database-app/
├── cmd/
│   └── app/
│       └── main.go          # Application entry point
├── internal/
│   ├── database/
│   │   ├── connection.go    # Database connection
│   │   ├── models.go        # Data models
│   │   ├── migrations.go    # Database migrations
│   │   └── queries.go       # Query operations
│   ├── api/
│   │   ├── handlers.go      # HTTP handlers
│   │   └── middleware.go    # Middleware
│   └── validation/
│       └── validator.go     # Data validation
├── migrations/
│   ├── 001_create_users.sql
│   ├── 002_create_products.sql
│   └── 003_create_orders.sql
├── Dockerfile
├── docker-compose.yml
└── go.mod
```

## Dependencies

Add these to your go.mod:
```
require (
    github.com/lib/pq v1.10.9
    github.com/go-sql-driver/mysql v1.7.1
    github.com/mattn/go-sqlite3 v1.14.17
    github.com/go-playground/validator/v10 v10.15.0
)
```

## Database Setup

### PostgreSQL
```sql
-- Create database
CREATE DATABASE myapp;

-- Create user table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create products table
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    stock INTEGER DEFAULT 0,
    active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### SQLite
```sql
-- Create users table
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Create products table
CREATE TABLE products (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT,
    price REAL NOT NULL,
    stock INTEGER DEFAULT 0,
    active BOOLEAN DEFAULT 1,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

## Next Steps

Start with Task 1 and work through each task systematically. Each task builds upon the previous ones, so make sure to complete them in order.