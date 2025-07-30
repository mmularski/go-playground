package main

// Task 1: Basic functions to benchmark
func SimpleFunction() int {
	// TODO: Return a simple value: return 42
	return 0
}

// Task 2: String operations
func StringConcatenation(a, b string) string {
	// TODO: Concatenate strings: return a + " " + b
	return ""
}

func StringBuilderConcatenation(a, b string) string {
	// TODO: Use strings.Builder:
	// var sb strings.Builder
	// sb.WriteString(a)
	// sb.WriteString(" ")
	// sb.WriteString(b)
	// return sb.String()
	return ""
}

// Task 3: Slice operations
func CreateSliceDynamic(size int) []int {
	// TODO: Create slice dynamically:
	// var slice []int
	// for i := 0; i < size; i++ {
	//     slice = append(slice, i)
	// }
	// return slice
	return nil
}

func CreateSlicePreallocated(size int) []int {
	// TODO: Create slice with preallocation:
	// slice := make([]int, 0, size)
	// for i := 0; i < size; i++ {
	//     slice = append(slice, i)
	// }
	// return slice
	return nil
}

// Task 4: Memory operations
func AllocateMemory(size int) []byte {
	// TODO: Allocate memory: return make([]byte, size)
	return nil
}

func ReuseMemory(data []byte, size int) {
	// TODO: Reuse existing memory:
	// for i := range data {
	//     data[i] = byte(i % 256)
	// }
}

// Task 5: Algorithm implementations
func BubbleSort(data []int) {
	// TODO: Implement bubble sort:
	// n := len(data)
	// for i := 0; i < n-1; i++ {
	//     for j := 0; j < n-i-1; j++ {
	//         if data[j] > data[j+1] {
	//             data[j], data[j+1] = data[j+1], data[j]
	//         }
	//     }
	// }
}

func QuickSort(data []int) {
	// TODO: Use sort.Ints: sort.Ints(data)
}
