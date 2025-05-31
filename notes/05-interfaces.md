# Go 接口设计模式和最佳实践

## 1. 接口基础

### 接口定义和实现

```go
// 定义接口
type Writer interface {
    Write([]byte) (int, error)
}

type Reader interface {
    Read([]byte) (int, error)
}

// 实现接口（隐式实现）
type FileWriter struct {
    filename string
}

func (f *FileWriter) Write(data []byte) (int, error) {
    // 实现写入逻辑
    return len(data), nil
}

// 接口组合
type ReadWriter interface {
    Reader
    Writer
}

type Closer interface {
    Close() error
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}
```

### 空接口和类型断言

```go
// 空接口可以接受任何类型
var anything interface{}
anything = 42
anything = "hello"
anything = []int{1, 2, 3}

// 类型断言
func processData(data interface{}) {
    // 安全的类型断言
    if str, ok := data.(string); ok {
        fmt.Printf("字符串: %s\n", str)
    }

    if num, ok := data.(int); ok {
        fmt.Printf("整数: %d\n", num)
    }
}

// 类型选择
func typeSwitch(data interface{}) {
    switch v := data.(type) {
    case string:
        fmt.Printf("字符串: %s\n", v)
    case int:
        fmt.Printf("整数: %d\n", v)
    case []int:
        fmt.Printf("整数切片: %v\n", v)
    default:
        fmt.Printf("未知类型: %T\n", v)
    }
}
```

## 2. 接口设计原则

### 接口应该小而专一

```go
// ❌ 过大的接口
type BadDatabase interface {
    Connect() error
    Disconnect() error
    CreateTable(string) error
    Insert(string, interface{}) error
    Update(string, interface{}) error
    Delete(string, string) error
    Select(string) ([]interface{}, error)
    Backup() error
    Restore() error
}

// ✅ 小而专一的接口
type Connector interface {
    Connect() error
    Disconnect() error
}

type TableCreator interface {
    CreateTable(string) error
}

type DataWriter interface {
    Insert(string, interface{}) error
    Update(string, interface{}) error
    Delete(string, string) error
}

type DataReader interface {
    Select(string) ([]interface{}, error)
}
```

### 接受接口，返回结构体

```go
// ✅ 函数参数使用接口
func ProcessData(r io.Reader) error {
    data, err := io.ReadAll(r)
    if err != nil {
        return err
    }
    // 处理数据...
    return nil
}

// ✅ 返回具体类型
func NewFileReader(filename string) *os.File {
    file, _ := os.Open(filename)
    return file
}

// 这样的设计允许调用者传入任何实现了 io.Reader 的类型
func example() {
    // 可以传入文件
    file, _ := os.Open("data.txt")
    ProcessData(file)

    // 可以传入字符串
    ProcessData(strings.NewReader("hello"))

    // 可以传入字节缓冲区
    ProcessData(bytes.NewBuffer([]byte("data")))
}
```

## 3. 常用接口模式

### 策略模式

```go
// 策略接口
type SortStrategy interface {
    Sort([]int)
}

// 具体策略实现
type BubbleSort struct{}
func (b BubbleSort) Sort(data []int) {
    // 冒泡排序实现
}

type QuickSort struct{}
func (q QuickSort) Sort(data []int) {
    // 快速排序实现
}

// 上下文
type Sorter struct {
    strategy SortStrategy
}

func (s *Sorter) SetStrategy(strategy SortStrategy) {
    s.strategy = strategy
}

func (s *Sorter) Sort(data []int) {
    s.strategy.Sort(data)
}

// 使用
func useStrategy() {
    sorter := &Sorter{}
    data := []int{3, 1, 4, 1, 5, 9}

    sorter.SetStrategy(BubbleSort{})
    sorter.Sort(data)

    sorter.SetStrategy(QuickSort{})
    sorter.Sort(data)
}
```

### 适配器模式

