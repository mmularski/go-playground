# Microservice Project - Overview

## What is a Microservice?

A microservice is a small, independent service that performs a specific business function. Microservices communicate with each other through well-defined APIs and can be developed, deployed, and scaled independently.

## Microservice Architecture

### Service Communication
```go
package main

import (
    "encoding/json"
    "net/http"
    "time"
)

type UserService struct {
    client *http.Client
}

func NewUserService() *UserService {
    return &UserService{
        client: &http.Client{
            Timeout: 30 * time.Second,
        },
    }
}

func (us *UserService) GetUser(id int) (*User, error) {
    resp, err := us.client.Get(fmt.Sprintf("http://user-service:8081/users/%d", id))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var user User
    decoder := json.NewDecoder(resp.Body)
    err = decoder.Decode(&user)

    return &user, err
}
```

## Service Discovery

### Simple Service Registry
```go
type ServiceRegistry struct {
    services map[string]string
    mu       sync.RWMutex
}

func NewServiceRegistry() *ServiceRegistry {
    return &ServiceRegistry{
        services: make(map[string]string),
    }
}

func (sr *ServiceRegistry) Register(serviceName, serviceURL string) {
    sr.mu.Lock()
    defer sr.mu.Unlock()
    sr.services[serviceName] = serviceURL
}

func (sr *ServiceRegistry) GetService(serviceName string) (string, bool) {
    sr.mu.RLock()
    defer sr.mu.RUnlock()
    url, exists := sr.services[serviceName]
    return url, exists
}
```

## API Gateway

### Basic Gateway
```go
type Gateway struct {
    userService    *UserService
    orderService   *OrderService
    productService *ProductService
}

func (g *Gateway) handleUserRequest(w http.ResponseWriter, r *http.Request) {
    // Extract user ID from URL
    userID := extractUserID(r.URL.Path)

    user, err := g.userService.GetUser(userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func (g *Gateway) handleOrderRequest(w http.ResponseWriter, r *http.Request) {
    // Route to order service
    // Implementation here
}
```

## Circuit Breaker Pattern

### Circuit Breaker Implementation
```go
type CircuitBreaker struct {
    failures   int
    lastError  time.Time
    state      string // "closed", "open", "half-open"
    threshold  int
    timeout    time.Duration
    mu         sync.RWMutex
}

func NewCircuitBreaker(threshold int, timeout time.Duration) *CircuitBreaker {
    return &CircuitBreaker{
        threshold: threshold,
        timeout:   timeout,
        state:     "closed",
    }
}

func (cb *CircuitBreaker) Execute(command func() error) error {
    cb.mu.Lock()
    defer cb.mu.Unlock()

    switch cb.state {
    case "open":
        if time.Since(cb.lastError) > cb.timeout {
            cb.state = "half-open"
        } else {
            return fmt.Errorf("circuit breaker is open")
        }
    }

    err := command()
    if err != nil {
        cb.failures++
        cb.lastError = time.Now()

        if cb.failures >= cb.threshold {
            cb.state = "open"
        }
        return err
    }

    cb.failures = 0
    cb.state = "closed"
    return nil
}
```

## Load Balancing

### Simple Load Balancer
```go
type LoadBalancer struct {
    servers []string
    current int
    mu      sync.Mutex
}

func NewLoadBalancer(servers []string) *LoadBalancer {
    return &LoadBalancer{
        servers: servers,
        current: 0,
    }
}

func (lb *LoadBalancer) GetNextServer() string {
    lb.mu.Lock()
    defer lb.mu.Unlock()

    server := lb.servers[lb.current]
    lb.current = (lb.current + 1) % len(lb.servers)
    return server
}
```

## Message Queues

