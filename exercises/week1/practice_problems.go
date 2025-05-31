package main

import (
	"errors"
	"fmt"
)

// 练习1：字符串操作
// TODO: 实现一个函数，检查字符串是否为回文
func isPalindrome(s string) bool {
	// TODO: 实现回文检查逻辑
	// 提示：忽略大小写和非字母字符
	return false
}

// 练习2：数组和切片
// TODO: 实现一个函数，找出数组中第二大的数
func findSecondLargest(nums []int) (int, error) {
	// TODO: 实现逻辑
	// 注意处理边界情况：空数组、只有一个元素、所有元素相同
	var max int
	var secondMax int
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			secondMax = max
			max = nums[i]
		} else if nums[i] > secondMax {
			secondMax = nums[i]
		}
	}
	return secondMax, nil
}

// 练习3：映射操作
// TODO: 实现一个函数，统计字符串中每个字符的出现次数
func countCharacters(s string) map[rune]int {
	// TODO: 实现字符计数逻辑
	return nil
}

// 练习4：结构体和方法
type BankAccount struct {
	// TODO: 定义银行账户的字段
	// 提示：账户号、余额、持有人姓名
}

// TODO: 实现存款方法
func (ba *BankAccount) Deposit(amount float64) error {
	// TODO: 实现存款逻辑，注意验证金额
	return errors.New("not implemented")
}

// TODO: 实现取款方法
func (ba *BankAccount) Withdraw(amount float64) error {
	// TODO: 实现取款逻辑，注意验证余额
	return errors.New("not implemented")
}

// TODO: 实现获取余额方法
func (ba *BankAccount) GetBalance() float64 {
	// TODO: 返回当前余额
	return 0
}

// 练习5：函数式编程
// TODO: 实现一个通用的过滤函数
func filter(slice []int, predicate func(int) bool) []int {
	// TODO: 实现过滤逻辑
	// 根据传入的判断函数过滤切片元素
	return nil
}

// TODO: 实现一个通用的映射函数
func mapInt(slice []int, transform func(int) int) []int {
	// TODO: 实现映射逻辑
	// 对切片中的每个元素应用变换函数
	return nil
}

// TODO: 实现一个通用的归约函数
func reduce(slice []int, initial int, accumulator func(int, int) int) int {
	// TODO: 实现归约逻辑
	// 将切片元素累积为单个值
	return 0
}

// 练习6：递归练习
// TODO: 实现阶乘函数（递归版本）
func factorial(n int) int {
	// TODO: 实现递归阶乘
	return 0
}

// TODO: 实现汉诺塔问题
func hanoi(n int, from, to, aux string) {
	// TODO: 实现汉诺塔递归解法
	// 打印移动步骤
}

// 练习7：算法实现
// TODO: 实现冒泡排序
func bubbleSort(arr []int) {
	// TODO: 实现冒泡排序算法
}

// TODO: 实现二分查找
func binarySearch(arr []int, target int) int {
	// TODO: 实现二分查找算法
	// 返回目标元素的索引，未找到返回-1
	return -1
}

// 练习8：数据结构
// TODO: 实现一个简单的栈
type Stack struct {
	// TODO: 定义栈的内部结构
}

// TODO: 实现栈的基本操作
func (s *Stack) Push(value int) {
	// TODO: 入栈操作
}

func (s *Stack) Pop() (int, error) {
	// TODO: 出栈操作
	return 0, errors.New("not implemented")
}

func (s *Stack) Peek() (int, error) {
	// TODO: 查看栈顶元素
	return 0, errors.New("not implemented")
}

func (s *Stack) IsEmpty() bool {
	// TODO: 检查栈是否为空
	return true
}

// 测试和验证的示例函数
func runTests() {
	fmt.Println("=== Week 1 练习题测试 ===")
	fmt.Println("请实现以上函数，然后取消注释进行测试")

	// TODO: 取消注释并实现函数后进行测试
	// fmt.Println("回文测试:", isPalindrome("A man a plan a canal Panama"))
	// nums := []int{1, 3, 4, 5, 2}
	// second, err := findSecondLargest(nums)
	// if err == nil {
	//     fmt.Println("第二大数:", second)
	// }
}
