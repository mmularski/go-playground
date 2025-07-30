package main

import (
	"fmt"
)

// Task 1: Basic channel operations
func basicChannel() {
	// TODO: Create an unbuffered channel: ch := make(chan int)
	// TODO: Send data to channel from goroutine:
	// go func() { ch <- 42 }()
	// TODO: Receive data from channel in main: value := <-ch
	// TODO: Print the value: fmt.Printf("Received: %d\n", value)
}

// Task 2: Buffered channels
func bufferedChannel() {
	// TODO: Create a buffered channel with capacity 3: ch := make(chan int, 3)
	// TODO: Send multiple values to the channel: ch <- 1, ch <- 2, ch <- 3
	// TODO: Receive all values from the channel: <-ch, <-ch, <-ch
	// TODO: Print each received value using fmt.Printf("Received: %d\n", value)
}

// Task 3: Channel direction
func sendOnly(ch chan<- int) {
	// TODO: Send numbers 1-5 to the channel using for loop
	// Use: for i := 1; i <= 5; i++ { ch <- i }
	// TODO: Close the channel when done using close(ch)
}

func receiveOnly(ch <-chan int) {
	// TODO: Receive and print all values from the channel using range loop
	// Use: for value := range ch { fmt.Printf("Received: %d\n", value) }
}

// Task 4: Channel closing
func channelClosing() {
	// TODO: Create an unbuffered channel: ch := make(chan int)
	// TODO: Launch goroutine that sends data and closes channel:
	// go func() {
	//     for i := 1; i <= 5; i++ {
	//         ch <- i
	//         time.Sleep(time.Millisecond * 100)
	//     }
	//     close(ch)
	// }()
	// TODO: Use range loop to receive all data: for value := range ch { ... }
	// TODO: Print each value: fmt.Printf("Received: %d\n", value)
}

// Task 5: Select statement
func selectExample() {
	// TODO: Create multiple channels: ch1 := make(chan string), ch2 := make(chan string)
	// TODO: Launch goroutines that send to different channels:
	// go func() { time.Sleep(time.Second); ch1 <- "from ch1" }()
	// go func() { time.Sleep(time.Second * 2); ch2 <- "from ch2" }()
	// TODO: Use select to handle different channel operations in a loop (0 to 1):
	// select {
	// case msg1 := <-ch1: fmt.Printf("Received: %s\n", msg1)
	// case msg2 := <-ch2: fmt.Printf("Received: %s\n", msg2)
	// case <-time.After(time.Second * 3): fmt.Println("Timeout")
	// }
}

// Task 6: Practical application
func pipeline() {
	// TODO: Create a pipeline with multiple stages
	// TODO: Stage 1: Generate numbers
	// numbers := make(chan int)
	// go func() {
	//     for i := 1; i <= 10; i++ { numbers <- i }
	//     close(numbers)
	// }()
	// TODO: Stage 2: Square numbers
	// squares := make(chan int)
	// go func() {
	//     for n := range numbers { squares <- n * n }
	//     close(squares)
	// }()
	// TODO: Stage 3: Print results
	// for square := range squares { fmt.Printf("Square: %d\n", square) }
}

func main() {
	fmt.Println("=== Task 1: Basic Channel Operations ===")
	basicChannel()

	fmt.Println("\n=== Task 2: Buffered Channels ===")
	bufferedChannel()

	fmt.Println("\n=== Task 3: Channel Direction ===")
	// TODO: Create channel: ch := make(chan int)
	// TODO: Launch sendOnly goroutine: go sendOnly(ch)
	// TODO: Call receiveOnly: receiveOnly(ch)

	fmt.Println("\n=== Task 4: Channel Closing ===")
	channelClosing()

	fmt.Println("\n=== Task 5: Select Statement ===")
	selectExample()

	fmt.Println("\n=== Task 6: Pipeline ===")
	pipeline()
}
