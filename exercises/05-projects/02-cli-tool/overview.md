# CLI Tool Project - Overview

## What is a CLI Tool?

A Command Line Interface (CLI) tool is a program that runs in the terminal and accepts commands and arguments to perform specific tasks.

## CLI Tool Structure

### Basic CLI Structure
```go
package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    // Parse command line flags
    name := flag.String("name", "World", "Name to greet")
    flag.Parse()

    // Execute the command
    fmt.Printf("Hello, %s!\n", *name)
}
```

## Command Line Arguments

### Using flag Package
```go
import "flag"

func main() {
    // String flags
    name := flag.String("name", "default", "description")
    age := flag.Int("age", 0, "age description")
    verbose := flag.Bool("verbose", false, "verbose mode")

    // Parse flags
    flag.Parse()

    // Use the values
    fmt.Printf("Name: %s, Age: %d, Verbose: %t\n", *name, *age, *verbose)
}
```

### Using os.Args
```go
import "os"

func main() {
    args := os.Args[1:] // Skip program name

    if len(args) == 0 {
        fmt.Println("No arguments provided")
        return
    }

    for i, arg := range args {
        fmt.Printf("Argument %d: %s\n", i, arg)
    }
}
```

## Subcommands Pattern

### Implementing Subcommands
```go
package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Expected 'create', 'read', 'update', or 'delete' subcommands")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "create":
        createCmd := flag.NewFlagSet("create", flag.ExitOnError)
        name := createCmd.String("name", "", "Name of the item")
        createCmd.Parse(os.Args[2:])

        if *name == "" {
            fmt.Println("Name is required")
            os.Exit(1)
        }

        fmt.Printf("Creating item: %s\n", *name)

    case "read":
        readCmd := flag.NewFlagSet("read", flag.ExitOnError)
        id := readCmd.String("id", "", "ID of the item")
        readCmd.Parse(os.Args[2:])

        if *id == "" {
            fmt.Println("ID is required")
            os.Exit(1)
        }

        fmt.Printf("Reading item with ID: %s\n", *id)

    default:
        fmt.Printf("Unknown subcommand: %s\n", os.Args[1])
        os.Exit(1)
    }
}
```

## File Operations

### Reading Files
```go
import (
    "bufio"
    "os"
    "strings"
)

func readFile(filename string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    return lines, scanner.Err()
}
```

### Writing Files
```go
import (
    "fmt"
    "os"
)

func writeFile(filename string, content string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = file.WriteString(content)
    return err
}
```

## Configuration Management

### Environment Variables
```go
import "os"

func getConfig() {
    // Get environment variables
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default value
    }

    databaseURL := os.Getenv("DATABASE_URL")
    if databaseURL == "" {
        databaseURL = "localhost:5432" // Default value
    }

    fmt.Printf("Port: %s, Database: %s\n", port, databaseURL)
}
```

### Configuration Files
```go
import (
    "encoding/json"
    "os"
)

type Config struct {
    Port     string `json:"port"`
    Database string `json:"database"`
    LogLevel string `json:"log_level"`
}

func loadConfig(filename string) (*Config, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var config Config
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&config)
    return &config, err
}
```

## Input/Output Handling

### Reading from stdin
```go
import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func readFromStdin() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Enter text (Ctrl+D to end):")

    for scanner.Scan() {
        line := scanner.Text()
        if strings.TrimSpace(line) == "" {
            break
        }
        fmt.Printf("You entered: %s\n", line)
    }
}
```

### Writing to stdout/stderr
```go
import (
    "fmt"
    "os"
)

func outputExample() {
    // Write to stdout
    fmt.Println("This goes to stdout")

    // Write to stderr
    fmt.Fprintf(os.Stderr, "This goes to stderr\n")

    // Write formatted output
    fmt.Printf("Formatted output: %s\n", "example")
}
```

## Progress Indicators

### Simple Progress Bar
```go
import (
    "fmt"
    "time"
)

func showProgress() {
    for i := 0; i <= 100; i += 10 {
        fmt.Printf("\rProgress: %d%%", i)
        time.Sleep(100 * time.Millisecond)
    }
    fmt.Println() // New line at the end
}
```

### Spinner
```go
import (
    "fmt"
    "time"
)

func showSpinner() {
    spinner := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

    for i := 0; i < 20; i++ {
        fmt.Printf("\r%s Processing...", spinner[i%len(spinner)])
        time.Sleep(100 * time.Millisecond)
    }
    fmt.Println("\rDone!")
}
```

## Color Output

### Using ANSI Colors
```go
import "fmt"

const (
    colorRed    = "\033[31m"
    colorGreen  = "\033[32m"
    colorYellow = "\033[33m"
    colorBlue   = "\033[34m"
    colorReset  = "\033[0m"
)

func coloredOutput() {
    fmt.Printf("%sError: %sSomething went wrong%s\n", colorRed, colorYellow, colorReset)
    fmt.Printf("%sSuccess: %sOperation completed%s\n", colorGreen, colorBlue, colorReset)
}
```

