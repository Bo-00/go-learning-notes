# Go 控制结构

## 1. 条件语句

### if 语句

```go
// 基本 if 语句
var age = 18
if age >= 18 {
    fmt.Println("成年人")
}

// if-else 语句
if age >= 18 {
    fmt.Println("成年人")
} else {
    fmt.Println("未成年人")
}

// if-else if-else 语句
var score = 85
if score >= 90 {
    fmt.Println("优秀")
} else if score >= 80 {
    fmt.Println("良好")
} else if score >= 60 {
    fmt.Println("及格")
} else {
    fmt.Println("不及格")
}

// if 语句的初始化
if err := someFunction(); err != nil {
    fmt.Printf("错误: %v\n", err)
    return
}
// err 的作用域仅在 if 块内

// 复合条件
if age >= 18 && hasLicense {
    fmt.Println("可以开车")
}
```

### switch 语句

```go
// 基本 switch
var day = "Monday"
switch day {
case "Monday":
    fmt.Println("星期一")
case "Tuesday":
    fmt.Println("星期二")
case "Wednesday", "Thursday", "Friday":
    fmt.Println("工作日")
default:
    fmt.Println("其他日期")
}

// switch 带初始化
switch hour := time.Now().Hour(); {
case hour < 12:
    fmt.Println("上午")
case hour < 18:
    fmt.Println("下午")
default:
    fmt.Println("晚上")
}

// 无表达式的 switch（相当于 if-else if）
var score = 85
switch {
case score >= 90:
    fmt.Println("A级")
case score >= 80:
    fmt.Println("B级")
case score >= 70:
    fmt.Println("C级")
default:
    fmt.Println("D级")
}

// 类型选择 switch
var value interface{} = 42
switch v := value.(type) {
case int:
    fmt.Printf("整数: %d\n", v)
case string:
    fmt.Printf("字符串: %s\n", v)
case bool:
    fmt.Printf("布尔值: %t\n", v)
default:
    fmt.Printf("未知类型: %T\n", v)
}

// fallthrough 关键字（谨慎使用）
switch day {
case "Saturday":
    fmt.Println("周末")
    fallthrough
case "Sunday":
    fmt.Println("休息日")
}
```

## 2. 循环语句

### for 循环

```go
// 传统 for 循环
for i := 0; i < 10; i++ {
    fmt.Printf("%d ", i)
}

// 条件循环（类似 while）
var i = 0
for i < 10 {
    fmt.Printf("%d ", i)
    i++
}

// 无限循环
for {
    fmt.Println("无限循环")
    if someCondition {
        break
    }
}

// 遍历数组/切片
var numbers = []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("索引: %d, 值: %d\n", index, value)
}

// 只要索引
for index := range numbers {
    fmt.Printf("索引: %d\n", index)
}

// 只要值（使用空白标识符）
for _, value := range numbers {
    fmt.Printf("值: %d\n", value)
}

// 遍历字符串
var str = "Hello, 世界"
for index, runeValue := range str {
    fmt.Printf("索引: %d, 字符: %c\n", index, runeValue)
}

// 遍历映射
var m = map[string]int{
    "apple":  5,
    "banana": 3,
    "orange": 8,
}
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}

// 遍历通道
ch := make(chan int, 3)
ch <- 1
ch <- 2
ch <- 3
close(ch)

for value := range ch {
    fmt.Printf("接收到: %d\n", value)
}
```

## 3. 控制跳转

### break 语句

```go
// 跳出当前循环
for i := 0; i < 10; i++ {
    if i == 5 {
        break
    }
    fmt.Printf("%d ", i)
}

// 标签跳转（跳出嵌套循环）
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 {
            break outer
        }
        fmt.Printf("(%d,%d) ", i, j)
    }
}

// switch 中的 break
switch day {
case "Monday":
    fmt.Println("星期一")
    break // 可选，switch 默认会 break
case "Tuesday":
    fmt.Println("星期二")
}
```

### continue 语句

```go
// 跳过当前迭代
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue // 跳过偶数
    }
    fmt.Printf("%d ", i)
}

// 标签跳转
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if j == 1 {
            continue outer
        }
        fmt.Printf("(%d,%d) ", i, j)
    }
}
```

