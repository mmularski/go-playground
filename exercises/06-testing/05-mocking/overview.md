# Mocking - Overview

## What is Mocking?

Mocking is a technique used in testing to replace real dependencies with fake implementations. This allows you to test code in isolation.

## Why Use Mocking?

- **Isolate code under test** - Test specific functionality without external dependencies
- **Control test environment** - Simulate different scenarios and conditions
- **Speed up tests** - Avoid slow external services or databases
- **Test error conditions** - Simulate failures and edge cases
- **Predictable tests** - Tests don't depend on external state
- **Faster feedback** - Tests run quickly without external calls

## Basic Mocking Patterns

### Interface-Based Mocking
```go
type Database interface {
    GetUser(id int) (*User, error)
    SaveUser(user *User) error
}

type MockDatabase struct {
    users map[int]*User
}

func (m *MockDatabase) GetUser(id int) (*User, error) {
    if user, exists := m.users[id]; exists {
        return user, nil
    }
    return nil, fmt.Errorf("user not found")
}

func (m *MockDatabase) SaveUser(user *User) error {
    m.users[user.ID] = user
    return nil
}
```

## Mocking with State Tracking

### Tracking Method Calls
```go
type MockDatabase struct {
    users map[int]*User
    calls []string  // Track method calls
}

func (m *MockDatabase) GetUser(id int) (*User, error) {
    m.calls = append(m.calls, fmt.Sprintf("GetUser(%d)", id))
    if user, exists := m.users[id]; exists {
        return user, nil
    }
    return nil, fmt.Errorf("user not found")
}

func (m *MockDatabase) SaveUser(user *User) error {
    m.calls = append(m.calls, fmt.Sprintf("SaveUser(%d)", user.ID))
    m.users[user.ID] = user
    return nil
}

// Verify mock was called correctly
func TestMockCalls(t *testing.T) {
    mockDB := &MockDatabase{
        users: make(map[int]*User),
        calls: make([]string, 0),
    }

    service := &UserService{db: mockDB}
    service.GetUser(1)

    if len(mockDB.calls) != 1 {
        t.Errorf("Expected 1 call, got %d", len(mockDB.calls))
    }
    if mockDB.calls[0] != "GetUser(1)" {
        t.Errorf("Expected 'GetUser(1)', got %s", mockDB.calls[0])
    }
}
```

## Mocking Error Conditions

### Configurable Error Mocks
```go
type ErrorMockDatabase struct {
    shouldError bool
    errorMessage string
}

func (m *ErrorMockDatabase) GetUser(id int) (*User, error) {
    if m.shouldError {
        return nil, fmt.Errorf(m.errorMessage)
    }
    return &User{ID: id, Name: "Test User"}, nil
}

func (m *ErrorMockDatabase) SaveUser(user *User) error {
    if m.shouldError {
        return fmt.Errorf(m.errorMessage)
    }
    return nil
}

// Test error scenarios
func TestDatabaseError(t *testing.T) {
    mockDB := &ErrorMockDatabase{
        shouldError: true,
        errorMessage: "connection failed",
    }

    service := &UserService{db: mockDB}
    _, err := service.GetUser(1)

    if err == nil {
        t.Error("Expected error, got nil")
    }
    if err.Error() != "connection failed" {
        t.Errorf("Expected 'connection failed', got %s", err.Error())
    }
}
```

## Advanced Mocking Patterns

### Mock with Expectations
```go
type MockWithExpectations struct {
    expectedCalls map[string]int
    actualCalls   map[string]int
    returnValues  map[string]interface{}
}

func (m *MockWithExpectations) GetUser(id int) (*User, error) {
    m.actualCalls["GetUser"]++

    if returnVal, exists := m.returnValues["GetUser"]; exists {
        if user, ok := returnVal.(*User); ok {
            return user, nil
        }
        if err, ok := returnVal.(error); ok {
            return nil, err
        }
    }

    return &User{ID: id, Name: "Default User"}, nil
}

func (m *MockWithExpectations) VerifyExpectations(t *testing.T) {
    for method, expected := range m.expectedCalls {
        actual := m.actualCalls[method]
        if actual != expected {
            t.Errorf("Expected %d calls to %s, got %d", expected, method, actual)
        }
    }
}
```

### Mock with Callbacks
```go
type MockWithCallbacks struct {
    getUserCallback func(id int) (*User, error)
    saveUserCallback func(user *User) error
}

func (m *MockWithCallbacks) GetUser(id int) (*User, error) {
    if m.getUserCallback != nil {
        return m.getUserCallback(id)
    }
    return nil, fmt.Errorf("no callback set")
}

func (m *MockWithCallbacks) SaveUser(user *User) error {
    if m.saveUserCallback != nil {
        return m.saveUserCallback(user)
    }
    return fmt.Errorf("no callback set")
}
```

