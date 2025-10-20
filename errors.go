package main

import "errors"

var (
	// ErrKeyNotFound 表示键未找到的错误
	ErrKeyNotFound = errors.New("key not found")

	// ErrConnection 表示连接错误
	ErrConnection = errors.New("connection error")

	// ErrInvalidParameter 表示参数无效错误
	ErrInvalidParameter = errors.New("invalid parameter")
)
