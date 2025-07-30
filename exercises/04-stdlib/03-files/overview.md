# File Operations - Overview

## What are File Operations?

File operations in Go are handled through the `os` and `io` packages. They provide functionality for reading, writing, creating, and managing files and directories.

## Basic File Operations

### Reading Files
```go
import (
    "fmt"
    "io"
    "os"
)

func readFile(filename string) {
    // Read entire file
    data, err := os.ReadFile(filename)
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
    fmt.Printf("File content: %s\n", string(data))
}
```

### Writing Files
```go
func writeFile(filename string, content string) {
    err := os.WriteFile(filename, []byte(content), 0644)
    if err != nil {
        fmt.Printf("Error writing file: %v\n", err)
        return
    }
    fmt.Println("File written successfully")
}
```

## File Handling with os.Open

```go
func readFileWithOpen(filename string) {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()

    // Read file in chunks
    buffer := make([]byte, 1024)
    for {
        n, err := file.Read(buffer)
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Printf("Error reading: %v\n", err)
            return
        }
        fmt.Printf("Read %d bytes: %s\n", n, buffer[:n])
    }
}
```

## File Creation and Appending

```go
func createAndAppendFile(filename string) {
    // Create new file
    file, err := os.Create(filename)
    if err != nil {
        fmt.Printf("Error creating file: %v\n", err)
        return
    }
    defer file.Close()

    // Write initial content
    file.WriteString("Hello, World!\n")

    // Append to file
    file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Printf("Error opening file for append: %v\n", err)
        return
    }
    defer file.Close()

    file.WriteString("This is appended content\n")
}
```

## Directory Operations

```go
import (
    "fmt"
    "os"
    "path/filepath"
)

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
```

## File Information

```go
func fileInfo(filename string) {
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
```

## File Path Operations

```go
import "path/filepath"

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
```

## Temporary Files

```go
import (
    "fmt"
    "os"
    "path/filepath"
)

func temporaryFiles() {
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
```

## File Copying

```go
import (
    "fmt"
    "io"
    "os"
)

func copyFile(src, dst string) error {
    sourceFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer sourceFile.Close()

    destFile, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer destFile.Close()

    _, err = io.Copy(destFile, sourceFile)
    return err
}
```

## Error Handling

```go
func robustFileOperations(filename string) {
    // Check if file exists
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        fmt.Printf("File %s does not exist\n", filename)
        return
    }

    // Try to open file
    file, err := os.Open(filename)
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()

    // Read with error handling
    data, err := io.ReadAll(file)
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }

    fmt.Printf("Successfully read %d bytes\n", len(data))
}
```

## Best Practices

1. **Always close files** - Use `defer file.Close()`
2. **Check errors** - Handle all error returns
3. **Use appropriate permissions** - Set correct file modes
4. **Clean up temporary files** - Remove temp files and dirs
5. **Use buffered I/O** - For large files, use buffered reading/writing
6. **Handle file existence** - Check if files exist before operations

## Example Program

```go
package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    // Create a test file
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

    // Get file info
    info, err := os.Stat(filename)
    if err != nil {
        fmt.Printf("Error getting file info: %v\n", err)
        return
    }

    fmt.Printf("File size: %d bytes\n", info.Size())

    // Clean up
    os.Remove(filename)
}
```

## Next Steps

Now that you understand file operations, try the tasks in `tasks.md`!