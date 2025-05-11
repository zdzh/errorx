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
- `func WithStack(err error, code ErrorCode) error`
- `func Wrap(err error, code ErrorCode, message string) error`
- `func Wrapf(err error, code ErrorCode, format string, args ...interface{}) error`
- `func WithMessage(err error, code ErrorCode, message string) error`
- `func WithMessagef(err error, code ErrorCode, format string, args ...interface{}) error`
- `func Cause(err error) error`

## 贡献指南

欢迎提交 Issue 和 PR，请遵循以下步骤：

1. Fork 项目
2. 创建分支 (`git checkout -b feature/YourFeature`)
3. 提交更改 (`git commit -am 'Add some feature'`)
4. 推送分支 (`git push origin feature/YourFeature`)
5. 创建 Pull Request

## 许可证

MIT 许可证，详情请见 LICENSE 文件。
```

这个README文件包含了项目的基本信息、安装方法、使用示例、API文档、贡献指南和许可证信息，可以帮助用户快速了解和使用ErrorX库。

        