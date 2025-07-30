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
	// TODO: Connect to SQLite database:
	// db, err := sql.Open("sqlite3", "./test.db")
	// if err != nil { return nil, fmt.Errorf("error opening database: %v", err) }
	// TODO: Test the connection: err = db.Ping()
	// TODO: Handle connection errors: if err != nil { return nil, fmt.Errorf("error connecting to database: %v", err) }
	// TODO: Return database connection: return db, nil
	return nil, nil
}

// Task 2: Creating tables
func createTables(db *sql.DB) error {
	// TODO: Create users table with appropriate columns:
	// createTableSQL := `
	// CREATE TABLE IF NOT EXISTS users (
	//     id INTEGER PRIMARY KEY AUTOINCREMENT,
	//     name TEXT NOT NULL,
	//     email TEXT UNIQUE,
	//     age INTEGER
	// );`
	// TODO: Handle table creation errors: _, err := db.Exec(createTableSQL)
	// TODO: Return error: return err
	return nil
}

// Task 3: Inserting data
func insertUser(db *sql.DB, name, email string, age int) error {
	// TODO: Insert single record:
	// insertSQL := "INSERT INTO users (name, email, age) VALUES (?, ?, ?)"
	// result, err := db.Exec(insertSQL, name, email, age)
	// TODO: Handle insertion errors: if err != nil { return fmt.Errorf("error inserting user: %v", err) }
	// TODO: Get the last insert ID: id, err := result.LastInsertId()
	// TODO: Print result: fmt.Printf("Inserted user with ID: %d\n", id)
	// TODO: Return nil: return nil
	return nil
}

// Task 4: Querying data
func getUserByID(db *sql.DB, id int) (*User, error) {
	// TODO: Query single row:
	// var user User
	// query := "SELECT id, name, email, age FROM users WHERE id = ?"
	// err := db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Age)
	// TODO: Handle query errors: if err != nil { return nil, fmt.Errorf("error querying user: %v", err) }
	// TODO: Return user: return &user, nil
	return nil, nil
}

func getAllUsers(db *sql.DB) ([]User, error) {
	// TODO: Query multiple rows:
	// query := "SELECT id, name, email, age FROM users ORDER BY id"
	// rows, err := db.Query(query)
	// if err != nil { return nil, fmt.Errorf("error querying users: %v", err) }
	// defer rows.Close()
	// TODO: Process query results:
	// var users []User
	// for rows.Next() {
	//     var user User
	//     err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
	//     if err != nil { return nil, fmt.Errorf("error scanning user: %v", err) }
	//     users = append(users, user)
	// }
	// TODO: Check for iteration errors: if err = rows.Err(); err != nil { return nil, fmt.Errorf("error iterating rows: %v", err) }
	// TODO: Return users: return users, nil
	return nil, nil
}

// Task 5: Updating and deleting
func updateUser(db *sql.DB, id int, name, email string, age int) error {
	// TODO: Update existing record:
	// updateSQL := "UPDATE users SET name = ?, email = ?, age = ? WHERE id = ?"
	// result, err := db.Exec(updateSQL, name, email, age, id)
	// TODO: Handle update errors: if err != nil { return fmt.Errorf("error updating user: %v", err) }
	// TODO: Check affected rows: rowsAffected, err := result.RowsAffected()
	// TODO: Print result: fmt.Printf("Updated %d rows\n", rowsAffected)
	// TODO: Return nil: return nil
	return nil
}

func deleteUser(db *sql.DB, id int) error {
	// TODO: Delete record:
	// deleteSQL := "DELETE FROM users WHERE id = ?"
	// result, err := db.Exec(deleteSQL, id)
	// TODO: Handle delete errors: if err != nil { return fmt.Errorf("error deleting user: %v", err) }
	// TODO: Check affected rows: rowsAffected, err := result.RowsAffected()
	// TODO: Print result: fmt.Printf("Deleted %d rows\n", rowsAffected)
	// TODO: Return nil: return nil
	return nil
}

// Task 6: Transactions
func transferUserData(db *sql.DB, fromID, toID int) error {
	// TODO: Begin transaction: tx, err := db.Begin()
	// TODO: Handle transaction errors: if err != nil { return fmt.Errorf("error beginning transaction: %v", err) }
	// TODO: Defer rollback: defer tx.Rollback()
	// TODO: Execute multiple operations:
	// _, err = tx.Exec("UPDATE users SET name = name || ' (transferred)' WHERE id = ?", fromID)
	// if err != nil { return fmt.Errorf("error updating source user: %v", err) }
	// _, err = tx.Exec("UPDATE users SET name = name || ' (received)' WHERE id = ?", toID)
	// if err != nil { return fmt.Errorf("error updating target user: %v", err) }
	// TODO: Commit transaction: err = tx.Commit()
	// TODO: Handle commit errors: if err != nil { return fmt.Errorf("error committing transaction: %v", err) }
	// TODO: Return nil: return nil
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
