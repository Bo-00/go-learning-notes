package main

import (
	"fmt"
)

// 练习1：自定义错误类型
// TODO: 实现一个银行账户系统的错误类型

// 基础错误类型
type BankError struct {
	AccountID string
	Operation string
	Reason    string
}

// TODO: 实现Error方法
func (e *BankError) Error() string {
	// TODO: 返回格式化的错误信息
	return ""
}

// TODO: 定义具体的错误类型
type InsufficientFundsError struct {
	// TODO: 定义字段
}

// TODO: 实现Error方法
func (e *InsufficientFundsError) Error() string {
	// TODO: 实现
	return ""
}

type InvalidAmountError struct {
	// TODO: 定义字段
}

// TODO: 实现Error方法
func (e *InvalidAmountError) Error() string {
	// TODO: 实现
	return ""
}

type AccountNotFoundError struct {
	// TODO: 定义字段
}

// TODO: 实现Error方法
func (e *AccountNotFoundError) Error() string {
	// TODO: 实现
	return ""
}

// 练习2：错误包装和解包
// TODO: 实现一个用户服务，演示错误包装

type UserService struct {
	users map[string]User
}

type User struct {
	ID      string
	Name    string
	Email   string
	Balance float64
}

// TODO: 实现查找用户的方法，返回适当的错误
func (us *UserService) FindUser(id string) (*User, error) {
	// TODO: 如果用户不存在，返回包装的错误
	return nil, nil
}

// TODO: 实现用户转账方法，演示多层错误包装
func (us *UserService) Transfer(fromID, toID string, amount float64) error {
	// TODO: 实现转账逻辑
	// 1. 验证金额
	// 2. 查找发送方用户
	// 3. 查找接收方用户
	// 4. 检查余额
	// 5. 执行转账
	// 在每一步都要正确处理和包装错误
	return nil
}

// 练习3：错误聚合
// TODO: 实现一个配置验证器，收集所有验证错误

type ConfigValidationError struct {
	errors []error
}

// TODO: 实现Error方法
func (e *ConfigValidationError) Error() string {
	// TODO: 返回所有错误的汇总信息
	return ""
}

// TODO: 实现添加错误的方法
func (e *ConfigValidationError) Add(err error) {
	// TODO: 添加错误到集合
}

// TODO: 实现检查是否有错误的方法
func (e *ConfigValidationError) HasErrors() bool {
	// TODO: 返回是否存在错误
	return false
}

// 配置结构
type Config struct {
	DatabaseURL string
	Port        int
	APIKey      string
	MaxUsers    int
}

// TODO: 实现配置验证函数
func validateConfig(config Config) error {
	// TODO: 验证所有配置项，收集所有错误
	// 验证规则：
	// - DatabaseURL 不能为空
	// - Port 必须在 1024-65535 之间
	// - APIKey 长度至少 32 字符
	// - MaxUsers 必须大于 0
	return nil
}

// 练习4：重试机制
// TODO: 实现一个带重试的操作执行器

type RetryableError struct {
	Err       error
	Retryable bool
}

// TODO: 实现Error方法
func (e *RetryableError) Error() string {
	// TODO: 实现
	return ""
}

// TODO: 实现判断是否可重试的方法
func (e *RetryableError) IsRetryable() bool {
	// TODO: 实现
	return false
}

type RetryConfig struct {
	MaxAttempts int
	BackoffFunc func(attempt int) int // 返回等待时间（毫秒）
}

// TODO: 实现重试执行器
func ExecuteWithRetry(operation func() error, config RetryConfig) error {
	// TODO: 实现重试逻辑
	// 1. 执行操作
	// 2. 如果成功，返回nil
	// 3. 如果失败且可重试，等待后重试
	// 4. 如果失败且不可重试，或达到最大重试次数，返回错误
	return nil
}

// 练习5：错误恢复
// TODO: 实现一个安全的数学计算器

type Calculator struct{}

// TODO: 实现安全除法，使用recover处理panic
func (c *Calculator) SafeDivide(a, b float64) (result float64, err error) {
	// TODO: 使用defer和recover捕获除零错误
	// 注意：Go中浮点除零不会panic，这里为了练习假设会panic
	defer func() {
		// TODO: 实现recover逻辑
	}()

	// TODO: 执行除法计算
	return 0, nil
}

// TODO: 实现安全的字符串转整数
func (c *Calculator) SafeStringToInt(s string) (result int, err error) {
	// TODO: 使用defer和recover捕获可能的panic
	defer func() {
		// TODO: 实现recover逻辑
	}()

	// TODO: 转换字符串为整数
	return 0, nil
}

