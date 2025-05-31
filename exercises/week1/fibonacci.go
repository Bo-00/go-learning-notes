package main

import "fmt"

// 练习1：斐波那契数列
// 编写多种实现方式，体验不同的算法思路

// 递归实现（效率较低）
func fibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

// 迭代实现（推荐）
func fibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// 使用切片缓存（动态规划）
func fibonacciDP(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// 生成器模式
func fibonacciGenerator() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

// 通道实现（并发）
func fibonacciChannel(n int, ch chan<- int) {
	defer close(ch)

	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
}

func main() {
	fmt.Println("=== 斐波那契数列练习 ===\n")

	n := 10

	// 1. 递归实现
	fmt.Printf("递归实现 fibonacci(%d) = %d\n", n, fibonacciRecursive(n))

	// 2. 迭代实现
	fmt.Printf("迭代实现 fibonacci(%d) = %d\n", n, fibonacciIterative(n))

	// 3. 动态规划
	fmt.Printf("动态规划 fibonacci(%d) = %d\n", n, fibonacciDP(n))

	// 4. 生成器模式
	fmt.Print("生成器模式前10项: ")
	gen := fibonacciGenerator()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", gen())
	}
	fmt.Println()

	// 5. 通道实现
	fmt.Print("通道实现前10项: ")
	ch := make(chan int)
	go fibonacciChannel(10, ch)

	for num := range ch {
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	fmt.Println("\n=== 练习完成 ===")
}

// 基准测试函数（需要在 *_test.go 文件中）
/*
import "testing"

func BenchmarkFibonacciRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciRecursive(20)
	}
}

func BenchmarkFibonacciIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciIterative(20)
	}
}

func BenchmarkFibonacciDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciDP(20)
	}
}
*/
