# Database Operations - Tasks

## Task 1: Database Connection
Create a database connection.

**Requirements:**
- Connect to a SQLite database
- Test the connection
- Handle connection errors
- Close the connection properly

## Task 2: Creating Tables
Create database tables.

**Requirements:**
- Create a users table with appropriate columns
- Handle table creation errors
- Use proper SQL syntax
- Demonstrate table structure

## Task 3: Inserting Data
Insert data into the database.

**Requirements:**
- Insert single records
- Handle insertion errors
- Get the last insert ID
- Insert multiple records efficiently

## Task 4: Querying Data
Query data from the database.

**Requirements:**
- Query single rows
- Query multiple rows
- Handle query errors
- Process query results

## Task 5: Updating and Deleting
Update and delete database records.

**Requirements:**
- Update existing records
- Delete records
- Handle update/delete errors
- Check affected rows

## Task 6: Transactions
Work with database transactions.

**Requirements:**
- Begin transactions
- Execute multiple operations
- Commit or rollback transactions
- Handle transaction errors

---

## How to complete these tasks:

1. Navigate to the `student/` directory
2. Create your code in `main.go`
3. Run your code: `go run main.go`

## Example structure for student/main.go:
```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/mattn/go-sqlite3"
)

type User struct {
    ID    int    `db:"id"`
    Name  string `db:"name"`
    Email string `db:"email"`
    Age   int    `db:"age"`
}

// Task 1: Database connection
func connectToDatabase() (*sql.DB, error) {
    // TODO: Connect to SQLite database
    // TODO: Test the connection
    // TODO: Handle connection errors
    return nil, nil
}

// Task 2: Creating tables
func createTables(db *sql.DB) error {
    // TODO: Create users table with appropriate columns
    // TODO: Handle table creation errors
    // TODO: Use proper SQL syntax
    return nil
}

// Task 3: Inserting data
func insertUser(db *sql.DB, name, email string, age int) error {
    // TODO: Insert single record
    // TODO: Handle insertion errors
    // TODO: Get the last insert ID
    return nil
}

// Task 4: Querying data
func getUserByID(db *sql.DB, id int) (*User, error) {
    // TODO: Query single row
    // TODO: Handle query errors
    // TODO: Process query results
    return nil, nil
}

func getAllUsers(db *sql.DB) ([]User, error) {
    // TODO: Query multiple rows
    // TODO: Handle query errors
    // TODO: Process query results
    return nil, nil
}

// Task 5: Updating and deleting
func updateUser(db *sql.DB, id int, name, email string, age int) error {
    // TODO: Update existing record
    // TODO: Handle update errors
    // TODO: Check affected rows
    return nil
}

func deleteUser(db *sql.DB, id int) error {
    // TODO: Delete record
    // TODO: Handle delete errors
    // TODO: Check affected rows
    return nil
}

// Task 6: Transactions
func transferUserData(db *sql.DB, fromID, toID int) error {
    // TODO: Begin transaction
    // TODO: Execute multiple operations
    // TODO: Commit or rollback transaction
    // TODO: Handle transaction errors
    return nil
}

func main() {
    // Connect to database
    db, err := connectToDatabase()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Create tables
    err = createTables(db)
    if err != nil {
        log.Fatal(err)
    }

    // Insert users
    err = insertUser(db, "John Doe", "john@example.com", 30)
    if err != nil {
        log.Fatal(err)
    }

    err = insertUser(db, "Jane Smith", "jane@example.com", 25)
    if err != nil {
        log.Fatal(err)
    }

    // Query users
    user, err := getUserByID(db, 1)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("User: %+v\n", user)

    users, err := getAllUsers(db)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("All users: %+v\n", users)

    // Update user
    err = updateUser(db, 1, "John Updated", "john.updated@example.com", 31)
    if err != nil {
        log.Fatal(err)
    }

    // Delete user
    err = deleteUser(db, 2)
    if err != nil {
        log.Fatal(err)
    }

    // Transaction example
    err = transferUserData(db, 1, 3)
    if err != nil {
        log.Fatal(err)
    }
}
```

## Testing Your Code

Run your program:
```bash
cd student/
go run main.go
```

You should see output demonstrating all database operation concepts.