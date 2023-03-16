package Logger

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
)

// 定义不同级别的logger
var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func Init() {
	// 创建一个lumberjack对象
	lj := &lumberjack.Logger{
		Filename:   "app.log", // 日志文件名
		MaxSize:    10,        // 每个日志文件最大10MB
		MaxBackups: 5,         // 最多保留5个备份
		MaxAge:     30,        // 最多保留30天
		LocalTime:  true,      // 使用本地时间戳
	}

	// 设置infoLogger输出到lumberjack对象和标准输出，并且带有时间和INFO前缀
	InfoLogger = log.New(lj, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger.SetOutput(io.MultiWriter(lj, os.Stdout))

	// 设置errorLogger输出到lumberjack对象和标准错误，并且带有时间和ERROR前缀
	ErrorLogger = log.New(lj, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger.SetOutput(io.MultiWriter(lj, os.Stderr))
}

func demo() {
	// 使用infoLogger记录一些普通信息
	InfoLogger.Println("程序启动")
	InfoLogger.Println("正在处理请求")

	// 使用errorLogger记录一些错误信息
	ErrorLogger.Println("发生了一个错误")
	ErrorLogger.Println("程序退出")
}
