
# 目录

- [序](#序)
- [并发控制](#并发控制)
- [内存管理：内存分、逃逸分析、垃圾回收](#内存管理：内存分配、逃逸分析、垃圾回收)
- [协程](#协程)
- [go语言中的锁控制](#go语言中的锁控制)
- [defer、chan、map、slice、select](#defer、chan、map、slice、select)
- [Writer、Reader接口](#Writer、Reader接口)
- [内置类型、引用类型、结构类型](#内置类型、引用类型、结构类型)
- [创建、声明、初始化](#创建、声明、初始化)
- [new、make](#new、make)
- [指针](#指针)
- [Byte、bit、Unicode、UTF-8](#Byte、bit、Unicode、UTF-8)
- [锁、sync.Mutex互斥锁、sync.RWMutex读写锁、sync.Once](#锁、sync.Mutex互斥锁、sync.RWMutex读写锁、sync.Once)
- [通过cgo调用c代码](#通过cgo调用c代码)
- [unsife包、context包](#unsife包、context包)
- [分布式系统](#分布式系统)
- [错误异常：fatal\panic\err](#错误异常：fatal\panic\err)

# 序

- go语言语法糖，最常用的语法糖莫过于赋值符 `:=` ，其次，表示函数变参的 `...`
- go语言定时器：一次性定时器，周期性定时器：Timer\Ticker
- 反射：反射是一种检查interface变量的底层类型和值的机制；
- map是一种hash表实现，每次遍历的顺序都可能不一样。

# 并发控制

- Channel: 使用channel控制子协程
- WaitGroup : 使用信号量机制控制子协程
- Context: 使用上下文控制子协程

> 三种方案各有优劣，比如Channel优点是实现简单，清晰易懂，WaitGroup优点是子协程个数动态可调整，Context优点是对子协程派生出来的孙子协程的控制。缺点是相对而言的，要结合实例应用场景进行选择。

# 内存管理：内存分配、逃逸分析、垃圾回收

# 协程

# go语言中的锁控制

# defer、chan、map、slice、select

recover捕获的是祖父级调用时的异常，直接调用时无效，必须在defer函数中直接调用才有效：

```go
func main() {
    defer func() {
        recover()
    }()
    panic(1)
}
```

# Writer、Reader接口

# 内置类型、引用类型、结构类型

# 创建、声明、初始化

# new、make

# 指针

# Byte、bit、Unicode、UTF-8

# 锁、sync.Mutex互斥锁、sync.RWMutex读写锁、sync.Once

# 通过cgo调用c代码

# unsife包、context包

# 分布式系统

- 分布式id生成器
- 分布式锁
- 延时任务系统
- 分布式搜索引擎
- 负载均衡
- 分布式配置管理
- 分布式爬虫


# 错误异常：fatal\panic\err