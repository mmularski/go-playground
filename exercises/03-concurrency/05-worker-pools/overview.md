# Worker Pools - Overview

## What are Worker Pools?

Worker pools are a concurrency pattern where a fixed number of goroutines (workers) process jobs from a queue. This pattern is useful for controlling resource usage and managing concurrent operations.

## Basic Worker Pool Structure

```go
type Job struct {
    ID   int
    Data string
}

type Worker struct {
    ID       int
    JobQueue chan Job
    Quit     chan bool
}

func (w *Worker) Start() {
    go func() {
        for {
            select {
            case job := <-w.JobQueue:
                // Process job
                fmt.Printf("Worker %d processing job %d\n", w.ID, job.ID)
            case <-w.Quit:
                return
            }
        }
    }()
}
```

## Simple Worker Pool

```go
func simpleWorkerPool(numWorkers int, jobs []Job) {
    jobQueue := make(chan Job, len(jobs))

    // Start workers
    for i := 0; i < numWorkers; i++ {
        go func(workerID int) {
            for job := range jobQueue {
                fmt.Printf("Worker %d processing job %d\n", workerID, job.ID)
                time.Sleep(time.Millisecond * 100) // Simulate work
            }
        }(i)
    }

    // Send jobs
    for _, job := range jobs {
        jobQueue <- job
    }
    close(jobQueue)
}
```

## Worker Pool with Results

```go
type Result struct {
    JobID  int
    Result string
}

func workerPoolWithResults(numWorkers int, jobs []Job) []Result {
    jobQueue := make(chan Job, len(jobs))
    resultQueue := make(chan Result, len(jobs))

    // Start workers
    for i := 0; i < numWorkers; i++ {
        go func(workerID int) {
            for job := range jobQueue {
                result := Result{
                    JobID:  job.ID,
                    Result: fmt.Sprintf("Processed by worker %d", workerID),
                }
                resultQueue <- result
            }
        }(i)
    }

    // Send jobs
    go func() {
        for _, job := range jobs {
            jobQueue <- job
        }
        close(jobQueue)
    }()

    // Collect results
    var results []Result
    for i := 0; i < len(jobs); i++ {
        results = append(results, <-resultQueue)
    }

    return results
}
```

## Worker Pool with Context

```go
func workerPoolWithContext(ctx context.Context, numWorkers int, jobs []Job) {
    jobQueue := make(chan Job, len(jobs))

    // Start workers
    for i := 0; i < numWorkers; i++ {
        go func(workerID int) {
            for {
                select {
                case job, ok := <-jobQueue:
                    if !ok {
                        return
                    }
                    fmt.Printf("Worker %d processing job %d\n", workerID, job.ID)
                case <-ctx.Done():
                    return
                }
            }
        }(i)
    }

    // Send jobs
    for _, job := range jobs {
        select {
        case jobQueue <- job:
        case <-ctx.Done():
            return
        }
    }
    close(jobQueue)
}
```

## Rate Limiting with Worker Pool

```go
func rateLimitedWorkerPool(numWorkers int, jobs []Job, rateLimit time.Duration) {
    jobQueue := make(chan Job, len(jobs))
    rateLimiter := time.NewTicker(rateLimit)
    defer rateLimiter.Stop()

    // Start workers
    for i := 0; i < numWorkers; i++ {
        go func(workerID int) {
            for job := range jobQueue {
                <-rateLimiter.C // Wait for rate limit
                fmt.Printf("Worker %d processing job %d\n", workerID, job.ID)
            }
        }(i)
    }

    // Send jobs
    for _, job := range jobs {
        jobQueue <- job
    }
    close(jobQueue)
}
```

## Best Practices

1. **Size the pool appropriately** - Too many workers can cause overhead
2. **Use buffered channels** - Prevent blocking on job submission
3. **Handle graceful shutdown** - Use context or quit channels
4. **Monitor performance** - Track processing times and throughput
5. **Error handling** - Handle worker failures gracefully

## Common Use Cases

- **Web scraping** - Process multiple URLs concurrently
- **Image processing** - Resize/compress images in parallel
- **Data processing** - Transform large datasets
- **API calls** - Make concurrent HTTP requests
- **File operations** - Process multiple files

## Next Steps

Now that you understand worker pools, try the tasks in `tasks.md`!