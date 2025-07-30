package main

import (
	"sync"
	"time"
)

// Task 1: Basic goroutines
func printMessage(msg string) {
	// TODO: Print the message
}

// Task 2: Goroutines with sleep
func delayedPrint(msg string, delay time.Duration) {
	// TODO: Sleep and print the delayed message
}

// Task 3: WaitGroup
func worker(id int, wg *sync.WaitGroup) {
	// TODO: Implement worker with WaitGroup
}

// Task 4: Shared data safety
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Increment() {
	// TODO: Safely increment the counter
}

func (c *SafeCounter) GetCount() int {
	// TODO: Safely get the count
	return 0
}

// Task 5: Goroutine communication
func producer(ch chan<- int) {
	// TODO: Send numbers to the channel
}

func consumer(ch <-chan int) {
	// TODO: Receive and print numbers from the channel
}

// Task 6: Practical application
func fetchURL(url string, wg *sync.WaitGroup, results chan<- string) {
	// TODO: Simulate fetching a URL
}

func main() {
	// Task 1: Basic goroutines
	// TODO: Launch printMessage as a goroutine

	// Task 2: Goroutines with sleep
	// TODO: Launch multiple delayedPrint goroutines

	// Task 3: WaitGroup
	// TODO: Use WaitGroup with multiple workers

	// Task 4: Shared data safety
	// TODO: Test SafeCounter with multiple goroutines

	// Task 5: Goroutine communication
	// TODO: Use channels for communication

	// Task 6: Practical application
	// TODO: Fetch multiple URLs concurrently
}
