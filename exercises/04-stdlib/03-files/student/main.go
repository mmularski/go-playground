package main

import (
	"fmt"
)

// Task 1: Basic file reading and writing
func basicFileOperations() {
	// TODO: Create a file with some content:
	// filename := "test.txt"
	// content := "Hello, this is a test file!\nSecond line of content."
	// TODO: Write file: err := os.WriteFile(filename, []byte(content), 0644)
	// TODO: Handle file reading/writing errors: if err != nil { fmt.Printf("Error writing file: %v\n", err); return }
	// TODO: Read file: data, err := os.ReadFile(filename)
	// TODO: Display file contents: fmt.Printf("File content:\n%s\n", string(data))
}

// Task 2: File information and statistics
func fileInfoOperations() {
	// TODO: Get file information (size, permissions, modification time):
	// filename := "test.txt"
	// info, err := os.Stat(filename)
	// if err != nil { fmt.Printf("Error getting file info: %v\n", err); return }
	// TODO: Display file statistics:
	// fmt.Printf("Name: %s\n", info.Name())
	// fmt.Printf("Size: %d bytes\n", info.Size())
	// fmt.Printf("Mode: %s\n", info.Mode())
	// fmt.Printf("ModTime: %s\n", info.ModTime())
	// fmt.Printf("IsDir: %t\n", info.IsDir())
}

// Task 3: Directory operations
func directoryOperations() {
	// TODO: Create directories and nested directories:
	// err := os.Mkdir("testdir", 0755)
	// if err != nil { fmt.Printf("Error creating directory: %v\n", err) }
	// err = os.MkdirAll("testdir/nested/deep", 0755)
	// if err != nil { fmt.Printf("Error creating nested directories: %v\n", err) }
	// TODO: List directory contents:
	// entries, err := os.ReadDir("testdir")
	// if err != nil { fmt.Printf("Error reading directory: %v\n", err); return }
	// for _, entry := range entries { fmt.Printf("Name: %s, IsDir: %t\n", entry.Name(), entry.IsDir()) }
}

// Task 4: File path operations
func pathOperations() {
	// TODO: Join and split file paths:
	// fullPath := filepath.Join("dir", "subdir", "file.txt")
	// fmt.Printf("Full path: %s\n", fullPath)
	// TODO: Get directory and filename:
	// dir := filepath.Dir(fullPath)
	// filename := filepath.Base(fullPath)
	// fmt.Printf("Directory: %s, Filename: %s\n", dir, filename)
	// TODO: Get file extension: ext := filepath.Ext(fullPath)
	// TODO: Clean path: cleanPath := filepath.Clean("/path//to///file.txt")
	// fmt.Printf("Extension: %s\n", ext)
	// fmt.Printf("Clean path: %s\n", cleanPath)
}

// Task 5: Temporary files and cleanup
func temporaryFileOperations() {
	// TODO: Create temporary files and directories:
	// tempFile, err := os.CreateTemp("", "prefix-*.txt")
	// if err != nil { fmt.Printf("Error creating temp file: %v\n", err); return }
	// defer os.Remove(tempFile.Name()) // Clean up
	// TODO: Write data to temporary files:
	// fmt.Printf("Temp file: %s\n", tempFile.Name())
	// tempFile.WriteString("Temporary content\n")
	// tempFile.Close()
	// TODO: Create temporary directory:
	// tempDir, err := os.MkdirTemp("", "tempdir-*")
	// if err != nil { fmt.Printf("Error creating temp dir: %v\n", err); return }
	// defer os.RemoveAll(tempDir) // Clean up
	// fmt.Printf("Temp directory: %s\n", tempDir)
}

// Task 6: File copying and moving
func fileCopyingOperations() {
	// TODO: Copy files from source to destination:
	// sourceFile := "source.txt"
	// destFile := "destination.txt"
	// TODO: Create source file: os.WriteFile(sourceFile, []byte("Source file content"), 0644)
	// defer os.Remove(sourceFile)
	// TODO: Handle file copying errors:
	// source, err := os.Open(sourceFile)
	// if err != nil { fmt.Printf("Error opening source file: %v\n", err); return }
	// defer source.Close()
	// destination, err := os.Create(destFile)
	// if err != nil { fmt.Printf("Error creating destination file: %v\n", err); return }
	// defer destination.Close()
	// TODO: Copy file: _, err = io.Copy(destination, source)
	// if err != nil { fmt.Printf("Error copying file: %v\n", err); return }
	// fmt.Printf("File copied from %s to %s\n", sourceFile, destFile)
	// defer os.Remove(destFile)
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