### goto 语句（不推荐使用）

```go
func example() {
    i := 0

loop:
    if i < 5 {
        fmt.Printf("%d ", i)
        i++
        goto loop
    }
}
```

## 4. defer 语句

### 基本用法

```go
func example() {
    defer fmt.Println("最后执行")
    fmt.Println("第一个")
    fmt.Println("第二个")
}
// 输出：
// 第一个
// 第二个
// 最后执行

// 多个 defer 语句（LIFO - 后进先出）
func multipleDefer() {
    defer fmt.Println("defer 1")
    defer fmt.Println("defer 2")
    defer fmt.Println("defer 3")
    fmt.Println("正常执行")
}
// 输出：
// 正常执行
// defer 3
// defer 2
// defer 1
```

### 实际应用场景

```go
// 文件操作
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close() // 确保文件会被关闭

    // 文件操作...
    return nil
}

// 互斥锁
func safeoperation() {
    mu.Lock()
    defer mu.Unlock() // 确保解锁

    // 临界区代码...
}

// 性能测量
func benchmark() {
    start := time.Now()
    defer func() {
        fmt.Printf("耗时: %v\n", time.Since(start))
    }()

    // 执行需要测量的代码...
}

// 错误恢复
func recoverExample() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("恢复自 panic: %v\n", r)
        }
    }()

    panic("这是一个 panic")
}
```

### defer 的陷阱

```go
// 陷阱1：defer 中的变量捕获
func deferTrap1() {
    for i := 0; i < 3; i++ {
        defer fmt.Printf("i = %d\n", i) // 立即求值
    }
}
// 输出：i = 2, i = 1, i = 0

// 陷阱2：defer 函数返回值
func deferTrap2() (result int) {
    defer func() {
        result++
    }()
    return 5
}
// 返回 6，不是 5

// 正确的用法
func correctDefer() {
    for i := 0; i < 3; i++ {
        defer func(val int) {
            fmt.Printf("i = %d\n", val)
        }(i)
    }
}
```

## 5. 异常处理

### panic 和 recover

```go
// panic 会停止正常的程序流程
func causPanic() {
    panic("出现了严重错误")
}

// recover 只能在 defer 函数中调用
func handlePanic() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("捕获到 panic: %v\n", r)
        }
    }()

    causPanic()
    fmt.Println("这行不会被执行")
}

// 实际应用：Web 服务器的中间件
func middleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("Panic: %v", err)
                http.Error(w, "Internal Server Error", 500)
            }
        }()

        next(w, r)
    }
}
```

## 6. 高级控制模式

### 状态机模式

```go
type State int

const (
    StateStart State = iota
    StateProcessing
    StateDone
    StateError
)

func stateMachine() {
    state := StateStart

    for {
        switch state {
        case StateStart:
            fmt.Println("开始处理")
            state = StateProcessing
        case StateProcessing:
            if processData() {
                state = StateDone
            } else {
                state = StateError
            }
        case StateDone:
            fmt.Println("处理完成")
            return
        case StateError:
            fmt.Println("处理错误")
            return
        }
    }
}
```

### select 语句（用于通道操作）

```go
func selectExample() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "来自 ch1"
    }()

    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "来自 ch2"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println(msg1)
        case msg2 := <-ch2:
            fmt.Println(msg2)
        case <-time.After(3 * time.Second):
            fmt.Println("超时")
            return
        default:
            fmt.Println("没有就绪的通道")
            time.Sleep(500 * time.Millisecond)
        }
    }
}
```

## 练习题

1. 编写一个函数，计算斐波那契数列的第 n 项
2. 实现一个简单的计算器，支持四则运算
3. 编写程序判断一个数是否为质数
4. 使用 defer 实现一个简单的函数调用计时器
5. 编写一个安全的除法函数，使用 recover 处理除零错误

## 最佳实践

1. **避免深层嵌套**：使用 early return 模式
2. **defer 的使用**：确保资源清理和错误恢复
3. **循环优化**：合理使用 break 和 continue
4. **错误处理**：优先使用 error 返回值，谨慎使用 panic
5. **select 超时**：避免无限阻塞
