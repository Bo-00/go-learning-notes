# Go 风格错误处理

## 1. 错误处理基础

### error 接口

```go
// Go 内置的 error 接口
type error interface {
    Error() string
}

// 最简单的错误创建
err := errors.New("something went wrong")
err := fmt.Errorf("value %d is invalid", value)

// 检查错误
if err != nil {
    // 处理错误
    return err
}
```

### 基本错误处理模式

```go
func readFile(filename string) ([]byte, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
    }
    defer file.Close()

    data, err := io.ReadAll(file)
    if err != nil {
        return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
    }

    return data, nil
}

// 使用
func main() {
    data, err := readFile("config.json")
    if err != nil {
        log.Printf("Error: %v", err)
        return
    }

    fmt.Printf("Read %d bytes\n", len(data))
}
```

## 2. 自定义错误类型

### 结构体错误

```go
// 自定义错误结构体
type ValidationError struct {
    Field   string
    Value   interface{}
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed for field '%s' with value '%v': %s",
        e.Field, e.Value, e.Message)
}

// 使用自定义错误
func validateAge(age int) error {
    if age < 0 {
        return &ValidationError{
            Field:   "age",
            Value:   age,
            Message: "age cannot be negative",
        }
    }
    if age > 150 {
        return &ValidationError{
            Field:   "age",
            Value:   age,
            Message: "age seems unrealistic",
        }
    }
    return nil
}

// 错误类型检查
func handleValidation(err error) {
    if validErr, ok := err.(*ValidationError); ok {
        fmt.Printf("Validation failed on field: %s\n", validErr.Field)
        fmt.Printf("Invalid value: %v\n", validErr.Value)
        fmt.Printf("Message: %s\n", validErr.Message)
    }
}
```

### 常量错误

```go
// 使用常量定义错误
var (
    ErrUserNotFound = errors.New("user not found")
    ErrInvalidInput = errors.New("invalid input")
    ErrUnauthorized = errors.New("unauthorized access")
)

// 使用
func getUser(id string) (*User, error) {
    if id == "" {
        return nil, ErrInvalidInput
    }

    user := findUserByID(id)
    if user == nil {
        return nil, ErrUserNotFound
    }

    return user, nil
}

// 错误比较
func main() {
    user, err := getUser("")
    if err == ErrInvalidInput {
        fmt.Println("Please provide a valid user ID")
        return
    }
    if err == ErrUserNotFound {
        fmt.Println("User does not exist")
        return
    }

    fmt.Printf("Found user: %+v\n", user)
}
```

## 3. 错误包装和解包

### 错误包装 (Go 1.13+)

```go
import "errors"

// 包装错误
func processFile(filename string) error {
    data, err := readFile(filename)
    if err != nil {
        return fmt.Errorf("processing file %s failed: %w", filename, err)
    }

    err = validate(data)
    if err != nil {
        return fmt.Errorf("validation failed for file %s: %w", filename, err)
    }

    return nil
}

// 错误解包
func handleError(err error) {
    // 检查是否包装了特定错误
    if errors.Is(err, ErrUserNotFound) {
        fmt.Println("User not found in the error chain")
    }

    // 解包到特定类型
    var validErr *ValidationError
    if errors.As(err, &validErr) {
        fmt.Printf("Found validation error: %s\n", validErr.Field)
    }

    // 手动解包
    unwrapped := errors.Unwrap(err)
    if unwrapped != nil {
        fmt.Printf("Underlying error: %v\n", unwrapped)
    }
}
```

### 多级错误包装

```go
// 创建错误链
func businessLogic() error {
    err := databaseOperation()
    if err != nil {
        return fmt.Errorf("business logic failed: %w", err)
    }
    return nil
}

func databaseOperation() error {
    err := networkCall()
    if err != nil {
        return fmt.Errorf("database operation failed: %w", err)
    }
    return nil
}

func networkCall() error {
    return fmt.Errorf("network connection failed: %w",
        errors.New("connection timeout"))
}

// 错误链分析
func analyzeError() {
    err := businessLogic()
    if err != nil {
        fmt.Printf("Top level error: %v\n", err)

        // 遍历错误链
        currentErr := err
        level := 0
        for currentErr != nil {
            fmt.Printf("Level %d: %v\n", level, currentErr)
            currentErr = errors.Unwrap(currentErr)
            level++
        }
    }
}
```

## 4. 错误处理模式

### 提前返回模式

