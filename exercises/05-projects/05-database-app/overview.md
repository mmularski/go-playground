# Database Application Project - Overview

## What is a Database Application?

A database application is a software system that manages data storage, retrieval, and manipulation using a database management system. It provides a structured way to store, organize, and access data efficiently.

## Database Types and Selection

### SQL Databases
```go
import (
    "database/sql"
    _ "github.com/lib/pq"     // PostgreSQL
    _ "github.com/go-sql-driver/mysql" // MySQL
    _ "github.com/mattn/go-sqlite3"    // SQLite
)
```

### NoSQL Databases
```go
import (
    "go.mongodb.org/mongo-driver/mongo"
    "github.com/go-redis/redis/v8"
)
```

## Database Connection Management

### Connection Pool Setup
```go
import (
    "database/sql"
    "time"
    _ "github.com/lib/pq"
)

type Database struct {
    db *sql.DB
}

func NewDatabase(databaseURL string) (*Database, error) {
    db, err := sql.Open("postgres", databaseURL)
    if err != nil {
        return nil, err
    }

    // Configure connection pool
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(5)
    db.SetConnMaxLifetime(5 * time.Minute)

    // Test connection
    if err := db.Ping(); err != nil {
        return nil, err
    }

    return &Database{db: db}, nil
}

func (d *Database) Close() error {
    return d.db.Close()
}
```

## Data Models and Structs

### User Model
```go
type User struct {
    ID        int       `json:"id" db:"id"`
    Username  string    `json:"username" db:"username"`
    Email     string    `json:"email" db:"email"`
    Password  string    `json:"-" db:"password_hash"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
```

### Product Model
```go
type Product struct {
    ID          int     `json:"id" db:"id"`
    Name        string  `json:"name" db:"name"`
    Description string  `json:"description" db:"description"`
    Price       float64 `json:"price" db:"price"`
    CategoryID  int     `json:"category_id" db:"category_id"`
    Stock       int     `json:"stock" db:"stock"`
    Active      bool    `json:"active" db:"active"`
}
```

## CRUD Operations

### Create (Insert)
```go
func (db *Database) CreateUser(user *User) error {
    query := `
        INSERT INTO users (username, email, password_hash, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id`

    now := time.Now()
    user.CreatedAt = now
    user.UpdatedAt = now

    return db.db.QueryRow(
        query,
        user.Username,
        user.Email,
        user.Password,
        user.CreatedAt,
        user.UpdatedAt,
    ).Scan(&user.ID)
}
```

### Read (Select)
```go
func (db *Database) GetUser(id int) (*User, error) {
    user := &User{}
    query := `
        SELECT id, username, email, password_hash, created_at, updated_at
        FROM users WHERE id = $1`

    err := db.db.QueryRow(query, id).Scan(
        &user.ID,
        &user.Username,
        &user.Email,
        &user.Password,
        &user.CreatedAt,
        &user.UpdatedAt,
    )

    if err != nil {
        return nil, err
    }

    return user, nil
}

