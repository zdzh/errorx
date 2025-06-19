# errorx

errorx 是一个用于处理错误码的 Go 语言库，提供了统一的错误码管理和错误处理机制。

## 功能特性

- 支持自定义错误码
- 提供错误包装和解包功能
- 支持错误链的追踪和解析
- 兼容标准库的 errors 包

## 安装

使用 Go 模块安装：

```bash
go get github.com/zdzh/errorx
```

## 快速开始

```go
package main

import (
	"fmt"
	"github.com/zdzh/errorx/errcode"
)

func main() {
	err := errcode.Wrap(errors.New("test error"), 500, "wrapped error")
	fmt.Println(err)
}
```

## API 文档

### 错误码定义

```go
type ErrorCode = int
```

### 错误结构体

```go
type Error struct {
	code ErrorCode
	error
}
```

### 主要方法
- `func WithCode(err error, code ErrorCode) error`
  - 为错误添加错误码,不修改原始错误信息
  - 参数:
    - err: 原始错误
    - code: 要添加的错误码
  - 返回: 包含错误码的新错误

- `func WithStack(err error, code ErrorCode) error` 
  - 为错误添加调用栈信息和错误码
  - 参数:
    - err: 原始错误
    - code: 要添加的错误码
  - 返回: 包含调用栈和错误码的新错误

- `func Wrap(err error, code ErrorCode, message string) error`
  - 包装错误,添加新的错误信息、错误码和调用栈
  - 参数:
    - err: 原始错误
    - code: 要添加的错误码
    - message: 新的错误信息
  - 返回: 包装后的新错误

- `func Wrapf(err error, code ErrorCode, format string, args ...interface{}) error`
  - 格式化包装错误,类似 Wrap 但支持格式化错误信息
  - 参数:
    - err: 原始错误
    - code: 要添加的错误码
    - format: 错误信息格式
    - args: 格式化参数
  - 返回: 包装后的新错误

- `func WithMessage(err error, code ErrorCode, message string) error`
  - 为错误添加新的错误信息和错误码,不添加调用栈
  - 参数: 
    - err: 原始错误
    - code: 要添加的错误码
    - message: 新的错误信息
  - 返回: 包含新信息的错误

- `func WithMessagef(err error, code ErrorCode, format string, args ...interface{}) error`
  - 为错误添加格式化的新错误信息和错误码,不添加调用栈
  - 参数:
    - err: 原始错误
    - code: 要添加的错误码
    - format: 错误信息格式
    - args: 格式化参数
  - 返回: 包含新信息的错误

- `func Cause(err error) error`
  - 获取错误链中的根本原因错误
  - 参数:
    - err: 要解析的错误
  - 返回: 最原始的错误

- `func SetDefaultCode(code ErrorCode)`
  - 设置默认的错误码
  - 参数:
    - code: 要设置的默认错误码
  - 说明: 当创建新的错误时，如果没有指定错误码，将使用此默认错误码

- `func Code(err error) ErrorCode`
  - 获取错误中的错误码
  - 参数:
    - err: 要获取错误码的错误
  - 返回: 
    - ErrorCode: 如果错误包含错误码则返回对应错误码，否则返回默认错误码

## 贡献指南

欢迎提交 Issue 和 PR，请遵循以下步骤：

1. Fork 项目
2. 创建分支 (`git checkout -b feature/YourFeature`)
3. 提交更改 (`git commit -am 'Add some feature'`)
4. 推送分支 (`git push origin feature/YourFeature`)
5. 创建 Pull Request

## 许可证

MIT 许可证，详情请见 LICENSE 文件。

这个README文件包含了项目的基本信息、安装方法、使用示例、API文档、贡献指南和许可证信息，可以帮助用户快速了解和使用ErrorX库。

