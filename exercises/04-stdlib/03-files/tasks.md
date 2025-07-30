# File Operations - Tasks

## Task 1: Basic File Reading and Writing
Create basic file operations.

**Requirements:**
- Create a file with some content
- Read the file and display its contents
- Handle file reading/writing errors
- Demonstrate basic file operations

## Task 2: File Information and Statistics
Work with file metadata.

**Requirements:**
- Get file information (size, permissions, modification time)
- Check if file exists
- List directory contents
- Display file statistics

## Task 3: Directory Operations
Work with directories.

**Requirements:**
- Create directories and nested directories
- List directory contents
- Navigate directory structure
- Handle directory operations

## Task 4: File Path Operations
Work with file paths.

**Requirements:**
- Join and split file paths
- Get file extensions
- Clean and normalize paths
- Work with relative and absolute paths

## Task 5: Temporary Files and Cleanup
Work with temporary files.

**Requirements:**
- Create temporary files and directories
- Write data to temporary files
- Clean up temporary resources
- Handle temporary file operations

## Task 6: File Copying and Moving
Implement file operations.

**Requirements:**
- Copy files from source to destination
- Handle file copying errors
- Implement file moving functionality
- Demonstrate file operations

---

## How to complete these tasks:

1. Navigate to the `student/` directory
2. Create your code in `main.go`
3. Run your code: `go run main.go`

## Example structure for student/main.go:
```go
package main

import (
    "fmt"
    "io"
    "os"
    "path/filepath"
)

// Task 1: Basic file reading and writing
func basicFileOperations() {
    filename := "test.txt"
    content := "Hello, this is a test file!\nSecond line of content."

    // Write file
    err := os.WriteFile(filename, []byte(content), 0644)
    if err != nil {
        fmt.Printf("Error writing file: %v\n", err)
        return
    }

    // Read file
    data, err := os.ReadFile(filename)
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }

    fmt.Printf("File content:\n%s\n", string(data))
}

// Task 2: File information and statistics
func fileInfoOperations() {
    filename := "test.txt"

    // Get file info
    info, err := os.Stat(filename)
    if err != nil {
        fmt.Printf("Error getting file info: %v\n", err)
        return
    }

    fmt.Printf("Name: %s\n", info.Name())
    fmt.Printf("Size: %d bytes\n", info.Size())
    fmt.Printf("Mode: %s\n", info.Mode())
    fmt.Printf("ModTime: %s\n", info.ModTime())
    fmt.Printf("IsDir: %t\n", info.IsDir())
}

// Task 3: Directory operations
func directoryOperations() {
    // Create directory
    err := os.Mkdir("testdir", 0755)
    if err != nil {
        fmt.Printf("Error creating directory: %v\n", err)
    }

    // Create nested directories
    err = os.MkdirAll("testdir/nested/deep", 0755)
    if err != nil {
        fmt.Printf("Error creating nested directories: %v\n", err)
    }

    // List directory contents
    entries, err := os.ReadDir("testdir")
    if err != nil {
        fmt.Printf("Error reading directory: %v\n", err)
        return
    }

    for _, entry := range entries {
        fmt.Printf("Name: %s, IsDir: %t\n", entry.Name(), entry.IsDir())
    }
}

// Task 4: File path operations
func pathOperations() {
    // Join paths
    fullPath := filepath.Join("dir", "subdir", "file.txt")
    fmt.Printf("Full path: %s\n", fullPath)

    // Get directory and filename
    dir := filepath.Dir(fullPath)
    filename := filepath.Base(fullPath)
    fmt.Printf("Directory: %s, Filename: %s\n", dir, filename)

    // Get file extension
    ext := filepath.Ext(fullPath)
    fmt.Printf("Extension: %s\n", ext)

    // Clean path
    cleanPath := filepath.Clean("/path//to///file.txt")
    fmt.Printf("Clean path: %s\n", cleanPath)
}

// Task 5: Temporary files and cleanup
func temporaryFileOperations() {
    // Create temporary file
    tempFile, err := os.CreateTemp("", "prefix-*.txt")
    if err != nil {
        fmt.Printf("Error creating temp file: %v\n", err)
        return
    }
    defer os.Remove(tempFile.Name()) // Clean up

    fmt.Printf("Temp file: %s\n", tempFile.Name())
    tempFile.WriteString("Temporary content\n")
    tempFile.Close()

    // Create temporary directory
    tempDir, err := os.MkdirTemp("", "tempdir-*")
    if err != nil {
        fmt.Printf("Error creating temp dir: %v\n", err)
        return
    }
    defer os.RemoveAll(tempDir) // Clean up

    fmt.Printf("Temp directory: %s\n", tempDir)
}

// Task 6: File copying and moving
func fileCopyingOperations() {
    sourceFile := "source.txt"
    destFile := "destination.txt"

    // Create source file
    os.WriteFile(sourceFile, []byte("Source file content"), 0644)
    defer os.Remove(sourceFile)

    // Copy file
    source, err := os.Open(sourceFile)
    if err != nil {
        fmt.Printf("Error opening source file: %v\n", err)
        return
    }
    defer source.Close()

    destination, err := os.Create(destFile)
    if err != nil {
        fmt.Printf("Error creating destination file: %v\n", err)
        return
    }
    defer destination.Close()

    _, err = io.Copy(destination, source)
    if err != nil {
        fmt.Printf("Error copying file: %v\n", err)
        return
    }

    fmt.Printf("File copied from %s to %s\n", sourceFile, destFile)
    defer os.Remove(destFile)
}

func main() {
    fmt.Println("=== Task 1: Basic File Operations ===")
    basicFileOperations()

    fmt.Println("\n=== Task 2: File Information ===")
    fileInfoOperations()

    fmt.Println("\n=== Task 3: Directory Operations ===")
    directoryOperations()

    fmt.Println("\n=== Task 4: Path Operations ===")
    pathOperations()

    fmt.Println("\n=== Task 5: Temporary Files ===")
    temporaryFileOperations()

    fmt.Println("\n=== Task 6: File Copying ===")
    fileCopyingOperations()
}
```

## Testing Your Code

Run your program:
```bash
cd student/
go run main.go
```

You should see output demonstrating all file operation concepts.