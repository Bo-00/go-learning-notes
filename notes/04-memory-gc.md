# Go 内存模型和垃圾回收机制

## 1. Go 内存模型概述

### 内存分配区域

Go 运行时将内存分为几个主要区域：

```
┌─────────────┐
│    代码区    │  <- 程序指令
├─────────────┤
│   全局区    │  <- 全局变量、常量
├─────────────┤
│    栈区     │  <- 函数调用、局部变量
├─────────────┤
│    堆区     │  <- 动态分配的内存
└─────────────┘
```

### 栈 vs 堆

```go
func stackVsHeap() {
    // 栈分配：编译器能确定生命周期
    var stackVar = 42           // 通常在栈上
    var array [100]int         // 小数组，在栈上

    // 堆分配：需要动态分配或逃逸分析决定
    var slice = make([]int, 100)    // 切片在堆上
    var ptr = new(int)              // 指针指向堆内存

    // 逃逸到堆的情况
    var local = getPointer()        // 返回指针，逃逸到堆
}

func getPointer() *int {
    x := 42
    return &x  // x 逃逸到堆，因为返回了它的地址
}
```

## 2. 逃逸分析 (Escape Analysis)

### 什么是逃逸分析

逃逸分析是编译器的一种优化技术，决定变量应该分配在栈上还是堆上。

```go
// 使用 go build -gcflags="-m" 查看逃逸分析结果

// 不逃逸 - 分配在栈上
func noEscape() {
    x := 42
    fmt.Println(x)  // x 不会逃逸
}

// 逃逸情况1：返回局部变量的地址
func escape1() *int {
    x := 42
    return &x  // x 逃逸到堆
}

// 逃逸情况2：发送到通道
func escape2() {
    ch := make(chan *int, 1)
    x := 42
    ch <- &x  // x 逃逸到堆
}

// 逃逸情况3：赋值给接口
func escape3() interface{} {
    x := 42
    return x  // x 逃逸到堆（接口存储）
}

// 逃逸情况4：切片扩容
func escape4() {
    slice := make([]int, 0, 4)
    for i := 0; i < 10; i++ {
        slice = append(slice, i)  // 扩容时可能逃逸
    }
}

// 逃逸情况5：大对象
func escape5() {
    var big [10000]int  // 大数组逃逸到堆
    _ = big
}
```

### 逃逸分析实践

```bash
# 查看逃逸分析结果
go build -gcflags="-m" main.go

# 查看更详细的分析
go build -gcflags="-m -m" main.go

# 禁用优化查看
go build -gcflags="-N -l -m" main.go
```

## 3. 内存分配器

### Go 内存分配器设计

Go 使用 **TCMalloc** 的思想，主要包含：

1. **Size Classes**: 预定义的对象大小类别
2. **Spans**: 连续的页集合
3. **Cache**: 每个线程的本地缓存

```go
// 不同大小对象的分配策略

// 微对象 (< 16 bytes)
func microAlloc() {
    var x int8      // 微对象分配器
    var y bool
    var z byte
    _, _, _ = x, y, z
}

// 小对象 (16 bytes - 32KB)
func smallAlloc() {
    slice := make([]int, 100)  // 小对象分配器
    _ = slice
}

// 大对象 (> 32KB)
func largeAlloc() {
    slice := make([]int, 10000) // 直接从堆分配
    _ = slice
}
```

### 内存池和对象复用

```go
import "sync"

// 使用 sync.Pool 复用对象
var bufferPool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 1024)
    },
}

func usePool() {
    // 从池中获取对象
    buffer := bufferPool.Get().([]byte)
    defer bufferPool.Put(buffer)  // 使用完毕归还

    // 使用 buffer...
}

// 自定义对象池
type Object struct {
    data []byte
}

var objectPool = sync.Pool{
    New: func() interface{} {
        return &Object{
            data: make([]byte, 1024),
        }
    },
}

func useObjectPool() {
    obj := objectPool.Get().(*Object)
    defer func() {
        // 重置对象状态
        obj.data = obj.data[:0]
        objectPool.Put(obj)
    }()

    // 使用 obj...
}
```

## 4. 垃圾回收器 (GC)

### GC 演进历史

1. **Go 1.0-1.4**: Stop-the-World 标记清除
2. **Go 1.5**: 三色并发标记
3. **Go 1.8+**: 混合写屏障

### 三色标记算法

```
白色：未访问的对象（待回收）
灰色：已访问但子对象未完全扫描
黑色：已访问且子对象已完全扫描
```

```go
// GC 的工作过程示例
func gcExample() {
    // 1. 标记阶段：从根对象开始标记
    root := &Node{value: 1}
    root.child = &Node{value: 2}
    root.child.child = &Node{value: 3}

    // 2. 清除阶段：回收未标记的白色对象
    root.child = nil  // 断开引用，子树变为可回收

    // 3. 触发 GC
    runtime.GC()
}

type Node struct {
    value int
    child *Node
}
```

### GC 调优参数