## Error Handling

### Proper Error Handling
```go
import (
    "fmt"
    "os"
)

func handleErrors() {
    // Check if file exists
    if _, err := os.Stat("nonexistent.txt"); os.IsNotExist(err) {
        fmt.Fprintf(os.Stderr, "Error: File does not exist\n")
        os.Exit(1)
    }

    // Handle other errors
    if err := someOperation(); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}
```

## Interactive CLI

### Confirming Actions
```go
import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func confirmAction(message string) bool {
    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("%s (y/N): ", message)

    response, err := reader.ReadString('\n')
    if err != nil {
        return false
    }

    response = strings.ToLower(strings.TrimSpace(response))
    return response == "y" || response == "yes"
}
```

### Menu Selection
```go
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func showMenu() int {
    fmt.Println("Select an option:")
    fmt.Println("1. Create")
    fmt.Println("2. Read")
    fmt.Println("3. Update")
    fmt.Println("4. Delete")
    fmt.Println("5. Exit")

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your choice: ")

    choice, err := reader.ReadString('\n')
    if err != nil {
        return 0
    }

    choice = strings.TrimSpace(choice)
    num, err := strconv.Atoi(choice)
    if err != nil {
        return 0
    }

    return num
}
```

## Data Processing

### CSV Processing
```go
import (
    "encoding/csv"
    "os"
)

func processCSV(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return err
    }

    for i, record := range records {
        if i == 0 {
            continue // Skip header
        }
        fmt.Printf("Row %d: %v\n", i, record)
    }

    return nil
}
```

### JSON Processing
```go
import (
    "encoding/json"
    "fmt"
    "os"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func processJSON(filename string) error {
    data, err := os.ReadFile(filename)
    if err != nil {
        return err
    }

    var people []Person
    err = json.Unmarshal(data, &people)
    if err != nil {
        return err
    }

    for _, person := range people {
        fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
    }

    return nil
}
```

## Example CLI Tool

```go
package main

import (
    "flag"
    "fmt"
    "os"
    "strings"
)

type CLI struct {
    verbose bool
    output  string
}

func NewCLI() *CLI {
    cli := &CLI{}

    flag.BoolVar(&cli.verbose, "verbose", false, "Enable verbose output")
    flag.StringVar(&cli.output, "output", "stdout", "Output destination")
    flag.Parse()

    return cli
}

func (c *CLI) Run() error {
    args := flag.Args()

    if len(args) == 0 {
        return fmt.Errorf("no command specified")
    }

    command := args[0]
    commandArgs := args[1:]

    switch command {
    case "echo":
        return c.handleEcho(commandArgs)
    case "count":
        return c.handleCount(commandArgs)
    case "reverse":
        return c.handleReverse(commandArgs)
    default:
        return fmt.Errorf("unknown command: %s", command)
    }
}

func (c *CLI) handleEcho(args []string) error {
    if c.verbose {
        fmt.Fprintf(os.Stderr, "Executing echo command\n")
    }

    output := strings.Join(args, " ")
    fmt.Println(output)
    return nil
}

func (c *CLI) handleCount(args []string) error {
    if c.verbose {
        fmt.Fprintf(os.Stderr, "Executing count command\n")
    }

    for _, arg := range args {
        fmt.Printf("%s: %d characters\n", arg, len(arg))
    }
    return nil
}

func (c *CLI) handleReverse(args []string) error {
    if c.verbose {
        fmt.Fprintf(os.Stderr, "Executing reverse command\n")
    }

    for _, arg := range args {
        runes := []rune(arg)
        for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
            runes[i], runes[j] = runes[j], runes[i]
        }
        fmt.Println(string(runes))
    }
    return nil
}

func main() {
    cli := NewCLI()

    if err := cli.Run(); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}
```

## Best Practices

### 1. Use Proper Exit Codes
```go
const (
    ExitSuccess = 0
    ExitError   = 1
    ExitUsage   = 2
)

func main() {
    if err := run(); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(ExitError)
    }
    os.Exit(ExitSuccess)
}
```

### 2. Provide Help Information
```go
func printUsage() {
    fmt.Println("Usage: mycli [flags] <command> [args]")
    fmt.Println()
    fmt.Println("Commands:")
    fmt.Println("  echo <text>     Echo the provided text")
    fmt.Println("  count <text>    Count characters in text")
    fmt.Println("  reverse <text>  Reverse the provided text")
    fmt.Println()
    fmt.Println("Flags:")
    flag.PrintDefaults()
}
```

### 3. Handle Signals Gracefully
```go
import (
    "os"
    "os/signal"
    "syscall"
)

func handleSignals() {
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        sig := <-sigChan
        fmt.Printf("\nReceived signal: %v\n", sig)
        os.Exit(0)
    }()
}
```

## Next Steps

Now that you understand CLI tools, try the tasks in `tasks.md`!