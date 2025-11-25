// Package ghelper
//
// ----------------develop info----------------
//
//	@Author Jimmy
//	@DateTime 2025-11-25 11:02
//
// --------------------------------------------
package ghelper

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// GetStartPath 获取查找的起始路径：
// - 如果运行的是编译后的二进制（非临时路径），使用二进制所在目录
// - 如果是 go run 产生的临时二进制，则使用当前源文件所在目录
func GetStartPath() string {
	exePath, err := os.Executable()
	if err == nil && !IsTempExecutable(exePath) {
		// 正常二进制：使用其所在目录
		return filepath.Dir(exePath)
	}

	// go run 场景：回退到源文件目录
	_, srcFile, _, ok := runtime.Caller(0)
	if ok {
		return filepath.Dir(srcFile)
	}
	return ""
}

// IsTempExecutable 判断是否为 go run 生成的临时二进制
func IsTempExecutable(path string) bool {
	// 临时二进制通常包含 "go-build" 或位于系统临时目录
	return strings.Contains(path, "go-build") ||
		strings.HasPrefix(path, os.TempDir())
}