## Mocking HTTP Clients

### HTTP Client Mock
```go
type MockHTTPClient struct {
    responses map[string]*http.Response
    requests  []*http.Request
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
    m.requests = append(m.requests, req)

    key := req.Method + " " + req.URL.String()
    if response, exists := m.responses[key]; exists {
        return response, nil
    }

    return &http.Response{
        StatusCode: 404,
        Body:       io.NopCloser(strings.NewReader("not found")),
    }, nil
}

func TestHTTPClientMock(t *testing.T) {
    mockClient := &MockHTTPClient{
        responses: make(map[string]*http.Response),
        requests:  make([]*http.Request, 0),
    }

    // Set up expected response
    mockClient.responses["GET https://api.example.com/users/1"] = &http.Response{
        StatusCode: 200,
        Body:       io.NopCloser(strings.NewReader(`{"id":1,"name":"John"}`)),
    }

    service := &UserService{httpClient: mockClient}
    user, err := service.FetchUser(1)

    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if user.Name != "John" {
        t.Errorf("Expected 'John', got %s", user.Name)
    }
}
```

## Mocking Best Practices

### 1. Use Interfaces
```go
// Good: Code depends on interface
type UserService struct {
    db Database  // Interface, not concrete type
}

// Bad: Code depends on concrete type
type UserService struct {
    db *sql.DB  // Concrete type, hard to mock
}
```

### 2. Keep Mocks Simple
```go
// Good: Simple mock
type MockDatabase struct {
    users map[int]*User
}

// Avoid: Complex mock with too much logic
type ComplexMockDatabase struct {
    users map[int]*User
    calls []string
    timing time.Duration
    errors map[string]error
    // ... too many responsibilities
}
```

### 3. Verify Mock Behavior
```go
func TestUserServiceWithMock(t *testing.T) {
    mockDB := &MockDatabase{users: make(map[int]*User)}
    service := &UserService{db: mockDB}

    // Test the service
    user, err := service.GetUser(1)

    // Verify the result
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }

    // Verify the mock was called correctly
    if len(mockDB.calls) != 1 {
        t.Error("Expected mock to be called once")
    }
}
```

### 4. Test Different Scenarios
```go
func TestUserServiceScenarios(t *testing.T) {
    tests := []struct {
        name        string
        shouldError bool
        errorMsg    string
        expectError bool
    }{
        {"successful get", false, "", false},
        {"database error", true, "connection failed", true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockDB := &ErrorMockDatabase{
                shouldError: tt.shouldError,
                errorMessage: tt.errorMsg,
            }

            service := &UserService{db: mockDB}
            _, err := service.GetUser(1)

            if tt.expectError && err == nil {
                t.Error("Expected error, got nil")
            }
            if !tt.expectError && err != nil {
                t.Errorf("Expected no error, got %v", err)
            }
        })
    }
}
```

## Mocking Tools and Libraries

### Manual Mocking
- Write mocks by hand (what we've shown)
- Full control over mock behavior
- Good for simple cases

### Code Generation
```bash
# Using mockgen
go install github.com/golang/mock/mockgen@latest
mockgen -source=interfaces.go -destination=mocks.go
```

### Mock Libraries
- `testify/mock` - Popular mocking library
- `gomock` - Google's mocking framework
- `mockery` - Mock generation tool

## Example Program

```go
package main

import (
    "fmt"
    "testing"
)

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

// Mock implementation
type MockDatabase struct {
    users map[int]*User
    calls []string
}

func (m *MockDatabase) GetUser(id int) (*User, error) {
    m.calls = append(m.calls, fmt.Sprintf("GetUser(%d)", id))
    if user, exists := m.users[id]; exists {
        return user, nil
    }
    return nil, fmt.Errorf("user not found")
}

func (m *MockDatabase) SaveUser(user *User) error {
    m.calls = append(m.calls, fmt.Sprintf("SaveUser(%d)", user.ID))
    m.users[user.ID] = user
    return nil
}

func TestUserServiceWithMock(t *testing.T) {
    mockDB := &MockDatabase{
        users: make(map[int]*User),
        calls: make([]string, 0),
    }

    service := &UserService{db: mockDB}

    // Test creating user
    user, err := service.CreateUser("Alice")
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if user.Name != "Alice" {
        t.Errorf("Expected 'Alice', got %s", user.Name)
    }

    // Verify mock was called
    if len(mockDB.calls) != 1 {
        t.Errorf("Expected 1 call, got %d", len(mockDB.calls))
    }
}
```

## Next Steps

Now that you understand mocking, try the tasks in `tasks.md`!