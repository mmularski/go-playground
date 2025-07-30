package main

import (
	"testing"
)

// Task 1: Basic benchmark
func BenchmarkSimpleFunction(b *testing.B) {
	// TODO: Write a basic benchmark function:
	// for i := 0; i < b.N; i++ {
	//     result := SimpleFunction()
	//     _ = result
	// }
}

// Task 2: String operation benchmarks
func BenchmarkStringConcatenation(b *testing.B) {
	// TODO: Benchmark string concatenation:
	// for i := 0; i < b.N; i++ {
	//     result := StringConcatenation("hello", "world")
	//     _ = result
	// }
}

func BenchmarkStringBuilderConcatenation(b *testing.B) {
	// TODO: Benchmark StringBuilder approach:
	// for i := 0; i < b.N; i++ {
	//     result := StringBuilderConcatenation("hello", "world")
	//     _ = result
	// }
}

// Task 3: Slice operation benchmarks
func BenchmarkSliceDynamic(b *testing.B) {
	// TODO: Benchmark dynamic slice growth:
	// b.ReportAllocs()
	// for i := 0; i < b.N; i++ {
	//     result := CreateSliceDynamic(1000)
	//     _ = result
	// }
}

func BenchmarkSlicePreallocated(b *testing.B) {
	// TODO: Benchmark preallocated slice:
	// b.ReportAllocs()
	// for i := 0; i < b.N; i++ {
	//     result := CreateSlicePreallocated(1000)
	//     _ = result
	// }
}

// Task 4: Memory allocation benchmarks
func BenchmarkAllocation(b *testing.B) {
	// TODO: Benchmark memory allocation:
	// b.ReportAllocs()
	// for i := 0; i < b.N; i++ {
	//     result := AllocateMemory(1024)
	//     _ = result
	// }
}

func BenchmarkReuseAllocation(b *testing.B) {
	// TODO: Benchmark memory reuse:
	// b.ReportAllocs()
	// data := make([]byte, 1024)
	// for i := 0; i < b.N; i++ {
	//     ReuseMemory(data, 1024)
	// }
}

// Task 5: Algorithm comparison benchmarks
func BenchmarkBubbleSort(b *testing.B) {
	// TODO: Benchmark bubble sort:
	// for i := 0; i < b.N; i++ {
	//     data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	//     BubbleSort(data)
	//     _ = data
	// }
}

func BenchmarkQuickSort(b *testing.B) {
	// TODO: Benchmark quick sort:
	// for i := 0; i < b.N; i++ {
	//     data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	//     QuickSort(data)
	//     _ = data
	// }
}

// Task 6: Parameterized benchmarks
func BenchmarkStringOperations(b *testing.B) {
	// TODO: Create parameterized benchmarks:
	// tests := []struct {
	//     name string
	//     size int
	// }{
	//     {"small", 10},
	//     {"medium", 100},
	//     {"large", 1000},
	// }
	//
	// for _, tt := range tests {
	//     b.Run(tt.name, func(b *testing.B) {
	//         for i := 0; i < b.N; i++ {
	//             var sb strings.Builder
	//             for j := 0; j < tt.size; j++ {
	//                 sb.WriteString("a")
	//             }
	//             _ = sb.String()
	//         }
	//     })
	// }
}