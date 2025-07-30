# Worker Pools - Tasks

## Task 1: Basic Worker Pool
Create a basic worker pool.

**Requirements:**
- Create a Job struct with ID and Data
- Create a simple worker pool with 3 workers
- Process jobs concurrently
- Demonstrate basic worker pool functionality

## Task 2: Worker Pool with Results
Create a worker pool that returns results.

**Requirements:**
- Create a Result struct
- Modify worker pool to collect results
- Return results from all workers
- Demonstrate result collection

## Task 3: Worker Pool with Context
Use context for graceful shutdown.

**Requirements:**
- Create a worker pool that accepts context
- Handle context cancellation
- Gracefully shutdown workers
- Demonstrate context-based control

## Task 4: Rate Limited Worker Pool
Create a worker pool with rate limiting.

**Requirements:**
- Implement rate limiting using time.Ticker
- Control job processing rate
- Demonstrate rate limiting behavior
- Show controlled throughput

## Task 5: Worker Pool with Error Handling
Handle errors in worker pool.

**Requirements:**
- Create jobs that may fail
- Handle worker errors gracefully
- Collect error information
- Demonstrate error handling

## Task 6: Practical Application
Create a web scraper using worker pool.

**Requirements:**
- Create a worker pool for web scraping
- Simulate fetching URLs
- Collect results from all workers
- Demonstrate practical usage

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

You should see output demonstrating all worker pool concepts.