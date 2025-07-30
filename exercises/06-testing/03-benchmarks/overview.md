# Benchmarking - Overview

## What is Benchmarking?

Benchmarking in Go is a way to measure the performance of your code. The `testing` package provides built-in support for writing and running benchmarks.

## Why Benchmark?

- **Measure performance** - Compare different implementations
- **Identify bottlenecks** - Find slow parts of your code
- **Optimize code** - Make informed decisions about improvements
- **Regression testing** - Ensure performance doesn't degrade

## Basic Benchmark Structure

### Simple Benchmark
```go
func BenchmarkFunction(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // Code to benchmark
        Function()
    }
}
```

### Benchmark with Setup
```go
func BenchmarkWithSetup(b *testing.B) {
    // Setup code (not measured)
    data := generateTestData()

    b.ResetTimer() // Start measuring from here

    for i := 0; i < b.N; i++ {
        // Code to benchmark
        ProcessData(data)
    }
}
```

## Benchmarking String Operations

### String Concatenation
```go
func BenchmarkStringConcatenation(b *testing.B) {
    for i := 0; i < b.N; i++ {
        result := "hello" + " " + "world"
        _ = result // Prevent compiler optimization
    }
}

func BenchmarkStringBuilder(b *testing.B) {
    for i := 0; i < b.N; i++ {
        var sb strings.Builder
        sb.WriteString("hello")
        sb.WriteString(" ")
        sb.WriteString("world")
        result := sb.String()
        _ = result
    }
}
```

### String Formatting
```go
func BenchmarkSprintf(b *testing.B) {
    for i := 0; i < b.N; i++ {
        result := fmt.Sprintf("Value: %d", i)
        _ = result
    }
}

func BenchmarkStringBuilderFormat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        var sb strings.Builder
        sb.WriteString("Value: ")
        sb.WriteString(strconv.Itoa(i))
        result := sb.String()
        _ = result
    }
}
```

## Benchmarking Slice Operations

### Slice Allocation
```go
func BenchmarkSlicePreallocated(b *testing.B) {
    for i := 0; i < b.N; i++ {
        slice := make([]int, 0, 1000)
        for j := 0; j < 1000; j++ {
            slice = append(slice, j)
        }
        _ = slice
    }
}

func BenchmarkSliceDynamic(b *testing.B) {
    for i := 0; i < b.N; i++ {
        var slice []int
        for j := 0; j < 1000; j++ {
            slice = append(slice, j)
        }
        _ = slice
    }
}
```

### Slice Copying
```go
func BenchmarkSliceCopy(b *testing.B) {
    source := make([]int, 1000)
    for i := range source {
        source[i] = i
    }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        dest := make([]int, len(source))
        copy(dest, source)
        _ = dest
    }
}

func BenchmarkSliceAppend(b *testing.B) {
    source := make([]int, 1000)
    for i := range source {
        source[i] = i
    }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        dest := append([]int(nil), source...)
        _ = dest
    }
}
```

## Benchmarking Map Operations

### Map Access Patterns
```go
func BenchmarkMapAccess(b *testing.B) {
    m := make(map[string]int)
    for i := 0; i < 1000; i++ {
        m[fmt.Sprintf("key%d", i)] = i
    }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        key := fmt.Sprintf("key%d", i%1000)
        _ = m[key]
    }
}

func BenchmarkMapWithValue(b *testing.B) {
    m := make(map[string]int)
    for i := 0; i < 1000; i++ {
        m[fmt.Sprintf("key%d", i)] = i
    }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        key := fmt.Sprintf("key%d", i%1000)
        if value, exists := m[key]; exists {
            _ = value
        }
    }
}
```

## Benchmarking Memory Allocation

### Memory Allocation Patterns
```go
func BenchmarkAllocation(b *testing.B) {
    b.ReportAllocs() // Report memory allocations

    for i := 0; i < b.N; i++ {
        data := make([]byte, 1024)
        _ = data
    }
}

func BenchmarkReuseAllocation(b *testing.B) {
    b.ReportAllocs()

    data := make([]byte, 1024)
    for i := 0; i < b.N; i++ {
        // Reuse the same slice
        for j := range data {
            data[j] = byte(i + j)
        }
        _ = data
    }
}
```

## Benchmarking Different Algorithms

### Sorting Algorithms
```go
func BenchmarkSortInts(b *testing.B) {
    for i := 0; i < b.N; i++ {
        data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
        sort.Ints(data)
        _ = data
    }
}

func BenchmarkSortSlice(b *testing.B) {
    for i := 0; i < b.N; i++ {
        data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
        sort.Slice(data, func(i, j int) bool {
            return data[i] < data[j]
        })
        _ = data
    }
}
```

## Benchmarking I/O Operations

### File Operations
```go
func BenchmarkFileRead(b *testing.B) {
    // Create test file
    data := []byte("Hello, World!")
    tmpfile, err := os.CreateTemp("", "benchmark")
    if err != nil {
        b.Fatal(err)
    }
    defer os.Remove(tmpfile.Name())
    tmpfile.Write(data)
    tmpfile.Close()

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        content, err := os.ReadFile(tmpfile.Name())
        if err != nil {
            b.Fatal(err)
        }
        _ = content
    }
}
```

## Benchmarking Network Operations

### HTTP Client
```go
func BenchmarkHTTPGet(b *testing.B) {
    // Start test server
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    }))
    defer server.Close()

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        resp, err := http.Get(server.URL)
        if err != nil {
            b.Fatal(err)
        }
        io.Copy(io.Discard, resp.Body)
        resp.Body.Close()
    }
}
```

## Benchmarking with Parameters

### Parameterized Benchmarks
```go
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

## Benchmarking Best Practices

1. **Use b.ResetTimer()** - Reset timer after setup
2. **Prevent optimization** - Use the result to prevent compiler optimization
3. **Report allocations** - Use b.ReportAllocs() for memory benchmarks
4. **Run multiple times** - Use -benchtime flag for longer runs
5. **Compare implementations** - Benchmark different approaches
6. **Consider context** - Benchmarks may not reflect real-world performance

## Running Benchmarks

```bash
# Run all benchmarks
go test -bench=.

# Run specific benchmark
go test -bench=BenchmarkString

# Run with longer time
go test -bench=. -benchtime=5s

# Run with memory allocation info
go test -bench=. -benchmem

# Run with CPU profiling
go test -bench=. -cpuprofile=cpu.prof

# Run with memory profiling
go test -bench=. -memprofile=mem.prof
```

## Example Program

```go
package main

import (
    "testing"
    "strings"
)

func BenchmarkStringConcatenation(b *testing.B) {
    for i := 0; i < b.N; i++ {
        result := "hello" + " " + "world"
        _ = result
    }
}

func BenchmarkStringBuilder(b *testing.B) {
    for i := 0; i < b.N; i++ {
        var sb strings.Builder
        sb.WriteString("hello")
        sb.WriteString(" ")
        sb.WriteString("world")
        result := sb.String()
        _ = result
    }
}

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

## Next Steps

Now that you understand benchmarking, try the tasks in `tasks.md`!