```go
// ✅ 推荐的提前返回模式
func processData(data []byte) error {
    if len(data) == 0 {
        return errors.New("data is empty")
    }

    parsed, err := parseData(data)
    if err != nil {
        return fmt.Errorf("parsing failed: %w", err)
    }

    err = validateParsedData(parsed)
    if err != nil {
        return fmt.Errorf("validation failed: %w", err)
    }

    err = saveData(parsed)
    if err != nil {
        return fmt.Errorf("saving failed: %w", err)
    }

    return nil
}

// ❌ 避免深层嵌套
func processDataBad(data []byte) error {
    if len(data) > 0 {
        parsed, err := parseData(data)
        if err == nil {
            err = validateParsedData(parsed)
            if err == nil {
                err = saveData(parsed)
                if err != nil {
                    return fmt.Errorf("saving failed: %w", err)
                }
            } else {
                return fmt.Errorf("validation failed: %w", err)
            }
        } else {
            return fmt.Errorf("parsing failed: %w", err)
        }
    } else {
        return errors.New("data is empty")
    }
    return nil
}
```

### 错误聚合

```go
// 收集多个错误
type MultiError struct {
    errors []error
}

func (m *MultiError) Error() string {
    if len(m.errors) == 0 {
        return ""
    }

    var builder strings.Builder
    builder.WriteString("multiple errors occurred:\n")
    for i, err := range m.errors {
        builder.WriteString(fmt.Sprintf("  %d: %v\n", i+1, err))
    }
    return builder.String()
}

func (m *MultiError) Add(err error) {
    if err != nil {
        m.errors = append(m.errors, err)
    }
}

func (m *MultiError) HasErrors() bool {
    return len(m.errors) > 0
}

// 批量处理
func processFiles(filenames []string) error {
    var multiErr MultiError

    for _, filename := range filenames {
        err := processFile(filename)
        multiErr.Add(err)
    }

    if multiErr.HasErrors() {
        return &multiErr
    }
    return nil
}
```

### 重试机制

```go
import "time"

// 重试配置
type RetryConfig struct {
    MaxRetries int
    Delay      time.Duration
    Backoff    func(attempt int) time.Duration
}

// 可重试的错误
type RetryableError struct {
    Err       error
    Retryable bool
}

func (r *RetryableError) Error() string {
    return r.Err.Error()
}

func (r *RetryableError) IsRetryable() bool {
    return r.Retryable
}

// 重试函数
func retry(fn func() error, config RetryConfig) error {
    var lastErr error

    for attempt := 0; attempt <= config.MaxRetries; attempt++ {
        err := fn()
        if err == nil {
            return nil
        }

        lastErr = err

        // 检查是否可重试
        if retryErr, ok := err.(*RetryableError); ok && !retryErr.IsRetryable() {
            return err
        }

        if attempt < config.MaxRetries {
            delay := config.Delay
            if config.Backoff != nil {
                delay = config.Backoff(attempt)
            }
            time.Sleep(delay)
        }
    }

    return fmt.Errorf("failed after %d attempts: %w", config.MaxRetries+1, lastErr)
}

// 使用重试
func unreliableOperation() error {
    // 模拟不稳定的操作
    if rand.Float32() < 0.7 {
        return &RetryableError{
            Err:       errors.New("temporary network error"),
            Retryable: true,
        }
    }
    return nil
}

func main() {
    config := RetryConfig{
        MaxRetries: 3,
        Delay:      time.Second,
        Backoff: func(attempt int) time.Duration {
            return time.Duration(attempt) * time.Second
        },
    }

    err := retry(unreliableOperation, config)
    if err != nil {
        log.Printf("Operation failed: %v", err)
    }
}
```

## 5. 上下文中的错误处理

### 超时和取消

```go
import "context"

func operationWithTimeout(ctx context.Context) error {
    // 创建子上下文，设置超时
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    // 模拟长时间运行的操作
    select {
    case <-time.After(3 * time.Second):
        return nil // 操作完成
    case <-ctx.Done():
        return fmt.Errorf("operation cancelled: %w", ctx.Err())
    }
}

// 检查上下文错误
func handleContextError(err error) {
    if errors.Is(err, context.Canceled) {
        fmt.Println("Operation was cancelled")
    } else if errors.Is(err, context.DeadlineExceeded) {
        fmt.Println("Operation timed out")
    } else {
        fmt.Printf("Other error: %v\n", err)
    }
}
```

### 上下文传递错误信息

```go
type contextKey string

const (
    RequestIDKey contextKey = "request_id"
    UserIDKey    contextKey = "user_id"
)

// 带上下文的错误
type ContextualError struct {
    RequestID string
    UserID    string
    Operation string
    Err       error
}

func (e *ContextualError) Error() string {
    return fmt.Sprintf("[req:%s][user:%s][op:%s] %v",
        e.RequestID, e.UserID, e.Operation, e.Err)
}

func newContextualError(ctx context.Context, operation string, err error) *ContextualError {
    requestID, _ := ctx.Value(RequestIDKey).(string)
    userID, _ := ctx.Value(UserIDKey).(string)

    return &ContextualError{
        RequestID: requestID,
        UserID:    userID,
        Operation: operation,
        Err:       err,
    }
}

// 使用
func businessOperation(ctx context.Context) error {
    err := someOperation()
    if err != nil {
        return newContextualError(ctx, "business_operation", err)
    }
    return nil
}
```

