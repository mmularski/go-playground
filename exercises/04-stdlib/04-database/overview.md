# Database Operations - Overview

## What are Database Operations?

Database operations in Go are handled through the `database/sql` package and various database drivers. This provides a standard interface for working with SQL databases.

## Basic Database Connection

### SQLite Example
```go
import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func connectToDatabase() (*sql.DB, error) {
    db, err := sql.Open("sqlite3", "./test.db")
    if err != nil {
        return nil, fmt.Errorf("error opening database: %v", err)
    }

    // Test the connection
    err = db.Ping()
    if err != nil {
        return nil, fmt.Errorf("error connecting to database: %v", err)
    }

    return db, nil
}
```

## Creating Tables

```go
func createTables(db *sql.DB) error {
    createTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        email TEXT UNIQUE,
        age INTEGER,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );`

    _, err := db.Exec(createTableSQL)
    return err
}
```

## Inserting Data

```go
func insertUser(db *sql.DB, name, email string, age int) error {
    insertSQL := "INSERT INTO users (name, email, age) VALUES (?, ?, ?)"

    result, err := db.Exec(insertSQL, name, email, age)
    if err != nil {
        return fmt.Errorf("error inserting user: %v", err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return fmt.Errorf("error getting last insert id: %v", err)
    }

    fmt.Printf("Inserted user with ID: %d\n", id)
    return nil
}
```

## Querying Data

### Single Row Query
```go
type User struct {
    ID    int    `db:"id"`
    Name  string `db:"name"`
    Email string `db:"email"`
    Age   int    `db:"age"`
}

func getUserByID(db *sql.DB, id int) (*User, error) {
    var user User
    query := "SELECT id, name, email, age FROM users WHERE id = ?"

    err := db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Age)
    if err != nil {
        return nil, fmt.Errorf("error querying user: %v", err)
    }

    return &user, nil
}
```

### Multiple Rows Query
```go
func getAllUsers(db *sql.DB) ([]User, error) {
    query := "SELECT id, name, email, age FROM users ORDER BY id"

    rows, err := db.Query(query)
    if err != nil {
        return nil, fmt.Errorf("error querying users: %v", err)
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
        if err != nil {
            return nil, fmt.Errorf("error scanning user: %v", err)
        }
        users = append(users, user)
    }

    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating rows: %v", err)
    }

    return users, nil
}
```

## Updating Data

```go
func updateUser(db *sql.DB, id int, name, email string, age int) error {
    updateSQL := "UPDATE users SET name = ?, email = ?, age = ? WHERE id = ?"

    result, err := db.Exec(updateSQL, name, email, age, id)
    if err != nil {
        return fmt.Errorf("error updating user: %v", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error getting rows affected: %v", err)
    }

    fmt.Printf("Updated %d rows\n", rowsAffected)
    return nil
}
```

## Deleting Data

```go
func deleteUser(db *sql.DB, id int) error {
    deleteSQL := "DELETE FROM users WHERE id = ?"

    result, err := db.Exec(deleteSQL, id)
    if err != nil {
        return fmt.Errorf("error deleting user: %v", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error getting rows affected: %v", err)
    }

    fmt.Printf("Deleted %d rows\n", rowsAffected)
    return nil
}
```

## Transactions

```go
func transferMoney(db *sql.DB, fromID, toID int, amount float64) error {
    tx, err := db.Begin()
    if err != nil {
        return fmt.Errorf("error beginning transaction: %v", err)
    }
    defer tx.Rollback() // Will be ignored if tx.Commit() is called

    // Deduct from first account
    _, err = tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, fromID)
    if err != nil {
        return fmt.Errorf("error deducting from account: %v", err)
    }

    // Add to second account
    _, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, toID)
    if err != nil {
        return fmt.Errorf("error adding to account: %v", err)
    }

    // Commit the transaction
    err = tx.Commit()
    if err != nil {
        return fmt.Errorf("error committing transaction: %v", err)
    }

    return nil
}
```

## Prepared Statements

```go
func preparedStatementExample(db *sql.DB) error {
    // Prepare the statement
    stmt, err := db.Prepare("INSERT INTO users (name, email, age) VALUES (?, ?, ?)")
    if err != nil {
        return fmt.Errorf("error preparing statement: %v", err)
    }
    defer stmt.Close()

    // Execute multiple times
    users := []struct {
        name  string
        email string
        age   int
    }{
        {"Alice", "alice@example.com", 25},
        {"Bob", "bob@example.com", 30},
        {"Charlie", "charlie@example.com", 35},
    }

    for _, user := range users {
        _, err := stmt.Exec(user.name, user.email, user.age)
        if err != nil {
            return fmt.Errorf("error executing prepared statement: %v", err)
        }
    }

    return nil
}
```

## Error Handling

```go
func robustDatabaseOperations(db *sql.DB) {
    // Check for specific errors
    var user User
    err := db.QueryRow("SELECT id, name FROM users WHERE id = ?", 999).Scan(&user.ID, &user.Name)
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("No user found with that ID")
            return
        }
        fmt.Printf("Database error: %v\n", err)
        return
    }

    fmt.Printf("Found user: %s\n", user.Name)
}
```

## Connection Pooling

```go
func configureConnectionPool(db *sql.DB) {
    // Set connection pool settings
    db.SetMaxOpenConns(25)    // Maximum number of open connections
    db.SetMaxIdleConns(5)     // Maximum number of idle connections
    db.SetConnMaxLifetime(5 * time.Minute) // Maximum lifetime of a connection
}
```

## Best Practices

1. **Always close connections** - Use `defer db.Close()`
2. **Use prepared statements** - For repeated queries
3. **Handle transactions properly** - Always commit or rollback
4. **Check for specific errors** - Handle `sql.ErrNoRows` appropriately
5. **Use connection pooling** - Configure pool settings
6. **Validate input** - Prevent SQL injection
7. **Use context for timeouts** - Set query timeouts

## Example Program

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

func main() {
    // Connect to database
    db, err := sql.Open("sqlite3", "./test.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Create table
    createTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        email TEXT UNIQUE,
        age INTEGER
    );`

    _, err = db.Exec(createTableSQL)
    if err != nil {
        log.Fatal(err)
    }

    // Insert user
    result, err := db.Exec("INSERT INTO users (name, email, age) VALUES (?, ?, ?)",
        "John Doe", "john@example.com", 30)
    if err != nil {
        log.Fatal(err)
    }

    id, _ := result.LastInsertId()
    fmt.Printf("Inserted user with ID: %d\n", id)

    // Query user
    var user User
    err = db.QueryRow("SELECT id, name, email, age FROM users WHERE id = ?", id).
        Scan(&user.ID, &user.Name, &user.Email, &user.Age)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("User: %+v\n", user)
}
```

## Next Steps

Now that you understand database operations, try the tasks in `tasks.md`!