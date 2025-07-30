package main

import (
	"fmt"
)

// Task 1: Basic worker pool
type Job struct {
	ID   int
	Data string
}

func basicWorkerPool() {
	// TODO: Create jobs slice:
	// jobs := []Job{
	//     {ID: 1, Data: "Job 1"},
	//     {ID: 2, Data: "Job 2"},
	//     {ID: 3, Data: "Job 3"},
	//     {ID: 4, Data: "Job 4"},
	//     {ID: 5, Data: "Job 5"},
	// }
	// TODO: Create a simple worker pool with 3 workers:
	// numWorkers := 3
	// jobQueue := make(chan Job, len(jobs))
	// TODO: Process jobs concurrently:
	// var wg sync.WaitGroup
	// for i := 0; i < numWorkers; i++ {
	//     wg.Add(1)
	//     go func(workerID int) {
	//         defer wg.Done()
	//         for job := range jobQueue {
	//             fmt.Printf("Worker %d processing job %d: %s\n", workerID, job.ID, job.Data)
	//             time.Sleep(time.Millisecond * 100)
	//         }
	//     }(i)
	// }
	// TODO: Send jobs and wait:
	// for _, job := range jobs { jobQueue <- job }
	// close(jobQueue)
	// wg.Wait()
}

// Task 2: Worker pool with results
type Result struct {
	JobID  int
	Result string
}

func workerPoolWithResults() {
	// TODO: Create a worker pool that returns results:
	// jobs := []Job{{ID: 1, Data: "Job 1"}, {ID: 2, Data: "Job 2"}, {ID: 3, Data: "Job 3"}}
	// numWorkers := 2
	// jobQueue := make(chan Job, len(jobs))
	// resultQueue := make(chan Result, len(jobs))
	// TODO: Modify worker pool to collect results:
	// var wg sync.WaitGroup
	// for i := 0; i < numWorkers; i++ {
	//     wg.Add(1)
	//     go func(workerID int) {
	//         defer wg.Done()
	//         for job := range jobQueue {
	//             result := Result{JobID: job.ID, Result: fmt.Sprintf("Processed by worker %d", workerID)}
	//             resultQueue <- result
	//             time.Sleep(time.Millisecond * 100)
	//         }
	//     }(i)
	// }
	// TODO: Send jobs and collect results:
	// go func() { for _, job := range jobs { jobQueue <- job }; close(jobQueue) }()
	// go func() { wg.Wait(); close(resultQueue) }()
	// for result := range resultQueue { fmt.Printf("Result: Job %d - %s\n", result.JobID, result.Result) }
}

// Task 3: Worker pool with context
func workerPoolWithContext() {
	// TODO: Create a worker pool that accepts context:
	// jobs := []Job{{ID: 1, Data: "Job 1"}, {ID: 2, Data: "Job 2"}, {ID: 3, Data: "Job 3"}}
	// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// defer cancel()
	// numWorkers := 2
	// jobQueue := make(chan Job, len(jobs))
	// TODO: Handle context cancellation:
	// var wg sync.WaitGroup
	// for i := 0; i < numWorkers; i++ {
	//     wg.Add(1)
	//     go func(workerID int) {
	//         defer wg.Done()
	//         for {
	//             select {
	//             case job, ok := <-jobQueue:
	//                 if !ok { return }
	//                 fmt.Printf("Worker %d processing job %d\n", workerID, job.ID)
	//                 time.Sleep(time.Millisecond * 500)
	//             case <-ctx.Done():
	//                 return
	//             }
	//         }
	//     }(i)
	// }
	// TODO: Send jobs with context check:
	// go func() {
	//     for _, job := range jobs {
	//         select {
	//         case jobQueue <- job:
	//         case <-ctx.Done():
	//             return
	//         }
	//     }
	//     close(jobQueue)
	// }()
	// wg.Wait()
}