### Using Channels as Message Queue
```go
type MessageQueue struct {
    messages chan Message
    workers  int
}

type Message struct {
    ID      string                 `json:"id"`
    Type    string                 `json:"type"`
    Data    map[string]interface{} `json:"data"`
    Created time.Time              `json:"created"`
}

func NewMessageQueue(workers int) *MessageQueue {
    mq := &MessageQueue{
        messages: make(chan Message, 100),
        workers:  workers,
    }

    // Start workers
    for i := 0; i < workers; i++ {
        go mq.worker()
    }

    return mq
}

func (mq *MessageQueue) Publish(msg Message) {
    mq.messages <- msg
}

func (mq *MessageQueue) worker() {
    for msg := range mq.messages {
        mq.processMessage(msg)
    }
}

func (mq *MessageQueue) processMessage(msg Message) {
    // Process message based on type
    switch msg.Type {
    case "user_created":
        // Handle user creation
    case "order_placed":
        // Handle order placement
    default:
        log.Printf("Unknown message type: %s", msg.Type)
    }
}
```

## Configuration Management

### Environment-based Configuration
```go
type Config struct {
    ServiceName    string `json:"service_name"`
    Port           int    `json:"port"`
    DatabaseURL    string `json:"database_url"`
    RedisURL       string `json:"redis_url"`
    LogLevel       string `json:"log_level"`
    Environment    string `json:"environment"`
}

func LoadConfig() (*Config, error) {
    config := &Config{
        ServiceName: getEnv("SERVICE_NAME", "unknown"),
        Port:        getEnvAsInt("PORT", 8080),
        DatabaseURL: getEnv("DATABASE_URL", ""),
        RedisURL:    getEnv("REDIS_URL", ""),
        LogLevel:    getEnv("LOG_LEVEL", "info"),
        Environment: getEnv("ENVIRONMENT", "development"),
    }

    return config, nil
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
    if value := os.Getenv(key); value != "" {
        if intValue, err := strconv.Atoi(value); err == nil {
            return intValue
        }
    }
    return defaultValue
}
```

## Health Checks

### Health Check Endpoint
```go
type HealthChecker struct {
    services map[string]HealthCheck
}

type HealthCheck func() error

func NewHealthChecker() *HealthChecker {
    return &HealthChecker{
        services: make(map[string]HealthCheck),
    }
}

func (hc *HealthChecker) AddCheck(name string, check HealthCheck) {
    hc.services[name] = check
}

func (hc *HealthChecker) CheckHealth() map[string]string {
    results := make(map[string]string)

    for name, check := range hc.services {
        if err := check(); err != nil {
            results[name] = "unhealthy"
        } else {
            results[name] = "healthy"
        }
    }

    return results
}

func (hc *HealthChecker) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
    results := hc.CheckHealth()

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(results)
}
```

## Logging and Monitoring

### Structured Logging
```go
import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

type Logger struct {
    zap *zap.Logger
}

func NewLogger(level string) (*Logger, error) {
    config := zap.NewProductionConfig()

    switch level {
    case "debug":
        config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
    case "info":
        config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
    case "warn":
        config.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
    case "error":
        config.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
    }

    logger, err := config.Build()
    if err != nil {
        return nil, err
    }

    return &Logger{zap: logger}, nil
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
    l.zap.Info(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
    l.zap.Error(msg, fields...)
}
```

## Database Integration

### Database Connection Pool
```go
import (
    "database/sql"
    _ "github.com/lib/pq"
)

type Database struct {
    db *sql.DB
}

func NewDatabase(databaseURL string) (*Database, error) {
    db, err := sql.Open("postgres", databaseURL)
    if err != nil {
        return nil, err
    }

    // Configure connection pool
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(5)
    db.SetConnMaxLifetime(5 * time.Minute)

    // Test connection
    if err := db.Ping(); err != nil {
        return nil, err
    }

    return &Database{db: db}, nil
}

func (d *Database) Close() error {
    return d.db.Close()
}
```

## Caching

