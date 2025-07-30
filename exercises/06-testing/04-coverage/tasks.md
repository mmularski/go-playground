# Test Coverage - Tasks

## Task 1: Basic Coverage
Run tests with coverage analysis.

**Requirements:**
- Run tests with coverage reporting
- Generate coverage profile
- View coverage in HTML format
- Analyze coverage results

## Task 2: Improve Coverage
Write tests to improve coverage.

**Requirements:**
- Identify uncovered code paths
- Write tests for missing coverage
- Achieve higher coverage percentage
- Test edge cases and error conditions

## Task 3: Coverage Analysis
Analyze coverage data.

**Requirements:**
- Use coverage tools effectively
- Interpret coverage reports
- Identify coverage gaps
- Plan test improvements

---

## How to complete these tasks:

1. Navigate to the `student/` directory
2. Create your code in `main.go`
3. Create tests in `main_test.go`
4. Run coverage analysis: `go test -cover`

## Example structure for student/main.go:
```go
package main

func ProcessData(data []int) string {
    if len(data) == 0 {
        return "empty"
    }

    sum := 0
    for _, value := range data {
        if value > 0 {
            sum += value
        }
    }

    if sum > 100 {
        return "high"
    } else if sum > 50 {
        return "medium"
    } else {
        return "low"
    }
}
```

## Testing Your Code

Run with coverage:
```bash
cd student/
go test -cover
```

Generate coverage profile:
```bash
go test -coverprofile=coverage.out
```

View HTML coverage:
```bash
go tool cover -html=coverage.out
```