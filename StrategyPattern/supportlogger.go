package StrategyPattern

import "fmt"

// LogManager 实现一个日志记录器
type LogManager struct {
	Logging
}

func NewLogManager(logging Logging) *LogManager {
	return &LogManager{logging}
}

// Logging 抽象的日志
type Logging interface {
	Info()
	Error()
}

// FileLogging 实现具体的日志: 文件方式的日志
type FileLogging struct {
}

func (fl *FileLogging) Info() {
	fmt.Println("文件记录Info日志")
}

func (fl *FileLogging) Error() {
	fmt.Println("文件记录Error日志")
}

// DBLogging 实现具体的日志: 数据库方式的日志
type DBLogging struct {
}

func (dbl *DBLogging) Info() {
	fmt.Println("数据库记录Info日志")
}

func (dbl *DBLogging) Error() {
	fmt.Println("数据库记录Error日志")
}
