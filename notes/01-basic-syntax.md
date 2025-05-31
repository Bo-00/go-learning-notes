# Go 基础语法

## 1. 程序结构

### 基本程序结构

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### 关键要点

- 每个 Go 文件都属于一个包 (package)
- `main` 包是程序的入口点
- `main()` 函数是程序执行的起点
- 使用 `import` 导入其他包

## 2. 变量声明

### 显式声明

```go
var name string = "Go"
var age int = 14
var isActive bool = true
```

### 类型推断

```go
var name = "Go"        // 自动推断为 string
var age = 14           // 自动推断为 int
```

### 短变量声明 (只能在函数内使用)

```go
name := "Go"
age := 14
isActive := true
```

### 多变量声明

```go
var (
    name     string = "Go"
    age      int    = 14
    isActive bool   = true
)

// 或者
var name, age, isActive = "Go", 14, true
```

## 3. 常量

### 基本常量

```go
const PI = 3.14159
const name = "Go"
```

### 常量组

```go
const (
    RED   = 0
    GREEN = 1
    BLUE  = 2
)
```

### iota 枚举器

```go
const (
    Sunday = iota    // 0
    Monday           // 1
    Tuesday          // 2
    Wednesday        // 3
    Thursday         // 4
    Friday           // 5
    Saturday         // 6
)
```

## 4. 函数

### 基本函数

```go
func add(a, b int) int {
    return a + b
}
```

### 多返回值

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

### 命名返回值

```go
func rectangle(length, width float64) (area, perimeter float64) {
    area = length * width
    perimeter = 2 * (length + width)
    return // 自动返回命名的返回值
}
```

### 可变参数

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
```

## 5. 作用域和可见性

### 包级别作用域

- 大写字母开头的标识符是**公开的** (exported)
- 小写字母开头的标识符是**私有的** (unexported)

```go
var PublicVar = "可以被其他包访问"
var privateVar = "只能在当前包内访问"

func PublicFunc() {
    // 公开函数
}

func privateFunc() {
    // 私有函数
}
```

### 块级作用域

```go
func example() {
    x := 10        // 函数作用域

    if true {
        y := 20    // if 块作用域
        fmt.Println(x, y) // 可以访问 x 和 y
    }

    // fmt.Println(y) // 错误：y 超出作用域
}
```

## 6. 注释

### 单行注释

```go
// 这是单行注释
var x int // 行尾注释
```

### 多行注释

```go
/*
这是多行注释
可以跨越多行
*/
```

### 文档注释

```go
// Package math 提供基本的数学函数
package math

// Add 计算两个整数的和
// 参数 a, b 是要相加的整数
// 返回 a + b 的结果
func Add(a, b int) int {
    return a + b
}
```

## 7. 包管理

### 导入包

```go
import "fmt"                    // 单个包
import "math"

// 或者
import (
    "fmt"                       // 分组导入
    "math"
    "strings"
)
```

### 包别名

```go
import (
    f "fmt"                     // 别名
    . "math"                    // 点导入，可直接使用 Sin, Cos 等
    _ "database/sql/driver"     // 仅执行包的 init 函数
)
```

## 8. init 函数

```go
package main

import "fmt"

// init 函数在 main 函数之前自动执行
func init() {
    fmt.Println("初始化包")
}

func main() {
    fmt.Println("主函数")
}

// 输出:
// 初始化包
// 主函数
```

## 练习题

1. 编写一个函数 `greeting(name string) string`，返回 "Hello, {name}!"
2. 使用 iota 创建一个表示星期的常量组
3. 编写一个可变参数函数计算多个数字的平均值
4. 创建一个包含多个返回值的函数，同时返回最大值和最小值