// Task 4: Rate limited worker pool
func rateLimitedWorkerPool() {
	// TODO: Implement rate limiting using time.Ticker:
	// jobs := []Job{{ID: 1, Data: "Job 1"}, {ID: 2, Data: "Job 2"}, {ID: 3, Data: "Job 3"}, {ID: 4, Data: "Job 4"}}
	// numWorkers := 2
	// rateLimit := time.Millisecond * 200
	// jobQueue := make(chan Job, len(jobs))
	// rateLimiter := time.NewTicker(rateLimit)
	// defer rateLimiter.Stop()
	// TODO: Control job processing rate:
	// var wg sync.WaitGroup
	// for i := 0; i < numWorkers; i++ {
	//     wg.Add(1)
	//     go func(workerID int) {
	//         defer wg.Done()
	//         for job := range jobQueue {
	//             <-rateLimiter.C // Wait for rate limit
	//             fmt.Printf("Worker %d processing job %d at %v\n", workerID, job.ID, time.Now())
	//         }
	//     }(i)
	// }
	// TODO: Send jobs and wait:
	// for _, job := range jobs { jobQueue <- job }
	// close(jobQueue)
	// wg.Wait()
}

// Task 5: Worker pool with error handling
type JobResult struct {
	JobID int
	Error error
	Data  string
}

func workerPoolWithErrors() {
	// TODO: Create jobs that may fail:
	// jobs := []Job{{ID: 1, Data: "Job 1"}, {ID: 2, Data: "Job 2"}, {ID: 3, Data: "Job 3"}}
	// numWorkers := 2
	// jobQueue := make(chan Job, len(jobs))
	// resultQueue := make(chan JobResult, len(jobs))
	// TODO: Handle worker errors gracefully:
	// var wg sync.WaitGroup
	// for i := 0; i < numWorkers; i++ {
	//     wg.Add(1)
	//     go func(workerID int) {
	//         defer wg.Done()
	//         for job := range jobQueue {
	//             var result JobResult
	//             result.JobID = job.ID
	//             if job.ID == 2 {
	//                 result.Error = fmt.Errorf("simulated error for job %d", job.ID)
	//             } else {
	//                 result.Data = fmt.Sprintf("Processed by worker %d", workerID)
	//             }
	//             resultQueue <- result
	//             time.Sleep(time.Millisecond * 100)
	//         }
	//     }(i)
	// }
	// TODO: Send jobs and collect results:
	// go func() { for _, job := range jobs { jobQueue <- job }; close(jobQueue) }()
	// go func() { wg.Wait(); close(resultQueue) }()
	// for result := range resultQueue {
	//     if result.Error != nil {
	//         fmt.Printf("Error processing job %d: %v\n", result.JobID, result.Error)
	//     } else {
	//         fmt.Printf("Success: Job %d - %s\n", result.JobID, result.Data)
	//     }
	// }
}

// Task 6: Practical application
func webScraperWorkerPool() {
	// TODO: Create a worker pool for web scraping:
	// urls := []string{"https://example1.com", "https://example2.com", "https://example3.com", "https://example4.com"}
	// numWorkers := 2
	// urlQueue := make(chan string, len(urls))
	// resultQueue := make(chan string, len(urls))
	// TODO: Simulate fetching URLs:
	// var wg sync.WaitGroup
	// for i := 0; i < numWorkers; i++ {
	//     wg.Add(1)
	//     go func(workerID int) {
	//         defer wg.Done()
	//         for url := range urlQueue {
	//             time.Sleep(time.Millisecond * 200)
	//             result := fmt.Sprintf("Worker %d scraped %s", workerID, url)
	//             resultQueue <- result
	//         }
	//     }(i)
	// }
	// TODO: Send URLs and collect results:
	// go func() { for _, url := range urls { urlQueue <- url }; close(urlQueue) }()
	// go func() { wg.Wait(); close(resultQueue) }()
	// for result := range resultQueue { fmt.Printf("Scraping result: %s\n", result) }
}

func main() {
	fmt.Println("=== Task 1: Basic Worker Pool ===")
	basicWorkerPool()

	fmt.Println("\n=== Task 2: Worker Pool with Results ===")
	workerPoolWithResults()

	fmt.Println("\n=== Task 3: Worker Pool with Context ===")
	workerPoolWithContext()

	fmt.Println("\n=== Task 4: Rate Limited Worker Pool ===")
	rateLimitedWorkerPool()

	fmt.Println("\n=== Task 5: Worker Pool with Error Handling ===")
	workerPoolWithErrors()

	fmt.Println("\n=== Task 6: Web Scraper Worker Pool ===")
	webScraperWorkerPool()
}