// 练习6：上下文错误
// TODO: 实现带上下文信息的错误处理

type ContextualError struct {
	RequestID string
	UserID    string
	Operation string
	Err       error
}

// TODO: 实现Error方法
func (e *ContextualError) Error() string {
	// TODO: 返回包含上下文信息的错误
	return ""
}

// TODO: 实现Unwrap方法
func (e *ContextualError) Unwrap() error {
	// TODO: 返回原始错误
	return nil
}

// 请求上下文
type RequestContext struct {
	RequestID string
	UserID    string
}

// TODO: 实现一个业务操作，添加上下文错误信息
func ProcessBusinessLogic(ctx RequestContext, data string) error {
	// TODO: 模拟业务逻辑处理
	// 如果出错，包装为上下文错误

	// 模拟可能的错误
	if data == "invalid" {
		// TODO: 返回包装的上下文错误
	}

	return nil
}

// 练习7：错误分类和处理
// TODO: 实现不同类型错误的分类处理

// 错误类型枚举
type ErrorType int

const (
	ErrorTypeValidation ErrorType = iota
	ErrorTypeNetwork
	ErrorTypeDatabase
	ErrorTypePermission
	ErrorTypeUnknown
)

type ClassifiedError struct {
	Type    ErrorType
	Message string
	Cause   error
}

// TODO: 实现Error方法
func (e *ClassifiedError) Error() string {
	// TODO: 实现
	return ""
}

// TODO: 实现错误分类器
func ClassifyError(err error) *ClassifiedError {
	// TODO: 根据错误内容或类型进行分类
	// 可以使用errors.Is和errors.As进行判断
	return nil
}

// TODO: 实现错误处理器
func HandleError(err error) {
	// TODO: 根据错误类型执行不同的处理策略
	classified := ClassifyError(err)

	switch classified.Type {
	case ErrorTypeValidation:
		// TODO: 处理验证错误
	case ErrorTypeNetwork:
		// TODO: 处理网络错误
	case ErrorTypeDatabase:
		// TODO: 处理数据库错误
	case ErrorTypePermission:
		// TODO: 处理权限错误
	default:
		// TODO: 处理未知错误
	}
}

// 练习8：错误测试
// TODO: 编写错误处理的测试用例

// 测试辅助函数
func expectError(t interface{}, err error, expectedErrType interface{}) {
	// TODO: 实现错误测试的辅助函数
	// 1. 检查错误是否为nil（应该不为nil）
	// 2. 检查错误类型是否匹配
	// 3. 如果不匹配，报告测试失败
}

func expectNoError(t interface{}, err error) {
	// TODO: 实现无错误测试的辅助函数
	// 检查错误是否为nil（应该为nil）
}

// 模拟测试函数
func TestBankOperations() {
	fmt.Println("=== 银行操作错误测试 ===")

	// TODO: 测试各种错误情况
	// 例如：
	// 1. 测试余额不足错误
	// 2. 测试无效金额错误
	// 3. 测试账户不存在错误
	// 4. 测试正常操作（无错误）

	fmt.Println("请实现TODO函数后进行测试")
}

// 综合测试函数
func runErrorTests() {
	fmt.Println("=== 错误处理练习测试 ===")

	// TODO: 取消注释进行测试

	// 测试配置验证
	// invalidConfig := Config{
	//     DatabaseURL: "",
	//     Port: 80,
	//     APIKey: "short",
	//     MaxUsers: -1,
	// }
	// err := validateConfig(invalidConfig)
	// if err != nil {
	//     fmt.Printf("配置验证错误: %v\n", err)
	// }

	// 测试重试机制
	// retryConfig := RetryConfig{
	//     MaxAttempts: 3,
	//     BackoffFunc: func(attempt int) int { return attempt * 1000 },
	// }
	// err = ExecuteWithRetry(func() error {
	//     return &RetryableError{Err: errors.New("temporary error"), Retryable: true}
	// }, retryConfig)
	// fmt.Printf("重试结果: %v\n", err)

	// 测试安全计算
	// calc := &Calculator{}
	// result, err := calc.SafeDivide(10, 0)
	// fmt.Printf("安全除法结果: %.2f, 错误: %v\n", result, err)

	TestBankOperations()

	fmt.Println("请实现TODO函数后进行完整测试")
}

// 使函数可被调用
var _ = runErrorTests
