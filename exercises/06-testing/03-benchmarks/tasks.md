# Benchmarking - Tasks

## Task 1: Basic Benchmark
Create a simple benchmark function.

**Requirements:**
- Write a basic benchmark function
- Use proper benchmark naming conventions
- Prevent compiler optimization
- Run the benchmark successfully

## Task 2: String Operation Benchmarks
Compare different string operations.

**Requirements:**
- Benchmark string concatenation vs StringBuilder
- Use b.ResetTimer() appropriately
- Compare performance of different approaches
- Analyze benchmark results

## Task 3: Slice Operation Benchmarks
Benchmark slice operations.

**Requirements:**
- Compare preallocated vs dynamic slice growth
- Benchmark slice copying methods
- Use b.ReportAllocs() for memory analysis
- Demonstrate performance differences

## Task 4: Memory Allocation Benchmarks
Benchmark memory allocation patterns.

**Requirements:**
- Compare allocation vs reuse patterns
- Use b.ReportAllocs() for memory tracking
- Benchmark different allocation strategies
- Analyze memory usage patterns

## Task 5: Algorithm Comparison
Benchmark different algorithms.

**Requirements:**
- Compare different sorting approaches
- Benchmark search algorithms
- Use parameterized benchmarks
- Demonstrate algorithm performance differences

## Task 6: I/O Operation Benchmarks
Benchmark I/O operations.

**Requirements:**
- Benchmark file operations
- Compare different I/O approaches
- Use proper setup and teardown
- Analyze I/O performance characteristics

---

## How to complete these tasks:

1. Navigate to the `student/` directory
2. Create your code in `main.go`
3. Create benchmarks in `main_test.go`
4. Run your benchmarks: `go test -bench=.`


## Example structure for student/main_test.go:
```go
package main

import (
    "testing"
    "strings"
)

// Task 1: Basic benchmark
func BenchmarkSimpleFunction(b *testing.B) {
    for i := 0; i < b.N; i++ {
        result := SimpleFunction()
        _ = result
    }
}

// Task 2: String operation benchmarks
func BenchmarkStringConcatenation(b *testing.B) {
    for i := 0; i < b.N; i++ {
        result := StringConcatenation("hello", "world")
        _ = result
    }
}

func BenchmarkStringBuilderConcatenation(b *testing.B) {
    for i := 0; i < b.N; i++ {
        result := StringBuilderConcatenation("hello", "world")
        _ = result
    }
}

// Task 3: Slice operation benchmarks
func BenchmarkSliceDynamic(b *testing.B) {
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        result := CreateSliceDynamic(1000)
        _ = result
    }
}

func BenchmarkSlicePreallocated(b *testing.B) {
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        result := CreateSlicePreallocated(1000)
        _ = result
    }
}

// Task 4: Memory allocation benchmarks
func BenchmarkAllocation(b *testing.B) {
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        result := AllocateMemory(1024)
        _ = result
    }
}

func BenchmarkReuseAllocation(b *testing.B) {
    b.ReportAllocs()
    data := make([]byte, 1024)
    for i := 0; i < b.N; i++ {
        ReuseMemory(data, 1024)
    }
}

// Task 5: Algorithm comparison benchmarks
func BenchmarkBubbleSort(b *testing.B) {
    for i := 0; i < b.N; i++ {
        data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
        BubbleSort(data)
        _ = data
    }
}

func BenchmarkQuickSort(b *testing.B) {
    for i := 0; i < b.N; i++ {
        data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
        QuickSort(data)
        _ = data
    }
}

// Task 6: Parameterized benchmarks
func BenchmarkStringOperations(b *testing.B) {
    tests := []struct {
        name string
        size int
    }{
        {"small", 10},
        {"medium", 100},
        {"large", 1000},
    }

    for _, tt := range tests {
        b.Run(tt.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                var sb strings.Builder
                for j := 0; j < tt.size; j++ {
                    sb.WriteString("a")
                }
                _ = sb.String()
            }
        })
    }
}
```

## Testing Your Code

Run your benchmarks:
```bash
cd student/
go test -bench=.
```

Run with memory allocation info:
```bash
go test -bench=. -benchmem
```

Run with longer time:
```bash
go test -bench=. -benchtime=5s
```

You should see output demonstrating all benchmarking concepts.