```go
import "runtime"

func gcTuning() {
    // 设置 GC 目标百分比（默认 100）
    debug.SetGCPercent(50)  // 当堆增长 50% 时触发 GC

    // 获取 GC 统计信息
    var m runtime.MemStats
    runtime.ReadMemStats(&m)

    fmt.Printf("分配的堆内存: %d KB\n", m.Alloc/1024)
    fmt.Printf("GC 次数: %d\n", m.NumGC)
    fmt.Printf("GC 总暂停时间: %v\n", time.Duration(m.PauseTotalNs))

    // 手动触发 GC
    runtime.GC()

    // 设置内存限制（Go 1.19+）
    debug.SetMemoryLimit(100 << 20)  // 100MB
}
```

### GC 触发条件

```go
// GC 触发的情况：
// 1. 堆内存达到上次 GC 后的目标大小
// 2. 超过 2 分钟未执行 GC
// 3. 手动调用 runtime.GC()

func monitorGC() {
    // 监控 GC 事件
    go func() {
        var lastGC uint32
        for {
            var m runtime.MemStats
            runtime.ReadMemStats(&m)

            if m.NumGC != lastGC {
                fmt.Printf("GC #%d: %v\n", m.NumGC,
                    time.Duration(m.PauseNs[(m.NumGC+255)%256]))
                lastGC = m.NumGC
            }

            time.Sleep(100 * time.Millisecond)
        }
    }()
}
```

## 5. 内存优化实践

### 减少内存分配

```go
// ❌ 频繁分配
func badConcat(strs []string) string {
    result := ""
    for _, s := range strs {
        result += s  // 每次都创建新字符串
    }
    return result
}

// ✅ 使用 strings.Builder
func goodConcat(strs []string) string {
    var builder strings.Builder
    builder.Grow(estimateSize(strs))  // 预分配容量
    for _, s := range strs {
        builder.WriteString(s)
    }
    return builder.String()
}

// ✅ 预分配切片容量
func goodSlice() []int {
    slice := make([]int, 0, 100)  // 预分配容量
    for i := 0; i < 100; i++ {
        slice = append(slice, i)
    }
    return slice
}
```

### 避免内存泄漏

```go
// ❌ 切片内存泄漏
func badSlice(data []byte) []byte {
    return data[10:20]  // 保留了整个底层数组的引用
}

// ✅ 复制需要的部分
func goodSlice(data []byte) []byte {
    result := make([]byte, 10)
    copy(result, data[10:20])
    return result
}

// ❌ Goroutine 泄漏
func badGoroutine() {
    ch := make(chan int)
    go func() {
        for {
            select {
            case <-ch:
                // 处理数据
            // 缺少退出条件，造成 Goroutine 泄漏
            }
        }
    }()
}

// ✅ 正确的 Goroutine 管理
func goodGoroutine(ctx context.Context) {
    ch := make(chan int)
    go func() {
        for {
            select {
            case <-ch:
                // 处理数据
            case <-ctx.Done():
                return  // 正确退出
            }
        }
    }()
}

// ❌ 映射内存泄漏
func badMap() {
    cache := make(map[string]*largeObject)
    // 持续添加，但从不删除过期项
}

// ✅ 定期清理
func goodMap() {
    cache := make(map[string]*cacheItem)
    go func() {
        ticker := time.NewTicker(time.Minute)
        defer ticker.Stop()

        for range ticker.C {
            now := time.Now()
            for key, item := range cache {
                if now.Sub(item.timestamp) > time.Hour {
                    delete(cache, key)
                }
            }
        }
    }()
}
```

### 内存分析工具

```go
import _ "net/http/pprof"

func enablePprof() {
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
}

// 使用方法：
// 1. 启动程序后访问 http://localhost:6060/debug/pprof/
// 2. 使用 go tool pprof 分析：
//    go tool pprof http://localhost:6060/debug/pprof/heap
//    go tool pprof http://localhost:6060/debug/pprof/allocs
```

## 6. 性能测试和基准

```go
// 内存分配基准测试
func BenchmarkStringConcat(b *testing.B) {
    strs := []string{"hello", "world", "go", "programming"}

    b.ResetTimer()
    b.ReportAllocs()  // 报告内存分配

    for i := 0; i < b.N; i++ {
        goodConcat(strs)
    }
}

// 运行基准测试
// go test -bench=. -benchmem
```

## 实践建议

### 1. 编码规范

- 尽量在栈上分配变量
- 避免不必要的指针传递
- 预分配已知大小的切片和映射
- 及时释放不再需要的大对象引用

### 2. 监控和诊断

- 使用 `runtime.MemStats` 监控内存使用
- 利用 pprof 进行内存分析
- 定期检查是否存在内存泄漏
- 在生产环境中设置适当的内存限制

### 3. GC 优化

- 根据应用特点调整 `GOGC` 参数
- 减少 GC 暂停时间的影响
- 避免在性能敏感路径上频繁分配

这些知识点帮助你深入理解 Go 的内存管理机制，编写更高效的 Go 程序。