### Redis Cache
```go
import (
    "context"
    "encoding/json"
    "time"

    "github.com/go-redis/redis/v8"
)

type Cache struct {
    client *redis.Client
}

func NewCache(redisURL string) (*Cache, error) {
    opt, err := redis.ParseURL(redisURL)
    if err != nil {
        return nil, err
    }

    client := redis.NewClient(opt)

    // Test connection
    ctx := context.Background()
    if err := client.Ping(ctx).Err(); err != nil {
        return nil, err
    }

    return &Cache{client: client}, nil
}

func (c *Cache) Set(key string, value interface{}, expiration time.Duration) error {
    data, err := json.Marshal(value)
    if err != nil {
        return err
    }

    ctx := context.Background()
    return c.client.Set(ctx, key, data, expiration).Err()
}

func (c *Cache) Get(key string, dest interface{}) error {
    ctx := context.Background()
    data, err := c.client.Get(ctx, key).Bytes()
    if err != nil {
        return err
    }

    return json.Unmarshal(data, dest)
}
```

## Example Microservice

```go
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
    "time"

    "go.uber.org/zap"
)

type UserService struct {
    logger *zap.Logger
    cache  *Cache
    db     *Database
}

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func NewUserService() (*UserService, error) {
    // Initialize logger
    logger, err := zap.NewProduction()
    if err != nil {
        return nil, err
    }

    // Initialize cache
    cache, err := NewCache(os.Getenv("REDIS_URL"))
    if err != nil {
        return nil, err
    }

    // Initialize database
    db, err := NewDatabase(os.Getenv("DATABASE_URL"))
    if err != nil {
        return nil, err
    }

    return &UserService{
        logger: logger,
        cache:  cache,
        db:     db,
    }, nil
}

func (us *UserService) GetUser(w http.ResponseWriter, r *http.Request) {
    // Extract user ID from URL
    userID := extractUserID(r.URL.Path)

    // Try cache first
    var user User
    cacheKey := fmt.Sprintf("user:%d", userID)

    if err := us.cache.Get(cacheKey, &user); err == nil {
        us.logger.Info("User found in cache", zap.Int("user_id", userID))
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(user)
        return
    }

    // Get from database
    user, err := us.db.GetUser(userID)
    if err != nil {
        us.logger.Error("Failed to get user from database",
            zap.Int("user_id", userID),
            zap.Error(err))
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Cache the result
    us.cache.Set(cacheKey, user, 5*time.Minute)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func (us *UserService) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Create user in database
    createdUser, err := us.db.CreateUser(&user)
    if err != nil {
        us.logger.Error("Failed to create user", zap.Error(err))
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(createdUser)
}

func main() {
    service, err := NewUserService()
    if err != nil {
        log.Fatal(err)
    }

    // Set up routes
    http.HandleFunc("/users/", service.GetUser)
    http.HandleFunc("/users", service.CreateUser)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("User service starting on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
```

## Docker Integration

### Dockerfile
```dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o user-service ./cmd/user-service

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/user-service .
EXPOSE 8080

CMD ["./user-service"]
```

### Docker Compose
```yaml
version: '3.8'

services:
  user-service:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DATABASE_URL=postgres://user:password@postgres:5432/users
      - REDIS_URL=redis://redis:6379
    depends_on:
      - postgres
      - redis

  postgres:
    image: postgres:13
    environment:
      - POSTGRES_DB=users
      - POSTGRES_USER=user
      - POSTGRES_PASSWORD=password
    volumes:
      - postgres_data:/var/lib/postgresql/data

  redis:
    image: redis:6-alpine
    ports:
      - "6379:6379"

volumes:
  postgres_data:
```

## Best Practices

### 1. Service Independence
- Each service should be independently deployable
- Services should have their own databases
- Minimize inter-service dependencies

### 2. API Design
- Use RESTful APIs with consistent patterns
- Implement proper error handling
- Use versioning for API changes

### 3. Monitoring and Observability
- Implement health checks
- Add comprehensive logging
- Use distributed tracing
- Monitor service metrics

### 4. Security
- Implement authentication and authorization
- Use HTTPS for all communications
- Validate all inputs
- Implement rate limiting

## Next Steps

Now that you understand microservices, try the tasks in `tasks.md`!