package main

import (
	"fmt"
)

// Task 1: Basic select
func basicSelect() {
	// TODO: Create a channel: ch := make(chan int)
	// TODO: Use select to check if data is available:
	// select {
	// case <-ch: fmt.Println("Received from ch")
	// default: fmt.Println("No data available")
	// }
}

// Task 2: Select with multiple channels
func multipleChannels() {
	// TODO: Create two channels: ch1 := make(chan string), ch2 := make(chan string)
	// TODO: Launch goroutines that send to different channels:
	// go func() { time.Sleep(time.Millisecond * 100); ch1 <- "from ch1" }()
	// go func() { time.Sleep(time.Millisecond * 200); ch2 <- "from ch2" }()
	// TODO: Use select to handle both channels in a loop (0 to 1):
	// select {
	// case msg1 := <-ch1: fmt.Printf("Received: %s\n", msg1)
	// case msg2 := <-ch2: fmt.Printf("Received: %s\n", msg2)
	// }
}

// Task 3: Select with timeout
func selectWithTimeout() {
	// TODO: Create a channel: ch := make(chan string)
	// TODO: Launch goroutine that sends data after delay:
	// go func() { time.Sleep(time.Second * 2); ch <- "result" }()
	// TODO: Use time.After() for timeout:
	// select {
	// case result := <-ch: fmt.Printf("Received: %s\n", result)
	// case <-time.After(time.Second): fmt.Println("Timeout")
	// }
}

// Task 4: Non-blocking operations
func nonBlockingSelect() {
	// TODO: Create channels: ch1 := make(chan int), ch2 := make(chan int)
	// TODO: Use select with default to avoid blocking:
	// select {
	// case ch1 <- 42: fmt.Println("Sent to ch1")
	// default: fmt.Println("ch1 is not ready")
	// }
	// TODO: Try to receive from ch2:
	// select {
	// case value := <-ch2: fmt.Printf("Received from ch2: %d\n", value)
	// default: fmt.Println("ch2 has no data")
	// }
}

// Task 5: Select in loop
func selectInLoop() {
	// TODO: Create multiple channels: ch1 := make(chan int), ch2 := make(chan int)
	// TODO: Launch goroutines that send data and close channels:
	// go func() {
	//     for i := 1; i <= 3; i++ { ch1 <- i; time.Sleep(time.Millisecond * 200) }
	//     close(ch1)
	// }()
	// go func() {
	//     for i := 10; i <= 12; i++ { ch2 <- i; time.Sleep(time.Millisecond * 300) }
	//     close(ch2)
	// }()
	// TODO: Use select in a loop to monitor all channels with channel closing check:
	// ch1Closed, ch2Closed := false, false
	// for !ch1Closed || !ch2Closed {
	//     select {
	//     case value, ok := <-ch1:
	//         if ok { fmt.Printf("Received from ch1: %d\n", value) } else { ch1Closed = true }
	//     case value, ok := <-ch2:
	//         if ok { fmt.Printf("Received from ch2: %d\n", value) } else { ch2Closed = true }
	//     }
	// }
}

// Task 6: Practical application
func loadBalancer() {
	// TODO: Create multiple worker channels: worker1 := make(chan string), worker2 := make(chan string)
	// TODO: Launch worker goroutines:
	// go func() {
	//     for i := 1; i <= 3; i++ {
	//         time.Sleep(time.Millisecond * 100)
	//         worker1 <- fmt.Sprintf("Worker1 processed job %d", i)
	//     }
	// }()
	// go func() {
	//     for i := 1; i <= 3; i++ {
	//         time.Sleep(time.Millisecond * 150)
	//         worker2 <- fmt.Sprintf("Worker2 processed job %d", i)
	//     }
	// }()
	// TODO: Use select to distribute work in a loop (0 to 5):
	// select {
	// case result := <-worker1: fmt.Printf("Load balancer: %s\n", result)
	// case result := <-worker2: fmt.Printf("Load balancer: %s\n", result)
	// }
}

func main() {
	fmt.Println("=== Task 1: Basic Select ===")
	basicSelect()

	fmt.Println("\n=== Task 2: Select with Multiple Channels ===")
	multipleChannels()

	fmt.Println("\n=== Task 3: Select with Timeout ===")
	selectWithTimeout()

	fmt.Println("\n=== Task 4: Non-blocking Operations ===")
	nonBlockingSelect()

	fmt.Println("\n=== Task 5: Select in Loop ===")
	selectInLoop()

	fmt.Println("\n=== Task 6: Load Balancer ===")
	loadBalancer()
}
