package main

import (
	"testing"
)

// Task 1: Basic coverage tests
func TestProcessDataEmpty(t *testing.T) {
	// TODO: Test empty data:
	// result := ProcessData([]int{})
	// if result != "empty" {
	//     t.Errorf("ProcessData([]int{}) = %s, want 'empty'", result)
	// }
}

func TestProcessDataHigh(t *testing.T) {
	// TODO: Test high sum:
	// result := ProcessData([]int{50, 60})
	// if result != "high" {
	//     t.Errorf("ProcessData([50, 60]) = %s, want 'high'", result)
	// }
}

func TestProcessDataMedium(t *testing.T) {
	// TODO: Test medium sum:
	// result := ProcessData([]int{30, 40})
	// if result != "medium" {
	//     t.Errorf("ProcessData([30, 40]) = %s, want 'medium'", result)
	// }
}

func TestProcessDataLow(t *testing.T) {
	// TODO: Test low sum:
	// result := ProcessData([]int{10, 20})
	// if result != "low" {
	//     t.Errorf("ProcessData([10, 20]) = %s, want 'low'", result)
	// }
}

func TestProcessDataWithNegative(t *testing.T) {
	// TODO: Test with negative numbers:
	// result := ProcessData([]int{10, -5, 20})
	// if result != "low" {
	//     t.Errorf("ProcessData([10, -5, 20]) = %s, want 'low'", result)
	// }
}

func TestProcessDataAllNegative(t *testing.T) {
	// TODO: Test all negative numbers:
	// result := ProcessData([]int{-10, -20, -30})
	// if result != "low" {
	//     t.Errorf("ProcessData([-10, -20, -30]) = %s, want 'low'", result)
	// }
}

// Task 2: Edge case tests for better coverage
func TestProcessDataSinglePositive(t *testing.T) {
	// TODO: Test single positive number:
	// result := ProcessData([]int{60})
	// if result != "medium" {
	//     t.Errorf("ProcessData([60]) = %s, want 'medium'", result)
	// }
}

func TestProcessDataSingleNegative(t *testing.T) {
	// TODO: Test single negative number:
	// result := ProcessData([]int{-10})
	// if result != "low" {
	//     t.Errorf("ProcessData([-10]) = %s, want 'low'", result)
	// }
}

func TestProcessDataExactThresholds(t *testing.T) {
	// TODO: Test exact threshold values:
	// result := ProcessData([]int{50})
	// if result != "medium" {
	//     t.Errorf("ProcessData([50]) = %s, want 'medium'", result)
	// }
	//
	// result = ProcessData([]int{100})
	// if result != "high" {
	//     t.Errorf("ProcessData([100]) = %s, want 'high'", result)
	// }
}

// Task 3: Table-driven tests for comprehensive coverage
func TestProcessDataComprehensive(t *testing.T) {
	// TODO: Create table-driven tests for comprehensive coverage:
	// tests := []struct {
	//     name     string
	//     input    []int
	//     expected string
	// }{
	//     {"empty", []int{}, "empty"},
	//     {"single positive", []int{60}, "medium"},
	//     {"single negative", []int{-10}, "low"},
	//     {"multiple positive high", []int{50, 60}, "high"},
	//     {"multiple positive medium", []int{30, 40}, "medium"},
	//     {"multiple positive low", []int{10, 20}, "low"},
	//     {"mixed positive negative", []int{10, -5, 20}, "low"},
	//     {"all negative", []int{-10, -20, -30}, "low"},
	//     {"exact medium threshold", []int{50}, "medium"},
	//     {"exact high threshold", []int{100}, "high"},
	// }
	//
	// for _, tt := range tests {
	//     t.Run(tt.name, func(t *testing.T) {
	//         result := ProcessData(tt.input)
	//         if result != tt.expected {
	//             t.Errorf("ProcessData(%v) = %s, want %s", tt.input, result, tt.expected)
	//         }
	//     })
	// }
}