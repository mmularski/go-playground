# Channels - Tasks

## Task 1: Basic Channel Operations
Create and use basic channels.

**Requirements:**
- Create an unbuffered channel
- Send data to the channel from a goroutine
- Receive data from the channel in main
- Demonstrate basic send/receive operations

## Task 2: Buffered Channels
Work with buffered channels.

**Requirements:**
- Create a buffered channel with capacity 3
- Send multiple values to the channel
- Receive all values from the channel
- Demonstrate the difference from unbuffered channels

## Task 3: Channel Direction
Use directional channels in functions.

**Requirements:**
- Create a function that accepts a send-only channel
- Create a function that accepts a receive-only channel
- Pass channels to these functions
- Demonstrate channel direction safety

## Task 4: Channel Closing
Work with channel closing and range loops.

**Requirements:**
- Create a goroutine that sends data and closes the channel
- Use a range loop to receive all data
- Handle channel closing properly
- Demonstrate the comma ok idiom

## Task 5: Select Statement
Use select to handle multiple channels.

**Requirements:**
- Create multiple channels
- Use select to handle different channel operations
- Include a default case
- Include a timeout case using time.After()

## Task 6: Practical Application
Create a simple pipeline using channels.

**Requirements:**
- Create a pipeline with multiple stages
- Use channels to pass data between stages
- Implement a simple data processing pipeline
- Demonstrate channel-based communication

---

## How to complete these tasks:

1. Navigate to the `student/` directory
2. Create your code in `main.go`
3. Run your code: `go run main.go`


## Testing Your Code

Run your program:
```bash
cd student/
go run main.go
```

You should see output demonstrating all channel concepts.