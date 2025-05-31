# Go 数据类型

## 1. 基本数据类型

### 整型

```go
// 有符号整型
var i8 int8 = -128          // -2^7 到 2^7-1
var i16 int16 = -32768      // -2^15 到 2^15-1
var i32 int32 = -2147483648 // -2^31 到 2^31-1
var i64 int64 = -9223372036854775808 // -2^63 到 2^63-1

// 无符号整型
var u8 uint8 = 255          // 0 到 2^8-1
var u16 uint16 = 65535      // 0 到 2^16-1
var u32 uint32 = 4294967295 // 0 到 2^32-1
var u64 uint64 = 18446744073709551615 // 0 到 2^64-1

// 平台相关
var i int = 42              // 32位或64位
var u uint = 42             // 32位或64位
var ptr uintptr = 0x123     // 存储指针的整型
```

### 浮点型

```go
var f32 float32 = 3.14159   // 32位浮点数
var f64 float64 = 3.141592653589793 // 64位浮点数

// 科学记数法
var scientific = 1.23e4    // 12300.0
var negative = 1.23e-4     // 0.000123
```

### 复数型

```go
var c64 complex64 = 1 + 2i          // 32位实数和虚数
var c128 complex128 = 1 + 2i        // 64位实数和虚数

// 使用 complex 函数创建
var c = complex(1.0, 2.0)           // 1+2i

// 获取实数和虚数部分
real := real(c)                     // 1.0
imag := imag(c)                     // 2.0
```

### 布尔型

```go
var isTrue bool = true
var isFalse bool = false

// 逻辑运算
var result = isTrue && !isFalse     // true
```

### 字符串

```go
var str string = "Hello, 世界"
var multiline = `这是一个
跨越多行的
原始字符串`

// 字符串是不可变的
// str[0] = 'h' // 编译错误

// 字符串操作
length := len(str)                  // 字符串长度（字节数）
substr := str[0:5]                  // 子字符串
concat := str + " Go"               // 字符串连接
```

## 2. 复合数据类型

### 数组 (Array)

```go
// 声明和初始化
var arr1 [5]int                     // 零值初始化
var arr2 = [5]int{1, 2, 3, 4, 5}    // 指定初始值
var arr3 = [...]int{1, 2, 3}        // 自动推断长度

// 多维数组
var matrix [3][4]int

// 数组是值类型，赋值会复制整个数组
arr4 := arr2                        // 复制数组
arr4[0] = 10                        // 不影响 arr2
```

### 切片 (Slice)

```go
// 创建切片
var slice1 []int                    // nil 切片
var slice2 = []int{1, 2, 3, 4, 5}   // 初始化切片
var slice3 = make([]int, 5)         // 长度为5的切片
var slice4 = make([]int, 5, 10)     // 长度5，容量10

// 从数组或切片创建
arr := [5]int{1, 2, 3, 4, 5}
slice5 := arr[1:4]                  // [2, 3, 4]
slice6 := arr[:3]                   // [1, 2, 3]
slice7 := arr[2:]                   // [3, 4, 5]

// 切片操作
length := len(slice2)               // 长度
capacity := cap(slice2)             // 容量
slice2 = append(slice2, 6)          // 追加元素
slice2 = append(slice2, 7, 8, 9)    // 追加多个元素

// 切片是引用类型
slice8 := slice2                    // 共享底层数组
slice8[0] = 100                     // 影响 slice2
```

### 映射 (Map)

```go
// 创建映射
var map1 map[string]int             // nil 映射
var map2 = make(map[string]int)     // 空映射
var map3 = map[string]int{          // 初始化映射
    "apple":  5,
    "banana": 3,
    "orange": 8,
}

// 映射操作
map2["key1"] = 10                   // 设置值
value := map2["key1"]               // 获取值
value, ok := map2["key2"]           // 检查键是否存在
delete(map2, "key1")                // 删除键值对

// 遍历映射
for key, value := range map3 {
    fmt.Printf("%s: %d\n", key, value)
}
```

### 通道 (Channel)

```go
// 创建通道
var ch1 chan int                    // nil 通道
var ch2 = make(chan int)            // 无缓冲通道
var ch3 = make(chan int, 10)        // 有缓冲通道

// 发送和接收
ch2 <- 42                           // 发送值
value := <-ch2                      // 接收值
value, ok := <-ch2                  // 接收值并检查通道是否关闭

// 关闭通道
close(ch2)

// 方向性通道
var sendCh chan<- int = ch3          // 只能发送的通道
var recvCh <-chan int = ch3          // 只能接收的通道
```

