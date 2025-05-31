package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 全局常量
const (
	AppName    = "Go学习示例"
	AppVersion = "1.0.0"
)

// 自定义类型
type Student struct {
	Name   string
	Age    int
	Scores []float64
}

// 方法：计算平均分
func (s Student) Average() float64 {
	if len(s.Scores) == 0 {
		return 0
	}

	total := 0.0
	for _, score := range s.Scores {
		total += score
	}
	return total / float64(len(s.Scores))
}

// 方法：添加分数
func (s *Student) AddScore(score float64) {
	s.Scores = append(s.Scores, score)
}

// 函数：创建学生
func NewStudent(name string, age int) *Student {
	return &Student{
		Name:   name,
		Age:    age,
		Scores: make([]float64, 0),
	}
}

// 函数：多返回值示例
func divideWithRemainder(a, b int) (quotient, remainder int, err error) {
	if b == 0 {
		return 0, 0, fmt.Errorf("division by zero")
	}
	quotient = a / b
	remainder = a % b
	return // 命名返回值，可以直接 return
}

// 可变参数函数
func calculateSum(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// 高阶函数示例
func applyOperation(numbers []int, operation func(int) int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = operation(num)
	}
	return result
}

func main() {
	fmt.Printf("=== %s v%s ===\n\n", AppName, AppVersion)

	// 1. 变量声明示例
	fmt.Println("1. 变量声明:")
	var name string = "张三"
	age := 20
	isStudent := true

	fmt.Printf("姓名: %s, 年龄: %d, 是学生: %t\n\n", name, age, isStudent)

	// 2. 数组和切片
	fmt.Println("2. 数组和切片:")
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	slice := []string{"Go", "Python", "Java", "JavaScript"}

	fmt.Printf("数组: %v\n", arr)
	fmt.Printf("切片: %v\n", slice)

	// 切片操作
	slice = append(slice, "Rust", "C++")
	fmt.Printf("添加元素后: %v\n\n", slice)

	// 3. 映射 (Map)
	fmt.Println("3. 映射 (Map):")
	scoreMap := map[string]int{
		"数学": 95,
		"英语": 88,
		"物理": 92,
	}

	scoreMap["化学"] = 90
	fmt.Printf("成绩单: %v\n", scoreMap)

	// 检查键是否存在
	if score, exists := scoreMap["生物"]; exists {
		fmt.Printf("生物成绩: %d\n", score)
	} else {
		fmt.Println("没有生物成绩")
	}
	fmt.Println()

	// 4. 结构体示例
	fmt.Println("4. 结构体示例:")
	student1 := NewStudent("李四", 19)
	student1.AddScore(95.5)
	student1.AddScore(88.0)
	student1.AddScore(92.5)

	fmt.Printf("学生信息: %+v\n", student1)
	fmt.Printf("平均分: %.2f\n\n", student1.Average())

	// 5. 控制结构示例
	fmt.Println("5. 控制结构:")

	// if-else
	if student1.Average() >= 90 {
		fmt.Println("优秀学生!")
	} else if student1.Average() >= 80 {
		fmt.Println("良好学生!")
	} else {
		fmt.Println("需要努力!")
	}

	// switch
	grade := "A"
	switch grade {
	case "A":
		fmt.Println("成绩等级: 优秀")
	case "B":
		fmt.Println("成绩等级: 良好")
	case "C":
		fmt.Println("成绩等级: 及格")
	default:
		fmt.Println("成绩等级: 不及格")
	}

	// for 循环
	fmt.Print("循环输出: ")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// range 循环
	fmt.Print("遍历切片: ")
	for index, value := range slice[:3] { // 只遍历前3个
		fmt.Printf("[%d]%s ", index, value)
	}
	fmt.Println("\n")

	// 6. 函数示例
	fmt.Println("6. 函数示例:")

	// 多返回值
	quotient, remainder, err := divideWithRemainder(17, 5)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("17 ÷ 5 = %d ... %d\n", quotient, remainder)
	}

	// 可变参数
	sum := calculateSum(1, 2, 3, 4, 5)
	fmt.Printf("求和结果: %d\n", sum)

	// 高阶函数
	numbers := []int{1, 2, 3, 4, 5}
	squared := applyOperation(numbers, func(x int) int {
		return x * x
	})
	fmt.Printf("原数组: %v\n", numbers)
	fmt.Printf("平方后: %v\n\n", squared)

	// 7. 字符串处理
	fmt.Println("7. 字符串处理:")
	text := "Go语言学习笔记"
	fmt.Printf("原字符串: %s\n", text)
	fmt.Printf("字符串长度: %d 字节\n", len(text))
	fmt.Printf("字符数量: %d\n", len([]rune(text)))

	// 字符串操作
	words := strings.Fields("Go is awesome and powerful")
	fmt.Printf("分词结果: %v\n", words)

	joined := strings.Join(words, "-")
	fmt.Printf("连接结果: %s\n\n", joined)

	// 8. 类型转换
	fmt.Println("8. 类型转换:")

	// 字符串转数字
	numStr := "123"
	if num, err := strconv.Atoi(numStr); err == nil {
		fmt.Printf("字符串 '%s' 转换为数字: %d\n", numStr, num)
	}

	// 数字转字符串
	number := 456
	str := strconv.Itoa(number)
	fmt.Printf("数字 %d 转换为字符串: '%s'\n", number, str)

	// 浮点数转换
	floatStr := "3.14159"
	if f, err := strconv.ParseFloat(floatStr, 64); err == nil {
		fmt.Printf("字符串 '%s' 转换为浮点数: %.3f\n", floatStr, f)
	}

	fmt.Println("\n=== 示例结束 ===")
}

// init 函数，在 main 之前执行
func init() {
	fmt.Println("程序初始化...")
}
