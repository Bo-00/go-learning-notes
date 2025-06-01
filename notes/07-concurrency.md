# Go 语言并发编程

## 目录

- [并发基础概念](#并发基础概念)
- [CSP 并发模型](#csp-并发模型)
- [Goroutine](#goroutine)
- [Channel](#channel)
- [Channel 内部实现机制](#channel-内部实现机制)
- [常见并发模式](#常见并发模式)
- [并发安全与同步](#并发安全与同步)
- [性能优化与调试](#性能优化与调试)

---

## 并发基础概念

### 并发 vs 并行

- **并发（Concurrency）**：同时处理多个任务的能力，不一定同时执行
- **并行（Parallelism）**：同时执行多个任务的能力，需要多核支持

### Go 的并发哲学

```
Don't communicate by sharing memory; share memory by communicating.
不要通过共享内存来通信，要通过通信来共享内存。
```

### 并发 vs 异步

- **异步编程**：回调、Promise、async/await
- **Go 并发**：goroutine + channel，同步式编程模型

---

## CSP 并发模型

### CSP（Communicating Sequential Processes）原理

- **Sequential Processes**：顺序执行的进程
- **Communication**：通过消息传递进行通信
- **No Shared State**：避免共享状态

### CSP 在 Go 中的实现

```go
// CSP 模型的核心：Process + Channel
Process1 ---Message---> Channel ---Message---> Process2
```

### CSP vs Actor 模型对比

| 特性     | CSP (Go)  | Actor (Erlang) |
| -------- | --------- | -------------- |
| 通信方式 | Channel   | Mailbox        |
| 发送模式 | 同步/异步 | 异步           |
| 状态管理 | 无状态    | 有状态         |
| 错误处理 | 显式      | 监督树         |

---

## Goroutine

### 基本概念

- **轻量级线程**：栈空间初始 2KB，可动态增长
- **M:N 调度**：M 个 goroutine 映射到 N 个系统线程
- **协作式调度**：在函数调用、channel 操作、系统调用时让出控制权

### Goroutine 调度器（GPM 模型）

```
G: Goroutine，用户级轻量级线程
P: Processor，处理器，管理 goroutine 队列
M: Machine，系统线程，执行 goroutine
```

#### 调度器工作原理

1. **本地队列**：每个 P 维护一个本地 runqueue
2. **全局队列**：当本地队列满时，部分 G 会放入全局队列
3. **工作窃取**：当 P 的本地队列为空时，会从其他 P 或全局队列窃取 G
4. **系统调用处理**：发生系统调用时，M 与 P 分离，避免阻塞其他 goroutine

### Goroutine 生命周期

```
Runnable -> Running -> Waiting/Dead
    ^          |         |
    |          v         |
    +-------- Blocked <--+
```

#### 状态转换

- **Runnable**：等待被调度
- **Running**：正在执行
- **Waiting**：等待某个条件（channel、锁、系统调用等）
- **Dead**：执行完毕

### 创建和管理 Goroutine

```go
// 基本创建
go func() {
    // goroutine 代码
}()

// 带参数
go func(name string) {
    fmt.Printf("Hello, %s\n", name)
}("World")
```

### Goroutine 最佳实践

1. **避免 goroutine 泄露**：确保 goroutine 能正常退出
2. **使用 context 控制生命周期**
3. **合理控制 goroutine 数量**：避免创建过多 goroutine
4. **panic 处理**：goroutine 中的 panic 不会被其他 goroutine 捕获

---

## Channel

### 基本概念

- **类型安全的管道**：在 goroutine 间传递特定类型的值
- **同步机制**：无缓冲 channel 提供同步保证
- **CSP 的核心**：实现 "通过通信来共享内存"

### Channel 类型

#### 无缓冲 Channel

```go
ch := make(chan int)        // 无缓冲
ch <- 42                    // 发送（阻塞）
value := <-ch               // 接收（阻塞）
```

特点：

- **同步语义**：发送和接收必须同时就绪
- **零容量**：无法存储数据
- **握手机制**：发送方和接收方直接交换数据

#### 缓冲 Channel

```go
ch := make(chan int, 3)     // 缓冲大小为 3
ch <- 1                     // 不阻塞（缓冲未满）
ch <- 2
ch <- 3
ch <- 4                     // 阻塞（缓冲已满）
```

特点：

- **异步语义**：缓冲区未满时发送不阻塞
- **FIFO 队列**：先进先出
- **容量固定**：创建时确定容量

### Channel 操作

#### 发送和接收

```go
// 发送
ch <- value

// 接收
value := <-ch
value, ok := <-ch    // ok 表示 channel 是否关闭

// 单向 channel
var sendOnly chan<- int = ch    // 只能发送
var recvOnly <-chan int = ch    // 只能接收
```

#### 关闭 Channel

```go
close(ch)

// 检查是否关闭
value, ok := <-ch
if !ok {
    // channel 已关闭
}

// 使用 range 遍历
for value := range ch {
    // 处理 value
    // channel 关闭后自动退出循环
}
```

#### Select 语句

```go
select {
case value := <-ch1:
    // 处理 ch1 的数据
case ch2 <- value:
    // 向 ch2 发送数据
case <-timeout:
    // 超时处理
default:
    // 非阻塞操作
}
```

---

## Channel 内部实现机制

### Channel 数据结构

```go
type hchan struct {
    qcount   uint           // 队列中数据个数
    dataqsiz uint           // 循环队列大小
    buf      unsafe.Pointer // 指向缓冲区
    elemsize uint16         // 元素大小
    closed   uint32         // 关闭标志
    elemtype *_type         // 元素类型
    sendx    uint           // 发送索引
    recvx    uint           // 接收索引
    recvq    waitq          // 接收等待队列
    sendq    waitq          // 发送等待队列
    lock     mutex          // 保护所有字段的锁
}
```

### 发送操作流程

1. **检查接收者**：如果有等待的接收者，直接传递数据
2. **缓冲区处理**：如果缓冲区有空间，将数据放入缓冲区
3. **阻塞等待**：将发送者加入发送等待队列

### 接收操作流程

1. **检查缓冲区**：如果缓冲区有数据，直接取出
2. **检查发送者**：如果有等待的发送者，直接接收数据
3. **阻塞等待**：将接收者加入接收等待队列

### 关闭操作

1. **设置关闭标志**
2. **唤醒所有等待者**：接收者接收零值，发送者触发 panic
3. **释放资源**

### 内存模型

- **happens-before 关系**：channel 操作建立内存同步点
- **无缓冲 channel**：发送 happens-before 对应的接收
- **缓冲 channel**：第 k 次接收 happens-before 第 k+C 次发送完成

---

## 常见并发模式

### 1. Generator Pattern（生成器模式）

```go
func generator() <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch)
        for i := 0; i < 10; i++ {
            ch <- i
        }
    }()
    return ch
}
```

### 2. Fan-out Pattern（扇出模式）

将一个输入分发给多个处理器：

```go
func fanOut(input <-chan int, workers int) []<-chan int {
    outputs := make([]<-chan int, workers)
    for i := 0; i < workers; i++ {
        output := make(chan int)
        outputs[i] = output

        go func() {
            defer close(output)
            for value := range input {
                output <- process(value)
            }
        }()
    }
    return outputs
}
```

### 3. Fan-in Pattern（扇入模式）

将多个输入合并到一个输出：

```go
func fanIn(inputs ...<-chan int) <-chan int {
    output := make(chan int)
    var wg sync.WaitGroup

    wg.Add(len(inputs))
    for _, input := range inputs {
        go func(ch <-chan int) {
            defer wg.Done()
            for value := range ch {
                output <- value
            }
        }(input)
    }

    go func() {
        wg.Wait()
        close(output)
    }()

    return output
}
```

### 4. Pipeline Pattern（管道模式）

```go
// Stage 1: 生成数据
func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums {
            out <- n
        }
    }()
    return out
}

// Stage 2: 处理数据
func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            out <- n * n
        }
    }()
    return out
}

// Pipeline 组合
func main() {
    // 设置管道
    c := generate(2, 3, 4)
    out := square(c)

    // 消费结果
    for result := range out {
        fmt.Println(result)
    }
}
```

### 5. Worker Pool Pattern（工作池模式）

```go
type Job struct {
    ID   int
    Data interface{}
}

type Result struct {
    Job    Job
    Output interface{}
    Error  error
}

func workerPool(jobs <-chan Job, results chan<- Result, numWorkers int) {
    var wg sync.WaitGroup

    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for job := range jobs {
                output, err := processJob(job)
                results <- Result{
                    Job:    job,
                    Output: output,
                    Error:  err,
                }
            }
        }()
    }

    go func() {
        wg.Wait()
        close(results)
    }()
}
```

### 6. Publish-Subscribe Pattern（发布订阅模式）

```go
type PubSub struct {
    mu          sync.RWMutex
    subscribers map[string][]chan interface{}
}

func (ps *PubSub) Subscribe(topic string) <-chan interface{} {
    ps.mu.Lock()
    defer ps.mu.Unlock()

    ch := make(chan interface{}, 1)
    ps.subscribers[topic] = append(ps.subscribers[topic], ch)
    return ch
}

func (ps *PubSub) Publish(topic string, data interface{}) {
    ps.mu.RLock()
    defer ps.mu.RUnlock()

    for _, ch := range ps.subscribers[topic] {
        select {
        case ch <- data:
        default:
            // 订阅者处理太慢，跳过
        }
    }
}
```

---

## 并发安全与同步

### Data Race 问题

```go
// 数据竞争示例
var counter int

func increment() {
    counter++  // 非原子操作，存在竞争
}

// 检测工具
// go run -race main.go
```

### 同步原语

#### 1. Mutex（互斥锁）

```go
var mu sync.Mutex
var counter int

func safeIncrement() {
    mu.Lock()
    counter++
    mu.Unlock()
}

// 或使用 defer
func safeIncrement() {
    mu.Lock()
    defer mu.Unlock()
    counter++
}
```

#### 2. RWMutex（读写锁）

```go
var rwmu sync.RWMutex
var data map[string]int = make(map[string]int)

func readData(key string) int {
    rwmu.RLock()
    defer rwmu.RUnlock()
    return data[key]
}

func writeData(key string, value int) {
    rwmu.Lock()
    defer rwmu.Unlock()
    data[key] = value
}
```

#### 3. Once（一次性执行）

```go
var once sync.Once
var instance *Singleton

func getInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{}
    })
    return instance
}
```

#### 4. WaitGroup（等待组）

```go
var wg sync.WaitGroup

func main() {
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Printf("Worker %d\n", id)
        }(i)
    }
    wg.Wait()
    fmt.Println("All workers done")
}
```

#### 5. Atomic（原子操作）

```go
var counter int64

func atomicIncrement() {
    atomic.AddInt64(&counter, 1)
}

func getCounter() int64 {
    return atomic.LoadInt64(&counter)
}
```

### Context（上下文）

```go
// 带超时的 context
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// 带取消的 context
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// 在 goroutine 中使用
go func(ctx context.Context) {
    select {
    case <-ctx.Done():
        fmt.Println("Context cancelled")
        return
    case <-time.After(time.Second):
        fmt.Println("Work completed")
    }
}(ctx)
```

### 内存模型和 Happens-Before

```go
// Channel 同步
var a, b int

func f() {
    a = 1           // 1
    b = 2           // 2
    ch <- struct{}{}    // 3
}

func g() {
    <-ch                // 4
    fmt.Print(a, b)     // 5
}

// happens-before 关系：1 -> 2 -> 3 -> 4 -> 5
// 保证输出 "12"
```

---

## 性能优化与调试

### 性能分析工具

#### 1. go tool trace

```bash
# 生成 trace 文件
go test -trace=trace.out

# 分析 trace
go tool trace trace.out
```

#### 2. go tool pprof

```go
import _ "net/http/pprof"

func main() {
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    // 应用代码
}
```

#### 3. runtime 包

```go
// 获取 goroutine 数量
fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine())

// 获取 CPU 数量
fmt.Printf("CPUs: %d\n", runtime.NumCPU())

// 设置最大 CPU 使用数
runtime.GOMAXPROCS(runtime.NumCPU())
```

### 常见性能问题

#### 1. Goroutine 泄露

```go
// 问题：goroutine 无法退出
func leak() {
    ch := make(chan int)
    go func() {
        for {
            // 永远不会收到数据
            <-ch
        }
    }()
}

// 解决：使用 context 控制
func fixed(ctx context.Context) {
    ch := make(chan int)
    go func() {
        for {
            select {
            case <-ch:
                // 处理数据
            case <-ctx.Done():
                return
            }
        }
    }()
}
```

#### 2. Channel 容量设计

```go
// 问题：缓冲区太小导致阻塞
ch := make(chan int, 1)

// 解决：合理设置缓冲区大小
ch := make(chan int, 100)

// 或使用无缓冲 channel 实现同步
ch := make(chan int)
```

#### 3. 锁竞争优化

```go
// 问题：粗粒度锁
var mu sync.Mutex
var data map[string]int

func updateData(key string, value int) {
    mu.Lock()
    defer mu.Unlock()
    data[key] = value
    // 其他耗时操作
}

// 优化：细粒度锁
func updateDataOptimized(key string, value int) {
    mu.Lock()
    data[key] = value
    mu.Unlock()
    // 其他操作移到锁外
}
```

### 最佳实践

#### 1. 资源管理

- 总是关闭不再使用的 channel
- 使用 defer 确保资源释放
- 合理控制 goroutine 数量

#### 2. 错误处理

- Goroutine 中的 panic 要恰当处理
- 使用 errgroup 管理多个 goroutine 的错误

#### 3. 测试

```go
func TestConcurrent(t *testing.T) {
    // 并发测试
    const numGoroutines = 100
    var wg sync.WaitGroup

    for i := 0; i < numGoroutines; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            // 测试逻辑
        }()
    }
    wg.Wait()
}
```

---

## 总结

### 核心要点

1. **CSP 模型**：通过通信来共享内存
2. **Goroutine**：轻量级、高效的并发单元
3. **Channel**：类型安全的通信机制
4. **并发模式**：解决常见并发问题的设计模式
5. **同步原语**：保证并发安全的工具
6. **性能优化**：避免常见陷阱，提高并发性能

### 学习建议

1. 理解 CSP 模型的哲学
2. 掌握 goroutine 和 channel 的基本用法
3. 学会使用常见并发模式
4. 重视并发安全和性能分析
5. 多练习、多调试、多测试

### 进阶方向

- 深入理解 Go 调度器实现
- 学习分布式并发模式
- 掌握高性能并发编程技巧
- 研究其他语言的并发模型对比
