# Go 语言学习指南

## 🎯 学习路径

### Week 1: 基础入门 (Day 1-7)

**Day 1-2: 环境搭建和基础语法**

- [ ] 安装 Go 环境
- [ ] 学习基本语法（变量、常量、函数）
- [ ] 练习：编写 Hello World 程序
- [ ] 阅读：`notes/01-basic-syntax.md`
- [ ] 运行：`go run examples/basic/main.go`

**Day 3-4: 数据类型和结构**

- [ ] 掌握基本数据类型
- [ ] 学习复合类型（数组、切片、映射、结构体）
- [ ] 练习：创建学生管理结构体
- [ ] 阅读：`notes/02-data-types.md`

**Day 5-6: 控制结构**

- [ ] 掌握条件语句和循环
- [ ] 学习 defer、panic、recover
- [ ] 练习：实现斐波那契数列
- [ ] 阅读：`notes/03-control-structures.md`
- [ ] 运行：`go run exercises/week1/fibonacci.go`

**Day 7: 综合练习**

- [ ] 完成 Week 1 所有练习
- [ ] 编写一个简单的命令行工具
- [ ] 复习和总结

### Week 2: 高级特性 (Day 8-14)

**Day 8-9: 内存模型和 GC**

- [ ] 理解 Go 内存模型
- [ ] 学习垃圾回收机制
- [ ] 练习：内存优化技巧
- [ ] 阅读：`notes/04-memory-gc.md`

**Day 10-11: 接口设计**

- [ ] 掌握接口定义和实现
- [ ] 学习接口设计模式
- [ ] 练习：设计可扩展的系统
- [ ] 阅读：`notes/05-interfaces.md`

**Day 12-13: 错误处理**

- [ ] 掌握 Go 风格的错误处理
- [ ] 学习错误包装和自定义错误
- [ ] 练习：健壮的错误处理
- [ ] 阅读：`notes/06-error-handling.md`

**Day 14: 项目实战**

- [ ] 完成学生管理系统项目
- [ ] 运行：`go run examples/projects/student-manager/main.go`
- [ ] 扩展功能：添加排序、统计等

## 🛠️ 实践项目

### 基础项目

1. **回文检查器** - 练习字符串操作和算法思维

   - 文件：`exercises/week1/practice_problems.go`
   - 要求：忽略大小写和特殊字符

2. **银行账户系统** - 练习结构体和方法

   - 文件：`exercises/week1/practice_problems.go`
   - 功能：存款、取款、余额查询、错误处理

3. **数据结构实现** - 练习基础数据结构
   - 文件：`exercises/week1/practice_problems.go`
   - 实现：栈、队列的基本操作

### 进阶项目

1. **内存优化练习** - 学习 Go 内存管理

   - 文件：`exercises/week2/memory_optimization.go`
   - 内容：字符串构建、切片优化、对象池、逃逸分析

2. **接口设计模式** - 掌握 Go 接口最佳实践

   - 文件：`exercises/week2/interface_patterns.go`
   - 模式：策略、适配器、装饰器、中间件

3. **错误处理系统** - 学习 Go 风格错误处理
   - 文件：`exercises/week2/error_handling.go`
   - 技术：自定义错误、错误包装、重试机制、错误分类

### 高级项目

1. **Web 日志分析器** - 综合练习项目

   - 文件：`exercises/week2/mini_project.go`
   - 技术栈：接口设计、内存优化、错误处理、性能监控
   - 功能：多格式解析、流式处理、统计分析

2. **学生管理系统** - 完整应用示例
   - 文件：`examples/projects/student-manager/main.go`
   - 功能：CRUD 操作、JSON 持久化、CLI 界面

## 📚 推荐资源

### 官方资源

- [Go 官方文档](https://golang.org/doc/)
- [Go 语言规范](https://golang.org/ref/spec)
- [Effective Go](https://golang.org/doc/effective_go.html)

### 书籍推荐

- 《Go 语言学习笔记》- 雨痕（你正在学习的）
- 《Go 语言实战》
- 《Go 并发编程实战》
- 《Go 语言高级编程》

### 在线资源

- [Go by Example](https://gobyexample.com/)
- [A Tour of Go](https://tour.golang.org/)
- [Go Playground](https://play.golang.org/)

## 🧪 测试和验证

### 运行示例代码

```bash
# 进入项目目录
cd go-learning

# 运行基础示例
go run examples/basic/main.go

# 运行斐波那契练习
go run exercises/week1/fibonacci.go

# 运行学生管理系统
go run examples/projects/student-manager/main.go

# 运行测试（如果有）
go test ./...

# 查看测试覆盖率
go test -cover ./...
```

### 性能测试

```bash
# 运行基准测试
go test -bench=. -benchmem

# 生成性能分析文件
go test -bench=. -cpuprofile=cpu.prof -memprofile=mem.prof

# 分析性能
go tool pprof cpu.prof
```

## 🎯 学习检查清单

### 基础知识

- [ ] 能够独立编写 Go 程序
- [ ] 理解 Go 的类型系统
- [ ] 掌握控制流和函数
- [ ] 能够使用标准库

### 进阶技能

- [ ] 理解内存管理和 GC
- [ ] 能够设计和使用接口
- [ ] 掌握错误处理最佳实践
- [ ] 能够编写可测试的代码

### 实战能力

- [ ] 能够构建完整的应用
- [ ] 理解性能优化技巧
- [ ] 掌握调试和分析工具
- [ ] 能够阅读和维护他人代码

## 🚀 下一步学习

完成基础学习后，建议继续深入：

1. **并发编程**

   - Goroutines 和 Channels
   - 并发模式和最佳实践
   - 锁和同步原语

2. **网络编程**

   - HTTP 服务开发
   - TCP/UDP 编程
   - WebSocket 和 gRPC

3. **数据库集成**

   - SQL 数据库操作
   - NoSQL 数据库
   - ORM 框架使用

4. **微服务架构**

   - 服务发现和注册
   - 负载均衡
   - 监控和日志

5. **DevOps 实践**
   - Docker 容器化
   - Kubernetes 部署
   - CI/CD 流水线

## 💡 学习建议

1. **理论与实践结合**：每学完一个概念，立即编写代码验证
2. **阅读优秀代码**：研究开源项目的实现
3. **参与社区**：加入 Go 语言社区，参与讨论
4. **定期复习**：巩固已学知识，查漏补缺
5. **实际项目**：将所学应用到真实项目中

## ⚡ 快速开始

```bash
# 克隆或下载学习材料
git clone <your-repo>
cd go-learning

# 确保 Go 环境正确
go version

# 开始第一个示例
go run examples/basic/main.go

# 查看学习笔记
cat notes/01-basic-syntax.md
```

祝你学习顺利！记住，编程是一门实践性很强的技能，多写代码是提高的最佳途径。 🎉
