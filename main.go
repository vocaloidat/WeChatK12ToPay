package main

import (
	"K12_P/Logger"
	"K12_P/router"
)

func main() {
	Logger.Init() // 初始化日志系统
	r := router.SetupRouter()
	r.Run(":1088")
}
