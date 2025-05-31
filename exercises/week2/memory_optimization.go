package main

import (
	"fmt"
	"runtime"
	"time"
)

// 练习1：内存分配优化
// TODO: 优化这个函数，减少内存分配次数
func inefficientStringBuilder(words []string) string {
	// 当前实现：每次 += 都会创建新字符串
	result := ""
	for _, word := range words {
		result += word + " "
	}
	return result
}

// TODO: 实现高效版本
func efficientStringBuilder(words []string) string {
	// TODO: 使用 strings.Builder 或预分配容量的方式优化
	return ""
}

// 练习2：切片容量管理
// TODO: 优化这个函数，避免多次内存重新分配
func inefficientSliceGrowth(n int) []int {
	var result []int
	for i := 0; i < n; i++ {
		result = append(result, i*i)
	}
	return result
}

// TODO: 实现优化版本
func efficientSliceGrowth(n int) []int {
	// TODO: 预分配适当的容量
	return nil
}

// 练习3：内存池模式
// TODO: 实现一个对象池，复用 []byte 切片
type ByteBufferPool struct {
	// TODO: 定义池的结构
}

// TODO: 实现获取缓冲区的方法
func (p *ByteBufferPool) Get() []byte {
	// TODO: 从池中获取或创建新的缓冲区
	return nil
}

// TODO: 实现归还缓冲区的方法
func (p *ByteBufferPool) Put(buf []byte) {
	// TODO: 将缓冲区归还到池中
	// 注意：需要重置缓冲区状态
}

// 练习4：减少逃逸分析
// TODO: 修改这个函数，让变量在栈上分配而不是堆上
func createEscapingPointer() *int {
	x := 42
	return &x // 这会导致 x 逃逸到堆
}

// TODO: 实现一个不会导致逃逸的版本
func createNonEscapingValue() int {
	// TODO: 返回值而不是指针
	return 0
}

// 练习5：内存使用监控
// TODO: 实现一个函数，监控和报告内存使用情况
func monitorMemoryUsage(fn func()) {
	// TODO: 在函数执行前后获取内存统计信息
	// 使用 runtime.MemStats
	// 计算并报告内存使用差异
}

// 练习6：大对象处理
// TODO: 实现一个处理大数据的函数，最小化内存占用
func processLargeData(data [][]byte) map[string]int {
	// TODO: 统计所有数据中每个字符串的出现次数
	// 注意：数据可能很大，需要考虑内存效率
	return nil
}

// 练习7：内存泄漏检测
// TODO: 找出并修复这个函数中的内存泄漏
func leakyFunction() {
	// 创建一个会持续增长的映射
	cache := make(map[string][]byte)

	// 模拟持续添加数据但从不清理
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("key_%d", i)
		cache[key] = make([]byte, 1024*1024) // 1MB per entry
	}

	// TODO: 这个函数有内存泄漏问题，请修复
	// 提示：考虑数据的生命周期和清理策略
}

// TODO: 实现修复版本
func fixedFunction() {
	// TODO: 实现一个不会泄漏内存的版本
}

// 基准测试框架
func benchmarkFunction(name string, fn func()) {
	var m1, m2 runtime.MemStats

	runtime.GC()
	runtime.ReadMemStats(&m1)

	start := time.Now()
	fn()
	duration := time.Since(start)

	runtime.GC()
	runtime.ReadMemStats(&m2)

	fmt.Printf("%s:\n", name)
	fmt.Printf("  时间: %v\n", duration)
	fmt.Printf("  内存分配: %d bytes\n", m2.TotalAlloc-m1.TotalAlloc)
	fmt.Printf("  分配次数: %d\n", m2.Mallocs-m1.Mallocs)
	fmt.Println()
}

// 测试函数
func runMemoryTests() {
	fmt.Println("=== 内存优化练习测试 ===")

	// words := []string{"hello", "world", "go", "programming", "memory", "optimization"}

	// TODO: 取消注释进行性能对比测试
	// words := []string{"hello", "world", "go", "programming", "memory", "optimization"}
	// benchmarkFunction("低效字符串构建", func() {
	//     inefficientStringBuilder(words)
	// })

	// benchmarkFunction("高效字符串构建", func() {
	//     efficientStringBuilder(words)
	// })

	// benchmarkFunction("低效切片增长", func() {
	//     inefficientSliceGrowth(10000)
	// })

	// benchmarkFunction("高效切片增长", func() {
	//     efficientSliceGrowth(10000)
	// })

	fmt.Println("请实现TODO函数后进行测试")
}
