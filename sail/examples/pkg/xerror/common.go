// Copyright 2024 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package xerror

var (
	Success       = NewError(0, "成功")
	ServerError   = NewError(10000, "服务内部错误")
	InvalidParams = NewError(10001, "入参错误")
	NotFound      = NewError(10002, "找不到")
)