func (db *Database) GetUsers(limit, offset int) ([]*User, error) {
    query := `
        SELECT id, username, email, created_at, updated_at
        FROM users
        ORDER BY created_at DESC
        LIMIT $1 OFFSET $2`

    rows, err := db.db.Query(query, limit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []*User
    for rows.Next() {
        user := &User{}
        err := rows.Scan(
            &user.ID,
            &user.Username,
            &user.Email,
            &user.CreatedAt,
            &user.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, rows.Err()
}
```

### Update
```go
func (db *Database) UpdateUser(user *User) error {
    query := `
        UPDATE users
        SET username = $1, email = $2, updated_at = $3
        WHERE id = $4`

    user.UpdatedAt = time.Now()

    result, err := db.db.Exec(
        query,
        user.Username,
        user.Email,
        user.UpdatedAt,
        user.ID,
    )
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected == 0 {
        return sql.ErrNoRows
    }

    return nil
}
```

### Delete
```go
func (db *Database) DeleteUser(id int) error {
    query := `DELETE FROM users WHERE id = $1`

    result, err := db.db.Exec(query, id)
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected == 0 {
        return sql.ErrNoRows
    }

    return nil
}
```

## Transactions

### Transaction Management
```go
func (db *Database) CreateOrderWithItems(order *Order, items []OrderItem) error {
    tx, err := db.db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    // Insert order
    orderQuery := `
        INSERT INTO orders (user_id, total_amount, status, created_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id`

    now := time.Now()
    order.CreatedAt = now

    err = tx.QueryRow(
        orderQuery,
        order.UserID,
        order.TotalAmount,
        order.Status,
        order.CreatedAt,
    ).Scan(&order.ID)
    if err != nil {
        return err
    }

    // Insert order items
    itemQuery := `
        INSERT INTO order_items (order_id, product_id, quantity, price)
        VALUES ($1, $2, $3, $4)`

    for _, item := range items {
        item.OrderID = order.ID
        _, err = tx.Exec(
            itemQuery,
            item.OrderID,
            item.ProductID,
            item.Quantity,
            item.Price,
        )
        if err != nil {
            return err
        }
    }

    return tx.Commit()
}
```

## Prepared Statements

### Using Prepared Statements
```go
type Database struct {
    db *sql.DB
    stmts map[string]*sql.Stmt
}

func (db *Database) prepareStatements() error {
    queries := map[string]string{
        "get_user": `
            SELECT id, username, email, created_at, updated_at
            FROM users WHERE id = $1`,
        "create_user": `
            INSERT INTO users (username, email, password_hash, created_at, updated_at)
            VALUES ($1, $2, $3, $4, $5)
            RETURNING id`,
        "update_user": `
            UPDATE users
            SET username = $1, email = $2, updated_at = $3
            WHERE id = $4`,
        "delete_user": `DELETE FROM users WHERE id = $1`,
    }

    db.stmts = make(map[string]*sql.Stmt)

    for name, query := range queries {
        stmt, err := db.db.Prepare(query)
        if err != nil {
            return err
        }
        db.stmts[name] = stmt
    }

    return nil
}

func (db *Database) GetUser(id int) (*User, error) {
    user := &User{}
    stmt := db.stmts["get_user"]

    err := stmt.QueryRow(id).Scan(
        &user.ID,
        &user.Username,
        &user.Email,
        &user.CreatedAt,
        &user.UpdatedAt,
    )

    if err != nil {
        return nil, err
    }

    return user, nil
}
```

## Database Migrations

### Migration System
```go
type Migration struct {
    ID        int       `db:"id"`
    Name      string    `db:"name"`
    AppliedAt time.Time `db:"applied_at"`
}

func (db *Database) RunMigrations() error {
    // Create migrations table if it doesn't exist
    createMigrationsTable := `
        CREATE TABLE IF NOT EXISTS migrations (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )`

    _, err := db.db.Exec(createMigrationsTable)
    if err != nil {
        return err
    }

    // Define migrations
    migrations := []struct {
        name string
        sql  string
    }{
        {
            name: "001_create_users_table",
            sql: `
                CREATE TABLE IF NOT EXISTS users (
                    id SERIAL PRIMARY KEY,
                    username VARCHAR(50) UNIQUE NOT NULL,
                    email VARCHAR(100) UNIQUE NOT NULL,
                    password_hash VARCHAR(255) NOT NULL,
                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
                )`,
        },
        {
            name: "002_create_products_table",
            sql: `
                CREATE TABLE IF NOT EXISTS products (
                    id SERIAL PRIMARY KEY,
                    name VARCHAR(100) NOT NULL,
                    description TEXT,
                    price DECIMAL(10,2) NOT NULL,
                    category_id INTEGER,
                    stock INTEGER DEFAULT 0,
                    active BOOLEAN DEFAULT true,
                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
                )`,
        },
    }

    // Apply migrations
    for _, migration := range migrations {
        if !db.isMigrationApplied(migration.name) {
            err := db.applyMigration(migration.name, migration.sql)
            if err != nil {
                return err
            }
        }
    }

    return nil
}

func (db *Database) isMigrationApplied(name string) bool {
    var count int
    query := `SELECT COUNT(*) FROM migrations WHERE name = $1`
    err := db.db.QueryRow(query, name).Scan(&count)
    return err == nil && count > 0
}

func (db *Database) applyMigration(name, sql string) error {
    tx, err := db.db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    // Execute migration SQL
    _, err = tx.Exec(sql)
    if err != nil {
        return err
    }

    // Record migration
    _, err = tx.Exec(`INSERT INTO migrations (name) VALUES ($1)`, name)
    if err != nil {
        return err
    }

    return tx.Commit()
}
```

## Query Building

### Dynamic Query Builder
```go
type QueryBuilder struct {
    table     string
    selectCols []string
    where     []string
    args      []interface{}
    orderBy   string
    limit     int
    offset    int
}

func NewQueryBuilder(table string) *QueryBuilder {
    return &QueryBuilder{
        table:     table,
        selectCols: []string{"*"},
        where:     []string{},
        args:      []interface{}{},
    }
}

func (qb *QueryBuilder) Select(cols ...string) *QueryBuilder {
    qb.selectCols = cols
    return qb
}

func (qb *QueryBuilder) Where(condition string, args ...interface{}) *QueryBuilder {
    qb.where = append(qb.where, condition)
    qb.args = append(qb.args, args...)
    return qb
}

func (qb *QueryBuilder) OrderBy(orderBy string) *QueryBuilder {
    qb.orderBy = orderBy
    return qb
}

func (qb *QueryBuilder) Limit(limit int) *QueryBuilder {
    qb.limit = limit
    return qb
}

func (qb *QueryBuilder) Offset(offset int) *QueryBuilder {
    qb.offset = offset
    return qb
}

func (qb *QueryBuilder) Build() (string, []interface{}) {
    query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(qb.selectCols, ", "), qb.table)

    if len(qb.where) > 0 {
        query += " WHERE " + strings.Join(qb.where, " AND ")
    }

    if qb.orderBy != "" {
        query += " ORDER BY " + qb.orderBy
    }

    if qb.limit > 0 {
        query += fmt.Sprintf(" LIMIT %d", qb.limit)
    }

    if qb.offset > 0 {
        query += fmt.Sprintf(" OFFSET %d", qb.offset)
    }

    return query, qb.args
}
```

## Connection Pooling

### Advanced Connection Pool
```go
type ConnectionPool struct {
    db *sql.DB
    maxOpen int
    maxIdle int
    maxLifetime time.Duration
}

func NewConnectionPool(databaseURL string, maxOpen, maxIdle int, maxLifetime time.Duration) (*ConnectionPool, error) {
    db, err := sql.Open("postgres", databaseURL)
    if err != nil {
        return nil, err
    }

    db.SetMaxOpenConns(maxOpen)
    db.SetMaxIdleConns(maxIdle)
    db.SetConnMaxLifetime(maxLifetime)

    // Test connection
    if err := db.Ping(); err != nil {
        return nil, err
    }

    return &ConnectionPool{
        db:          db,
        maxOpen:     maxOpen,
        maxIdle:     maxIdle,
        maxLifetime: maxLifetime,
    }, nil
}

func (cp *ConnectionPool) GetStats() sql.DBStats {
    return cp.db.Stats()
}

func (cp *ConnectionPool) Close() error {
    return cp.db.Close()
}
```

## Error Handling

### Database Error Handling
```go
import (
    "errors"
    "github.com/lib/pq"
)

func handleDatabaseError(err error) error {
    if err == nil {
        return nil
    }

    // Check for specific PostgreSQL errors
    if pqErr, ok := err.(*pq.Error); ok {
        switch pqErr.Code {
        case "23505": // unique_violation
            return errors.New("duplicate entry")
        case "23503": // foreign_key_violation
            return errors.New("referenced record not found")
        case "23502": // not_null_violation
            return errors.New("required field is missing")
        case "42P01": // undefined_table
            return errors.New("table does not exist")
        default:
            return fmt.Errorf("database error: %s", pqErr.Message)
        }
    }

    // Check for sql.ErrNoRows
    if errors.Is(err, sql.ErrNoRows) {
        return errors.New("record not found")
    }

    return err
}
```

## Example Complete Application

```go
package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "os"
    "time"

    _ "github.com/lib/pq"
)

type App struct {
    db *Database
}

type User struct {
    ID        int       `json:"id" db:"id"`
    Username  string    `json:"username" db:"username"`
    Email     string    `json:"email" db:"email"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Database struct {
    db *sql.DB
}

func NewDatabase(databaseURL string) (*Database, error) {
    db, err := sql.Open("postgres", databaseURL)
    if err != nil {
        return nil, err
    }

    // Configure connection pool
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(5)
    db.SetConnMaxLifetime(5 * time.Minute)

    // Test connection
    if err := db.Ping(); err != nil {
        return nil, err
    }

    return &Database{db: db}, nil
}

func (d *Database) CreateUser(user *User) error {
    query := `
        INSERT INTO users (username, email, created_at, updated_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id`

    now := time.Now()
    user.CreatedAt = now
    user.UpdatedAt = now

    return d.db.QueryRow(
        query,
        user.Username,
        user.Email,
        user.CreatedAt,
        user.UpdatedAt,
    ).Scan(&user.ID)
}

func (d *Database) GetUser(id int) (*User, error) {
    user := &User{}
    query := `
        SELECT id, username, email, created_at, updated_at
        FROM users WHERE id = $1`

    err := d.db.QueryRow(query, id).Scan(
        &user.ID,
        &user.Username,
        &user.Email,
        &user.CreatedAt,
        &user.UpdatedAt,
    )

    if err != nil {
        return nil, err
    }

    return user, nil
}

func (d *Database) GetUsers(limit, offset int) ([]*User, error) {
    query := `
        SELECT id, username, email, created_at, updated_at
        FROM users
        ORDER BY created_at DESC
        LIMIT $1 OFFSET $2`

    rows, err := d.db.Query(query, limit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []*User
    for rows.Next() {
        user := &User{}
        err := rows.Scan(
            &user.ID,
            &user.Username,
            &user.Email,
            &user.CreatedAt,
            &user.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, rows.Err()
}

func (app *App) handleCreateUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := app.db.CreateUser(&user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func (app *App) handleGetUser(w http.ResponseWriter, r *http.Request) {
    // Extract user ID from URL
    // Implementation here

    user, err := app.db.GetUser(1) // Example
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func (app *App) handleGetUsers(w http.ResponseWriter, r *http.Request) {
    users, err := app.db.GetUsers(10, 0) // Example: limit 10, offset 0
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func main() {
    databaseURL := os.Getenv("DATABASE_URL")
    if databaseURL == "" {
        databaseURL = "postgres://user:password@localhost:5432/mydb?sslmode=disable"
    }

    db, err := NewDatabase(databaseURL)
    if err != nil {
        log.Fatal(err)
    }
    defer db.db.Close()

    app := &App{db: db}

    // Set up routes
    http.HandleFunc("/users", app.handleCreateUser)
    http.HandleFunc("/users/", app.handleGetUser)
    http.HandleFunc("/users", app.handleGetUsers)

    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Best Practices

### 1. Use Connection Pooling
- Configure appropriate pool sizes
- Monitor pool statistics
- Handle connection timeouts

### 2. Use Prepared Statements
- Prepare statements once, execute many times
- Prevent SQL injection
- Improve performance

### 3. Handle Transactions Properly
- Use transactions for multi-step operations
- Always defer Rollback()
- Commit only when all operations succeed

### 4. Validate Input Data
- Validate data before database operations
- Use parameterized queries
- Handle database constraints

### 5. Implement Proper Error Handling
- Check for specific database errors
- Provide meaningful error messages
- Log errors for debugging

## Next Steps

Now that you understand database applications, try the tasks in `tasks.md`!