# go-study

学习 Go 语言期间写的练习代码集合，按主题分目录保存，每个目录是一个独立的 Go module，可以单独运行。

仓库创建于 2026-03，属于阶段性学习记录，不包含生产代码。保留的目的是回看基础语法、并发、泛型和文件读写这些练习过程。

## 目录索引

| 目录 | 内容 |
| --- | --- |
| `anonymous/` | 匿名函数 |
| `arrays/` | 数组基础；`arrays/lists/` 切片列表、`arrays/maps/` 映射 |
| `bank/` | 银行存取练习：读写余额文件、`fileops/` 封装文件操作、含第三方依赖 |
| `concurrency/` | 并发（goroutine / channel） |
| `generics/` | 泛型 |
| `profit-calculator/` | 按利润率计算收益，结果写入 `results.txt` |
| `rate-calculator/` | 汇率换算练习，拆分为 `cmdmanager/` `conversion/` `filemanager/` `iomanager/` `prices/` 多个包，结果输出为 JSON |
| `recursion/` | 递归 |
| `structs/` | 结构体基础，`structs/user/` 为用户结构练习 |
| `structs-practice/` | 结构体综合练习：`note/` 笔记、`todo/` 待办，数据以 JSON 读写 |
| `types/` | 基本类型 |
| `variadic-function/` | 可变参数函数 |

## 运行方式

每个练习目录都是独立 module，进入对应目录后执行：

```bash
cd arrays
go run .
```

入口文件不统一叫 `main.go`（例如 `structs/structs.go`、`types/types.go`、`profit-calculator/profit-calculator.go`），直接 `go run .` 即可，不需要指定文件名。

部分程序依赖同目录下的数据文件，需要在对应目录内执行，例如 `bank/balance.txt`、`rate-calculator/prices.txt`。

## 已知情况

- 根目录没有 `go.mod`，各练习目录各自成模块，无法在仓库根目录整体构建
- `arrays/`、`arrays/lists/`、`arrays/maps/` 三个目录下的 `go.mod` 内容相同，看起来是复制模板时一并留下的
- `rate-calculator/.DS_Store` 是 macOS 元数据文件，属于误提交，已在本次整理中移除

## 说明

本仓库按主题保留为学习记录，不再继续扩展。
