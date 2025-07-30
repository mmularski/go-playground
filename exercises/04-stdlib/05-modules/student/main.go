package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Task 1: Initialize module (done via command line)
// Run: go mod init myproject

// Task 2: Add dependencies (done via command line)
// Run: go get github.com/fatih/color github.com/spf13/cobra

// Task 3: Use color library
func useColorLibrary() {
	// TODO: Experiment with the color library
	fmt.Println("Color library demo:")
	fmt.Println("Hello")
	fmt.Println("World")
	fmt.Println("Go Modules")
	fmt.Println("Success!")
}

// Task 4: Create CLI application
func createCLI() {
	// TODO: Build a CLI application with cobra
	fmt.Println("CLI application demo:")
	fmt.Println("Commands: version, greet")
}

// Task 5: Module management (done via command line)
// Run these commands:
// go mod tidy
// go mod download
// go mod verify
// go mod vendor

// Task 6: Version management (done via command line)
// Run these commands:
// go get github.com/fatih/color@v1.15.0
// go get -u github.com/fatih/color
// go get github.com/fatih/color@v1.14.0

// Helper function for CLI version command
func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: Print version information
			fmt.Println("Version: 1.0.0")
		},
	}
}

// Helper function for CLI greet command
func greetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "greet [name]",
		Short: "Greet someone by name",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: Greet the person by name
			name := args[0]
			fmt.Printf("Hello, %s!\n", name)
		},
	}
}

func main() {
	fmt.Println("Go Modules Practice")

	// Task 1: Initialize module (done via command line)
	// go mod init myproject

	// Task 2: Add dependencies (done via command line)
	// go get github.com/fatih/color github.com/spf13/cobra

	// Task 3: Use color library
	useColorLibrary()

	// Task 4: Create CLI application
	createCLI()

	// Task 5: Module management (done via command line)
	// go mod tidy
	// go mod download
	// go mod verify
	// go mod vendor

	// Task 6: Version management (done via command line)
	// go get github.com/fatih/color@v1.15.0
	// go get -u github.com/fatih/color
	// go get github.com/fatih/color@v1.14.0
}
