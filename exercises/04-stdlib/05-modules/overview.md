# Go Modules - Overview

## What are Go Modules?

Go Modules is the official dependency management system for Go, introduced in Go 1.11. It provides a way to manage dependencies, version them, and ensure reproducible builds.

## Why Go Modules?

### Before Go Modules (GOPATH)
- All code had to be in `$GOPATH/src/`
- No versioning of dependencies
- Difficult to manage multiple projects
- Problems with dependency conflicts

### With Go Modules
- Project can be anywhere on your system
- Precise version control of dependencies
- Reproducible builds
- Better security with checksums

## Basic Module Structure

### go.mod File
```go
module go-playground

go 1.24

require (
    github.com/fatih/color v1.15.0
    github.com/spf13/cobra v1.7.0
)

require (
    github.com/inconshreveable/mousetrap v1.1.0 // indirect
)
```

### go.sum File
```
github.com/fatih/color v1.15.0 h1:3o6QX21zSF8jQfPiQOyPV6lV6+6LyctPg6K2wxW7H0=
github.com/fatih/color v1.15.0/go.mod h1:6DVwH7UNuAR8otU/p5TLpV9vcqMQ9zqYvYLzD2eEHlM=
github.com/spf13/cobra v1.7.0 h1:hyqWnYt1ZQShIddO5kBpj3vu05/++x6tJ6dg8ECdwI=
github.com/spf13/cobra v1.7.0/go.mod h1:iW0MLG+71LWzyVZ7rhIr0SXkfasgFdfSp/7YDHfrYs=
```

## Module Commands

### Initialize a Module
```bash
# Create a new module
go mod init myproject

# Initialize in existing project
go mod init github.com/username/project
```

### Add Dependencies
```bash
# Add a dependency
go get github.com/fatih/color

# Add specific version
go get github.com/fatih/color@v1.15.0

# Add multiple dependencies
go get github.com/spf13/cobra github.com/fatih/color
```

### Manage Dependencies
```bash
# Download all dependencies
go mod download

# Clean up unused dependencies
go mod tidy

# Verify dependencies
go mod verify

# Vendor dependencies (copy to vendor/ directory)
go mod vendor
```

## Module Paths

### Module Naming
```go
// Local module
module myproject

// GitHub module
module github.com/username/project

// Company module
module company.com/team/project
```

### Import Paths
```go
import (
    "fmt"                    // Standard library
    "github.com/fatih/color" // External dependency
    "myproject/internal"     // Local module
)
```

## Version Management

### Semantic Versioning
```go
require (
    github.com/fatih/color v1.15.0    // Exact version
    github.com/spf13/cobra v1.7.0     // Exact version
)
```

## Working with Modules

### Creating a New Project
```bash
# Create project directory
mkdir myproject
cd myproject

# Initialize module
go mod init myproject

# Create main.go
echo 'package main

import (
    "fmt"
    "github.com/fatih/color"
)

func main() {
    color.Green("Hello, Modules!")
}' > main.go

# Run the program
go run main.go
```

### Adding Dependencies
```bash
# Add color library
go get github.com/fatih/color

# Add CLI framework
go get github.com/spf13/cobra

# Check go.mod
cat go.mod
```

### Building and Running
```bash
# Build the project
go build

# Run the project
go run main.go

# Test the project
go test ./...
```

## Best Practices

### 1. Module Naming
- Use descriptive module names
- Follow the convention: `domain.com/username/project`
- Avoid generic names like `project` or `app`

### 2. Dependency Management
- Use `go mod tidy` regularly
- Pin specific versions for production
- Review `go.sum` for security

### 3. Project Structure
```
myproject/
├── go.mod          # Module definition
├── go.sum          # Dependency checksums
├── main.go         # Main application
├── internal/       # Private packages
├── pkg/           # Public packages
└── cmd/           # Application entry points
```

### 4. Version Control
- Commit both `go.mod` and `go.sum`
- Use semantic versioning for releases
- Tag releases with `git tag v1.0.0`

## Common Issues and Solutions

### Issue: Module not found
```bash
# Solution: Initialize module
go mod init myproject
```

### Issue: Dependencies not downloaded
```bash
# Solution: Download dependencies
go mod download
```

### Issue: Unused dependencies
```bash
# Solution: Clean up
go mod tidy
```

### Issue: Version conflicts
```bash
# Solution: Update dependencies
go get -u ./...
```

## Next Steps

Now that you understand Go Modules, try the tasks in `tasks.md`!