```go
// 目标接口
type MediaPlayer interface {
    Play(filename string) error
}

// 需要适配的接口
type AdvancedMediaPlayer interface {
    PlayVlc(filename string) error
    PlayMp4(filename string) error
}

// 具体的高级播放器
type VlcPlayer struct{}
func (v *VlcPlayer) PlayVlc(filename string) error {
    fmt.Printf("Playing vlc file: %s\n", filename)
    return nil
}

type Mp4Player struct{}
func (m *Mp4Player) PlayMp4(filename string) error {
    fmt.Printf("Playing mp4 file: %s\n", filename)
    return nil
}

// 适配器
type MediaAdapter struct {
    player AdvancedMediaPlayer
}

func NewMediaAdapter(audioType string) *MediaAdapter {
    switch audioType {
    case "vlc":
        return &MediaAdapter{&VlcPlayer{}}
    case "mp4":
        return &MediaAdapter{&Mp4Player{}}
    }
    return nil
}

func (m *MediaAdapter) Play(filename string) error {
    if vlc, ok := m.player.(*VlcPlayer); ok {
        return vlc.PlayVlc(filename)
    }
    if mp4, ok := m.player.(*Mp4Player); ok {
        return mp4.PlayMp4(filename)
    }
    return fmt.Errorf("unsupported format")
}

// 音频播放器
type AudioPlayer struct {
    adapter *MediaAdapter
}

func (a *AudioPlayer) Play(audioType, filename string) error {
    switch audioType {
    case "mp3":
        fmt.Printf("Playing mp3 file: %s\n", filename)
        return nil
    case "vlc", "mp4":
        a.adapter = NewMediaAdapter(audioType)
        return a.adapter.Play(filename)
    default:
        return fmt.Errorf("invalid media. %s format not supported", audioType)
    }
}
```

### 装饰器模式

```go
// 基础组件接口
type Coffee interface {
    Cost() float64
    Description() string
}

// 基础实现
type SimpleCoffee struct{}

func (s SimpleCoffee) Cost() float64 {
    return 2.0
}

func (s SimpleCoffee) Description() string {
    return "Simple coffee"
}

// 装饰器基类
type CoffeeDecorator struct {
    coffee Coffee
}

func (c CoffeeDecorator) Cost() float64 {
    return c.coffee.Cost()
}

func (c CoffeeDecorator) Description() string {
    return c.coffee.Description()
}

// 具体装饰器
type MilkDecorator struct {
    CoffeeDecorator
}

func NewMilkDecorator(coffee Coffee) *MilkDecorator {
    return &MilkDecorator{CoffeeDecorator{coffee}}
}

func (m MilkDecorator) Cost() float64 {
    return m.coffee.Cost() + 0.5
}

func (m MilkDecorator) Description() string {
    return m.coffee.Description() + ", milk"
}

type SugarDecorator struct {
    CoffeeDecorator
}

func NewSugarDecorator(coffee Coffee) *SugarDecorator {
    return &SugarDecorator{CoffeeDecorator{coffee}}
}

func (s SugarDecorator) Cost() float64 {
    return s.coffee.Cost() + 0.2
}

func (s SugarDecorator) Description() string {
    return s.coffee.Description() + ", sugar"
}

// 使用装饰器
func useDecorator() {
    coffee := SimpleCoffee{}
    fmt.Printf("%s: $%.2f\n", coffee.Description(), coffee.Cost())

    milkCoffee := NewMilkDecorator(coffee)
    fmt.Printf("%s: $%.2f\n", milkCoffee.Description(), milkCoffee.Cost())

    sweetMilkCoffee := NewSugarDecorator(milkCoffee)
    fmt.Printf("%s: $%.2f\n", sweetMilkCoffee.Description(), sweetMilkCoffee.Cost())
}
```

## 4. 函数式接口

### 函数类型实现接口

```go
// HTTP 处理器接口
type Handler interface {
    ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// 函数类型
type HandlerFunc func(http.ResponseWriter, *http.Request)

// 让函数类型实现接口
func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    f(w, r)
}

// 使用
func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to home page!")
}

func main() {
    // 将普通函数转换为接口实现
    http.Handle("/", HandlerFunc(homeHandler))

    // 或者使用匿名函数
    http.Handle("/about", HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "About page")
    }))
}
```