## 6. 测试中的错误处理

### 错误测试模式

```go
func TestValidateAge(t *testing.T) {
    tests := []struct {
        name        string
        age         int
        expectError bool
        errorType   error
    }{
        {
            name:        "valid age",
            age:         25,
            expectError: false,
        },
        {
            name:        "negative age",
            age:         -1,
            expectError: true,
            errorType:   &ValidationError{},
        },
        {
            name:        "unrealistic age",
            age:         200,
            expectError: true,
            errorType:   &ValidationError{},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := validateAge(tt.age)

            if tt.expectError {
                if err == nil {
                    t.Error("expected error but got none")
                    return
                }

                if tt.errorType != nil {
                    if !errors.As(err, &tt.errorType) {
                        t.Errorf("expected error type %T, got %T", tt.errorType, err)
                    }
                }
            } else {
                if err != nil {
                    t.Errorf("unexpected error: %v", err)
                }
            }
        })
    }
}
```

### Mock 错误

```go
// 错误注入用于测试
type ErrorInjector struct {
    shouldFail map[string]bool
    errors     map[string]error
}

func (e *ErrorInjector) ShouldFail(operation string) bool {
    return e.shouldFail[operation]
}

func (e *ErrorInjector) GetError(operation string) error {
    return e.errors[operation]
}

// 可测试的服务
type UserService struct {
    injector *ErrorInjector
}

func (s *UserService) CreateUser(user *User) error {
    if s.injector != nil && s.injector.ShouldFail("create_user") {
        return s.injector.GetError("create_user")
    }

    // 正常的创建逻辑
    return nil
}

// 测试
func TestUserService_CreateUser_Error(t *testing.T) {
    injector := &ErrorInjector{
        shouldFail: map[string]bool{"create_user": true},
        errors:     map[string]error{"create_user": errors.New("database error")},
    }

    service := &UserService{injector: injector}

    err := service.CreateUser(&User{Name: "Test"})
    if err == nil {
        t.Error("expected error but got none")
    }
}
```

## 7. 最佳实践

### 错误消息设计

```go
// ✅ 好的错误消息
func goodError() error {
    return fmt.Errorf("failed to connect to database at %s:%d: %w",
        host, port, originalError)
}

// ❌ 不好的错误消息
func badError() error {
    return errors.New("error")  // 信息太少
}

// ✅ 结构化错误信息
type DatabaseError struct {
    Host      string    `json:"host"`
    Port      int       `json:"port"`
    Database  string    `json:"database"`
    Operation string    `json:"operation"`
    Timestamp time.Time `json:"timestamp"`
    Cause     error     `json:"-"`
}

func (e *DatabaseError) Error() string {
    return fmt.Sprintf("database error [%s:%d/%s] during %s at %s: %v",
        e.Host, e.Port, e.Database, e.Operation,
        e.Timestamp.Format(time.RFC3339), e.Cause)
}
```

### 错误级别分类

```go
type ErrorLevel int

const (
    ErrorLevelInfo ErrorLevel = iota
    ErrorLevelWarning
    ErrorLevelError
    ErrorLevelCritical
)

type LeveledError struct {
    Level   ErrorLevel
    Message string
    Cause   error
}

func (e *LeveledError) Error() string {
    levels := map[ErrorLevel]string{
        ErrorLevelInfo:     "INFO",
        ErrorLevelWarning:  "WARN",
        ErrorLevelError:    "ERROR",
        ErrorLevelCritical: "CRITICAL",
    }

    return fmt.Sprintf("[%s] %s", levels[e.Level], e.Message)
}

// 创建不同级别的错误
func NewInfoError(message string) error {
    return &LeveledError{Level: ErrorLevelInfo, Message: message}
}

func NewCriticalError(message string, cause error) error {
    return &LeveledError{
        Level:   ErrorLevelCritical,
        Message: message,
        Cause:   cause,
    }
}
```

## 核心原则

1. **明确性**：错误消息应该清楚说明发生了什么
2. **可操作性**：提供足够信息让调用者知道如何处理
3. **上下文**：包含相关的上下文信息
4. **分层**：在每层添加适当的上下文，不要丢失原始错误
5. **一致性**：在整个项目中保持错误处理的一致性
