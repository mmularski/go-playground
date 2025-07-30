package main

import "testing"

// Task 1: Test Add function
func TestAdd(t *testing.T) {
	// TODO: Write tests for Add function
	// Test with at least 3 different input combinations
}

// Task 2: Test Divide function
func TestDivide(t *testing.T) {
	// TODO: Write tests for Divide function
	// Test successful division
	// Test division by zero error
}

// Task 3: Test string functions
func TestReverse(t *testing.T) {
	// TODO: Write tests for Reverse function
	// Test with various strings including edge cases
}

func TestCountVowels(t *testing.T) {
	// TODO: Write tests for CountVowels function
	// Test with strings containing different numbers of vowels
}

// Task 4: Test boolean functions
func TestIsPalindrome(t *testing.T) {
	// TODO: Write tests for IsPalindrome function
	// Test palindromes and non-palindromes
}

func TestIsPrime(t *testing.T) {
	// TODO: Write tests for IsPrime function
	// Test prime and non-prime numbers
}

// Task 5: Test slice functions
func TestSumSlice(t *testing.T) {
	// TODO: Write tests for SumSlice function
	// Test with various slices
}

func TestFindMax(t *testing.T) {
	// TODO: Write tests for FindMax function
	// Test with various slices and empty slice
}

// Task 6: Helper functions
func assertEqual(t *testing.T, got, want int) {
	// TODO: Implement helper function
	// Compare integers and report error if different
}

func assertStringEqual(t *testing.T, got, want string) {
	// TODO: Implement helper function
	// Compare strings and report error if different
}

func assertError(t *testing.T, err error, expectedMsg string) {
	// TODO: Implement helper function
	// Check if error exists and has expected message
}
