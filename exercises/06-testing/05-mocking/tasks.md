# Mocking - Tasks

## Task 1: Basic Mocking
Create a simple mock implementation.

**Requirements:**
- Create an interface
- Implement a mock struct
- Test with mock dependency
- Verify mock behavior

## Task 2: Mock with State
Create a mock that maintains state.

**Requirements:**
- Mock with internal state
- Track method calls
- Verify state changes
- Test different scenarios

## Task 3: Mock Error Conditions
Test error handling with mocks.

**Requirements:**
- Mock error conditions
- Test error handling
- Verify error messages
- Test recovery scenarios

---

## How to complete these tasks:

1. Navigate to the `student/` directory
2. Create your code in `main.go`
3. Create tests in `main_test.go`
4. Run your tests: `go test`

## Example structure for student/main.go:
```go
package main

type Database interface {
    GetUser(id int) (*User, error)
    SaveUser(user *User) error
}

type User struct {
    ID   int
    Name string
}

type UserService struct {
    db Database
}

func (s *UserService) GetUser(id int) (*User, error) {
    return s.db.GetUser(id)
}

func (s *UserService) CreateUser(name string) (*User, error) {
    user := &User{Name: name}
    err := s.db.SaveUser(user)
    if err != nil {
        return nil, err
    }
    return user, nil
}
```

## Testing Your Code

Run your tests:
```bash
cd student/
go test
```