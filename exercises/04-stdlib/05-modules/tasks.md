# Go Modules - Tasks

## Task 1: Initialize a Module
Create and initialize a Go module.

**Requirements:**
- Create a new directory for your project
- Initialize a Go module with `go mod init`
- Create a simple main.go file
- Run the program to verify it works

## Task 2: Add Simple Dependencies
Add external dependencies to your module.

**Requirements:**
- Add a color library (github.com/fatih/color)
- Add a CLI framework (github.com/spf13/cobra)
- Use `go get` to download dependencies
- Verify dependencies in go.mod

## Task 3: Use Color Library
Create a colorful console application.

**Requirements:**
- Use fatih/color for colored output
- Print different messages in different colors
- Use various color methods (Green, Red, Yellow, Blue)
- Create a simple color demo

## Task 4: Create CLI Application
Build a basic CLI using cobra.

**Requirements:**
- Use spf13/cobra for CLI framework
- Create a root command
- Add subcommands (version, help)
- Handle command-line arguments

## Task 5: Module Management
Practice module management commands.

**Requirements:**
- Use `go mod tidy` to clean dependencies
- Use `go mod download` to download dependencies
- Use `go mod verify` to verify checksums
- Use `go mod vendor` to vendor dependencies

## Task 6: Version Management
Work with different dependency versions.

**Requirements:**
- Add a dependency with specific version
- Update dependencies to latest versions
- Downgrade a dependency to older version
- Handle version conflicts

---

## How to complete these tasks:

1. Navigate to the `student/` directory
2. Create your code in `main.go`
3. Run your code: `go run main.go`


## Commands to run:

### Task 1: Initialize module
```bash
cd student/
go mod init myproject
```

### Task 2: Add dependencies
```bash
go get github.com/fatih/color
go get github.com/spf13/cobra
```

### Task 3-4: Run your code
```bash
go run main.go
```

### Task 5: Module management
```bash
go mod tidy
go mod download
go mod verify
go mod vendor
```

### Task 6: Version management
```bash
go get github.com/fatih/color@v1.15.0
go get -u github.com/fatih/color
go get github.com/fatih/color@v1.14.0
```