## 3. 指针 (Pointer)

```go
var x int = 42
var p *int = &x                     // 获取 x 的地址
var value = *p                      // 解引用，获取指针指向的值

*p = 100                            // 通过指针修改值
fmt.Println(x)                      // 输出: 100

// 指针的零值是 nil
var p2 *int
if p2 == nil {
    fmt.Println("p2 is nil")
}

// Go 不支持指针运算
// p++  // 编译错误
```

## 4. 结构体 (Struct)

```go
// 定义结构体
type Person struct {
    Name    string
    Age     int
    Email   string
    private int  // 私有字段
}

// 创建结构体实例
var p1 Person                       // 零值初始化
var p2 = Person{
    Name:  "Alice",
    Age:   30,
    Email: "alice@example.com",
}
var p3 = Person{"Bob", 25, "bob@example.com", 0} // 按顺序初始化

// 结构体指针
var p4 = &Person{Name: "Charlie", Age: 35}
p4.Name = "Charlie Brown"           // 自动解引用

// 匿名结构体
var p5 = struct {
    Name string
    Age  int
}{
    Name: "Anonymous",
    Age:  20,
}

// 结构体嵌入（组合）
type Employee struct {
    Person                          // 嵌入 Person
    Department string
    Salary     float64
}

var emp = Employee{
    Person: Person{
        Name: "David",
        Age:  28,
    },
    Department: "Engineering",
    Salary:     75000,
}

// 可以直接访问嵌入字段
fmt.Println(emp.Name)               // 等同于 emp.Person.Name
```

## 5. 接口 (Interface)

```go
// 定义接口
type Writer interface {
    Write([]byte) (int, error)
}

type Reader interface {
    Read([]byte) (int, error)
}

// 组合接口
type ReadWriter interface {
    Reader
    Writer
}

// 空接口可以持有任何类型的值
var empty interface{}
empty = 42
empty = "hello"
empty = []int{1, 2, 3}

// 类型断言
if str, ok := empty.(string); ok {
    fmt.Println("empty is a string:", str)
}

// 类型选择
switch v := empty.(type) {
case int:
    fmt.Printf("Integer: %d\n", v)
case string:
    fmt.Printf("String: %s\n", v)
default:
    fmt.Printf("Unknown type: %T\n", v)
}
```

## 6. 类型别名和自定义类型

```go
// 类型别名（完全等价）
type String = string
var s String = "hello"

// 自定义类型（新类型）
type UserId int
type Temperature float64

var id UserId = 123
var temp Temperature = 36.5

// 为自定义类型定义方法
func (t Temperature) Celsius() float64 {
    return float64(t)
}

func (t Temperature) Fahrenheit() float64 {
    return float64(t)*9/5 + 32
}

// 使用
fmt.Printf("%.1f°C = %.1f°F\n", temp.Celsius(), temp.Fahrenheit())
```

## 7. 零值 (Zero Values)

```go
var i int           // 0
var f float64       // 0.0
var b bool          // false
var s string        // ""
var p *int          // nil
var slice []int     // nil
var m map[string]int // nil
var ch chan int     // nil
var fn func()       // nil
```

## 8. 类型转换

```go
var i int = 42
var f float64 = float64(i)          // 显式转换
var s string = fmt.Sprintf("%d", i) // 转换为字符串

// 字符串转换
import "strconv"

str := "123"
num, err := strconv.Atoi(str)       // 字符串转整数
if err != nil {
    fmt.Println("转换错误:", err)
}

floatStr := "3.14"
floatNum, err := strconv.ParseFloat(floatStr, 64)
```

## 练习题

1. 创建一个学生结构体，包含姓名、年龄、成绩（切片）
2. 编写函数计算学生的平均成绩
3. 使用映射存储多个学生的信息
4. 实现一个简单的栈（使用切片）
5. 定义一个接口并实现不同的类型

## 内存布局要点

- **值类型**：数组、结构体（直接存储值）
- **引用类型**：切片、映射、通道、指针（存储引用）
- **字符串**：不可变的字节序列
- **接口**：存储类型信息和值的指针
