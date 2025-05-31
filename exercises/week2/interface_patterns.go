package main

import (
	"fmt"
)

// 练习1：设计小而专一的接口
// TODO: 将这个大接口拆分为多个小接口

// ❌ 不好的设计：接口太大
type MediaProcessor interface {
	LoadFile(filename string) error
	SaveFile(filename string) error
	Play() error
	Pause() error
	Stop() error
	GetVolume() float64
	SetVolume(volume float64) error
	GetDuration() int
	Seek(position int) error
}

// TODO: 设计更好的小接口
// 提示：可以按功能分组，如 FileHandler, Player, VolumeController 等

// 练习2：策略模式实现
// TODO: 实现一个可以使用不同排序算法的排序器

// 排序策略接口
type SortStrategy interface {
	// TODO: 定义排序方法
}

// TODO: 实现快速排序策略
type QuickSortStrategy struct{}

// TODO: 实现冒泡排序策略
type BubbleSortStrategy struct{}

// TODO: 实现排序上下文
type Sorter struct {
	// TODO: 定义字段
}

// TODO: 实现设置策略的方法
func (s *Sorter) SetStrategy(strategy SortStrategy) {
	// TODO: 实现
}

// TODO: 实现执行排序的方法
func (s *Sorter) Sort(data []int) []int {
	// TODO: 实现
	return nil
}

// 练习3：适配器模式
// TODO: 实现一个适配器，让旧的接口兼容新的接口

// 新接口
type ModernPrinter interface {
	PrintDocument(content string, format string) error
}

// 旧接口
type LegacyPrinter interface {
	Print(text string) error
}

// 旧的打印机实现
type OldPrinter struct{}

func (p *OldPrinter) Print(text string) error {
	fmt.Printf("Old Printer: %s\n", text)
	return nil
}

// TODO: 实现适配器，让 OldPrinter 实现 ModernPrinter 接口
type PrinterAdapter struct {
	// TODO: 定义字段
}

// TODO: 实现适配器的方法
func (a *PrinterAdapter) PrintDocument(content string, format string) error {
	// TODO: 将新接口调用转换为旧接口调用
	return nil
}

// 练习4：装饰器模式
// TODO: 实现一个HTTP客户端的装饰器，添加日志和重试功能

// 基础HTTP客户端接口
type HTTPClient interface {
	Get(url string) (string, error)
}

// 基础实现
type SimpleHTTPClient struct{}

func (c *SimpleHTTPClient) Get(url string) (string, error) {
	// 模拟HTTP请求
	return fmt.Sprintf("Response from %s", url), nil
}

// TODO: 实现日志装饰器
type LoggingHTTPClient struct {
	// TODO: 定义字段
}

// TODO: 实现带日志的Get方法
func (c *LoggingHTTPClient) Get(url string) (string, error) {
	// TODO: 添加日志记录功能
	return "", nil
}

// TODO: 实现重试装饰器
type RetryHTTPClient struct {
	// TODO: 定义字段
}

// TODO: 实现带重试的Get方法
func (c *RetryHTTPClient) Get(url string) (string, error) {
	// TODO: 添加重试逻辑
	return "", nil
}

// 练习5：接口组合
// TODO: 设计一个文件处理系统，使用接口组合

// 基础接口
type Reader interface {
	Read([]byte) (int, error)
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

// TODO: 组合接口
type ReadWriter interface {
	// TODO: 组合 Reader 和 Writer
}

type ReadWriteCloser interface {
	// TODO: 组合所有三个接口
}

// TODO: 实现一个文件处理器，实现 ReadWriteCloser 接口
type FileHandler struct {
	// TODO: 定义字段
}

// TODO: 实现所有必要的方法

// 练习6：类型断言和类型选择
// TODO: 实现一个通用的数据处理器

type DataProcessor struct{}

// TODO: 实现一个方法，根据不同的数据类型执行不同的处理
func (dp *DataProcessor) Process(data interface{}) string {
	// TODO: 使用类型选择处理不同类型的数据
	// 支持: string, int, []int, map[string]int
	return ""
}

// 练习7：接口作为函数参数
// TODO: 实现一个通用的数据验证框架

// 验证器接口
type Validator interface {
	Validate(data interface{}) error
}

// TODO: 实现邮箱验证器
type EmailValidator struct{}

// TODO: 实现邮箱验证方法
func (ev *EmailValidator) Validate(data interface{}) error {
	// TODO: 验证邮箱格式
	return nil
}

// TODO: 实现年龄验证器
type AgeValidator struct {
	Min int
	Max int
}

// TODO: 实现年龄验证方法
func (av *AgeValidator) Validate(data interface{}) error {
	// TODO: 验证年龄范围
	return nil
}

// TODO: 实现验证管理器
type ValidationManager struct {
	validators []Validator
}

// TODO: 实现添加验证器的方法
func (vm *ValidationManager) AddValidator(validator Validator) {
	// TODO: 实现
}

// TODO: 实现执行所有验证的方法
func (vm *ValidationManager) ValidateAll(data interface{}) error {
	// TODO: 执行所有验证器
	return nil
}

// 练习8：函数接口
// TODO: 实现一个中间件系统

// HTTP处理函数类型
type HandlerFunc func(request string) string

// TODO: 让 HandlerFunc 实现 Handler 接口
type Handler interface {
	Handle(request string) string
}

// TODO: 实现HandlerFunc的Handle方法
func (f HandlerFunc) Handle(request string) string {
	// TODO: 实现
	return ""
}

// TODO: 实现中间件类型
type Middleware func(Handler) Handler

// TODO: 实现日志中间件
func LoggingMiddleware(next Handler) Handler {
	// TODO: 返回一个包装了日志功能的Handler
	return nil
}

// TODO: 实现认证中间件
func AuthMiddleware(next Handler) Handler {
	// TODO: 返回一个包装了认证功能的Handler
	return nil
}

// TODO: 实现中间件链
func ChainMiddleware(handler Handler, middlewares ...Middleware) Handler {
	// TODO: 将多个中间件串联起来
	return nil
}

// 测试函数
func runInterfaceTests() {
	fmt.Println("=== 接口设计模式练习测试 ===")
	fmt.Println("请实现TODO函数后进行测试")

	// TODO: 取消注释进行测试

	// 测试排序策略
	// sorter := &Sorter{}
	// data := []int{64, 34, 25, 12, 22, 11, 90}
	// sorter.SetStrategy(&QuickSortStrategy{})
	// sorted := sorter.Sort(data)
	// fmt.Println("快速排序结果:", sorted)

	// 测试适配器模式
	// oldPrinter := &OldPrinter{}
	// adapter := &PrinterAdapter{legacy: oldPrinter}
	// adapter.PrintDocument("Hello World", "PDF")

	// 测试装饰器模式
	// client := &SimpleHTTPClient{}
	// loggedClient := &LoggingHTTPClient{client: client}
	// retryClient := &RetryHTTPClient{client: loggedClient, maxRetries: 3}
	// response, _ := retryClient.Get("https://example.com")
	// fmt.Println(response)
}

// 使 runInterfaceTests 可以被调用
var _ = runInterfaceTests
