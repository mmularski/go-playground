# Select Statement - Tasks

## Task 1: Basic Select
Create a basic select statement.

**Requirements:**
- Create a channel
- Use select to check if data is available
- Include a default case
- Demonstrate non-blocking behavior

## Task 2: Select with Multiple Channels
Use select with multiple channels.

**Requirements:**
- Create two channels
- Launch goroutines that send to different channels
- Use select to handle both channels
- Demonstrate random selection when both are ready

## Task 3: Select with Timeout
Use select with timeout functionality.

**Requirements:**
- Create a channel that receives data after a delay
- Use time.After() for timeout
- Handle both successful receive and timeout cases
- Demonstrate timeout behavior

## Task 4: Non-blocking Operations
Use select for non-blocking channel operations.

**Requirements:**
- Create channels that may or may not have data
- Use select with default to avoid blocking
- Handle multiple channel operations
- Demonstrate graceful handling of empty channels

## Task 5: Select in Loop
Use select in a loop for continuous monitoring.

**Requirements:**
- Create multiple channels
- Use select in a loop to monitor all channels
- Handle channel closing
- Demonstrate continuous operation

## Task 6: Practical Application
Create a simple load balancer using select.

**Requirements:**
- Create multiple worker channels
- Use select to distribute work
- Handle worker availability
- Demonstrate load balancing behavior

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

You should see output demonstrating all select concepts.