### 错误处理接口

```go
// 错误接口
type Error interface {
    Error() string
}

// 自定义错误类型
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}

// 包装错误
type WrappedError struct {
    Cause   error
    Message string
}

func (e WrappedError) Error() string {
    return fmt.Sprintf("%s: %v", e.Message, e.Cause)
}

func (e WrappedError) Unwrap() error {
    return e.Cause
}

// 使用
func validateUser(user User) error {
    if user.Name == "" {
        return ValidationError{
            Field:   "name",
            Message: "name cannot be empty",
        }
    }
    return nil
}
```

## 5. 最佳实践

### 接口命名约定

```go
// 单方法接口通常以 -er 结尾
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

type Closer interface {
    Close() error
}

type Stringer interface {
    String() string
}

// 多方法接口使用描述性名称
type Database interface {
    Connect() error
    Execute(query string) error
    Close() error
}
```

### 接口版本演化

```go
// 版本1：基础接口
type FileProcessorV1 interface {
    Process(filename string) error
}

// 版本2：扩展功能，保持向后兼容
type FileProcessorV2 interface {
    FileProcessorV1
    ProcessWithOptions(filename string, opts Options) error
}

// 检查接口实现
func processFile(processor FileProcessorV1, filename string) error {
    // 尝试使用新版本功能
    if v2, ok := processor.(FileProcessorV2); ok {
        return v2.ProcessWithOptions(filename, DefaultOptions)
    }

    // 回退到基础功能
    return processor.Process(filename)
}
```

### 依赖注入

```go
// 服务接口
type UserService interface {
    GetUser(id string) (*User, error)
    CreateUser(user *User) error
}

type EmailService interface {
    SendEmail(to, subject, body string) error
}

// 控制器
type UserController struct {
    userService  UserService
    emailService EmailService
}

func NewUserController(userSvc UserService, emailSvc EmailService) *UserController {
    return &UserController{
        userService:  userSvc,
        emailService: emailSvc,
    }
}

func (c *UserController) CreateUser(userData UserData) error {
    user := &User{
        Name:  userData.Name,
        Email: userData.Email,
    }

    if err := c.userService.CreateUser(user); err != nil {
        return err
    }

    return c.emailService.SendEmail(
        user.Email,
        "Welcome!",
        "Welcome to our service!",
    )
}
```

### 测试中的接口

```go
// 使用接口便于测试
type PaymentService interface {
    ProcessPayment(amount float64) error
}

type OrderService struct {
    payment PaymentService
}

func (o *OrderService) CreateOrder(amount float64) error {
    return o.payment.ProcessPayment(amount)
}

// 测试时使用模拟实现
type MockPaymentService struct {
    shouldFail bool
}

func (m *MockPaymentService) ProcessPayment(amount float64) error {
    if m.shouldFail {
        return errors.New("payment failed")
    }
    return nil
}

func TestOrderService(t *testing.T) {
    mockPayment := &MockPaymentService{shouldFail: false}
    orderService := &OrderService{payment: mockPayment}

    err := orderService.CreateOrder(100.0)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }

    mockPayment.shouldFail = true
    err = orderService.CreateOrder(100.0)
    if err == nil {
        t.Error("Expected error, got nil")
    }
}
```

## 6. 性能考虑

### 接口调用开销

```go
// 直接调用
type ConcreteType struct {
    value int
}

func (c *ConcreteType) Method() int {
    return c.value
}

// 接口调用
type Interface interface {
    Method() int
}

func BenchmarkDirectCall(b *testing.B) {
    obj := &ConcreteType{value: 42}
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        _ = obj.Method()
    }
}

func BenchmarkInterfaceCall(b *testing.B) {
    var obj Interface = &ConcreteType{value: 42}
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        _ = obj.Method()
    }
}
```

## 总结

Go 接口的核心设计哲学：

1. **隐式实现** - 降低耦合
2. **组合优于继承** - 通过接口组合实现复杂功能
3. **小接口** - 易于理解和实现
4. **鸭子类型** - 关注行为而非类型
